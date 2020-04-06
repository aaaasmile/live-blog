export default {
  data() {
    return {
      drawer: false,
      links: [
        { icon: 'dashboard', text: 'Dashboard', route: '/' },
      ]
    }
  },
  template: `
  <nav>
    <v-app-bar dense flat>
      <v-btn text color="grey"><v-icon>menu</v-icon></v-btn>
      <v-toolbar-title class="text-uppercase grey--text">
        <span class="font-weight-light">Live</span>
        <span>Blog</span>
      </v-toolbar-title>
      <v-spacer></v-spacer>
      <v-btn text color="grey">
        <span>Sign Out</span>
        <v-icon right>exit_to_app</v-icon>
      </v-btn>
    </v-app-bar>
  </nav>`
}