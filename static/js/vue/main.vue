<template>
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
  </v-app>
</template>