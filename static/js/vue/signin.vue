<template>
  <div>
    <h2>Sign In</h2>
    <div>
      <div>
        <label for="username">Username</label>
        <input id="username" v-model="username" type="text" name="username" />
      </div>
      <div>
        <label for="password">Password</label>
        <input id="password" v-model="password" type="password" name="password" />
      </div>
      <button v-on:click="SignIn">Sign In</button>
    </div>
  </div>
</template>