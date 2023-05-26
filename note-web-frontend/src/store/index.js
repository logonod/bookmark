import Vue from 'vue';
import Vuex from 'vuex';
import myModule from './modules/my';

Vue.use(Vuex);

export default new Vuex.Store({
  modules: {
    myModule
  }
});
