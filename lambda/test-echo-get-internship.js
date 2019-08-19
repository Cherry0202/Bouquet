exports.handler = (event, context, callback) => {
    var response = {
        statusCode : 200,
        headers: {
            "Access-Control-Allow-Origin" : "*"
        },
        body: JSON.stringify({"message" : ""})
    };

    //TODO: 変数qnameにクエリストリングのnameに該当する値を代入してください。

    //TODO: responseオブジェクトのbodyプロパティに変数qnameを代入してください。

    //コールバック関数でレスポンスを返す
    callback(null, response);
};
