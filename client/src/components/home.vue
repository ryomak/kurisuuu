<template>
    <div class="home">
    <HeaderComponent :path="this.$route.path"/>
    <router-view></router-view>
    <FooterComponent/>
    </div>
</template>

<script lang="ts">
    import HeaderComponent from "./layouts/header_component.vue";
    import FooterComponent from "./layouts/footer_component.vue";
    import {createRouter} from "./../router";
    import { Vue, Component, Prop } from "vue-property-decorator";
    @Component({
        components:{HeaderComponent, FooterComponent},
        router:createRouter(),
    })
    export default class Home extends Vue{

    };
</script>

<style scoped>
.home{
    margin: 0;         /* 余白の削除 */
    padding: 0;        /* 余白の削除 */
    width:100%;
    height:100%;
}
</style>