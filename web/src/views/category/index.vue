<template>
  <div class="page">
    <div class="page-header category-header">
      <section class="category-search" 
        @click="handleSearch('/search')">
        <input 
          class="category-search-input" 
          type="text" 
          readonly 
          placeholder="搜索和小店的商品" />
      </section>
    </div>
    <div class="page-main">
      <div class="page-scroller category-scroller">
        <ul class="category-menu">
          <li :class="`category-menu-item ${index == visible ?  'on' : ''}`"
            v-for="item, index in categoryList" 
            @click="handleTab(index, item.id)">
            {{item.categoryName}}
          </li>
        </ul>
        <div class="category-panel">
          <a herf="javascript:;" class="category-panel-hd">
            <img src="uploads/images/banner2.jpg" />
          </a>
          <div class="category-panel-bd">
            <div class="category-panel-item" 
              v-for="item, index in goodsList">
              <a class="category-panel-pic" 
                :href="`/product?id=${item.id}`">
                <img :src="link + item.small" />
              </a>
              <div class="category-panel-text">
                <div class="category-panel-name">
                  {{item.goodsName}}
                </div>
                <div class="category-panel-row">规格：{{item.unit}}</div>
                <div class="category-panel-row">价格：
                  <span class="category-panel-price">￥{{item.goodsPrice}}</span>
                </div>
                <div class="category-panel-cart">
                  快速订购
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import {
  mapActions
} from 'vuex'

import { URIS } from '@/api/config'

export default {
  name: 'category',
  data() {
    return {
      link: URIS.Goods,
      categoryList: [],
      goodsList: [],
      listQuery: {
        pageNumber: 1,
        pageSize: 100,
        id: 0
      },
      visible: 0
    }
  },
  methods: {
    ...mapActions(['getCategory', 'getCategoryGoods']),
    get() {
      const app = this
      app.getCategory(app.listQuery).then(response => {
        const res = response.data.data
        app.categoryList = res.categoryList
        app.goodsList = res.goodsList
      });
    },
    handleTab(index, id) {
      const app = this
      app.visible = index
      app.listQuery.id = id
      app.getCategoryGoods(app.listQuery).then(response => {
        const res = response.data.data
        app.goodsList = res.rows
      });
    },
    handleSearch(url) {
      window.location.href = url
    }
  },
  created() {
    this.get();
  }
}
</script>
<style type="text/css">
@import '../../styles/category.css'
</style>
