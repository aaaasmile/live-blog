import { Login } from './login.js'
import { Upload } from './upload.js'

const routes = [
    { path: '/login', component: Login },
    { path: '/cloud', component: Upload }
]

export const app = new Vue({
    el: '#app',
    router: new VueRouter({ routes }),
    vuetify: new Vuetify(),
    data() {
        return {
            Buildnr: "",
            username: ""
        }
    },
    mounted() {
        this.Buildnr = window.buildnr
    },
    methods:{
        Logout () {
            console.log('Execute logout.')
            let req = { Username: this.username}
            this.$http.post("Logout", JSON.stringify(req), { headers: { "content-type": "application/json" } }).then(result => {
              this.response = "Status: " + result.data.Status + "\n";
              console.log('Call terminated ', result.data)
            }, error => {
              console.error(error);
            });
          }
    },
    template: `
  <v-app>
    <h2>Live Blog</h2>
    <h3>Actions</h3>
    <ul>
      <li><router-link to="/login">Login</router-link></li>
      <li><router-link to="/cloud">Cloud</router-link></li>
      <li>Blog live</li>
      <v-btn @click="Logout">Logout</v-btn>
    </ul>
    <router-view></router-view>
    <div>
      <p>Buildnr: {{Buildnr}}</p>
    </div>
  </v-app>
`
})

console.log('Main is here!')