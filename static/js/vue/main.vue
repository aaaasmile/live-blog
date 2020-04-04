<template>
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
  </v-app>
</template>