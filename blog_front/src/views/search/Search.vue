<template>
  <!--中间内容-->
  <div class="m-container-small m-padded-tb-big">
    <div class="ui container">
      <!--header-->
      <div class="ui top attached segment">
        <div class="ui middle aligned two column grid">
          <div class="column">
            <h3 class="ui teal header">搜索 {{$route.query.key}} 的结果</h3>
          </div>
          <div class="right aligned column">
            共
            <h2 class="ui orange header m-inline-block m-text-thin">{{pageBlog.totalElements}}</h2>条
          </div>
        </div>
      </div>

      <blog-list :pageJump="pageJump" :pageData="pageData" :pageBlog="pageBlog"></blog-list>
    </div>
  </div>
</template>

<script>
import BlogList from "components/blogList/BlogList";
import request from "ajax/base.js";
export default {
  name: "Search",
  components: { BlogList },
  data() {
    return {
      pageData: {
        // sort: "",
        limit: 5,
        page: 0,
        search: ""
      },
      pageBlog: {}
    };
  },
  activated() {
    if (this.$route.query.key == null || this.$route.query.key == "") {
      this.$router.push("/");
    } else {
      this.getPageBlog();
    }
  },
  computed: {},
  watch: {
    $route(newRoute) {
      this.getPageBlog();
    }
  },
  methods: {
    //获取博客分页查询列表
    getPageBlog() {
      this.pageData.search = this.$route.query.key
      request({
        url: "/post/list",
        method: 'get',
        params: this.pageData,
        }).then( res => {
          this.pageBlog = res;
        })
    },
    //跳页
    pageJump(a) {
      this.pageData.page += a;
      this.getPageBlog();
    }
  }
};
</script>

<style scoped>
</style>