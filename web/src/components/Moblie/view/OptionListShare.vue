<template>
  <van-share-sheet
      v-model:show="ShareList"
      title="立即分享给好友"
      :options="OptionListShare"
      @select="ClickShareSelect"
  />
</template>

<script>
// 分享组件内容
import Expen from "@/assets/js/expen";
import {Toast} from "vant";

export default {
  name: "OptionListShare",
  data() {
    return {
      ShareList: false,
      Comment: false,
      QRCodeBool: false,
      QRCode: "",
      CommentText: "",
      OptionListShare: [[
        {name: '微信', icon: 'wechat'},
        {name: '朋友圈', icon: 'wechat-moments'},
        {name: '微博', icon: 'weibo'},
        {name: 'QQ', icon: 'qq'},
      ],
        [
          {name: '复制链接', icon: 'link'},
          {name: '海报', icon: 'poster'},
        ]
      ],
    }
  },
  methods: {
    // 点击分享后的样式
    ClickShareSelect: async function (option) {
      let LocationURL = window.location.href
      switch (option.name) {
        case '复制链接':
          await Expen.CopySrc(LocationURL)
          Toast("已复制链接到剪切板")
          break
      }
    },
  }
}
</script>

<style scoped>

</style>