<template>
  <div class="m-container-small m-padded-tb-big animated fadeIn">
    <div class="ui container">
      <!--header-->
      <div v-if="categoryList" class="ui top attached segment">
        <div class="ui middle aligned two column grid">
          <div class="column">
            <h3 class="ui teal header">分类</h3>
          </div>
          <div class="right aligned column">
            共
            <h2 class="ui orange header m-inline-block m-text-thin">{{ categoryList.length }}</h2>个
          </div>
        </div>
      </div>

      <div class="ui attached segment m-padded-tb-large" v-if="categoryList">
        <router-link
          tag="div"
          @click.native="getPageBlog(item.id)"
          class="ui labeled button m-margin-tb-tiny"
          v-for="(item,index) in categoryList"
          :to="'/categoryblog/' + item.id"
          :key="'categoryList'+index"
        >
          <span class="ui basic button" :class="{teal:$route.params.tid==item.id}">{{item.name}}</span>
          <div
            class="ui basic left pointing label"
            :class="{teal:$route.params.tid==item.id}"
            th:text="${#arrays.length(type.blogs)}"
          >{{item.post_count}}</div>
        </router-link>
      </div>

      <blog-list :pageJump="pageJump" :pageData="pageData" :pageBlog="pageBlog"></blog-list>
    </div>
  </div>
</template>

<script>
import BlogList from "components/blogList/BlogList";
import request from "ajax/base.js";

export default {
  name: "CategoryBlog",
  components: { BlogList },
  data() {
    return {
      pageData: {
        // sort: "",
        limit: 5,
        page: 0,
        category_id: 0
      },
      categoryList: null,
      pageBlog: {
        code: 0,
        msg: '',
        data:[],
        pagination: {total_page: 0}
      },
    };
  },
  computed: {},
  mounted() {},
  activated() {
    //获取分类列表
    request({
      url: "/category/list",
      method: 'get',
      params: {},
      }).then( res => {
        this.categoryList = res.data;
        //如果没有tid则跳转至第一个tid
        if (!this.$route.params.tid) {
          this.$router.push("/categoryblog/" + this.categoryList[0].id);
          this.getPageBlog(this.categoryList[0].id);
        } else {
          this.getPageBlog(this.$route.params.tid);
        }
      })
  },
  methods: {
    //获取博客分页查询列表
    getPageBlog(category_id) {
      if (category_id != -1) {
        this.pageData.category_id = category_id;
      }
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
      this.getPageBlog(-1);
    },
  }
};
</script>

<style scoped>
</style>