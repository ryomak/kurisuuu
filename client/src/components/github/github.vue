<template>
    <div class="container">
        <transition name="fade">
            <div class="loading" v-show="showLoading">Loading</div>
        </transition>
        <transition-group name="list" tag="p">
        <div class="card" :key="item.name" v-on:click="goGithub(item.url)" v-for="item in list" >
                <div class="card-title">{{item.name}}</div>
                <div class="card-foot">
                    <img :src="getIconPath(item.language)">
                    <div class="card-time">{{item.updatedAt}}</div>
                </div>
        </div>
        </transition-group>
    </div>
</template>

<script lang="ts">
    import { Vue, Component, Prop } from "vue-property-decorator";
    import {Fetch} from "../../api/fetch";
    import GithubList from './../../models/GithubList';
    import * as Const from "../../const/const";
    @Component
    export default class Github extends Vue{
        list  = [] as GithubList[];
        showLoading:boolean;
        c = Const;
        constructor(){
            super();
            this.showLoading = true;
            let f = new Fetch;
            f.fetchGithub().then((res:any)=>{
                res.data.map((el:any)=> {
                    this.list.push(new GithubList(el.name,el.language,el.html_url,el.updated_at))
                })
            }).then(()=>{
                this.closeLoading();
            });
        }
        getIconPath(name:string){
            let githubIcon = this.c.GITHUB_ICON;
           switch (name){
               case "":
                   return  githubIcon;
               case "C":
                   return  githubIcon;
               case "C++":
                   return  githubIcon;
               case "Objective-C":
                   return  githubIcon;
               case "Vue":
                   return githubIcon;
               default:
                   return "/static/img/icon/language/"+name+".png";
           }
        }
        goGithub(url :string){
            location.href = url;
        }
        closeLoading(){
            this.showLoading = false;
        }

    };
</script>

<style scoped>
    .container {
        display: flex;
        flex-wrap: wrap;
        justify-content:center ;
    }
    .card {
        display: flex;
        width: 250px;
        background: #fff;
        border-radius: 10px;
        box-shadow: 7px 7px rgb(49,103,69);
        flex-direction: column;
        margin: 20px;
        border:solid 2px rgb(49,103,69);
    }

    .card-title {
        font-size: 30px;
        text-align: center;
        padding-top: 20px;
        color: rgb(49,103,69);
        text-decoration: none;
    }

    .card-time{
        font-size: 15px;
        color: rgb(49,103,69);
        text-decoration: none;
    }
    .card-foot{
        display: flex;
        justify-content: center;
        padding: 10px;
    }

    .card-link a {
        text-decoration: none;
        justify-content: space-around;
        margin: 0 10px;
    }

    .card-link a:hover {
        color:rgb(49,103,69);
    }

    html, css{
        touch-action:none;
    }
    /*transition*/
    .list-enter-active, .list-leave-active {
        transition: all 1s;
    }
    .list-enter, .list-leave-to /* .list-leave-active for below version 2.1.8 */ {
        opacity: 0;
        transform: translateY(30px);
    }
    .fade-enter-active, .fade-leave-active {
        transition: opacity .5s;
    }
    .fade-enter, .fade-leave-to /* .fade-leave-active below version 2.1.8 */ {
        opacity: 0;
    }

    p{
        display: flex;
        flex-wrap: wrap;
        justify-content:center ;
    }
    .loading{
        font-size: 30px;
        color:rgb(49,103,69);
        position: fixed;
    }
</style>