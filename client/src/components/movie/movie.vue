<template>
    <div class="container">
        <section class="card" v-for="item in list">
                <video class="card-img" :name="item.name" preload="auto" controls :poster="item.imgPath">
                    <source :src="item.path" type="video/mp4"/>
                    <p>ご利用のブラウザはvideo タグによる動画の再生に対応していません。</p>
                </video>
            <div class="card-content">
                <h1 class="card-title">{{item.name}}</h1>
                <img :src="/static/">
            </div>
            <div class="card-link">
                <a href="http://webcreatorbox.com/about">About</a>
                <a href="http://webcreatorbox.com/">Website</a>
            </div>
        </section>
    </div>
</template>

<script lang="ts">
    import { Vue, Component, Prop } from "vue-property-decorator";
    import {Fetch} from "../../api/fetch"
    import MovieList  from './../../models/MovieList'
    @Component
    export default class Movie extends Vue{
        list  = [] as MovieList[];
        constructor(){
            super();
            let f = new Fetch;
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
        getIconPath(name:string){
            return "/static/"
        }
    };
</script>

<style scoped lang="scss">
    .container{
        display: flex;
        flex-flow: column wrap;
        justify-content:center;
    }
    .card {
        width: 350px;
        background: #fff;
        border-radius: 5px;
        box-shadow: 0 2px 5px #ccc;
        display: inline-flex;
    }
    .card-img {
        border-radius: 5px 5px 0 0;
        max-width: 100%;
        height: auto;
    }
    .card-content {
        padding: 20px;
    }
    .card-title {
        font-size: 20px;
        margin-bottom: 20px;
        text-align: center;
        color: #333;
    }
    .card-text {
        color: #777;
        font-size: 14px;
        line-height: 1.5;
    }
    .card-link {
        text-align: center;
        border-top: 1px solid #eee;
        padding: 20px;
    }
    .card-link a {
        text-decoration: none;
        color: #0bd;
        margin: 0 10px;
    }
    .card-link a:hover {
        color: #0090aa;
    }

</style>