// 重要的接口
import server from "@/assets/js/server";

const MajorApi = {
    // 新建文章（帖子），评论接口
    insertArticle: async (article) => {
        let data
        await server.post("/api/major/insertArticle", {
            "Title": article.Title,
            "Img": article.Img,
            "AuthorUID": article.AuthorUID,
            "ClassificationUID": article.ClassificationUID,
            "Content": article.Content,
            "ViewPermissions": article.ViewPermissions,
            "Keywords": article.Keywords
        }).then(function (response) {
            data = response.data
        })
        console.log(data)
    }
}
export default MajorApi