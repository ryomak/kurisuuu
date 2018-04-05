import Vue from 'vue'
import VueRouter, { Location, Route, RouteConfig } from 'vue-router'
import ReadmeComponent from "./components/readme/readme.vue"
import MovieComponent from "./components/movie/movie.vue"
import GitHubComponent from "./components/github/github.vue"
import QiitaComponent from "./components/qiita/qiita.vue"

Vue.use(VueRouter);

export const createRoutes: () => RouteConfig[] = () => [
    {
        path: '/',
        component: ReadmeComponent
    },
    {
        path: '/movie',
        component: MovieComponent
    },
    {
        path: '/github',
        component: GitHubComponent
    },
    {
        path: '/qiita',
        component: QiitaComponent
    }
];

export const createRouter = () => new VueRouter({ mode: 'history', routes: createRoutes() })