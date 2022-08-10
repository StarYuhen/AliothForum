<template>

  <van-nav-bar title="推荐"/>
  <van-search
      v-model="Config.Search"
      placeholder="请输入搜索关键词"
      input-align="center"
  />

  <ArticleList :Type="true" :Uid="Config.Search"></ArticleList>

</template>

<script>
// 论坛推荐界面

import ArticleList from "@/components/Moblie/view/ArticleList";


export default {
  name: "RecomMendMoblie",
  components: {ArticleList},
  // 引用vuex状态树
  // store,
  data() {
    return {
      Config: {
        Search: "",
      },
    }
  },
  mounted() {
  }
}
</script>

<style scoped>


</style>