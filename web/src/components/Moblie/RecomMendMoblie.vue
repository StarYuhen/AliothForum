<template>

  <van-nav-bar title="推荐"/>
  <van-search
      v-model="config.value"
      placeholder="请输入搜索关键词"
      input-align="center"
  />
  <!--  内容使用列表刷新-->
  <van-list
      v-model:loading="loading"
      :finished="finished"
      finished-text="没有更多了"
      @load="onLoad"
  >
    <div v-for="i in src.ListSrc" v-bind:key="i.id" @click="this.$router.push('/article/1657156821922318800')">
      <div style="height: 0.7rem;width: 100%;background-color: #f4f4f5;"></div>
      <div style="margin-left: 1rem;margin-right: 1rem;margin-top: 1rem;">
        <!--帖子基本内容-->
        <!--标题-->
        <text style="font-weight:bolder;font-size: 1rem">有哪些知识，你知道后你的世界马上就不一样了？如何做到更好的自己？</text>
        <div>
          <div style="height: 2rem">
            <!--头像及用户名字-->
            <van-image
                round
                width="1.6rem"
                height="1.6rem"
                src="https://q1.qlogo.cn/g?b=qq&nk=3446623843&s=640"
            />

            <text class="recomName">StarYuhen</text>
          </div>

          <!--          帖子内容-->
          <div class="van-multi-ellipsis--l3">
            <text style="font-size: 0.8rem">某公司没钱用了,他就随便找来一张纸,写上100万,然后卖给你(免不了要吹一通说这纸怎么怎么值钱,要不然你也不会去买),这张纸就是股票
              如果他赚了钱,你要凭这张纸来领钱,这叫派息
              如果他亏损了,不要紧,他说明年我们会赚的,这叫预期收益 当然,每年都要告诉赚没赚,出个文件,这叫年报表 怕你老担心,年中出一次,这叫半年报 还不行?一个季度汇报一次,叫季报 你还是拿不准,要求退钱,
            </text>
          </div>
          <!--  赞同，浏览量，评论-->
          <div style="margin-top: 0.7rem;margin-bottom: 0.7rem">
            <text style="font-size: 0.6rem;color: #c8c9cc">200000浏览·15000赞同·840评论</text>
          </div>

        </div>
      </div>

    </div>

  </van-list>

</template>

<script>
// 论坛推荐界面

import {ref} from "vue";


export default {
  name: "RecomMendMoblie",
  // 引用vuex状态树
  // store,
  data() {
    const list = ref([]);
    const loading = ref(false);
    const finished = ref(false);

    const onLoad = () => {
      // 异步更新数据
      // setTimeout 仅做示例，真实场景中一般为 ajax 请求
      setTimeout(() => {
        for (let i = 0; i < 10; i++) {
          list.value.push(list.value.length + 1);
        }

        // 加载状态结束
        loading.value = false;

        // 数据全部加载完成
        if (list.value.length >= 40) {
          finished.value = true;
        }
      }, 1000);
    };

    return {
      config: {},
      list,
      onLoad,
      loading,
      finished,
      src: {
        ListSrc: [{id: 0}, {id: 1}, {id: 2}, {id: 3}, {id: 4}, {id: 5}]
      }
    }
  },
  mounted() {
  }
}
</script>

<style scoped>


</style>