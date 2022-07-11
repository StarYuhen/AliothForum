const {defineConfig} = require('@vue/cli-service')
// webpack按需引用插件
const CompressionPlugin = require('compression-webpack-plugin')
const AutoImport = require('unplugin-auto-import/webpack')
const Components = require('unplugin-vue-components/webpack')
const {ElementPlusResolver} = require('unplugin-vue-components/resolvers')
const path = require("path");
// const {InjectManifest} = require('workbox-webpack-plugin');
// vue seo 预渲染 https://learnku.com/articles/46637
// const PrerenderSPAPlugin = require('prerender-spa-plugin')
// const Renderer = PrerenderSPAPlugin.PuppeteerRenderer
// const path = require('path')

module.exports = defineConfig({
    transpileDependencies: true,
    productionSourceMap: false,// 删除打包生成后的js的map文件，如 chunk-vendors.0c45bfa2.js.map
    configureWebpack: {
        plugins: [
            // 前端压缩gzip功能
            new CompressionPlugin({
                test: /\.(woff|js|css|png|jpg|gif)?$/i, // 哪些文件要压缩
                filename: '[path][base].gz',// 压缩后的文件名--此插件的版本更新为这样了，原始的是 '[path].gz[query]'
                algorithm: 'gzip',// 使用gzip压缩
                minRatio: 1,// 压缩率小于1才会压缩
                deleteOriginalAssets: false // 删除未压缩的文件，谨慎设置，如果希望提供非gzip的资源，可不设置或者设置为false
            }),
            // 实现element-plus的按需引用功能
            AutoImport({
                resolvers: [ElementPlusResolver()],
            }),
            Components({
                resolvers: [ElementPlusResolver()],
            }),
            // // 调用vue的server worker缓存 用于减轻服务器请求的压力 https://juejin.cn/post/6859311191505534989
            // new GenerateSW({
            //     clientsClaim: true, //service worker是否应该在任何现有客户端激活后立即开始控制它
            //     skipWaiting: true,//service worker是否应该跳过等待生命周期阶段，用于清除缓存，强制等待中的service-worker被激活
            //     globPatterns: ['**/*.{html,js,css,png.jpg,woff2}'], // 匹配的文件
            //     globIgnores: ['service-worker.js'], // 忽略的文件
            //     runtimeCaching: [
            //         {
            //             urlPattern: new RegExp('/api'),//相关接口正则配置，跨域接口必须以 ‘^’开头，配置完整域名
            //             handler: 'NetworkFirst',//缓存策略，网络请求优先。
            //             options: {
            //                 cacheableResponse: {
            //                     statuses: [200]
            //                 }
            //             }
            //         }
            //     ]
            // })

            // // 预渲染
            // new PrerenderSPAPlugin({
            //     // 生成文件的路径，这个目录只能有一级。若目录层次大于一级，在生成的时候不会有任何错误提示，在预渲染的时候只会卡着不动
            //     staticDir: path.join(__dirname, './dist'),
            //     // 对应自己的路由文件
            //     routes: ['/'],
            //     // 若没有这段则不会进行预编译
            //     renderer: new Renderer({
            //         inject: {
            //             foo: 'bar'
            //         },
            //         headless: false,
            //         // 在 main.js 中 document.dispatchEvent(new Event('render-event'))，两者的事件名称要对应上。
            //         renderAfterDocumentEvent: 'render-event'
            //     })
            // }),
        ],
        module: {
            rules: [
                // 配置读取 *.md 文件的规则
                {
                    test: /\.md$/,
                    use: [
                        {loader: "html-loader"},
                        {loader: "markdown-loader", options: {}}
                    ]
                }
            ]
        }
    },

    // // 配置PWA 本地测试不起效果，可以使用browser-sync dist 来使用
    // pwa: {
    //   workboxOptions: {
    //     skipWaiting: true,
    //     clientsClaim: true,
    //     importWorkboxFrom: 'local',
    //     importsDirectory: 'js',
    //     navigateFallback: '/',
    //     navigateFallbackBlacklist: [
    //       /\/api\//
    //     ]
    //   },
    //   workboxPluginMode: 'InjectManifest'
    // }
})
