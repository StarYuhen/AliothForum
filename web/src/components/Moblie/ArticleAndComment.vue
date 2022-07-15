<template>
  <van-nav-bar title="请输入内容"/>
  <van-cell-group inset>
    <van-field v-model="CreateArticle.Title" placeholder="请输入文章标题" label="标题" v-if="this.Config.Type" required/>
    <van-field v-model="CreateArticle.Keywords" placeholder="游戏,二次元,Cosplay(这样编写)" label="标签" v-if="this.Config.Type"
               required/>
  </van-cell-group>
  <div style="margin-top: 10px" id="vditor"/>
  <van-button style="margin-top: 10px" type="primary" size="large"
              @click="this.Config.Type?
              PostArticle(this.CreateArticle,this.CommentContent):
              CreateComment(this.CreateArticle,this.CommentContent)">
    提交内容
  </van-button>
</template>

<script>
// 通用评论和文章组件
import Vditor from 'vditor'
import "vditor/src/assets/less/index.less"
import {Toast} from "vant";
import MajorApi from "@/assets/js/RouterMajor";

// 通过父组件传参进行判断是写的帖子（文章）还是评论
export default {
  name: "ArticleAndComment",
  // true 代表着是文章(帖子),false 代表评论
  props: ['Type', 'AuthorUID', 'ClassificationUID', 'CommentType', 'CommentUID'],
  data() {
    return {
      Config: {
        Type: false,
        UploadFile: "http://localhost:47/api/user/ArticleUploadFile",
      },
      CreateArticle: {
        ID: "",
        Title: "",
        Img: localStorage.ArticleImg,
        AuthorUID: "",
        ClassificationUID: "",
        ViewPermissions: false,
        Content: "",
        Keywords: "",
        comment: "",
        CommentType: "",
        CommentUID: "",
      },
      CommentContent: new Vditor('vditor'),
      Title: "",
    }
  },
  mounted() {
    this.CreateArticle.AuthorUID = this.AuthorUID
    this.CreateArticle.ClassificationUID = this.ClassificationUID
    this.CreateArticle.CommentType = this.CommentType
    this.CreateArticle.CommentUID = this.CommentUID
    console.log("请求的值", this.CreateArticle)
    this.CreateArticle.ID = this.$route.path.split("/")[2]
    let id = this.CreateArticle.ID
    // vditor编辑器配置内容
    this.CommentContent = new Vditor('vditor', {
      height: 600,
      mode: "wysiwyg", // 所见即所得
      toolbarConfig: {
        pin: true,
      },
      // 本地缓存
      cache: {
        enable: true,
        id: id,
      },
      // 文件上传
      upload: {
        accept: 'image/jpg,image/png,image/jpg',
        url: this.Config.UploadFile,
        headers: {"Authorization": "Bearer " + localStorage.jwt},
        linkToImgUrl: this.Config.UploadFile,
        fieldName: "file",
        CreateArticleUpload: this.CreateArticle,
        AddArticleFun: (txt) => {
          this.CommentContent.insertValue(txt);
        },
        max: 3 * 1024 * 1024,
        success(editor, msg) {
          console.log(editor, msg)
          let Msg = JSON.parse(msg)
          // 将图片插入数据中
          if (Msg.code === 200) {
            this.AddArticleFun("![image.png](" + Msg.data + ")")
            // 储存进浏览器
            localStorage.ArticleImg = Msg.data
          }
          Toast(Msg.msg)
        },

      },
    })

  },
  methods: {
    // 创建文章
    PostArticle: (CreateArticle, CommentContent) => {
      if (CreateArticle.Title === "" || CreateArticle.Keywords === "") {
        Toast("参数不完全")
        return
      }
      CreateArticle.Content = CommentContent.getValue()
      let data = MajorApi.insertArticle(CreateArticle)
      if (data.code === 200) {
        console.log(data)
      }
    },
    // 创建评论
    CreateComment: (CreateArticle, CommentContent) => {
      let comment = {
        "ArticleUID": CreateArticle.ID,
        "ClassificationUID": CreateArticle.ClassificationUID,
        "Type": CreateArticle.CommentType,
        "CommentUID": CreateArticle.CommentUID,
        "Text": CommentContent.getValue()
      }
      let data = MajorApi.insertComment(comment)
      console.log(data)

    }

  }
}
</script>

<style scoped>

</style>