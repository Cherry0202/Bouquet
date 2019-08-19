exports.handler = (event, context, callback) => {
    var response = {
        statusCode : 200,
        headers: {
            "Access-Control-Allow-Origin" : "*"
        },
        body: JSON.stringify({"message" : ""})
    };

    //TODO: 変数rbodyにリクエストボディのオブジェクトを代入してください。

    //TODO: responseオブジェクトのbodyプロパティに変数rbodyを代入

    //コールバック関数でレスポンスを返す
    callback(null, response);
};
