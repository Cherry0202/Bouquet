// ビューオブジェクト生成
var vm = new Vue({
    el: "#app", // Vue.jsを使うタグのIDを指定
    data: {
    // Vue.jsで使う変数はここに記述する
        user: {
            user_id:null,
            height: null,
            position: null,
            wedding_day: null,
            weight: null
        }
    },
    watch: {
        user_date: function () {
        // データの保存
            localStorage.setItem('wedding_day', vm.user_date);
        },
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
            // ログイン処理はここに
            // APIにPOSTリクエストを送る
            fetch(url + "/bouquet/user/personal", {
                method: "POST",
                body: JSON.stringify({
                    "user_id":localStorage.getItem('user_id', vm.user.user_id),
                    "height":Number(vm.user.height),
                    "position":vm.user.part,
                    "wedding_day":vm.user_date,
                    "weight": Number(vm.user.weight),
                    "goal_weight": Number(vm.user.best_weight)
                })
            })
                .then(function(response) {
                    if (response.status == 200) {
                        return response.json();
                    }
                    // 200番以外のレスポンスはエラーを投げる
                    return response.json().then(function(json) {
                        throw new Error(json.message);
                    });
                })
                .then(function(json) {
                // レスポンスが200番で返ってきたときの処理はここに記述する
                    var content = JSON.stringify(json, null, 2);
                    //var content = JSON.stringify(json);
                    console.log(content);
                    console.log(json);
                    // localStorage.setItem('user_id', vm.user.user_id);
                    // localStorage.setItem('height', Number(vm.user.height));
                    // localStorage.setItem('position', vm.user.part);
                    // localStorage.setItem('wedding_day', vm.user.date);

                    // localStorage.setItem('wedding_day', vm.user.date);
                    // console.log(localStorage.getItem('wedding_day'));

                    // localStorage.setItem('wedding_day', vm.user.date);
                    // localStorage.setItem('wedding_day', vm.user_date);
                    // // localStorage.setItem('weight', Number(vm.user.weight));

                    // カレンダーへ遷移
                    // location.href = "./calendar.html";

                })
                .catch(function(err) {
                // レスポンスがエラーで返ってきたときの処理はここに記述する
                    console.log(err);

                    return;
                });
        }
    }
});
