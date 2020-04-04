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
              :append-icon="showPsw ? 'visibility' : 'visibility_off'"
              :type="showPsw ? 'text' : 'password'"
              name="password"
              id="password"
              v-model="password"
              label="Password"
              hint="At least 6 characters"
              class="input-group--focused"
              @click:append="showPsw = !showPsw"
            ></v-text-field>
          </v-col>
        </v-row>
         <v-btn class="mr-4" v-on:click="SignIn">Sign In</v-btn>
         <v-btn class="mr-4" v-on:click="RefreshToken">Refresh Token</v-btn>
      </v-container>
     
    </v-form>
  </div>
</template>