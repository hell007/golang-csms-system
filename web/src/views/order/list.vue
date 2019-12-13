<template>
<div class="page">
  <div class="page-header">
    <header class="header">
      <div class="header-left">
        <a href="javascript:history.go(-1);" 
          class="header-btn header-btn-back">
          返回
        </a>
      </div>
      <div class="header-title">
        我的订单
      </div>
    </header>

    <ul class="tab order-tab js-tab">
      <li :class="`tab-item ${item.value == listQuery.state ? 'tab-item-active' : ''}`"
        v-for="item in states"
        @click="handleState(item)">
        <span class="tab-inner">{{item.label}}</span>
      </li>
    </ul>
  </div>

  <div class="page-main">
    <div class="page-scroller js-scroller">
      <section class="js-list">
        <div class="order-panel"
          v-for="item in list">
          <div class="order-panel-hd">
            <div class="order-panel-store">
              <span>{{item.ordersn}}</span>
            </div>
            <span class="order-panel-state">
              <b v-if="item.orderState==1">待付款</b>
              <b v-if="item.orderState==2">待发货</b>
              <b v-if="item.orderState==3">已发货</b>
              <b v-if="item.orderState==4">已取消</b>
              <b v-if="item.orderState==5">已完成</b>
            </span>
          </div>
          <div class="order-panel-bd">
            <a :href="`/product?id=${it.goodsId}`" class="order-panel-item"
              v-for="it in item.list">
              <div class="order-panel-picture">
                <img
                  :src="link + it.small"/>
              </div>
              <div class="order-panel-text">
                <div class="order-panel-info">
                  名称：{{it.goodsName}} {{it.color}}
                </div>
                <div class="order-panel-info">
                  单价：{{it.goodsPrice}}元
                </div>
                <div class="order-panel-info">
                  级别：{{it.goodsSku}}
                  <span class="order-panel-num">
                    x {{it.goodsNum}}扎
                  </span>
                </div>
                <div class="order-panel-info">
                  总价：
                  <span class="order-panel-price">
                    {{(it.goodsPrice * it.goodsNum).toFixed(2)}}元
                  </span>
                </div>
              </div>
            </a>
          </div>
          <div class="order-panel-ft">
            <div class="order-panel-total">
              <span>运费：<b>{{item.shipPrice}}</b>元</span>,
               <span>总金额：<b>{{item.orderPrice}}</b></span>
            </div>
            <div class="order-panel-opt">
              <a class="btn btn-pill btn-gray" 
                :href="`detail?id=${item.ordersn}`">&nbsp;详情&nbsp;</a>
              <!-- <button class="btn btn-pill btn-red">立即付款</button> -->
            </div>
          </div>
        </div>
      </section>
      <div class="no-entry" v-if="list.length == 0">暂无订单信息</div>
    </div>
  </div>
</div>
</template>
<script>
import {
  mapActions
} from 'vuex'

import * as storage from '@/utils/storage'

import {
  CSMSKEY
} from '@/api/config'
import { URIS } from '@/api/config'

export default {
  name: 'order',
  components: {
  },
  data() {
    return {
      link: URIS.Goods,
      listQuery: {
        pageNumber: 1,
        pageSize: 10,
        state: 0,
        token: storage.get(CSMSKEY.Token)
      },
      states: [{
        value: 0,
        label: '全部'
      }, {
        value: 1,
        label: '待付款'
      }, {
        value: 2,
        label: '待发货'
      }, {
        value: 3,
        label: '已发货'
      }, {
        value: 4,
        label: '已取消'
      }, {
        value: 5,
        label: '已完成'
      }],
      list: []
    }
  },
  computed: {
    isEdit() {
      this.listQuery.state = this.$route.query.state ? this.$route.query.state : 0
      return this.listQuery.state
    }
  },
  methods: {
    ...mapActions(['getOrderList']),
    get() {
      const self = this
      self.getOrderList(self.listQuery).then(response => {
        const res = response.data.data.rows
        const state = response.data.state
        const message = response.data.msg
        if (state) {
          self.list = res
        } else {
          self.$toast(message, 'error')
        }
      })
    },
    handleState(item) {
      const self = this
      self.listQuery.state = item.value
      self.listQuery.pageNumber = 1
      self.get()
    }
  },
  created() {
    if(this.isEdit){
      this.get()
    }
  }
}
</script>
<style type="text/css">
@import '../../styles/order.css'
</style>
