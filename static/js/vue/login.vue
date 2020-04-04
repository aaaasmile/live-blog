<template>
  <div>
    <h2>Login</h2>
    <div>
      <div>
        <label for="username">Username</label>
        <input id="username" v-model="username" type="text" name="username" />
      </div>
      <div>
        <label for="password">Password</label>
        <input id="password" v-model="password" type="password" name="password" />
      </div>
      <button v-on:click="Login">Login</button>
    </div>
  </div>
</template>