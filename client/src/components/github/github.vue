<template>
    <div>
        <div v-for="item in list" >
            <a :href="item.url">
                <img :src="getIconPath(item.language)">
                <div>
                    {{item.name}}
                </div>
            </a>
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

</style>