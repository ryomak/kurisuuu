<template>
    <div class="container">
        <div class="head">
            <img class="window-button" :src="singou">
        </div>
        <div class="terminal">
            <div class="wrapper">
                <div class="time">Last login:{{timeCommandLineFormat()}} on tty0413</div>
                <div class="time">Welcome to fish, the friendly interactive shell</div>
                <div v-for="command in commands">
                    <div ><span class='command'>{{path}}</span>{{command[0]}}</div>
                    <div v-html="command[1]"/>
                </div>
                <span class='command'>{{path}}</span>vim hello.go <span class="cursor"></span>
            </div>
        </div>
    </div>
</template>

<script lang="ts">
    import { Vue, Component, Prop } from "vue-property-decorator";
    import * as Const from "../../const/const"
    @Component
    export default class Readme extends Vue{
        c = Const;
        //const
        singou:string = this.c.SINGOU;
        now:Date = new Date();
        path:string = this.c.TERMINAL_PATH;
        commands :[string,string][];
        constructor(){
            super();
            this.commands=[
                ["whoami","Kurisu"],
                ["birth","1996/04/13"],
                ["hobby","BasketBall,Reading,Programming"],
                ["language","Go,C,HTML,TypeScript"],
                ["intern","VOYAGE GROUP,Works Applications,Eureka,Livesense"],
                ["mail","ryooomaaa0413[at]gmail.com"]
            ]
        }

        timeCommandLineFormat(){
            const months = ["Jan","Feb","Mar","Apr","May","Jun","Jul","Aug","Sep","Oct","Nov","Dec"];
            const weeks = ["Sun","Mon","Tue","Wed","Thu","Fri","Sat"];
            let year = this.now.getFullYear();
            let month = months[this.now.getMonth()];
            let week = weeks[this.now.getDay()];
            let day = this.now.getDate();
            let hour = this.now.getHours()+1;
            let minutes = this.now.getMinutes()+1;
            let seconds = this.now.getSeconds()+1;
            return week+" "+month+" "+day+" "+year+" "+hour+":"+minutes+":"+seconds;
        }
    };
</script>

<style scoped>
    .container{
        display: flex;
        justify-content: center;
        flex-wrap: wrap;
        color: white;
        padding-top: 130px;
    }
    /*雑実装，本当はjsごと変える*/
    @media (max-width:1000px) {
        .head{
            width: 80%;
            height: 50px;
            background: lightgrey;
            border-top-left-radius: 15px;
            border-top-right-radius: 15px;
        }
        .terminal{
            font-size: 20px;
            line-height: 30px;
            background-color: black;
            top:0;
            border-bottom-left-radius: 15px;
            border-bottom-right-radius: 15px;
            width: 80%;
            height: auto;
        }
    }

    @media (min-width:1000px) {
        .head{
            width: 60%;
            height: 50px;
            background: lightgrey;
            border-top-left-radius: 15px;
            border-top-right-radius: 15px;
        }
        .terminal{
            font-size: 23px;
            line-height: 33px;
            background-color: black;
            top:0;
            border-bottom-left-radius: 15px;
            border-bottom-right-radius: 15px;
            width: 60%;
            height: auto;
        }
    }

    .wrapper{
        padding: 15px;
    }
    .command{
        font-size: 25px;
        color:rgb(49,103,69);
    }
    .window-button{
        width: auto;
        height: 30px;
        padding: 10px
    }

    @media (max-width:1000px) {

    }


    .cursor::after {
        content: "";
        top: 0;
        right: -15px;
        /* Remove display: inline-block if not required to be on the same line as text etc */
        display: inline-block;
        background-color: white;
        vertical-align: middle;
        width: 10px;
        /* Set height to the line height of .text */
        height: 24px;
        -webkit-animation: blink 1s step-end infinite;
        animation: blink 1s step-end infinite;
    }
    @-webkit-keyframes blink {
        0% { opacity: 1.0; }
        50% { opacity: 0.0; }
        100% { opacity: 1.0; }
    }

    @keyframes blink {
        0% { opacity: 1.0; }
        50% { opacity: 0.0; }
        100% { opacity: 1.0; }
    }


</style>