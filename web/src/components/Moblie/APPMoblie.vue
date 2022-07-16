<template>
  <!--  <keep-alive include="ForumUI">-->
  <!--    <router-view v-if="true"/>-->
  <!--  </keep-alive>-->
  <router-view v-if="true"/>

  <van-tabbar route v-show="this.$route.name!=='Article'">
    <van-tabbar-item replace to="/" icon="home-o">首页</van-tabbar-item>
    <van-tabbar-item replace n to="/forum" icon="friends-o">论坛</van-tabbar-item>
    <van-tabbar-item replace to="/setting" icon="setting-o">我的</van-tabbar-item>
  </van-tabbar>

</template>

<script>
// 初始化首页界面地址

export default {
  name: 'APPMoblie',
  data() {
    const active = 0;
    return {
      active,
      config: {
        value: ""
      }
    }
  },
  methods: {},
  mounted() {
    console.log("test")
  }
}
</script>

<style>

</style>
