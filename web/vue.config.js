const {defineConfig} = require('@vue/cli-service')
// webpack按需引用插件
const CompressionPlugin = require('compression-webpack-plugin')
const AutoImport = require('unplugin-auto-import/webpack')
const Components = require('unplugin-vue-components/webpack')
const {ElementPlusResolver} = require('unplugin-vue-components/resolvers')

// const {InjectManifest} = require('workbox-webpack-plugin');
// vue seo 预渲染 https://learnku.com/articles/46637
// const PrerenderSPAPlugin = require('prerender-spa-plugin')
// const Renderer = PrerenderSPAPlugin.PuppeteerRenderer
// const path = require('path')
// const {GenerateSW} = require('workbox-webpack-plugin');

module.exports = defineConfig({
    transpileDependencies: true,
    productionSourceMap: false,// 删除打包生成后的js的map文件，如 chunk-vendors.0c45bfa2.js.map
    configureWebpack: {
        plugins: [
            // 前端压缩gzip功能 nginx配置内容 https://better.blog.csdn.net/article/details/124481138?spm=1001.2101.3001.6650.1&utm_medium=distribute.pc_relevant.none-task-blog-2%7Edefault%7ECTRLIST%7Edefault-1-124481138-blog-120543171.pc_relevant_multi_platform_whitelistv1_exp2&depth_1-utm_source=distribute.pc_relevant.none-task-blog-2%7Edefault%7ECTRLIST%7Edefault-1-124481138-blog-120543171.pc_relevant_multi_platform_whitelistv1_exp2&utm_relevant_index=1
            new CompressionPlugin({
                test: /\.(woff|js|css|png|jpg|gif)?$/i, // 哪些文件要压缩
                filename: '[path][base].gz',// 压缩后的文件名--此插件的版本更新为这样了，原始的是 '[path].gz[query]'
                algorithm: 'gzip',// 使用gzip压缩
                minRatio: 1,// 压缩率小于1才会压缩
                deleteOriginalAssets: false // 删除未压缩的文件，谨慎设置，如果希望提供非gzip的资源，可不设置或者设置为false
            }),
            // 实现element-plus和vant的按需引用功能
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
        },
        optimization: {
            runtimeChunk: 'single',
            splitChunks: {
                chunks: 'all', // 表示哪些代码需要优化，有三个可选值：initial(初始块)、async(按需加载块)、all(全部块)，默认为async
                maxInitialRequests: Infinity, // 按需加载时候最大的并行请求数，默认为5
                minSize: 30000, // 依赖包超过300000bit将被单独打包
                // 缓存组
                // priority: 缓存组打包的先后优先级
                // minChunks: 表示被引用次数，默认为1
                cacheGroups: {
                    // 第三方库
                    libs: {
                        name: 'chunk-libs',
                        test: /[\\/]node_modules[\\/]/,
                        priority: 10,
                        chunks: 'initial', // only package third parties that are initially dependent
                        reuseExistingChunk: true,
                        enforce: true
                    },
                    echarts: {
                        name: 'chunk-echarts',
                        test: /[\\/]node_modules[\\/]echarts[\\/]/,
                        chunks: 'all',
                        priority: 12,
                        reuseExistingChunk: true,
                        enforce: true
                    }
                }
            }
        },
    },
// css设置可以使用less后缀类型
// css: {
//     loaderOptions: {
//         less: {
//             javascriptEnabled: true
//         }
//     }
// }

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

