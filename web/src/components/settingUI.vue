<template>
  <div>
    <component :is="componentIs"></component>
  </div>
</template>

<script>
import settingPC from "@/components/PC/settingPC";
import settingMoblie from "@/components/Moblie/settingMoblie";

export default {
  name: "settingUI",
  components: {
    settingPC,
    settingMoblie
  },
  data() {
    return {
      componentIs: "settingMoblie"
    }
  }
}
</script>

<style scoped>

</style>