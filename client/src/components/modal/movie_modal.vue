<template>
    <transition name="modal" >
        <div class="modal-mask" @click="$emit('close')">
            <div class="modal-wrapper">
                <div class="modal-container">
                    <div class="modal-header">
                        <img class="modal-img" :src="detail.imgPath" >
                    </div>

                    <div class="modal-body">
                        <slot name="body">
                            <div v-html="detail.description"></div>
                        </slot>
                    </div>
                </div>
            </div>
        </div>
    </transition>
</template>

<script lang="ts">
    import { Vue, Component, Prop } from "vue-property-decorator";
    @Component({
        components:{
            MovieModal
        },
        props:['detail']
    })
    export default class MovieModal extends Vue{
        constructor(){
            super();
        }
    }
</script>

<style scoped>
    .modal-mask {
        position: fixed;
        z-index: 9998;
        top: 0;
        left: 0;
        width: 100vw;
        height: 100vh;
        background-color: rgba(256, 256, 256, .9);
        display: table;
        transition: opacity .3s ease;
    }

    .modal-wrapper {
        display: table-cell;
        vertical-align: middle;
    }

    .modal-container {
        width: 400px;
        margin: 0px auto;
        padding: 20px 30px;
        background-color: #fff;
        border-radius: 10px;
        box-shadow: 0 2px 8px rgba(0, 0, 0, .33);
        border:solid 2px rgb(49,103,69);
        transition: all .3s ease;
        word-wrap: break-word;
        font-size: 25px;
        font-family: "Wawati SC",sans-serif;
    }

    .modal-header {
        margin-top: 0;
    }

    .modal-img{
        width: 100%;
        height: 100%;
        border-radius: 10px;
    }
    .modal-close{
        width: 80px;
        height: auto;
        float:right;
    }

    .modal-body {
        margin: 20px 0;
    }

    .modal-default-button {
        float: right;
    }

    .modal-enter {
        opacity: 0;
    }

    .modal-leave-active {
        opacity: 0;
    }

    .modal-enter .modal-container,
    .modal-leave-active .modal-container {
        -webkit-transform: scale(1.1);
        transform: scale(1.1);
    }
</style>