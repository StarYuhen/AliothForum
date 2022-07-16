<template>
  <van-nav-bar title="文章" @click-left="onClickLeft" left-arrow/>
  <div>
    <p class="text-left font-black text-black text-2xl">
      {{ ArticleContent.ArticleData.Title }}
    </p>
    <!-- 作者昵称，头像，关注按钮-->
    <div class="p-1 max-w-sm mx-auto bg-white rounded-xl shadow-md flex items-center space-x-4 shadow">
      <!-- ... -->
      <div class="flex-shrink-0">
        <el-avatar class="float-left  w-12" :src="ArticleContent.Article.AuthorIMG" :lazy="true"/>
      </div>

      <div>
        <div class="text-base font-medium text-red-600 max-w-lg ">{{ ArticleContent.Article.AuthorName }}</div>
        <div class="font-thin text-current text-sm">
          {{ ArticleContent.ArticleData.PageViews + "浏览·" + ArticleContent.ArticleData.Likes + "点赞" }}
        </div>

      </div>

      <!--      关注按钮-->
      <!--      <div class="float-right rounded w-30 border-8 bg-blue-300 mr-8 text-gray-500">关注</div>-->

    </div>
  </div>

  <!--    文章内容-->
  <div v-html="ArticleContent.Article.Content" class="markdown-body rounded-md pl-3 pr-3 pb-3 pt-3 shadow "></div>
  <!--  点赞，分享，收藏，评论-->
  <van-grid :clickable=true icon-size="22px">
    <van-grid-item icon="like-o" text="点赞"/>
    <van-grid-item icon="chat-o" text="评论" @click="Config.Comment=true"/>
    <van-grid-item icon="star-o" text="收藏"/>
    <van-grid-item icon="share-o" text="分享" @click="Config.ShareList=true"/>
    <OptionListShare
        v-model:show="Config.ShareList"
    />
  </van-grid>
  <van-divider :style="{ color: '#1989fa', borderColor: '#1989fa', padding: '0 16px' }">评论</van-divider>
  <!--  展示评论内容 使用下拉刷新功能-->

  <van-list
      v-model:loading="loading"
      :finished="finished"
      finished-text="没有更多了"
      @load="onLoad"
  >
    <!--    评论列表-->
    <div class="comment" v-for="comment in Config.CommentList" v-bind:key="comment">
      <van-cell-group inset>
        <div class="flex-shrink-0">
          <el-avatar class="comment-avatar float-left  w-12 " :src="comment.AuthImg" :lazy="true"/>
        </div>
        <div class="comment-name font-medium max-w-lg">
          <text class="text-sm font-normal">{{ comment.AuthorName }}</text>
          <br>
          <text class="font-thin text-current text-xs">
            {{ comment.CreatedAt.split("T")[0] }}
          </text>
        </div>
        <div id="CommentHtml" class="comment-content text-base " v-html="comment.CommentText "></div>
      </van-cell-group>
    </div>

  </van-list>

  <!--  评论圆角弹出层-->
  <van-popup
      v-model:show="Config.Comment"
      position="bottom"
      closeable
      :style="{ height: '100%' }"
  >
    <ArticleAndComment :Type=false
                       :AuthorUID=this.ArticleContent.ArticleData.AuthorUID
                       :ClassificationUID=this.ArticleContent.ArticleData.ClassificationUID
                       :CommentType=this.Config.CommentType
                       :CommentUID=this.Config.CommentUID
                       @FalseCommentPopup="FalseCommentPopup(this.Config)"
    ></ArticleAndComment>
  </van-popup>
  <van-popup v-model:show="Config.QRCodeBool">
    <el-image :src="Config.QRCode"/>
  </van-popup>

</template>

<script>
import TouristApi from "@/assets/js/RouterTourist";
import UserApi from "@/assets/js/RouterUser";

const onClickLeft = () => history.back();
export default {
  name: "ArticleMoblie",
  components: {
    OptionListShare: () => import("./OptionListShare"),
    ArticleAndComment: () => import("./ArticleAndComment"),
  },
  data() {
    return {
      Config: {
        ShareList: false,
        Comment: false,
        QRCodeBool: false,
        QRCode: "",
        CommentText: "",
        CommentType: true,
        // 文章评论UID
        CommentUID: "",
        OptionListShare: [[
          {name: '微信', icon: 'wechat'},
          {name: '朋友圈', icon: 'wechat-moments'},
          {name: '微博', icon: 'weibo'},
          {name: 'QQ', icon: 'qq'},
        ],
          [
            {name: '复制链接', icon: 'link'},
            {name: '海报', icon: 'poster'},
          ]
        ],
        // 评论列表
        CommentList: [{
          "ID": 2,
          "Uid": "1657156821922318800",
          "AuthorName": "StarYuhen",
          "AuthImg": "https://q1.qlogo.cn/g?b=qq&nk=3446623843&s=640",
          "CommentOneUid": "d28e3c19-2309-44ef-adfc-a810491a5b1e",
          "CommentTwoUid": "",
          "CommentText": "测试一级评论",
          "CreatedAt": "2022-07-07T11:20:41.745+08:00",
          "UpdatedAt": "2022-07-07T11:20:41.745+08:00",
          "DeletedAt": null
        }],
      },
      ArticleContent: {
        ID: "",
        Img: "",
        ArticleData: {
          Title: "",
          PageViews: 0,
          Likes: 0,
          AuthorUID: "",
          ViewPermissions: false,
          ClassificationUID: ""
        },
        Article: {
          Description: "",
          Content: "",
          AuthorName: "",
          ClassificationName: "",
          AuthorIMG: "",
          ClassificationIMG: "",
          CreateTime: "",
          Keywords: ""
        }
      },
      CommentContent: "",
      onClickLeft
    }
  },
  async mounted() {
    // 先关闭底部的导航栏
    console.log(this.$refs.TabBar)

    console.log(this.$route.path)
    let id = this.$route.path.split("/")[2]
    this.ArticleContent = await TouristApi.GetArticleContent(id)
    // 同步请求评论列表--先生成请求内容
    let Comment = {
      UID: this.ArticleContent.ID,
      Number: 1,
      Type: false,
      CommentUID: "",
      ClassificationUID: this.ArticleContent.ArticleData.ClassificationUID
    }
    this.Config.CommentList = await UserApi.PostCommentOne(Comment)
  },
  methods: {
    // 子组件传递的值更改父组件
    FalseCommentPopup: (config) => {
      config.Comment = false
    }
  }

}
</script>

<style>
@import "github-markdown-css/github-markdown.css";
</style>