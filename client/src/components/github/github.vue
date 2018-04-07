<template>
    <div>
        <div v-for="item in list" >
            <a :href="item.url">
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
    };
</script>

<style scoped>

</style>