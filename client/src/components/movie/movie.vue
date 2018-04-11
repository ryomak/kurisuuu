<template>
    <div class="container">
       <transition name="fade">
            <div class="loading" v-show="showLoading">Loading</div>
        </transition>
        <transition-group name="list" tag="p">
        <div class="card" v-for="item in list" :key="item.name">
                <img class="card-img" :name="item.name" :src="item.imgPath">
                <p class="card-title">{{item.name}}</p>
            <div class="card-link">
                <img class="card-link-img" :src="playButton" @click="goMovie(item.path)">
                <img class="card-link-img" :src="descriptionButton" @click="openModal(item)">
            </div>
        </div>
        </transition-group>
        <movie-modal v-show="showModal" @close="showModal = false" :detail="modalDetail">
        </movie-modal>
    </div>
</template>

<script lang="ts">
    import { Vue, Component, Prop } from "vue-property-decorator";
    import {Fetch} from "../../api/fetch"
    import MovieModal from "../modal/movie_modal.vue";
    import MovieList  from './../../models/MovieList';
    import * as Const from "../../const/const"
    @Component({
        components:{
            Movie,
            MovieModal
        }
    })
    export default class Movie extends Vue{
        list  = [] as MovieList[];
        showModal:boolean;
        showLoading:boolean;
        modalDetail:MovieList ;
        c = Const;
        playButton = this.c.PLAY_BUTTON;
        descriptionButton = this.c.DESCRIPTION_BUTTON;
        constructor(){
            super();
            this.showModal = false;
            this.showLoading = true;
            let f = new Fetch;
            f.fetchMovie().then((res:any)=>{
                res.data.map((el:any)=> {
                    this.list.push(new MovieList(el.name,el.movie_path,el.image_path,el.description));
                });
            }).then(()=>{
                this.closeLoading();
            });

            this.modalDetail = new MovieList("not",".mp4",".png","選択されていません");
        }
        openModal(item:MovieList){
            this.modalDetail = item;
            this.showModal = true;
        }
        closeLoading(){
            this.showLoading = false;
        }
        goMovie(url:string){
            location.href = url;
        }
    };
</script>

<style scoped lang="scss">
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

    .card-img {
        display: flex;
        border-radius: 10px;
        max-width: 100%;
        height: auto;
        margin: 10px;
        border: 2px solid #eee;
    }

    .card-content {
        padding: 20px;
    }

    .card-title {
        font-size: 20px;
        margin-bottom: 20px;
        text-align: center;
        color: rgb(49,103,69);
    }

    .card-text {
        color: #777;
        font-size: 14px;
        line-height: 1.5;
    }

    .card-link {
        justify-content: space-around;
        display: flex;
        border-top: 2px solid #eee;
        padding: 15px;
    }

    .card-link a {
        text-decoration: none;
        justify-content: space-around;
        color: #0bd;
        margin: 0 10px;
    }
    .card-link-img{
        width: 40px;
        height: 40px;
    }

    .card-link a:hover {
        color:rgb(49,103,69);
    }

    html, css{
        touch-action:none;
    }

    p{
        display: flex;
        flex-wrap: wrap;
        justify-content:center ;
    }
    .loading{
        font-size: 30px;
        color:rgb(49,103,69);
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

</style>