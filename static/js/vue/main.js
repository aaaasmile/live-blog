export const app = new Vue({
    el: '#app',
    data: {
        username: '',
        password: ''
    },
    methods: {
        Login: function (event) {
            console.log('Execute login.')
            let req = {Username: this.username, Password: this.password}
            this.$http.post("Login", JSON.stringify(req), { headers: { "content-type": "application/json" } }).then(result => {
                this.response = "Status: " + result.data.Status + "\n";
                console.log('Call terminated ', result.data)
            }, error => {
                console.error(error);
            });
        }
    },
    template: `
        <div>
            <h2>Live Blog</h2>
            <div>
                <div>
                    <label for="username">Username</label>
                    <input id="username" v-model="username" type="text" name="username">
                </div>
                <div>
                    <label for="password">Password</label>
                    <input id="password" v-model="password" type="password" name="password">
                </div>
                <button v-on:click="Login">Login</button>
            </div>
            <h3>Upload</h3>
            <form
                enctype="multipart/form-data"
                action="upload"
                method="post">
                <input type="file" name="myFile" />
                <input type="submit" value="upload" />
            </form>
        </div>
    `
})

console.log('Main is here!')