<template>
    <div class="container">
        <div class="card" v-for="item in list" :key="item.name">
                <video class="card-img" :name="item.name" preload="metadata" :poster="item.imgPath">
                    <source :src="item.path" type="video/mp4"/>
                    <p>ご利用のブラウザはvideo タグによる動画の再生に対応していません。</p>
                </video>
                <p class="card-title">{{item.name}}</p>
            <div class="card-link">
                <a :href="item.path">Play</a>
                <a href="http://webcreatorbox.com/">Website</a>
            </div>
        </div>
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
    };
</script>

<style scoped lang="scss">
    .container {
        display: flex;
        flex-wrap: wrap;
        justify-content: space-around;
    }
    .card {
        display: flex;
        width: 350px;
        background: #fff;
        border-radius: 10px;
        box-shadow: 10px 10px 10px darkgreen;
        flex-direction: column;
        padding: 6px;
    }

    .card-img {
        display: flex;
        border-radius: 10px;
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
        justify-content: space-around;
        color: #0bd;
        margin: 0 10px;
    }

    .card-link a:hover {
        color: #0090aa;
    }
    html, css{
        touch-action:none;
    }
</style>