<template>
    <div class="container">
        <transition name="fade">
            <div class="loading" v-show="showLoading">Loading</div>
        </transition>
        <transition-group name="list" tag="p">
            <div class="card" :key="item.id" v-on:click="goQiita(item.url)" v-for="item in list" >
                <div class="card-title">{{item.getTitle()}}</div>
                <div class="card-text">{{item.getSummary()}}</div>
                <div class="card-foot">
                    <div class="card-time">{{item.updatedAt}}</div>
                </div>
            </div>
        </transition-group>
    </div>
</template>

<script lang="ts">
    import { Vue, Component, Prop } from "vue-property-decorator";
    import {Fetch} from "../../api/fetch"
    import QiitaList from './../../models/QiitaList'
    @Component
    export default class Qiita extends Vue{
       list  = [] as QiitaList[];
       showLoading:boolean;
       constructor(){
           super();
           let f = new Fetch;
           this.showLoading = true;
           f.fetchQiita().then((res:any) => {
               res.data.map((el:any)=> {
                   this.list.push(new QiitaList(el.id,el.title,el.body,el.url,el.created_at))
               })
           }).then(()=>{
               this.closeLoading();
           });

       }

       goQiita(url:string){
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
        justify-content:center;
        padding-top: 130px;
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
        font-size: 20px;
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

    .card-text {
        color: #777;
        font-size: 14px;
        line-height: 1.5;
        border-top: 2px solid #eee;
        word-wrap: break-word;
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
        font-size: 50px;
        color:rgb(49,103,69);
    }
</style>