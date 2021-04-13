import { createRouter, createWebHashHistory, createWebHistory } from 'vue-router';
import Home from '/@/views/Home.vue';
import Posts from '/@/views/Posts.vue';
import About from '/@/views/About.vue';

export default createRouter({
  history: createWebHistory(),
  routes: [
    { path: '/', name: 'Home', component: Home },
    { path: '/posts/:slug', name: 'Posts', component: Posts , props: true},
    { path: '/about', name: 'About', component: About },
  ],
});
