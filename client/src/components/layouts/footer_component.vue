<template>
    <footer class="footer">
        <div class="icon" v-for="item in sns">
            <a :href="item[1]"><img class="icon-img" :src="getImgPath(item[0])"></a>
        </div>
    </footer>
</template>

<script lang="ts">
    import { Vue, Component, Prop } from "vue-property-decorator";
    @Component
    export default class FooterComponent extends Vue{
        sns:[string,string][];
        constructor(){
            super();
            this.sns =[["facebook","https://www.facebook.com/ryoma0413"],["github","https://github.com/ryomak"]]
        }

        getImgPath(name:string){
            return "/static/img/icon/footer/"+name+".png";
        }

    };
</script>

<style scoped>
    .footer {
        padding: 10px 0;
        display: flex;
        flex-direction : row-reverse;
        bottom: 0;
        width: 100%;
        height: 100px;
    }
    .icon{
      padding:10px;
    }
    .icon-img{
        width: 60px;
        height: 60px;
    }
</style>