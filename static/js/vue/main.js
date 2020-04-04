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
        <h2>Testing</h2>
        <v-btn color="pink">Rosa o bella</v-btn>
        <v-btn>Cliccami</v-btn>
        <v-btn class="pink white--text">
          <v-icon left small>email</v-icon>
        </v-btn>
        <v-btn fab dark small depressed color="purple">
          <v-icon dark>favorite</v-icon>
        </v-btn>
        <router-link to="/">Home</router-link>
        <router-link to="/signin">Sign In</router-link>
      </div>
       <router-view></router-view>
      <p>Buildnr: {{Buildnr}}</p>
    </v-content>
  </v-app>`
})

console.log('Main is here!')