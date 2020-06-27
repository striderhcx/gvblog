<template>
  <div class="m-container-small m-padded-tb-massive" style="max-width: 30em !important;">
    <div class="ur container">
      <div class="ui middle aligned center aligned grid">
        <div class="column">
          <h2 class="ui teal image header">
            <div class="content">管理后台登录</div>
          </h2>
          <form class="ui large form">
            <div class="ui segment">
              <div class="field">
                <div class="ui left icon input">
                  <i class="user icon"></i>
                  <input type="text" v-model="data.username" placeholder="用户名" />
                </div>
              </div>
              <div class="field">
                <div class="ui left icon input">
                  <i class="lock icon"></i>
                  <input type="password" v-model="data.password" placeholder="密码" />
                </div>
              </div>
              <button type="button" @click="login" class="ui fluid large teal button">登 录</button>
            </div>
            <div v-if="isShowMsg">
              <div class="ui error mini message"></div>
              <div class="ui mini negative message">用户名和密码错误</div>
            </div>
          </form>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import request from "ajax/base.js"
import store from "../../store/index.js"

export default {
  name: "Login",
  components: {},
  data() {
    return {
      isShowMsg: false,
      data: {
        username: "",
        password: ""
      }
    };
  },
  activated() {
    document.onkeydown = e => {
      if (e.keyCode == 13) {
        this.login();
      }
    };
  },
  deactivated() {
    document.onkeydown = null;
  },
  inject: ["setNav"],
  mounted() {},
  computed: {},
  methods: {
    login() {
      request({
        url: "/login",
        method: "post",
        data: this.data,
      }).then(res => {
        console.log("login resp", res)
        let code = res.code
        let data = res.data
        let msg = res.msg
        if (code == 0) {
            // window.localStorage.setItem("ihxn-blog-token", data["token"])
            let auth_payload = {auth_token: data["token"]}
            console.log("登录调试：　this.$store", this.$store, auth_payload)
            this.$store.commit("setAuthToken", auth_payload)
            this.$router.push("/")　// Todo 还有ｂｕｇ：登录了之后跳转过去首页不能正常显示，需要手动刷新
        } else {
          alert("账号密码错误！")
        }
      })
    }
  }
};
</script>

<style scoped>
</style>