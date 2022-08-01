// 封装axios拦截器
import axios from 'axios'
import {Toast} from "vant";
// import JSEncrypt from 'jsencrypt'
// import md5 from "js-md5";
//
// const PubKey = "-----BEGIN RSA Public Key-----\n" +
//     "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA5jQR8YGNhbbZgiLJLO2S\n" +
//     "7jvhCbAA4ye2pPjs+8Ml0I4ASAgDXl/v6cecD4vcDk4GFl+FCF87fQ3Ktyf3YUO+\n" +
//     "jjfX44+R9/zz4zIE6YS2C7C34EYCdxOLf0k/t6WeHCAX3yck2hUGuMI5rB2Khdxn\n" +
//     "XpbY/wf7KLSxMvRYagikwXJ2E+iR/GDDsr5puLlTAnV7rimjN61Ayqm1gnk8kHs9\n" +
//     "3+c4kx+HUTcWdLGNt4e80wx1F+iXbGNTywop+YUzUvLkUOV4kcunIB9NIEOZIwkB\n" +
//     "K6zqEzIDAHTBLN17M5dznlQVYh8qxqB+arypyiH8jkfbW5fxHWJAGehSgGEaxd1q\n" +
//     "PwIDAQAB\n" +
//     "-----END RSA Public Key-----"
//
// const UUID = crypto.randomUUID();
//
// const encryptor = new JSEncrypt()  // 创建加密对象实例
// encryptor.setPublicKey(PubKey)//设置公钥
//
//
// let TimeSecond = Date.parse(new Date()) / 1000

const service = axios.create({
    baseURL: "http://localhost:47",
    timeout: 1000,
    method: "POST",
    withCredentials: true
})

//
// function RSAEncrypt() {
//     let sign = encryptor.encrypt(TimeSecond + "|" + UUID + "|" + md5('StarYuhen'))  // 对内容进行加密
//     console.log(sign)
//     return sign
// }

// 定制拦截器内容
service.interceptors.request.use(
    config => {
        config.data = JSON.stringify(config.data)
        config.headers = {
            'Content-Type': 'application/json',
            "Authorization": "Bearer " + localStorage.jwt,
            "captcha": localStorage.captcha,
        }
        return config
    }
)

// 添加响应拦截器
service.interceptors.response.use(function (response) {
    // // 对响应数据做点什么
    // let DeleteMsg = ["查询uid成功", "欢迎查看文章"]
    // if (DeleteMsg.indexOf(response.data.msg) === -1) {
    //     Toast(response.data.msg)
    // }
    // 只提示请求错误的情况
    if (response.data.code !== 200) {
        Toast(response.data.msg)
        return
    }

    // 发送特地需要toast的api
    if (response.data.code === 444) {
        Toast(response.data.msg)
    }


    return response;
}, function (error) {
    // 对响应错误做点什么
    if (error.response.status === 429) {
        Toast("请求超过限制次数太多啦，等一秒哦~");
    } else if (error.response.status === 0) {
        Toast("网站出现问题，请等待管理员维护")
    }

    return Promise.reject(error);
});


export default service