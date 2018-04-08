<template>
    <div class="container">
        <div class="card" v-for="item in list" :key="item.name">
                <img class="card-img" :name="item.name" :src="item.imgPath">
                <p class="card-title">{{item.name}}</p>
            <div class="card-link">
                <a :href="item.path">Play</a>
                <div @click="modal.showModal"></div>
            </div>
        </div>
    </div>
</template>

<script lang="ts">
    import { Vue, Component, Prop } from "vue-property-decorator";
    import {Fetch} from "../../api/fetch"
    import MovieModal from "../modal/movie_modal.vue";
    import MovieList  from './../../models/MovieList'
    @Component
    export default class Movie extends Vue{
        list  = [] as MovieList[];
        constructor(){
            super(); 
            let f = new Fetch;
            let modal = new MovieModal;
            f.fetchMovie().then((res:any)=>{
                res.data.map((el:any)=> {
                    this.list.push(new MovieList(el.name,el.movie_path,el.image_path));
                })
            });
        }
        prev(){}
        next(){
            console.log("next");
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
        text-align: center;
        border-top: 2px solid #eee;
        padding: 20px;
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