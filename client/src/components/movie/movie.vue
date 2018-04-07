<template>
    <div>
        <div v-for="item in list" >
            <div>
                {{item.name}}
                <video :name="item.name" preload="auto" controls :poster="item.imgPath">
                    <source :src="item.path" type="video/mp4"/>
                    <p>ご利用のブラウザはvideo タグによる動画の再生に対応していません。</p>
                </video>
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
    };
</script>

<style scoped>
    .stack-wrapper{
        margin: 0 auto;
        position: relative;
        z-index: 1000;
        width: 320px;
        height: 320px;
        padding: 0;
        list-style: none;
        pointer-events: none;
    }
    .controls{
        position: relative;
        width: 200px;
        text-align: center;
        display:flex;/*Flex布局*/
        display: -webkit-flex; /* Safari */
        justify-content:space-around;
        margin: 0 auto;
        margin-top: 50px
    }
    .controls .button {
        border: none;
        background: none;
        position: relative;
        display: inline-block;
        cursor: pointer;
        font-size: 16px;
        width: 50px;
        height: 50px;
        z-index: 100;
        -webkit-tap-highlight-color: rgba(0,0,0,0);
        border-radius: 50%;
        background: #fff;
    }
    .button .next{
        display: inline-block;
        width: 10px;
        height:5px;
        background: rgb(129,212,125);
        line-height: 0;
        font-size:0;
        vertical-align: middle;
        -webkit-transform: rotate(45deg);
        left: -5px;
        top: 2px;
        position: relative;
    }
    .button .next:after{
        content:'/';
        display:block;
        width: 20px;
        height:5px;
        background: rgb(129,212,125);
        -webkit-transform: rotate(-90deg) translateY(-50%) translateX(50%);
    }
    .button .prev{
        display: inline-block;
        width: 20px;
        height:5px;
        background: rgb(230,104,104);
        line-height: 0;
        font-size:0;
        vertical-align: middle;
        -webkit-transform: rotate(45deg);
    }
    .button .prev:after{
        content:'/';
        display:block;
        width: 20px;
        height:5px;
        background: rgb(230,104,104);
        -webkit-transform: rotate(-90deg);
    }
    .controls .text-hidden {
        position: absolute;
        overflow: hidden;
        width: 0;
        height: 0;
        color: transparent;
        display: block;
    }
</style>