<template>
  <div>
    <component :is="componentIs"></component>
  </div>
</template>

<script>
import FourmPC from "@/components/PC/FourmPC";
import FourmMoblie from "@/components/Moblie/FourmMoblie";

export default {
  name: "ForumUI",
  components: {
    FourmPC,
    FourmMoblie
  },
  data() {
    return {
      componentIs: "FourmMoblie"
    }
  }

}
</script>

<style scoped>

</style>