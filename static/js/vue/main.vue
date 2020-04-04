<template>
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
  </v-app>
</template>