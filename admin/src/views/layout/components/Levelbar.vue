<template>
  <el-breadcrumb class="xa-block" separator="/">
    <el-breadcrumb-item 
      v-for="(item,index)  in levelList" 
      :key="`bread-${item.path}`">
      <!--最后一个面包屑不跳转-->
      <span v-if='item.redirect==="noredirect" || index==levelList.length-1' 
        class="xa-text-gray">
        {{item.name}}
      </span>
      <router-link v-else :to="item.redirect || item.path">{{item.name}}
      </router-link>
    </el-breadcrumb-item>
  </el-breadcrumb>
</template>

<script>
export default {
  name:'level-bar',
  data() {
    return {
      levelList: null
    }
  },
  methods: {
    getBreadcrumb() {
      let matched = this.$route.matched.filter(item => item.name)
      const first = matched[0]
        //如果第一个不包含首页，将首页添加到面包屑
      if (first && (first.name !== '首页' || first.path !== '')) {
        matched = [{
          name: '首页',
          path: '/'
        }].concat(matched)
      }
      this.levelList = matched
        //console.log('levelList==',this.levelList)
    }
  },
  watch: {
    $route() {
      this.getBreadcrumb()
    }
  },
  created() {
    this.getBreadcrumb()
  }
}
</script>
