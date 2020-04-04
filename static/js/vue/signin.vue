<template>
  <div>
    <h2>Sign In</h2>
    <v-form>
      <v-container>
        <v-row>
          <v-col cols="12" sm="6" md="3">
            <v-text-field
              label="Username"
              id="username"
              v-model="username"
              type="text"
              name="username"
            ></v-text-field>
          </v-col>
          <v-col cols="12" sm="6" md="3">
            <v-text-field
              :append-icon="show2 ? 'mdi-eye' : 'mdi-eye-off'"
              :type="show2 ? 'text' : 'password'"
              name="password"
              id="password"
              v-model="password"
              label="Password"
              hint="At least 6 characters"
              class="input-group--focused"
              @click:append="show2 = !show2"
            ></v-text-field>
          </v-col>
        </v-row>
      </v-container>
    </v-form>
    <v-icon dark>folder_open</v-icon>
    <v-btn v-on:click="SignIn">Sign In</v-btn>
  </div>
</template>