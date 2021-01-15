<template>
  <div class="page" v-cloak id="js-app">
    <div class="page-header">
      <header class="header">
        <div class="header-left">
          <a href="javascript:history.go(-1);" class="header-btn header-btn-back">返回</a>
        </div>
        <section class="search-panel">
          <div class="search-panel-box">
            <input v-model="listQuery.key" 
              class="search-panel-input" 
              type="text" 
              placeholder="请输入">
            <button class="search-panel-clear" 
              type="button" 
              v-show="listQuery.key" 
              @click="handleClear">
              清空
            </button>
          </div>
          <span class="search-panel-submit js-search-submit" @click="handleSearch">搜索</span>
        </section>
      </header>
    </div>
    <div class="page-main">
      <div class="page-scroller" 
        v-infinite-scroll="loadMore" 
        infinite-scroll-disabled="busy" 
        infinite-scroll-distance="0">
        <!--列表-->
        <div class="search-panel-list">
          <div class="search-panel-item" v-for="item,index in list">
            <a class="search-panel-picture" 
              :href="`/product?id=${item.id}`">
              <img :src="link + item.small" />
            </a>
            <div class="search-panel-text">
              <div class="search-panel-name">
                {{item.goodsName}}
              </div>
              <div class="search-panel-row">规格：20支/扎</div>
              <div class="search-panel-row">价格：
                <span class="search-panel-price">￥{{item.goodsPrice}}</span>
              </div>
              <div class="search-panel-cart">
                快速订购
              </div>
            </div>
          </div>
        </div>
        <div class="no-entry">{{tips}}</div>
      </div>
    </div>
  </div>
</template>

<script>
import {
  mapActions
} from 'vuex'

import infiniteScroll from 'vue-infinite-scroll'

import {URIS} from '@/api/config'

export default {
  name: 'search',
  components: {},
  directives: {
    infiniteScroll
  },
  data() {
    return {
      link: URIS.Goods,
      listQuery: {
        key: '',
        pageNumber: 1,
        pageSize: 2,
        sortName: "create_time",
        sortOrder: "desc"
      },
      list: [],
      busy: false,
      tips: '',
      totalPage: 0
    }
  },
  methods: {
    ...mapActions(['getSearch']),
    loadMore() {
      const app = this
      app.busy = true
      app.getSearch(app.listQuery).then(response => {
        const state = response.data.state
        const msg = response.data.msg
        const res = response.data.data
        if (state) {
          let list = res.rows
          app.totalPage = Math.ceil(res.total / app.listQuery.pageSize)
          if (list.length == 0 && app.totalPage == 0) {
            app.tips = '暂无数据'
            return
          } else {
            if (app.totalPage >= app.listQuery.pageNumber) {
              for (var item of list) {
                app.list.push(item)
              }
              app.listQuery.pageNumber++
              app.busy = false  
            } else {
              app.tips = "- 亲，已经到底了 -"
            }
          }
        } else {
          //utils.alert(msg)
        }
      });
    },
    handleSearch() {
      const app = this
      if (!app.listQuery.key) return
      app.listQuery.pageNumber = 1
      app.list = []
      app.loadMore()
    },
    handleClear() {
      const app = this
      app.listQuery.key = ''
      app.listQuery.pageNumber = 1
      app.list = []
      app.loadMore()
    },
    handleGo(url) {
      window.location.href = url
    }
  },
  created() {
  }
}
</script>
<style type="text/css">
@import '../../styles/search-list.css'
</style>
