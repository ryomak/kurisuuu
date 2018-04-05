<template>
    <div>
        <div v-for="item in list" >
            <a :href="item.url">
            <div>
                {{item.title}}
            </div>
            </a>
        </div>
    </div>
</template>

<script lang="ts">
    import { Vue, Component, Prop } from "vue-property-decorator";
    import {Fetch} from "../../api/fetch"
    import QiitaList from './../../models/QiitaList'
    @Component
    export default class Qiita extends Vue{
       list  = [] as QiitaList[];
       constructor(){
           super();
           let f = new Fetch;
           this.list = f.getQiita();
       }
    };
</script>

<style scoped>

</style>