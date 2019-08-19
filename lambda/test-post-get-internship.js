var AWS = require("aws-sdk");
var dynamo = new AWS.DynamoDB.DocumentClient();
var tableName = "post";

exports.handler = (event, context, callback) => {
    var response = {
        statusCode: 200,
        headers: {
            "Access-Control-Allow-Origin" : "*"
        },
        body: JSON.stringify({"message" : ""})
    };

    var userId = event.queryStringParameters.userId;
    var start = parseInt(event.queryStringParameters.start, 10);
    var end = parseInt(event.queryStringParameters.end, 10);
    var category = event.queryStringParameters.category;

    if (!userId) {
        callback(null, "userIdは必須です");
        return;
    }

    if (userId && start && !end) {
        callback(null, "終了時刻も入力してください");
        return;
    }

    var param;
    //TODO: query()に渡すparamを宣言
    if (!start) {
        param = {
            "TableName" : tableName,
            //キー、インデックスによる検索の定義
            "KeyConditionExpression" : "userId = :userId",
            //検索値のプレースホルダの定義
            "ExpressionAttributeValues" : {
                ":userId": userId
            }
        };
    } else if (category) {
        param = {
            "TableName" : tableName,
            //キー、インデックスによる検索の定義
            "KeyConditionExpression" : "userId = :userId AND (#timestamp BETWEEN :start AND :end)",
            //プライマリーキー以外の属性でのフィルタ
            "FilterExpression" : " #category = :category",
            //属性名のプレースホルダの定義
            "ExpressionAttributeNames" : {
                "#timestamp": "timestamp",
                "#category": "category"
            },
            //検索値のプレースホルダの定義
            "ExpressionAttributeValues" : {
                ":userId": userId,
                ":start": start,
                ":end": end,
                ":category": category
            }
        };
    } else {
        param = {
            "TableName" : tableName,
            //キー、インデックスによる検索の定義
            "KeyConditionExpression" : "userId = :userId AND (#timestamp BETWEEN :start AND :end)",
            //プライマリーキー以外の属性でのフィルタ
            // "FilterExpression" : "",
            //属性名のプレースホルダの定義
            "ExpressionAttributeNames" : {
                "#timestamp": "timestamp"
            },
            //検索値のプレースホルダの定義
            "ExpressionAttributeValues" : {
                ":userId": userId,
                ":start": start,
                ":end": end
            }
        };
    }


    //dynamo.query()を用いてuserIdとpasswordが一致するデータの検索
    dynamo.query(param, function(err, data){
        //userの取得に失敗
        if(err){
            console.log(err);
            response.statusCode = 500;
            response.body = JSON.stringify({"message" : "予期せぬエラーが発生しました"});
            callback(null, response);
            return;
        }
        //TODO: 該当するデータが見つからない場合の処理を記述(ヒント：data.Itemsの中身が空)
        if (!data.Items.length) {
            response.statusCode = 401;
            response.body = JSON.stringify({
                "message" : "条件に合致するものがありません"
            });
            callback(null, response);
            return;
        }

        //TODO: 認証が成功した場合のレスポンスボディとコールバックを記述
        response.body = JSON.stringify(data.Items);
        callback(null, response);
    });
};
