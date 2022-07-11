<template>
  <div>
    <component :is="componentIs"></component>
  </div>
</template>

<script>
// 初始化首页界面地址
import APPPC from "@/components/PC/APPPC";
import APPMoblie from "@/components/Moblie/APPMoblie";

export default {
  name: 'App',
  components: {
    APPPC,
    APPMoblie
  },
  data() {
    return {
      componentIs: 'APPMoblie'
    };
  },
}
</script>

<style>

</style>
