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
    <v-content>
      <div class="home">
        <h2>Live-blog</h2>
        <router-link to="/">Home</router-link>
        <router-link to="/signin">Sign In</router-link>
      </div>
       <router-view></router-view>
      <p>Buildnr: {{Buildnr}}</p>
    </v-content>
  </v-app>`
})

console.log('Main is here!')