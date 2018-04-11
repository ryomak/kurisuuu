<template>
    <div class="container">
        <div class="not-found">NotFound</div>
    </div>
</template>

<script lang="ts">
    import { Vue, Component, Prop } from "vue-property-decorator";
    @Component
    export default class NotFoundComponent extends Vue{
        constructor(){
            super();
        }
    };
</script>

<style scoped>
.container{
    width: 70%;
    height: auto;
    padding-top: 50px;
}
.not-found{
    font-size: 50px;
    text-align: center;
}
/* enter、 leave アニメーションで異なる間隔やタイミング関数を利用することができます */
.slide-fade-enter-active {
    transition: all .3s ease;
}
.slide-fade-leave-active {
    transition: all .8s cubic-bezier(1.0, 0.5, 0.8, 1.0);
}
.slide-fade-enter, .slide-fade-leave-to
    /* .slide-fade-leave-active below version 2.1.8 */ {
    transform: translateX(10px);
    opacity: 0;
}

</style>