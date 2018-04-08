<template>
    <div class="container">
        <div class="card" v-for="item in list" >
            <div :href="item.url">
                <div class="card-title">{{item.name}}</div>
                <div class="card-foot">
                    <img :src="getIconPath(item.language)">
                    <div class="card-time">{{item.updatedAt}}</div>
                </div>
            </div>
        </div>
    </div>
</template>

<script lang="ts">
    import { Vue, Component, Prop } from "vue-property-decorator";
    import {Fetch} from "../../api/fetch"
    import GithubList from './../../models/GithubList'
    @Component
    export default class Github extends Vue{
        list  = [] as GithubList[];
        constructor(){
            super();
            let f = new Fetch;
            f.fetchGithub().then((res:any)=>{
                res.data.map((el:any)=> {
                    this.list.push(new GithubList(el.name,el.language,el.html_url,el.updated_at))
                })
            });
        }
        getIconPath(name:string){
           switch (name){
               case "":
                   return  "/static/img/icon/language/Github.png";
               case "C":
                   return  "/static/img/icon/language/Github.png";
               case "C++":
                   return  "/static/img/icon/language/Github.png";
               case "Objective-C":
                   return  "/static/img/icon/language/Github.png";
               case "Vue":
                   return "/static/img/icon/language/Github.png";
               default:
                   return "/static/img/icon/language/"+name+".png";
           }
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
        color: #0bd;
        margin: 0 10px;
    }

    .card-link a:hover {
        color:rgb(49,103,69);
    }

    html, css{
        touch-action:none;
    }
</style>