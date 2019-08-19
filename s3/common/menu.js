Vue.component('common-menu', {
    props: ["current"],
    template: `
    <div class="ui secondary pointing inverted massive menu" id="header_menu">
      <a class="item" href="./index.html" v-bind:class="{active: current=='home'}">Top</a>
      <a class="item" href="./user.html" v-bind:class="{active: current=='users'}"></a>
      <a class="item" href="./profile.html" v-bind:class="{active: current=='profile'}">Profile</a>
      <h1 align="center">BOUQUET</h1>

    </div>
    `,
    methods: {
        logout: function () {
            localStorage.removeItem('token');
            localStorage.removeItem('userId');
            location.href = "./login.html";
        }
    }
});

// <h1>BOUQUET</h1>
// <div class="right menu"><button class="item" v-on:click="logout">Logout</button></div>
