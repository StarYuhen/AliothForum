import {createApp} from 'vue'
//绑定起始vue
import App from './App.vue'
// //按需引用vant组件
import Vant, {Lazyload} from 'vant'
//绑定css文件
import 'vant/lib/index.css';
import router from "@/assets/js/router";
// 绑定element框架
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
// // 绑定md解析库 这个库暂时不支持vue3
// import MavonEditor from 'mavon-editor'
// import 'mavon-editor/dist/css/index.css'
// import "vditor/src/assets/less/index.less"
// 自定义css文件
import "./assets/css/index.css"


// // 使用自定义Vue的PWA插件，减少资源使用  http://kmanong.top/kmn/qxw/form/article?id=70335&cate=52
// import './registerServiceWorker'
// import meta from 'vue-meta';
// import store from "@/assets/js/store";
// import {List, NavBar, Search, Skeleton, Tabbar, TabbarItem} from "vant";
//绑定元素与框架中间件
const app = createApp(App, {
    // 预渲染seo 暂时取消
    // mounted() {
    //     document.dispatchEvent(new Event('render-event'))
    // }
});


// // 按需引用框架资源 太麻烦了，直接引用全部了
// app.use(Tabbar);
// app.use(TabbarItem);
// app.use(NavBar)
// app.use(Skeleton)
// app.use(Search)
// app.use(List)


// // 绑定vant所有资源
app.use(Vant)
// 实现图片与组件懒加载
app.use(Lazyload, {
    lazyComponent: true,
})



// 绑定element资源
app.use(ElementPlus)
//绑定路由
app.use(router)
// 绑定页面mate标签,利于seo  https://juejin.cn/post/6930964642814836743
// app.use(meta)
// 需要动态加载标签，需要使用vuex的对象管理
// app.use(store)


app.mount('#app');

// 开发完后使用 https://tauri.app/zh/ 进行跨平台打包

// // 注册组件
// app.component("ArticleAndComment", {
//     props: ['Type'],
//     template: ArticleAndComment
// })
