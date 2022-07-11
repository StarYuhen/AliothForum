<template>
  <div>
    <component :is="componentIs"></component>
  </div>
</template>

<script>
import ArticlePC from "@/components/PC/ArticlePC";
import ArticleMoblie from "@/components/Moblie/ArticleMoblie";

export default {
  name: "ArticleContent",
  components: {
    ArticleMoblie,
    ArticlePC
  },
  data() {
    return {
      componentIs: "ArticleMoblie"
    }
  }
}
</script>

<style scoped>

</style>