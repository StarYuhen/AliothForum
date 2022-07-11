import server from "@/assets/js/server";
import {Toast} from "vant";
import expen from "@/assets/js/expen";
import {marked} from "marked"

const TouristApi = {
    // 请求图形验证码
    GetCaptcha: async () => {
        let data
        await server.get("/api/tourist/captcha").then(function (response) {
            data = response.data
        })
        console.log("请求验证码的值", data)
        return data.data
    },
    // 登录账号
    PostAccountLogin: async (account) => {
        let Bool = false
        await server.post("/api/tourist/loginAccount", {
            User: account.UserMail,
            Password: account.PassWord,
            Captcha: account.Captcha
        }).then(function (response) {
            if (response.data.code === 200) {
                // 同时写入本地缓存
                expen.InterFaceLogin(response.data.data)
                Bool = true
            }
        })
        return !Bool
    },
    // 注册账号
    PostRegisterAccount: async (account) => {
        let Bool = false
        if (account.PassWord !== account.PassWordS) {
            Toast("两次密码不一样,请自行检测")
            return !Bool
        }
        await server.post("/api/tourist/registerAccount", {
            User: account.UserMail,
            UserName: account.UserName,
            Password: account.PassWord,
            Captcha: account.MailCaptcha
        }).then(function (response) {
            if (response.data.code === 200) {
                // 同时写入本地缓存
                expen.InterFaceLogin(response.data.data)
                Bool = true
            }
        })
        return !Bool
    },
    // 发送邮件验证码
    PostRegisterAccountMail: async (mail) => {
        await server.get("/api/tourist/registerAccountMail?mail=" + mail)
    },
    // 请求文章内容
    GetArticleContent: async (id) => {
        let data
        await server.get("/api/tourist/article/" + id).then(function (response) {
            data = response.data.data
            console.log("文章返回的内容：",data)
            data.Article.Content = marked(data.Article.Content)
        })
        return data
    },
}

export default TouristApi