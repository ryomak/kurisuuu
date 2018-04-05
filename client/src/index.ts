import Vue from "vue";
import HeaderComponent from "./components/layouts/header_component.vue";
import FooterComponent from "./components/layouts/footer_component.vue";
import Readme from "./components/readme/readme.vue";
import {createRouter} from "./router";

new Vue({
    el: "#app",
    template: `<div>
                <HeaderComponent/>
                <router-view></router-view>
                <FooterComponent/>
               </div>`,
    router:createRouter(),
    components: {
        HeaderComponent,
        FooterComponent
    }
});
