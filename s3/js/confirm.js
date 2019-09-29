// ビューオブジェクト生成
let vm = new Vue({
    el: "#app", // Vue.jsを使うタグのIDを指定
    data: {
    // Vue.jsで使う変数はここに記述する
        user: {
            user_id:null,
            height: null,
            position: "選択してください",
            wedding_day: null,
            weight: null,
            goal_weight: null
        }
    },
    computed: {
    // 計算した結果を変数として利用したいときはここに記述する
    },
    created: function() {
    // Vue.jsの読み込みが完了したときに実行する処理はここに記述する
    },
    methods: {
    // Vue.jsで使う関数はここで記述する
        submit: function () {
            // APIにPOSTリクエストを送る
            fetch(url + "/bouquet/user/personal", {
                method: "POST",
                body: JSON.stringify({
                    "user_id":localStorage.getItem('user_id'),
                    "height":Number(vm.user.height),
                    "position":vm.user.position,
                    "wedding_day":vm.user.wedding_day,
                    "weight": Number(vm.user.weight),
                    "goal_weight": Number(vm.user.goal_weight)
                })
            })
                .then(function(response) {
                    if (response.status === 200) {
                        return response.json();
                    }
                    // 200番以外のレスポンスはエラーを投げる
                    return response.json().then(function(json) {
                        throw new Error(json.message);
                    });
                })
                .then(function(json) {
                // レスポンスが200番で返ってきたときの処理はここに記述する
                    let content = JSON.stringify(json, null, 2);
                    //var content = JSON.stringify(json);
                    console.log(content);
                    console.log(json);

                    // カレンダーへ遷移
                    location.href = "./calendar.html";

                })
                .catch(function(err) {
                // レスポンスがエラーで返ってきたときの処理はここに記述する
                    console.log(err);

                    return false;
                });
        }
    }
});
