<template>
  <van-nav-bar title="标题" left-arrow @click-left="onClickLeft">
    <template #right>
      <van-icon name="search" size="18"/>
    </template>
  </van-nav-bar>
</template>

<script>
// 从外界点击论坛进入的论坛主页面
const onClickLeft = () => history.back();
export default {
  name: "InsideTheForum",
  data() {
    return {
      onClickLeft
    }
  }
}
</script>

<style scoped>

</style>