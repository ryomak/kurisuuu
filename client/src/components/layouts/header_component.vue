<template>
    <div>
        <div class="container">
            <div v-for="(link,index) in links" :key="link.path" >
                <router-link class="hover"  :to="link.path" v-if="index === getCurrent(path)">{{link.name}}</router-link>
                <router-link class="menu" :to="link.path" v-else>{{link.name}}</router-link>
            </div>
        </div>
    </div>
</template>

<script lang="ts">
    import { Vue, Component, Prop } from "vue-property-decorator";
    import {Link} from "./link";
    @Component({
        props:['path']
    })
    export default class HeaderComponent extends Vue{
        links :Link[];
        current:number;
        constructor(){
            super();
            this.links = [
                new Link("りーどみ", "/"),
                new Link("きーた","/qiita"),
                new Link("ぎぶはぶ","/github"),
                new Link("ど-が","/movie")
            ];
            this.current = this.getCurrent(this.$parent.$route.path);
        }

        getCurrent(uri:string){
            switch (uri){
                case "/":
                    return 0;
                case "/qiita":
                    return 1;
                case "/github":
                    return 2;
                case "/movie":
                    return 3;
                default:
                    return 100;
            }
        }
    };
</script>

<style scoped>
    .container{
        padding-top: 20px;
        display: flex;
        justify-content:center;
    }

    .menu{
        padding:20px;
        font-size: 40px;
        border-radius: 10px;
        display: inline-flex;
        background-color: white;
        color: rgb(49,103,69);
        box-shadow: 5px 10px rgb(49,103,69);
        margin: 20px 10px;
    }

    .hover{
        padding:20px;
        font-size: 40px;
        border-radius: 10px;
        display: inline-flex;
        background-color:  rgb(49,103,69);
        color: white;
        margin: 20px 10px;
    }

</style>