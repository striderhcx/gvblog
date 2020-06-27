<template>
  <footer class="ui inverted vertical segment m-padded-tb-massive">
    <div class="ui center aligned container">
      <div class="ui inverted divided stackable grid">
        <div class="four wide column">
          <h4 class="ui inverted header m-text-thin m-text-spaced">友情链接</h4>
          <div class="ui inverted link list">
            <a href="http://ihxn.herokuapp.com/">本站老版本</a>
          </div>
        </div>
        <div class="four wide column">
          <h4 class="ui inverted header m-text-thin m-text-spaced">联系我</h4>
          <div class="ui inverted link list">
            <a class="item m-text-thin">email: 1225427292@qq.com</a>
            <a class="item m-text-thin">微信: cdcdlsls</a>
          </div>
        </div>
        <div class="four wide column">
          <br />
          <br />
          <p class="m-text-thin m-text-spaced m-opacity-mini">唯有时间对我们一视同仁</p>
        </div>
      </div>
      <div class="ui inverted section divider"></div>
      <p class="m-text-thin m-text-spaced m-opacity-tiny">
        <a
          style="color:rgba(255,255,255,.9)"
          href="http://www.beian.gov.cn/"
          target="_blank"
        >本站是golang学习的产物，主要是学习如何使用golang进行web开发以及熟悉golang的语法
      </p>
    </div>
  </footer>
</template>

<script>
export default {
  name: "Footer",
  components: {},
  data() {
    return {
      listContact: [],
      listLink: []
    };
  },
  mounted() {
    $.get({
      url: "contacts",
      success: res => {
        this.listContact = res[0];
      }
    });
    $.get({
      url: "links",
      success: res => {
        this.listLink = res;
      }
    });
  },
  computed: {},
  methods: {}
};
</script>

<style scoped>
</style>