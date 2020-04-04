import { SignIn } from './signin.js'
import { Upload } from './upload.js'

const routes = [
    { path: '/signin', component: SignIn },
    { path: '/cloud', component: Upload }
]

export const app = new Vue({
    el: '#app',
    router: new VueRouter({ routes }),
    vuetify: new Vuetify(),
    data() {
        return {
            Buildnr: ""
        }
    },
    mounted() {
        this.Buildnr = window.buildnr
    },
    methods:{
    
    },
    template: `
  <v-app>
    <h2>Live Blog</h2>
    <h3>Actions</h3>
    <ul>
      <li><router-link to="/">Home</router-link></li>
      <li><router-link to="/signin">Sign In</router-link></li>
      <li><router-link to="/cloud">Cloud</router-link></li>
      <li>Blog live</li>
    </ul>
    <router-view></router-view>
    <div>
      <p>Buildnr: {{Buildnr}}</p>
    </div>
  </v-app>`
})

console.log('Main is here!')