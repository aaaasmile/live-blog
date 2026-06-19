import api from "../apicaller.js"

export default {
  data() {
    return {
      username: '',
      password: '',
      showPsw: false
    }
  },
  methods: {
    SignIn: function (event) {
      console.log('Execute request auth token (Sign In).')
      let req = { Username: this.username, Password: this.password }
      api.CallTokenRequest(this, req)
    },
    RefreshToken() {
      console.log('Refresh token.')
      let req = { Token: localStorage.token_refresh }
      api.CallTokenRequest(this, req)
    }
  },
  template: `
  <div>
    <h2>Sign In</h2>
    
  </div>
`
}
