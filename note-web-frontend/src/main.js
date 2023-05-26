import Vue from 'vue';
import App from './App.vue';
import store from './store';
import router from './router';
import VueIconFont from 'vue-icon-font'
import './assets/iconfont/iconfont.js'
import infiniteScroll from 'vue-infinite-scroll'
import VueInputAutowidth from 'vue-input-autowidth'

Vue.use(VueIconFont);
Vue.use(infiniteScroll);
Vue.use(VueInputAutowidth);

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app');
