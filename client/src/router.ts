import Vue from 'vue'
import VueRouter, { Location, Route, RouteConfig } from 'vue-router'
import ReadmeComponent from "./components/readme/read_me.vue"
import MovieComponent from "./components/movie/movie.vue"
import GitHubComponent from "./components/github/github.vue"
import QiitaComponent from "./components/qiita/qiita.vue"
import NotFoundComponent from "./components/notfound/not_found.vue"

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
    },
    {
        path: '*',
        component: NotFoundComponent
    }
];

export const createRouter = () => new VueRouter({ mode: 'history', routes: createRoutes() });