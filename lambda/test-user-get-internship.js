var AWS = require("aws-sdk");
var dynamo = new AWS.DynamoDB.DocumentClient();
var tableName = "user";

exports.handler = (event, context, callback) => {
    var response = {
        statusCode : 200,
        headers: {
            "Access-Control-Allow-Origin" : "*"
        },
        body: JSON.stringify({"message" : ""})
    };

    // if (!event.headers || event.headers.Authorization !== "mti2019"){
    //     response.statusCode = 400;
    //     response.body = JSON.stringify({"message": "tokenが違います"});
    //     callback(null, response);
    //     return;
    // }

    var userId = event.queryStringParameters.userId;   //見たいユーザのuserId

    //TODO: 取得対象のテーブル名と検索に使うキーをparamに宣言
    var param = {
        TableName: tableName,
        Key: {
            userId: userId
        }
    };

    //dynamo.get()でDBからデータを取得
    dynamo.get(param, function(err, data){
        if(err){
            //TODO: 取得に失敗した時の処理を記述
            response.statusCode = 500;
            response.body = JSON.stringify({"message": "予期せぬエラーが発生しました"});
            callback(null, response);
            return;
        }
        //TODO: 条件に該当するデータがあればパスワードを隠蔽をする処理を記述
        if (data.Item){
            delete data.Item.password;
        }
        //TODO: レスポンスボディの設定とコールバックを記述
        response.body = JSON.stringify(data);
        callback(null, response);
        return;
    });
};
