
import Dashboard from './views/Dashboard.js'
import Signin from './views/Signin.js'
import Upload from './views/upload.js'
import Navbar from './components/Navbar.js'


const routes = [
  { path: '/', component: Dashboard },
  { path: '/signin', component: Signin },
  { path: '/cloud', component: Upload }
]

export const app = new Vue({
  el: '#app',
  router: new VueRouter({ routes }),
  components: { Navbar },
  vuetify: new Vuetify(),
  data() {
    return {
      Buildnr: ""
    }
  },
  mounted() {
    this.Buildnr = window.buildnr
  },
  methods: {

  },
  template: `
  <v-app class="grey lighten-4">
    <Navbar />
    <v-content class="mx-4 mb-4">
      <router-view></router-view>
      <p>Buildnr: {{Buildnr}}</p>
    </v-content>
  </v-app>
`
})

console.log('Main is here!')