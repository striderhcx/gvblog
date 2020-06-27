<template>
  <!--中间内容-->
  <div class="m-container m-padded-tb-big animated fadeIn">
    <div class="ui container">
      <div class="ui stackable grid">
        <!--左边博客列表-->
        <div class="eleven wide column">
          <div v-if="isLogin">
            <router-link to="/create" tag="button" class="ui teal button">写新博客</router-link>
          </div>
          <!--header-->
          <div class="ui top attached segment">
            <div class="ui middle aligned two column grid">
              <div class="column">
                <h3 class="ui teal header">文章</h3>
              </div>
              <div class="right aligned column">
                共
                <h2 class="ui orange header m-inline-block m-text-thin">{{pageBlog.totalElements}}</h2>篇
              </div>
            </div>
          </div>

          <blog-list :pageJump="pageJump" :pageData="pageData" :pageBlog="pageBlog"></blog-list>
        </div>

        <!--右边的top-->
        <div class="five wide column">
          <!--分类-->
          <div class="ui segments">
            <div class="ui secondary segment">
              <div class="ui two column grid">
                <div class="column">
                  <i class="idea icon"></i>分类
                </div>
                <div class="right aligned column">
                  <router-link to="/categoryblog/">
                    more
                    <i class="angle double right icon"></i>
                  </router-link>
                </div>
              </div>
            </div>
            <div class="ui teal segment">
              <div class="ui fluid vertical menu">
                <router-link
                  :to="'/categoryblog/' + item.id"
                  class="item"
                  v-for="(item,index) in listType"
                  :key="'getListType'+index"
                >
                  <span>{{item.name}}</span>
                  <div class="ui teal basic left pointing label">{{item.post_count}}</div>
                </router-link>
              </div>
            </div>
          </div>

          <!--标签-->
          <div class="ui segments m-margin-top-large">
            <div class="ui secondary segment">
              <div class="ui two column grid">
                <div class="column">
                  <i class="tags icon"></i>标签
                </div>
                <div class="right aligned column">
                  <router-link to="/tagblog/">
                    more
                    <i class="angle double right icon"></i>
                  </router-link>
                </div>
              </div>
            </div>
            <div class="ui teal segment">
              <router-link
                :to="'/tagblog/' + item.id"
                :style="{
                            'color':item.color+'!important',
                            'border-color':item.color+'!important'
                          }"
                class="ui teal basic left pointing label m-margin-tb-tiny"
                v-for="(item,index) in listTag"
                :key="'listTag'+index"
              >
                <span>{{item.name}}</span>
                <div class="detail">{{item.post_count}}</div>
              </router-link>
            </div>
          </div>

          <!--最新推荐-->
          <div class="ui segments m-margin-top-large">
            <div class="ui secondary segment">
              <i class="bookmark icon"></i>最新推荐
            </div>
            <div
              class="ui segment"
              v-for="(item,index) in listRecommend"
              :key="'listRecommend'+index"
            >
              <router-link tag="a" :to="'/blog/'+item.id" class="m-black m-text-thin">{{item.title}}</router-link>
            </div>
          </div>

          <!--广告-->
          <advertising></advertising>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Advertising from "components/advertising/Advertising";
import BlogList from "components/blogList/BlogList";
import request from "ajax/base.js";

export default {
  name: "IhxnBlog",
  components: { Advertising, BlogList },
  data() {
    return {
      pageData: {
        // sort: "",
        limit: 5,
        page: 0
      },
      // todo 这里为了console不报错,直接初始化一个pageBlog对象,应该有更好的方式
      pageBlog: {
        code: 0,
        msg: '',
        data:[],
        pagination: {total_page: 0}
      },
      listType: [],
      listTag: [],
      listRecommend: []
    };
  },
  mounted() {
    console.log("ihxnBlog mounted!!!!")
    this.getPageBlog();
    this.getListType();
    this.getListTag();
  },

  activated() {
    console.log("ihxnBlog activated!!!!")
    this.getPageBlog();
    this.getListType();
    this.getListTag();
  },

  computed: {
    isLogin() {
      console.log("islogin:", this.$store.state.auth_token)
      return this.$store.state.auth_token != ''
    }
  },
  methods: {
    //获取博客分页查询列表
    getPageBlog() {
      request({
        url: "/post/list",
        method: 'get',
        params: this.pageData,
        }).then( res => {
          this.pageBlog = res;
      })
    },
    //获取分类排名
    getListType() {
      request({
        url: "/category/list",
        method: 'get',
        params: {},
      }).then( res => {
        this.listType = res.data;
      })
    },
    //获取标签排名
    getListTag() {
      request({
        url: "/tag/list",
        method: 'get',
        params: {},
      }).then( res => {
        this.listTag = res.data;
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