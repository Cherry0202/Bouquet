var vm = new Vue({
    el: "#app", // Vue.jsを使うタグのIDを指定
    data: {
        // Vue.jsで使う変数はここに記述する
        users: [],
        user: {
            user_id: localStorage.getItem("user_id"),
            password: null,
            user_name: null,
            goal_weight: null,
            height: null,
            nail_and_extension :{
                title: null,
                start: null,
                end: null
            }
        }
    },
    computed: {
        // 計算した結果を変数として利用したいときはここに記述する
    },
    created: function() {

        // let hoge = localStorage.getItem('weight');
        // console.log(hoge);
        // let hoge2 = localStorage.getItem('user_id');
        // console.log(hoge2);
        // let hoge3 = localStorage.getItem('height');
        // console.log(hoge3);
        // let hoge4 = localStorage.getItem('position');
        // console.log(hoge4);
        // let hoge5 = localStorage.getItem('wedding_day');
        // console.log(hoge5);

        // Vue.jsの読み込みが完了したときに実行する処理はここに記述する
        // APIにGETリクエストを送る
        fetch(url + "/bouquet/user/calendar" +
                "?user_id=" + localStorage.getItem('user_id'), {
            method: "GET"
        })
            .then(function (response) {
                if (response.status == 200) {
                    return response.json();
                }
                // 200番以外のレスポンスはエラーを投げる
                return response.json().then(function (json) {
                    throw new Error(json.message);
                });
            })
            .then(function (json) {
                // レスポンスが200番で返ってきたときの処理はここに記述する
                // vm.user.userId = json.Item.userId;
                // vm.user.nickname = json.Item.nickname;
                // vm.user.age = json.Item.age;
                console.log("200OK");
                console.log("json");
                console.log(json);
                console.log("height");
                console.log(json.height);
                console.log("weight");
                console.log(json.weight);
                console.log("wedding_day");
                console.log(json.wedding_day);
                console.log("position");
                console.log(json.position);
                console.log("goal_weight");
                console.log(json.goal_weight);
                vm.users = json;
                var div_weight = Number(json.weight) - Number(json.goal_weight);
                var wedding_day_s = localStorage.getItem('wedding_day');
                var wedding_day = new Date(wedding_day_s);

                console.log(div_weight);
                vm.users.div_weight = div_weight;
                vm.users.wedding_day = wedding_day;
                console.log("vm.users.wedding_day");
                console.log(vm.users.wedding_day);

                var today = new Date();
                var timestamp = today.getTime();
                // var timestamp = today.getTime();
                var wedding_count = (wedding_day - timestamp);

                console.log("wedding_count");
                console.log(wedding_count);
                vm.users.wedding_count = wedding_count;
                // (day2 - day1) / 86400000

                // console.log(vm.users.Bridal_beauty_treatment_salon);
                // console.log(vm.users.Bridal_beauty_treatment_salon.start);
                // console.log(vm.users.Bridal_beauty_treatment_salon.title);
                // console.log(vm.users.Bridal_beauty_treatment_salon.end);
                // console.log(vm.Nail_and_extetiton);
                // console.log(vm.users.Nail_consideration);
                // console.log(vm.users.Salon_consideration);



            })
            .catch(function (err) {
                // レスポンスがエラーで返ってきたときの処理はここに記述する
            });
    },
    methods: {
    // Vue.jsで使う関数はここで記述する
        deleteUser: function () {
            console.log("delete");
            fetch(url + "/bouquet/user", {
                method: "DELETE",
                headers: new Headers({
                    "Authorization": localStorage.getItem('token')
                }),
                body: JSON.stringify({
                    // userId: localStorage.getItem('user_id')
                    user_id: 'aaa'
                    // password: vm.user.password
                })
            }).then(function (response) {
                if (response.status == 200) {
                    console.log(response);
                    return response.json();
                }
                // 200番以外のレスポンスはエラーを投げる
                return response.json().then(function (json) {
                    throw new Error(json.message);
                });
            }).then(function (json) {
                // レスポンスが200番で返ってきたときの処理はここに記述する
                console.log(json);
                window.alert("本当に退会しますか？");
                location.href = "./index.html";
            }).catch(function (err) {
                console.log(err);
                // レスポンスがエラーで返ってきたときの処理はここに記述する
            });
        }
    }
});
