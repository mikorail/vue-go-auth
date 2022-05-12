<template>
  <v-app id="inspire">
    <v-main>
    <v-switch
       class="d-inline-flex pa-2 float-right"
       v-model="$vuetify.theme.dark"
       prepend-icon=mdi-lamp-outline
       append-icon="mdi-lamp"
     ></v-switch>
      <v-dialog
        v-model="dialog"
        max-width="290"
      >
        <v-card>
          <v-card-title class="headline">Account created!</v-card-title>

          <v-card-text>
            Congratulation, you can now login with your login information.
          </v-card-text>

          <v-card-actions>
            <v-spacer></v-spacer>

            <v-btn
              color="green darken-1"
              text
              @click="handleLogin"
            >
              OK
            </v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>
      <v-container
        class="fill-height"
      >
        <v-row
          align="center"
          justify="center"
        >
          <v-col
            cols="12"
            sm="8"
            md="4"
          >
            <v-card class="elevation-12">
              <v-toolbar
                color="success"
                dark
                flat
              >
                <v-toolbar-title>Register form</v-toolbar-title>
                <v-spacer></v-spacer>
              </v-toolbar>
              <v-card-text>
                <v-alert v-if="message" type="error">
                  {{message.error}}
                </v-alert>
                <v-form @submit.prevent="handleRegister"
                  lazy-validation>
                  <v-text-field
                    v-model="user.email"
                    label="Email"
                    name="email"
                    prepend-icon="mdi-account"
                    type="text"
                    required
                    :rules="emailRules"
                  ></v-text-field>

                  <v-text-field
                    v-model="user.username"
                    id="sername"
                    label="Username"
                    name="username"
                    prepend-icon="mdi-at"
                    type="text"
                     required
                    :rules="usernameRules"
                  ></v-text-field>

                  <v-text-field
                    v-model="user.password"
                    id="password"
                    label="Password"
                    name="password"
                    prepend-icon="mdi-lock"
                    :append-icon="value ? 'mdi-eye' : 'mdi-eye-off'"
                    @click:append="() => (value = !value)"
                    :type="value ? 'password' : 'text'"
                     required
                    :rules="passwordRules"
                  ></v-text-field>

                  <span>Already have an account ?
                  <br> Login <router-link to="/">here.</router-link></span>
                  <v-card-actions>
                    <v-spacer></v-spacer>
                    <v-btn type="submit" color="success">Register</v-btn>
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
  name: 'RegisterForm',
  data() {
    return {
      value: String,
      user: new User('', '', ''),
      submitted: false,
      successful: false,
      message: '',
      dialog: false,
      
      emailRules: [
        v => !!v || 'E-mail is required',
        v => /.+@.+\..+/.test(v) || 'E-mail must be valid',
      ],
      usernameRules: [
        v => !!v || 'Username is required',
      ],
      passwordRules: [
        v => !!v || 'Password is required',
      ]

    };
  },
  computed: {
    loggedIn() {
      return this.$store.state.auth.status.loggedIn;
    }
  },
  created() {
    if (this.loggedIn) {
      this.$router.push('/home');
    }
  },
  methods: {
    handleRegister() {
      this.message = '';
      this.submitted = true;
      this.$store.dispatch('auth/register', this.user).then(
        () => {
          this.$router.push('/login');
        },
        error => {
          if (error.response === undefined) {
            this.dialog = true;
          }else {
            this.message =
              (error.response && error.response.data) ||
              error.message ||
              error.toString();
            this.successful = false;
          }
        }
      );
    },
    handleLogin() {
      this.dialog = false;
      this.$router.push('/login');
    }
  }
};
</script>