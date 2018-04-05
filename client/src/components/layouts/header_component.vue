<template>
    <div>
        <ul>
            <li v-for="link in links" :key="link.path">
                <router-link :to="link.path">{{link.name}}</router-link>
            </li>
        </ul>
    </div>
</template>

<script lang="ts">
    import { Vue, Component, Prop } from "vue-property-decorator";
    import {Link} from "./link";
    @Component
    export default class HeaderComponent extends Vue{
        links :Link[];
        constructor(){
            super();
            this.links = [
                new Link("りーどみ", "/"),
                new Link("きーた","/qiita"),
                new Link("ぎぶはぶ","/github"),
                new Link("ど-が","/movie")
            ];
        }
    };
</script>

<style scoped>

</style>