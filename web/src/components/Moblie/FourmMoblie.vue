<template>
  <van-nav-bar title="论坛"/>
  <!--  推荐的贴吧列表-->
  <div>
    <div class="ml-6 text-xl font-light">
      <van-icon name="fire"/>
      <text>推荐的贴吧</text>
    </div>
    <van-grid square>
      <van-grid-item v-for="value in 8" :key="value" icon="photo-o" text="文字"/>
    </van-grid>

    <div class="ml-6 text-xl font-light">
      <van-icon name="star" />
      <text>收藏的贴吧</text>
    </div>
  </div>
</template>

<script>
// 论坛页面，获取推荐论坛和已经关注的论坛

export default {
  name: "ForumUI",
  components: {},
  data() {
    return {}
  },
  mounted() {

  },
  methods: {},
}
</script>

<style scoped>
</style>