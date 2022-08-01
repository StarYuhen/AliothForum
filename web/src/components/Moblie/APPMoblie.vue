<template>
  <!--  <keep-alive include="ForumUI">-->
  <!--    <router-view v-if="true"/>-->
  <!--  </keep-alive>-->
  <router-view v-if="true"/>

  <van-tabbar route v-show="RouterShow(this.$route.name)">
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
  methods: {
    RouterShow: (name) => {
      // 这个函数是返回是否隐藏tabbar的
      let Msg = ["Article", "InsideForum"]
      // -1说明没有存在相同的
      console.log("测试是否隐藏tabbar：", name)
      return Msg.indexOf(name) === -1
    }
  },
  mounted() {
    console.log("test")
  }
}
</script>

<style>

</style>
