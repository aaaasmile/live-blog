<template>
  <v-app class="grey lighten-4">
    <Navbar />
    <v-content class="mx-4 mb-4">
      <router-view></router-view>
      <p>Buildnr: {{Buildnr}}</p>
    </v-content>
  </v-app>
</template>