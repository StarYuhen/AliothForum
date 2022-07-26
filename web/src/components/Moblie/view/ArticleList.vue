<template>
  <!--  内容使用列表刷新-->
  <van-list
      @load="OnLoadArticle(this.List)"
  >
    <div v-for="i in List" v-bind:key="i.id" @click="this.$router.push('/article/'+i.Article.Uid)">
      <div style="height: 0.7rem;width: 100%;background-color: #f4f4f5;"></div>
      <div style="margin-left: 1rem;margin-right: 1rem;margin-top: 1rem;">
        <!--帖子基本内容-->
        <!--标题-->
        <text style="font-weight:bolder;font-size: 1rem">
          {{ i.Article.Title }}
        </text>
        <div>
          <div style="height: 2rem">
            <!--头像及用户名字-->
            <van-image
                round
                width="1.6rem"
                height="1.6rem"
                :src=i.Article.AuthImg
            />

            <text class="recomName">{{ i.Article.AuthorName }}</text>
          </div>

          <!--          帖子内容-->
          <div class="van-multi-ellipsis--l3">
            <p style="font-size: 0.8rem">
              {{ i.Article.Content }}
            </p>


            <van-image
                :src="i.Article.Img"
                v-if="i.Article.Img!==''"
            />
          </div>
          <!--  赞同，浏览量，评论-->
          <div style="margin-top: 0.7rem;margin-bottom: 0.7rem">
            <text style="font-size: 0.6rem;color: #c8c9cc">{{ i.All.PageViews }}浏览·{{ i.All.Likes }}赞同</text>
          </div>

        </div>
      </div>

    </div>

  </van-list>
</template>

<script>
// 文章列表，将文章内容抽出来做成组件，判断是首页还是贴吧内部的
import TouristApi from "@/assets/js/RouterTourist";
import {Toast} from "vant";

export default {
  name: "ArticleList",
  // type=true 则是首页，type=false 则是论坛
  // todo 还在写首页的
  props: ['Type'],
  data() {
    return {
      List: []
    }
  },

  async mounted() {
    // 请求随机推荐内容,未真则是首页推荐内容
    if (this.Type) {
      // 首页推荐接口
      let data = await TouristApi.GetRandomIndexArticle()
      this.List = data.data
      if (data.data.length === 0) {
        Toast("压根没人写文章（哭）")
      }
    } else {
      // 论坛内的推荐文章
    }

  },
  methods: {
    // 再次请求加载内容
    OnLoadArticle: async (list) => {
      let data = await TouristApi.GetRandomIndexArticle()
      list.push(...data.data)
    }
  }
}
</script>

<style scoped>

</style>