import Vue from 'vue';
import VueRouter from 'vue-router';

import LoginComponent from '../../components/LoginForm'
import RegisterComponent from '../../components/RegisterForm'
import HelloComponent from '../../components/HelloWorld'

Vue.use(VueRouter)

export const router = new VueRouter({
  mode: 'history',
  routes: [
    { path: '/', component: LoginComponent },
    { path: '/home', component: HelloComponent },
    { path: '/register', component: RegisterComponent },
  ]
});

router.beforeEach((to, from, next) => {
  const publicPages = ['/', '/register'];
  const authRequired = !publicPages.includes(to.path);
  const loggedIn = localStorage.getItem('user');

  console.log("to.path",to.path)

  if (authRequired && !loggedIn) {
    next('/');
  } else {
    next();
  }
});