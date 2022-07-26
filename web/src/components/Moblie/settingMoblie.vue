<template>
  <van-nav-bar title="我的"/>
  <!--  检测是否登录-->
  <van-action-sheet v-model:show="Config.JwtBool" title="注册/登录账号" @cancel="ClickJwtDevice"
                    @click-overlay="ClickJwtDevice">
    <div style="">
      <van-cell-group inset>
        <van-field v-model="Account.UserMail" label="账号" placeholder="请输入账号邮箱"/>
        <van-field v-model="Account.UserName" v-if="Account.Register" label="昵称" placeholder="请输入账号昵称"/>
        <van-field v-model="Account.PassWord" label="密码" type="password" placeholder="请输入账号密码"/>
        <van-field v-model="Account.PassWordS" v-if="Account.Register" type="password" label="重复密码"
                   placeholder="请再次输入密码"/>
        <!--        普通的图像验证码输入框-->
        <van-field
            v-if="!Account.Register"
            v-model="Account.Captcha"
            center
            clearable
            label="图形验证码"
            placeholder="请输入图形验证码"
        >

          <template #right-icon>
            <img :src="Config.CaptchaImg" style="height: 2rem" @click="ClickImgCaptcha">
          </template>
        </van-field>


        <!--        当选择注册时弹出邮箱验证码对话框，当选择登陆时弹出普通图片验证码对话框-->
        <van-field
            v-if="Account.Register"
            v-model="Account.MailCaptcha"
            center
            clearable
            label="邮箱验证码"
            placeholder="请输入邮箱验证码"
        >
          <template #button>
            <van-button size="small" type="primary" @click="ClickPostMailNumber">
              发送验证码
              <van-count-down v-if="Account.PostMail" millisecond :time="MailTime" format="ss:SS"
                              @finish="DeleteTimeMail"/>
            </van-button>
          </template>
        </van-field>


        <van-cell center title="选择注册账号">
          <template #right-icon>
            <van-switch v-model="Account.Register" size="24"/>
          </template>
        </van-cell>

        <van-button type="primary" block @click="ClickAccount">确定</van-button>
      </van-cell-group>
    </div>


  </van-action-sheet>

  <div style="height: 0.7rem;width: 100%;background-color: #f4f4f5;"></div>
  <div class="p-6 max-w-sm mx-auto bg-white rounded-xl shadow-md flex items-center space-x-4">
    <div class="flex-shrink-0">
      <el-avatar class="float-left ...  w-11" :src="Config.ImgUrl" :lazy="true"/>
    </div>
    <p class="text-xl font-medium text-black">{{ Config.Name }}</p>
  </div>

  <div style="height: 0.7rem;width: 100%;background-color: #f4f4f5;"></div>
  <van-grid>
    <van-grid-item icon="manager-o" text="个人主页"/>
    <van-grid-item icon="balance-o" text="创作中心"/>
    <van-grid-item icon="guide-o" text="论坛管理"/>
    <van-grid-item icon="envelop-o" text="消息中心"/>
  </van-grid>
  <div style="height: 0.7rem;width: 100%;background-color: #f4f4f5;"></div>
  <van-cell-group inset>
    <van-cell title="网站介绍" is-link/>
    <van-cell title="更新日志" is-link/>
    <van-cell title="分享网站" is-link @click="Config.ShareList=true"/>
    <OptionListShare
        v-model:show="Config.ShareList"
    />
    <van-cell title="关于" is-link @click="ClickDialogAbout"/>
  </van-cell-group>

  <!--  生成二维码后弹出的结果-->
  <van-popup v-model:show="Config.QRCodeBool">
    <el-image :src="Config.QRCode"/>
  </van-popup>
</template>

<script>
import {Dialog, Toast} from "vant";
import UserApi from "@/assets/js/RouterUser";
import TouristApi from "@/assets/js/RouterTourist";
import {ref} from "vue";

export default {
  name: "settingMoblie",
  components: {
    OptionListShare: () => import("./view/OptionListShare")
  },
  setup() {
    const MailTime = ref(60 * 1000);
    return {MailTime}
  },
  data() {
    return {
      Config: {
        JwtBool: false,
        CaptchaImg: "",
        NumberKey: false,
        MailTimeClick: 0,
        ImgUrl: "",
        Name: "",
      },
      Account: {
        UserMail: "",
        UserName: "",
        PassWord: "",
        PassWordS: "",
        Captcha: "",
        MailCaptcha: "",
        Register: false,
        PostMail: false
      }
    }
  },
  methods: {
    ClickJwtDevice: async function () {
      Toast("您取消了登录，网站部分功能无法正常使用")
    },
    // 更换验证码
    ClickImgCaptcha: async function () {
      this.Config.CaptchaImg = await TouristApi.GetCaptcha()
    },
    // 进行登录或注册
    ClickAccount: async function () {
      this.Account.Register ? this.Config.JwtBool =
              await TouristApi.PostRegisterAccount(this.Account) :
          this.Config.JwtBool = await TouristApi.PostAccountLogin(this.Account)
      this.$router.go(0)

    },
    // 发送邮件
    ClickPostMailNumber: async function () {
      if (this.Account.PostMail === true && this.Config.MailTimeClick !== 0) {
        Toast("请等待倒计时结束后再发送邮箱")
        return
      }

      if (this.Account.UserMail !== "") {
        this.Account.PostMail = true
        this.Config.MailTimeClick++
        await TouristApi.PostRegisterAccountMail(this.Account.UserMail)
      } else {
        Toast("请输入账号邮箱后再请求发送邮箱")
      }
    },
    // 清除邮箱请求状态
    DeleteTimeMail: async function () {
      this.Account.PostMail = false
      this.Config.MailTimeClick = 0
    },
    // 点击关于的结果
    ClickDialogAbout: async function () {
      Dialog.alert({
        title: '关于',
        message: '由StarYuhen进行开发，商业合作请联系QQ:3446623843',
        theme: 'round-button',
        confirmButtonColor: '#6EE7B7',
      }).then(() => {
        // on close
      });
    },
  },
  async mounted() {
    // 设置账号状态
    this.Config.JwtBool = await UserApi.GetJwtDevice()
    if (this.Config.JwtBool) {
      this.Config.CaptchaImg = await TouristApi.GetCaptcha()
    }
    this.Config.ImgUrl = localStorage.img
    this.Config.Name = localStorage.name
    // // 查询头像地址并保存，倘若没有更新就不会刷新
    // if (this.Config.ImgUrl === undefined || this.Config.ImgUrl === "" || this.Config.ImgUrl === null) {
    //   await UserApi.GetImgUser()
    //   this.Config.ImgUrl = localStorage.ImgUrl
    // }
  },
}
</script>

<style scoped>

</style>