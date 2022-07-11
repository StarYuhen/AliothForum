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
    <van-share-sheet
        v-model:show="Config.ShareList"
        title="立即分享给好友"
        :options="Config.OptionListShare"
        @select="ClickShareSelect"
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
    <div class="comment" v-for="comment in Config.CommentList" v-bind:key="comment" @click="Config.Comment=true">
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
        <van-cell class="comment-content text-base">{{ comment.CommentText }}</van-cell>
      </van-cell-group>
      <van-divider/>
    </div>

  </van-list>


  <van-popup
      v-model:show="Config.Comment"
      round
      position="bottom"
      :style="{ height: '30%' }"
  >
    <!--    评论输入内容-->
    <van-cell-group inset>
      <van-field
          v-model="Config.CommentText"
          rows="2"
          autosize
          label="评论"
          maxlength="100"
          type="textarea"
          placeholder="请输入评论"
          show-word-limit
      />
      <van-button round class="float-right" type="primary">发送评论</van-button>
    </van-cell-group>

  </van-popup>
  <van-popup v-model:show="Config.QRCodeBool">
    <el-image :src="Config.QRCode"/>
  </van-popup>

</template>

<script>

import TouristApi from "@/assets/js/RouterTourist";
import Expen from "@/assets/js/expen";
import {Toast} from "vant";
import UserApi from "@/assets/js/RouterUser";

const onClickLeft = () => history.back();
export default {
  name: "ArticleMoblie",
  data() {
    return {
      Config: {
        ShareList: false,
        Comment: false,
        QRCodeBool: false,
        QRCode: "",
        CommentText: "",
        OptionListShare: [[
          {name: '微信', icon: 'wechat'},
          {name: '朋友圈', icon: 'wechat-moments'},
          {name: '微博', icon: 'weibo'},
          {name: 'QQ', icon: 'qq'},
        ],
          [
            {name: '复制链接', icon: 'link'},
            {name: '海报', icon: 'poster'},
            {name: '二维码', icon: 'qrcode'},
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
      onClickLeft
    }
  },
  async mounted() {
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
    // 点击分享后的样式
    ClickShareSelect: async function (option) {
      let LocationURL = window.location.href
      switch (option.name) {
        case '二维码':
          this.Config.QRCode = await UserApi.GetQrCode(LocationURL)
          console.log(this.Config.QRCode)
          this.Config.QRCodeBool = true
          break
        case '复制链接':
          await Expen.CopySrc(LocationURL)
          Toast("已复制链接到剪切板")
          break
      }
    },
  }

}
</script>

<style>
@import "github-markdown-css/github-markdown.css";
</style>