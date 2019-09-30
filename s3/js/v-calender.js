let vc = new Vue({
    el: "#calendar", // Vue.jsを使うタグのIDを指定
    data: {
        // Vue.jsで使う変数はここに記述する
        attrs: [
            {
                dates: [
                    new Date(),
                    {
                        start: new Date(2019, 9, 1),
                        end: new Date(2019, 9, 5)
                    },
                    {
                        start: new Date(),
                        span: 5
                    }
                ]
            }
        ],
    },
    computed: {
        // 計算した結果を変数として利用したいときはここに記述する
    },
    created: function() {
        // Vue.jsの読み込みが完了したときに実行する処理はここに記述する
    },
    methods: {
        // Vue.jsで使う関数はここで記述する
    },
});
