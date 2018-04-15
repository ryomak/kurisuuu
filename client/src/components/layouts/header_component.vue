<template>
    <div>
        <header class="container" v-bind:class="[position>0?'s-header': '']">
            <div v-for="(link,index) in links" :key="link.path" >
                <router-link class="hover"  :to="link.path" v-if="index === getCurrent(path)">{{link.name}}</router-link>
                <router-link class="menu" :to="link.path" v-else>{{link.name}}</router-link>
            </div>
        </header>
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
        position:any;
        constructor(){
            super();
            this.links = [
                new Link("りーどみ", "/"),
                new Link("きーた","/qiita"),
                new Link("ぎぶはぶ","/github"),
                new Link("ど-が","/movie")
            ];
            this.position = 0;
            window.addEventListener('scroll', this.handleScroll);
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

        handleScroll(){
            this.position = window.scrollY
        }
    };
</script>

<style scoped>
    .container{
        position: fixed;
        display: flex;
        justify-content:center;
        width: 100%;
        background-color: white;
    }
    .s-header{
        /*もしheaderしたスクロールで何か必要であれば*/
    }

    .menu{
        padding:20px;
        font-size: 20px;
        border-radius: 10px;
        display: inline-flex;
        background-color: white;
        color: rgb(49,103,69);
        box-shadow: 7px 7px rgb(49,103,69);
        margin: 20px 10px;
        border:solid 2px rgb(49,103,69);
        text-decoration: none;
    }

    .hover{
        padding:20px;
        font-size: 25px;
        border-radius: 10px;
        display: inline-flex;
        background-color:  rgb(49,103,69);
        color: white;
        margin: 20px 10px;
        text-decoration: none;
    }
</style>