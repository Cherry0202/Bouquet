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

    var param = {
        TableName: tableName
    };

    //dynamo.scan()で全件取得
    dynamo.scan(param, function (err, data) {
        if (err) {
            //TODO: //データの取得に失敗
            console.log(err);
            response.statusCode = 500;
            response.body = JSON.stringify({
                "message": "予期せぬエラーが発生しました"
            });
            callback(null, response);
            return;
        }

        //TODO: レスポンスボディの設定とコールバックの記述
        response.body = JSON.stringify(data.Items);
        callback(null, response);
    });
};
