<template>
  <div class="page">
    <div class="page-header home-header">
      <section class="home-search" 
        @click="handleGo('/search')">
        <input class="home-search-input" 
          type="text" 
          readonly 
          placeholder="搜索和小店的商品" />
      </section>
    </div>
    <div class="page-main">
      <div class="page-scroller">
        <!--banner-->
        <section class="home-banner">
          <div class="swiper-container">
            <div class="swiper-wrapper">
              <div class="swiper-slide">
                <img src="uploads/images/banner1.jpg" />
              </div>
              <div class="swiper-slide">
                <img src="uploads/images/banner2.jpg" />
              </div>
            </div>
            <!-- Add Pagination -->
            <div class="swiper-pagination"></div>
          </div>
        </section>
        <!--menu-->
        <section class="home-menu">
          <a class="home-menu-item" href="###">
            <div class="picture"><img src="http://temp.im/80x80" /></div>
            <span class="name">店铺</span>
          </a>
          <a class="home-menu-item" href="###">
            <div class="picture">
              <img src="uploads/images/pic1.jpg" />
            </div>
            <span class="name">手机</span>
          </a>
          <a class="home-menu-item" href="###">
            <div class="picture">
              <img src="uploads/images/pic2.jpg" />
            </div>
            <span class="name">热门业务</span>
          </a>
          <a class="home-menu-item" href="###">
            <div class="picture">
              <img src="uploads/images/pic3.jpg" />
            </div>
            <span class="name">智能家居</span>
          </a>
        </section>
        <!--热销商品-->
        <section class="home-panel">
          <div class="home-panel-hd">
            <h5>热销商品</h5>
            <span>探索你的专属系列</span>
          </div>
          <div class="home-panel-bd">
            <div class="home-panel-item" 
              v-for="item,index in hotList">
              <a class="home-panel-picture" 
                :href="'/product?id='+item.id">
                <img :src="link + item.small" />
              </a>
              <div class="home-panel-name">
                {{item.goodsName}}
              </div>
              <div class="home-panel-price">
                ￥{{item.goodsPrice}}
              </div>
            </div>
          </div>
        </section>
        <!--主推业务-->
        <section class="home-panel">
          <div class="home-panel-hd">
            <h5>主推业务</h5>
            <span>主推专属你口味的业务</span>
          </div>
          <div class="home-panel-bd">
            <a class="home-panel-item" :href="'/product?id='+item.id" v-for="item,index in firstList">
              <div class="home-panel-picture">
                <img :src="item.small" />
              </div>
              <div class="home-panel-name">
                {{item.goodsName}}
              </div>
              <div class="home-panel-price">￥{{item.goodsPrice}}</div>
            </a>
          </div>
        </section>
      </div>
    </div>
    <Footer></Footer>
  </div>
</template>
<script>
import {
  mapActions
} from 'vuex'

import { URIS } from '@/api/config'

import Footer from '@/components/Footer'

import Swiper from 'swiper'
import 'swiper/css/swiper.min.css'

export default {
  name: 'home',
  components: {
    Footer
  },
  data() {
    return {
      link: URIS.Goods,
      hotList: [],
      firstList: [],
      listQuery: {
        pageNumber: 1,
        pageSize: 4
      }
    }
  },
  methods: {
    ...mapActions(['getHome']),
    get() {
      const self = this
      self.getHome(self.listQuery).then(response => {
        const res = response.data.data
        self.hotList = res.hot
        self.firstList = res.first
      });
    },
    handleGo(url) {
      window.location.href = url
    }
  },
  created() {
    this.get();
  },
  mounted() {
    // swiper
    window.onload = function() {
      setTimeout(function(){
        var swiper = new Swiper('.swiper-container', {
          pagination: {
            el: '.swiper-pagination'
          }
        })
      },200)
    }
  },
}
</script>
<style type="text/css">
@import '../../styles/home.css'
</style>
