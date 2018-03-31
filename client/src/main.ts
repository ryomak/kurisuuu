import Vue from 'vue'
import Home from './components/home'
import router from './router'

Vue.config.productionTip = false;

/* eslint-disable no-new */
new Vue({
    el: '#app',
    router,
    components: { Home },
    template: '<App/>'
});

