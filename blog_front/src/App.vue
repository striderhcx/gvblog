<template>
  <div id="app" :class="isNight?'night':'day'">
    <nav-bar></nav-bar>
    <keep-alive>
      <router-view class="info" style="min-height:800px"></router-view>
    </keep-alive>
    <footer-bar></footer-bar>
    <div class="moon" v-if="isNight"></div>
  </div>
</template>

<script>
import "ajax/base";
import NavBar from "components/navBar/NavBar";
import FooterBar from "components/footer/FooterBar";

export default {
  name: "App",
  components: { NavBar, FooterBar },
  data() {
    return {
      isNight: false
    };
  },
  created() {
    //在页面加载时读取sessionStorage里的状态信息
    if (sessionStorage.getItem("store") ) {
      this.$store.replaceState(Object.assign({}, this.$store.state,JSON.parse(sessionStorage.getItem("store"))))
    }
  },
  mounted() {
    $.post({
      url: "trafficUp"
    });
  },
  provide() {
    return {
      setNight: this.setNight
    };
  },
  methods: {
    setNight() {
      this.isNight = !this.isNight;
    }
  }
};
</script>

<style>
@import "assets/css/base.css";
@import "assets/css/me.css";
@import "assets/css/background.css";
</style>
