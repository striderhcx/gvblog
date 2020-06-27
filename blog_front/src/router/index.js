import store from "../store/index.js"

const routes = [{
  path: '/',
  redirect: '/ihxnblog'
}, {
  path: '/ihxnblog',
  component: () => import('views/ihxnBlog/IhxnBlog')
}, {
  path: '/categoryblog',
  component: () => import('views/categoryBlog/CategoryBlog'),
}, {
  path: '/categoryblog/:tid',
  component: () => import('views/categoryBlog/CategoryBlog'),
}, {
  path: '/tagblog',
  component: () => import('views/tagBlog/TagBlog')
}, {
  path: '/tagblog/:gid',
  component: () => import('views/tagBlog/TagBlog')
}, {
  path: '/about',
  component: () => import('views/about/About')
}, {
  path: '/archives',
  component: () => import('views/archives/Archives')
}, {
  path: '/search',
  component: () => import('views/search/Search')
}, {
  path: '/blog',
  component: () => import('views/blog/Blog')
}, {
  path: '/blog/:bid',
  component: () => import('views/blog/Blog')
}, {
  path: '/traffic',
  component: () => import('views/traffic/Traffic')
}, {
    path: '/create',
    component: () => import('views/admin/CreatePost'),
    meta:{
      requireAuth:true,//添加该字段，表示进入这个路由是需要登录的。
    },
  },
  {
    path: '/bloglogin',
    component: () => import('views/login/Login'),
  },
]

const router = new VueRouter({
  routes,
  mode: "history",
})

router.beforeEach( (to,　from,　next) =>　{
  if(to.meta.requireAuth){//判断该路由是否需要登录权限。
      if(store.state.auth_token){//通过vuex state获取当前的token是否存在。
          next();
      }else{
          console.log(to, "需要登录")
          next({
              path:'/bloglogin',
              query:{redirect:to.fullPath}//将跳转的路由path作为参数，登陆成功后跳转到该路由
          })
      }
  }else{
      next();
  }
})

export default router