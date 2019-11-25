// ビューオブジェクト生成
let vm = new Vue({
    el: "#app", // Vue.jsを使うタグのIDを指定
    data: {
        // Vue.jsで使う変数はここに記述する
        mode: "login",
        submitText: "ログイン",
        toggleText: "新規登録",
        user: {
            user_id: null,
            password: null,
            user_name: null,
            re_password: null
        }
    },
    computed: {
        // 計算した結果を変数として利用したいときはここに記述する
    },
    created: function () {
        // Vue.jsの読み込みが完了したときに実行する処理はここに記述する
    },
    methods: {
        // Vue.jsで使う関数はここで記述する
        toggleMode: function () {
            if (vm.mode === "login") {
                vm.mode = "signup";
                vm.submitText = "次へ";
                vm.toggleText = "ログイン";
            } else if (vm.mode === "signup") {
                vm.mode = "login";
                vm.submitText = "ログイン";
                vm.toggleText = "新規登録";
            }
        },
        submit: function () {

            if (vm.mode === "login") {
                // ログイン処理はここに
                // APIにPOSTリクエストを送る
                fetch(url + "/bouquet/user/login", {
                    method: "POST",
                    body: JSON.stringify({
                        "user_id": vm.user.user_id,
                        "password": vm.user.password
                    })
                })
                    .then(function (response) {
                        if (response.status === 200) {
                            return response.json();
                        }
                        // 200番以外のレスポンスはエラーを投げる
                        return response.json().then(function (json) {
                            throw new Error(json.message);
                        });
                    })
                    .then(function (json) {
                        // レスポンスが200番で返ってきたときの処理はここに記述する
                        let content = JSON.stringify(json, null, 2);
                        console.log(content);
                        console.log(json);
                        localStorage.setItem('token', json.token);
                        localStorage.setItem('user_id', vm.user.user_id);
                        // トップページへ遷移
                        location.href = "./calendar.html";

                    })
                    .catch(function (err) {
                        console.log(err);
                        return false;
                    });
            } else if (vm.mode === "signup") {
                // 新規登録時
                // 同じパスワードが入力された時の警告
                if (vm.user.password !== vm.user.re_password) {
                    window.alert("正しいパスワードを入力してください");
                    // location.href = "./index.html";
                    return false;
                }
                // APIにPOSTリクエストを送る
                console.log(vm.user);
                fetch(url + "/bouquet/user/register", {
                    method: "POST",
                    body: JSON.stringify({
                        "user_id": vm.user.user_id,
                        "password": vm.user.password,
                        "user_name": vm.user.user_name
                    })
                })
                    .then(function (response) {
                        if (response.status === 200) {
                            return response.json();
                        }
                        // 200番以外のレスポンスはエラーを投げる
                        return response.json().then(function (json) {
                            throw new Error(json.message);
                        });
                    })
                    .then(function (json) {
                        // レスポンスが200番で返ってきたときの処理はここに記述する
                        let response = JSON.stringify(json);
                        console.log("レスポンス200番OK");
                        console.log(json);
                        console.log(response);
                        console.log(json.user_id);
                        localStorage.setItem('token', json.token);
                        localStorage.setItem('user_id', vm.user.user_id);
                        location.href = "./confirm.html";
                    })
                    .catch(function (err) {
                        // レスポンスがエラーで返ってきたときの処理はここに記述する
                        console.log(err);
                        return false;
                    });
            }
        }
    }
});
