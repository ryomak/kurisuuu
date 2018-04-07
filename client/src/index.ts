import Vue from "vue";
import Home from "./components/home.vue";
import {createRouter} from "./router";

new Vue({
    el: "#app",
    template: `<div>
                <home/>
               </div>`,
    router:createRouter(),
    components: {
       Home
    }
});
