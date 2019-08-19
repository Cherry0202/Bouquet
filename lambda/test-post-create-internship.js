var AWS = require("aws-sdk");
var dynamo = new AWS.DynamoDB.DocumentClient();
var tableName = "post";

exports.handler = (event, context, callback) => {
    var response = {
        statusCode: 200,
        headers: {
            "Access-Control-Allow-Origin": "*"
        },
        body: JSON.stringify({ "message": "" })
    };

    var body = JSON.parse(event.body);

    if (!body) {
        response.statusCode = 400;
        response.body = JSON.stringify({
            "message": "リクエストボディが空です"
        });
        callback(null, response);
        return;
    }

    if (!body.userId) {
        response.statusCode = 400;
        response.body = JSON.stringify({
            "message": "userIdを入力してください"
        });
        callback(null, response);
        return;
    }

    if (!body.text) {
        response.statusCode = 400;
        response.body = JSON.stringify({
            "message": "textを入力してください"
        });
        callback(null, response);
        return;
    }

    var date = new Date();
    var timestamp = date.getTime();

    //TODO: DBに登録するための情報をparamオブジェクトとして宣言する（中身を記述）
    var param = {
        TableName: tableName,
        Item: {
            userId: body.userId,
            timestamp: timestamp,
            category: body.category,
            text: body.text
        }
    };

    //dynamo.put()でDBにデータを登録
    dynamo.put(param, function (err, data) {
        if (err) {
            //TODO: 登録に失敗した場合の処理を記述
            console.log(err);
            response.statusCode = 500;
            response.body = JSON.stringify({
                "message": "予期せぬエラーが発生しました"
            });
            callback(null, response);
            return;
        } else {
            //TODO: 登録に成功した場合の処理を記述
            response.body = JSON.stringify(param);
            callback(null, response);
        }
    });
};
