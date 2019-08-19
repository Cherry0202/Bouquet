var AWS = require("aws-sdk");
var dynamo = new AWS.DynamoDB.DocumentClient();
var tableName = "user";

exports.handler = (event, context, callback) => {
    var response = {
        statusCode: 200,
        headers: {
            "Access-Control-Allow-Origin" : "*"
        },
        body: JSON.stringify({"message" : ""})
    };

    var body = JSON.parse(event.body);

    if (!body){
        response.statusCode = 400;
        response.body = JSON.stringify({"message": "bodyが空です"});
        callback(null, response);
        return;
    }

    var userId = body.userId;
    var age = body.age;
    var password = body.password;
    var nickname = body.nickname;

    if(!userId || !password || !nickname || !age){
        response.statusCode = 400;
        response.body = JSON.stringify ({"message" : "パラメータが足りません"});
        callback(null, response);
        return;
    }

    //TODO: DBに登録するための情報をparamオブジェクトとして宣言する（中身を記述）
    var param = {
        TableName: tableName,
        Item: {
            userId: userId,
            password: password,
            age: age,
            nickname: nickname
        }
    };

    //dynamo.put()でDBにデータを登録
    dynamo.put(param, function(err, data) {
        if (err) {
            //TODO: 登録に失敗した場合の処理を記述
            response.statusCode = 500;
            response.body = JSON.stringify({"message": "予期せぬエラーが発生しました"});
            callback(null, response);
            return;
        } else {
            //TODO: 登録に成功した場合の処理を記述
            param.Item["token"] = "mti2019";
            response.body = JSON.stringify(param.Item);
            callback(null, response);
            return;
        }
    });
};
