let vm = new Vue({
    el: "#calendar", // Vue.jsを使うタグのIDを指定
    data: {
        // Vue.jsで使う変数はここに記述する
        attrs: [
            {
                key: 'today',
                highlight: {
                    animated: true,
                    height: '1.8rem',
                    backgroundColor: 'red',
                    borderColor: null,
                    borderWidth: '1px',
                    borderStyle:'solid',
                    borderRadius:'1.8rem',
                    opacity: 1
                },
                dates: new Date(),
                popover: {
                    label: 'メッセージを表示できます',
                },
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
