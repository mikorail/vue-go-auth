<template>
  <v-app id="inspire">
    <v-main>
     <v-switch
       class="d-inline-flex pa-2 float-right"
       v-model="$vuetify.theme.dark"
       prepend-icon=mdi-lamp-outline
       append-icon="mdi-lamp"
     ></v-switch>
      <v-container class="fill-height">
        <v-row align="center" justify="center">
          <v-col cols="12" sm="8" md="4">
            <v-card class="elevation-12">
              <v-toolbar color="success" dark flat>
                <v-toolbar-title>Login form</v-toolbar-title>
                <v-spacer></v-spacer>
              </v-toolbar>
              <v-card-text>
               <v-alert v-if="message" type="error">
                  {{message.error}}
                </v-alert>
                <v-form @submit.prevent="handleLogin"
                  lazy-validation
                >
                  <v-text-field
                    v-model="user.email"
                    label="Email"
                    :rules="emailRules"
                    name="email"
                    prepend-icon="mdi-account"
                    type="text"
                    required
                  ></v-text-field>
                  <v-text-field
                    v-model="user.password"
                    id="password"
                    label="Password"
                    name="password"
                    :rules="passwordRules"
                    required
                    prepend-icon="mdi-lock"
                    :append-icon="value ? 'mdi-eye' : 'mdi-eye-off'"
                    @click:append="() => (value = !value)"
                    :type="value ? 'password' : 'text'"
                  ></v-text-field>
                  <span
                    >Don't have an account? 
                    <br>Register
                    <router-link to="/register"> here. </router-link></span
                  >
                  <v-card-actions>
                    <v-spacer></v-spacer>
                    <v-btn type="submit" color="success">Login</v-btn>
                  </v-card-actions>
                </v-form>
              </v-card-text>
            </v-card>
          </v-col>
        </v-row>
      </v-container>
    </v-main>
  </v-app>
</template>

<script>
import User from '../common/models/user';

export default {
  name: 'LoginForm',
  data() {
    return {
      value: String,
      user: new User('', ''),
      loading: false,
      message: '',

      emailRules: [
        v => !!v || 'E-mail is required',
        v => /.+@.+\..+/.test(v) || 'E-mail must be valid',
      ],
      passwordRules: [
        v => !!v || 'Password is required',
      ]


    };
  },
  computed: {
    loggedIn() {
      return this.$store.state.auth.status.loggedIn;
    },
  },
  created() {
    if (this.loggedIn) {
      this.$router.push('/home');
    }
  },
  methods: {
    handleLogin() {
      if (this.user.email && this.user.password) {
        this.$store.dispatch('auth/login', this.user).then(
          () => {
            this.$router.push('/home');
          },
          error => {
            this.loading = false;
            this.message =
              (error.response && error.response.data) ||
              error.message ||
              error.toString();
          }
        );
      }
      else if(!this.user.email && !this.user.password)
        this.message.error = 'Please fill the form';
      else if(!this.user.password)
        this.message.error = 'Please fill password';
      else
        this.message.error = 'Please fill email';
    },
  },
};
</script>
