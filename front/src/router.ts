import { createRouter, createWebHistory } from 'vue-router'
import HomeView from "./views/HomeView.vue";
import SearchView from "./views/SearchView.vue";


const routes = [
    { path: '/', name: "home",  component: HomeView },
    { path: '/search', name: "search", component: SearchView },
]

export const router = createRouter({
    history: createWebHistory(),
    routes,
})