<template>
  <div class="page-footer">
    <nav class="nav">
      <a v-for="item in navs"
        class="nav-item" 
        :key="`nav-${item.name}`"
        :href="item.url">
        <i :class="item.icon"></i>
        <span>{{item.name}}</span>
      </a>
    </nav>
  </div>
</template>

<script>
export default {
  name: 'Footer',
  data() {
    return {
      navs: [{
        id: 1,
        name: '首页',
        url: '/home',
        icon: 'icon-nav-home'
      },{
        id: 1,
        name: '分类',
        url: '/category',
        icon: 'icon-nav-category'
      },{
        id: 1,
        name: '订购单',
        url: '/cart',
        icon: 'icon-nav-cart'
      },{
        id: 1,
        name: '我的',
        url: '/profile',
        icon: 'icon-nav-me'
      }]
    }
  }
}
</script>
