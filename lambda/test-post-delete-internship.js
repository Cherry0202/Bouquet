var AWS = require("aws-sdk");
var dynamo = new AWS.DynamoDB.DocumentClient();
var tableName = "post";

exports.handler = (event, context, callback) => {
    var response = {
        statusCode: 200,
        headers: {
            "Access-Control-Allow-Origin" : "*"
        },
        body: JSON.stringify({"message" : "success"})
    };

    // if (!event.headers || event.headers.Authorization !== "mti2019"){
    //     response.statusCode = 400;
    //     response.body = JSON.stringify({"message": "tokenが違います"});
    //     callback(null, response);
    //     return;
    // }
    
    var body = JSON.parse(event.body);
    var userId = body.userId;
    var timestamp = parseInt(body.timestamp);

    //TODO: 削除対象のテーブル名と削除したいデータのkeyをparamに設定
    var param = {
        TableName: tableName,
        Key: {
            userId: userId,
            timestamp: timestamp
        }
    };

    console.log(body.userId);
    console.log(body.timestamp);
    
    

    //dynamo.delete()を用いてデータを削除
    dynamo.delete(param, function(err, data){
        if(err){
            //TODO: 削除に失敗した場合の処理を記述
            response.statusCode = 500;
            response.body = JSON.stringify({"message": "予期せぬエラーが発生しました"});
            callback(null, response);
            return;
        }else{
            //TODO: 削除に成功した場合の処理を記述
            callback(null, response);
            return;
        }
    });
};
