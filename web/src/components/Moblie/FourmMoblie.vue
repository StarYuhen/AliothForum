<template>
  <van-nav-bar title="论坛"/>
  <!--  推荐的贴吧列表-->
  <div>
    <div class="ml-6 text-xl font-light">
      <van-icon name="fire"/>
      <text>推荐的贴吧</text>
    </div>
    <van-grid square>
      <van-grid-item class="truncate " v-for="value in Config.RandomRecommend"
                     :key="value" :icon="value.ImgURL" :text="value.Name"/>
    </van-grid>

    <div class="ml-6 text-xl font-light">
      <van-icon name="star"/>
      <text>收藏的贴吧</text>
    </div>
  </div>
</template>

<script>
// 论坛页面，获取推荐论坛和已经关注的论坛

import TouristApi from "@/assets/js/RouterTourist";

export default {
  name: "ForumUI",
  components: {},
  data() {
    return {
      Config: {
        RandomRecommend: []
      }
    }
  },
  async mounted() {
    let Random = await TouristApi.GetRandomRecommendForum()
    // 赋值给随机推荐论坛列表
    this.Config.RandomRecommend = Random.data
    console.log("返回的数据内容长度", this.Config.RandomRecommend.length)
  },
  methods: {},
}
</script>

<style scoped>
</style>