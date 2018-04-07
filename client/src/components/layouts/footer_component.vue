<template>
    <div class="footer">
        <div class="icon" v-for="item in sns">
            <a :href="item[1]"><img class="icon-img" :src="getImgPath(item[0])"></a>
        </div>
    </div>
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
    }
    .footer a {
        color: #70726F;
        font-size: 20px;
        padding: 10px;
        border-right: 1px solid #70726F;
    }
    .footer a:first-child {
        border-left: 1px solid #70726F;
    }
    .footer a:hover {
        color: white;
    }
    .icon{
        width: 100px;
        height: 30px;
    }
    .icon-img{
        width:auto;
        height:auto;
        max-width:100%;
        max-height:100%;
    }
</style>