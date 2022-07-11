<template>
  <div>
    <component :is="componentIs"></component>
  </div>
</template>

<script>
// 论坛推荐界面
import RecomMendPC from "@/components/PC/RecomMendPC";
import RecomMendMoblie from "@/components/Moblie/RecomMendMoblie";

export default {
  name: "RecomMend",
  components: {
    RecomMendPC,
    RecomMendMoblie
  },
  data() {
    return {
      componentIs: "RecomMendMoblie"
    }
  }
}
</script>

<style scoped>


</style>