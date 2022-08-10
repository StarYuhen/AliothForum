<template>
  <van-nav-bar :title="Config.Forum.Name" left-arrow @click-left="onClickLeft">
    <template #right>
      <van-icon name="search" size="18"/>
    </template>
  </van-nav-bar>


  <ArticleList :Type="false" :Uid="this.Config.ID"></ArticleList>
</template>

<script>
// 从外界点击论坛进入的论坛主页面
import TouristApi from "@/assets/js/RouterTourist";
import ArticleList from "@/components/Moblie/view/ArticleList";

const onClickLeft = () => history.back();
export default {
  name: "InsideTheForum",
  components: {ArticleList},
  data() {
    return {
      Config: {
        Forum: {},
        ID: "",
      },
      onClickLeft
    }
  },
  async mounted() {
    // 请求当前页面路由的uid然后获取贴吧信息
    let id = this.$route.path.split("/")[2]
    this.Config.ID = id
    let Forum = await TouristApi.GetForumDetails(id)
    this.Config.Forum = Forum.data
  }
}
</script>

<style scoped>

</style>