export const Upload = Vue.component('upload', {
  data() {
    return {
    }
  },
  methods: {
    Upload: function (event) {
      console.log('Execute upload')
      let req = { Username: this.username, Password: this.password }
      this.$http.post("Login", JSON.stringify(req), { headers: { "content-type": "application/json" } }).then(result => {
        this.response = "Status: " + result.data.Status + "\n";
        console.log('Call terminated ', result.data)
      }, error => {
        console.error(error);
      });
    }
  },
  template: ``
})