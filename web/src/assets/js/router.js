import {createRouter, createWebHistory} from 'vue-router'


/*
   path: "/", component: () => {
        return import(RecomMend)
    }
    使用  component:()=>{return import(RecomMend)} 可以异步加载资源文件，否则只会打包成一个js
    https://www.webpackjs.com/guides/tree-shaking/
    https://juejin.cn/post/6844903998634328072

    component:()=>import(RecomMend) 新版分割页面js文件
 */
// 尝试路由懒加载 https://router.vuejs.org/zh/guide/advanced/lazy-loading.html
const RecomMend = () => import("../../components/RecomMend")
const settingUI = () => import("../../components/settingUI")
const ForumUI = () => import("../../components/ForumUI")
const Article = () => import("../../components/ArticleContent.vue")

const routes = [{
    path: "/", name: "RecomMend", component: RecomMend
}, {
    path: "/setting", name: "settingUI", component: settingUI
}, {
    path: "/forum", name: "ForumUI", component: ForumUI
}, {
    path: "/article/:id", name: "Article", component: Article
}]


// https://router.vuejs.org/zh/guide/#javascript
const router = createRouter({
    // createWebHashHistory, 这是使用hash路由表
    history: createWebHistory(), // 不使用hash路由表
    routes, // (缩写) 相当于 routes: routes
})


export default router