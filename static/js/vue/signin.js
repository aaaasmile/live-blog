export const SignIn = Vue.component('signin', {
  data() {
    return {
      username: '',
      password: ''
    }
  },
  methods: {
    SignIn: function (event) {
      console.log('Execute request auth token (Sign In).')
      let req = { Username: this.username, Password: this.password }
      this.$http.post("Token", JSON.stringify(req), { headers: { "content-type": "application/json" } }).then(result => {
        this.response = "Status: " + result.data.Status + "\n";
        console.log('Call terminated ', result.data)
      }, error => {
        console.error(error);
      });
    }
  },
  template: `
  <div>
    <h2>Sign In</h2>
    <div>
      <div>
        <label for="username">Username</label>
        <input id="username" v-model="username" type="text" name="username" />
      </div>
      <div>
        <label for="password">Password</label>
        <input id="password" v-model="password" type="password" name="password" />
      </div>
      <button v-on:click="SignIn">Sign In</button>
    </div>
  </div>`
})
