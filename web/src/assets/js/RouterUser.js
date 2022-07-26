// 请求是否登录接口
import server from "@/assets/js/server";

const UserApi = {
    // 请求设备jwt是否正确
    GetJwtDevice: async () => {
        let Bool
        await server.get("/api/user/JwtDevice").then(function (response) {
            Bool = response.data != null && response.data.data === true
        })
        console.log("登录账号", Bool)
        return !Bool
    }, // 获取当前用户头像
    GetImgUser: async () => {
        let url
        await server.get("/api/user/ImgUser").then(function (response) {
            url = response.data.data
        })
        localStorage.ImgUrl = url
        return url
    }, // 生成当前页面二维码
    GetQrCode: async (RouterUrl) => {
        let url
        await server.get("/api/user/UrlQrCode?qrcode=" + RouterUrl).then(function (response) {
            url = response.data.data
        })
        console.log("获取的url" + url)
        return url
    }, // 获取一级评论信息
    PostCommentOne: async (comment) => {
        let data
        await server.post("/api/user/ReadComment", {
            UID: comment.UID,
            Number: comment.Number,
            Type: comment.Type,
            CommentUID: comment.CommentUID,
            ClassificationUID: comment.ClassificationUID,
        }).then(function (response) {
            data = response.data.data
        })
        console.log(data)
        return data
    },
    // 文章点赞接口
    ArticleLike: async (uid) => {
        let data
        await server.get("/api/user/ArticleLike?article=" + uid).then(function (response) {
            console.log(response.data)
            data = response.data
        })
        return data
    }
}

export default UserApi
