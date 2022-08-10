<template>
  <van-nav-bar left-arrow title="文章" @click-left="onClickLeft"/>
  <div>
    <p class="text-left font-black text-black text-2xl">
      {{ ArticleContent.ArticleData.Title }}
    </p>
    <!-- 作者昵称，头像，关注按钮-->
    <div class="p-1 max-w-sm mx-auto bg-white rounded-xl shadow-md flex items-center space-x-4 shadow">
      <!-- ... -->
      <div class="flex-shrink-0">
        <el-avatar :lazy="true" :src="ArticleContent.Article.AuthorIMG" class="float-left  w-12"/>
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
  <div class="markdown-body rounded-md pl-3 pr-3 pb-3 pt-3 shadow " v-html="ArticleContent.Article.Content"></div>
  <!--  点赞，分享，收藏，评论-->
  <van-grid :clickable=true icon-size="22px">
    <van-grid-item icon="like-o" text="点赞" @click="ArticleLikeUID(this.ArticleContent.ID)"/>
    <van-grid-item icon="chat-o" text="评论" @click="Config.Comment=true"/>
    <van-grid-item icon="star-o" text="收藏"/>
    <van-grid-item icon="share-o" text="分享" @click="Config.ShareList=true"/>
    <OptionListShare
        v-model:show="Config.ShareList"
    />
  </van-grid>
  <van-divider :style="{ color: '#1989fa', borderColor: '#1989fa', padding: '0 16px' }">评论</van-divider>
  <!--  展示评论内容 使用下拉刷新功能-->

  <van-list v-if="this.CommentListPost.Number!==undefined"
            v-model:loading="this.ListTrue.loading"
            :finished="this.ListTrue.finished"
            finished-text="已全部加载完成辣！"
            @load="ArticleComment(this.CommentListPost,this.Config.CommentList,this.ListTrue)"
  >
    <!--    评论列表-->
    <div v-for="comment in Config.CommentList" v-bind:key="comment" class="comment">
      <van-cell-group inset>
        <div class="flex-shrink-0">
          <el-avatar :lazy="true" :src="comment.AuthImg" class="comment-avatar float-left  w-12 "/>
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
      :style="{ height: '100%' }"
      closeable
      position="bottom"
  >
    <ArticleAndComment :AuthorUID=this.ArticleContent.ArticleData.AuthorUID
                       :ClassificationUID=this.ArticleContent.ArticleData.ClassificationUID
                       :CommentType=this.Config.CommentType
                       :CommentUID=this.Config.CommentUID
                       :Type=false
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
import {Toast} from "vant";

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
        CommentList: [],
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
      CommentListPost: {},
      CommentContent: "",
      ListTrue: {
        loading: false,
        finished: false,
      },
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

    this.CommentListPost = {
      UID: this.ArticleContent.ID,
      Number: 0,
      Type: false,
      CommentUID: "",
      ClassificationUID: this.ArticleContent.ArticleData.ClassificationUID
    }
    // this.Config.CommentList = await UserApi.PostCommentOne(this.CommentListPost)
    console.log("储存的请求内容接口", this.CommentListPost)
  },
  methods: {
    // 子组件传递的值更改父组件
    FalseCommentPopup: (config) => {
      config.Comment = false
    },
    // 文章点赞
    ArticleLikeUID: async (uid) => {
      let data = await UserApi.ArticleLike(uid)
      Toast(data.msg)
    },
    // 请求文章评论
    ArticleComment: async (comment, list, ListTrue) => {
      console.log("触发了评论更新")
      comment.Number++
      let data = await UserApi.PostCommentOne(comment)
      if (data.length === 0) {
        ListTrue.finished = true
      }
      list.push(...data)
      ListTrue.loading = false
    }
  }

}
</script>

<style>
@import "github-markdown-css/github-markdown.css";
</style>