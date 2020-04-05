export default {
  CallTokenRequest(that, req) {
    that.$http.post("Token", JSON.stringify(req), { headers: { "content-type": "application/json" } }).then(result => {
      that.response = "Status: " + result.data.Status + "\n";
      console.log('Call terminated ', result.data)
      localStorage.setItem('token', result.data.Token.access_token)
      localStorage.setItem('token_refresh', result.data.Token.refresh_token)
      localStorage.setItem('username', result.data.Username)
      localStorage.setItem('expiry', result.data.Token.expiry)
    }, error => {
      console.error(error);
    });
  }
}