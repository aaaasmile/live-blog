import {Login} from './login.js'


const routes = [
    { path: '/login', component: Login }
]

const router = new VueRouter({
    routes // short for `routes: routes`
})

export const app = new Vue({
    el: '#app',
    router,
    template: `
  <div>
    <h2>Live Blog</h2>
    <h3>Actions</h3>
    <ul>
      <li>Login</li>
      <li>Cloud</li>
      <li>Blog live</li>
    </ul>
  </div>`
})

console.log('Main is here!')