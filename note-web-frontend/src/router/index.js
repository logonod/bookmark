import Vue from 'vue';
import VueRouter from 'vue-router';
import MyPage from '../pages/my/My.vue';

Vue.use(VueRouter);

const router = new VueRouter({
  mode: 'history',
  routes: [
    {
      path: '/',
      component: MyPage
    }
  ]
});

export default router;
