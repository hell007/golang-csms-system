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
        订单确认
      </div>
    </header>
  </div>

  <div class="page-main">
    <div class="page-scroller">
      <section class="order-user">
        <h3>收货信息</h3>
        <div class="order-user-name">
          <span>{{order.consignee}}</span>
          <a :href="`tel:${order.phone}`">{{order.phone}}</a>
        </div>
        <p>{{order.province}}{{order.city}}{{order.district}}{{order.address}}</p>
      </section>

      <section class="order-panel">
        <div class="order-panel-bd">
          <a :href="`/product?id=${it.goodsId}`" 
            class="order-panel-item"
            v-for="it in list">
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
                合计：
                <span class="order-panel-price">
                  {{(it.goodsPrice * it.goodsNum).toFixed(2)}}元
                </span>
              </div>
            </div>
          </a>
        </div>
      </section>

      <section class="order-card">
        <div class="order-card-item">
          <div class="order-card-key">商品总价</div>
          <div class="order-card-val">￥{{order.totalPrice}}</div>
        </div>
        <div class="order-card-item">
          <div class="order-card-key">运费</div>
          <div class="order-card-val">+￥{{order.shipPrice}}</div>
        </div>
      </section>

    </div>
  </div>

  <div class="page-footer">
    <div class="bar">
      <div class="bar-left bar-left-confirm">
        <div class="bar-price">
          <em>订单金额：￥</em>
          <b>{{order.orderPrice}}</b>
        </div>
      </div>
      <div class="bar-right bar-right-confirm">
        <button class="btn btn-block btn-primary" type="button"
          @clcik="handleSubmit">
          提交订单
        </button>
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
  name: 'order-confrim',
  components: {
  },
  data() {
    return {
      link: URIS.Goods,
      id: '141111111111111',
      order: {},
      user: {},
      list: []
    }
  },
  computed: {
    isEdit() {
      this.id = this.$route.query.id ? this.$route.query.id : null
      return this.id
    }
  },
  methods: {
    ...mapActions(['getOrder', 'saveOrder']),
    get() {
      const self = this
      self.getOrder(self.id).then(response => {
        const res = response.data.data
        const state = response.data.state
        const message = response.data.msg
        if (state) {
          self.order = res.order
          self.list = res.list
          self.user = res.member
        } else {
          self.$toast(message, 'error')
        }
      })
    },
    handleSubmit() {
      const self = this
      let title = `确定要取消此订单？`
      let confirm = self.$confirm({
        text: title,
        button: [{
          text: '取消',
          ontap() {
            confirm.close()
          }
        }, {
          text: '确定',
          ontap() {
            confirm.close().then((res) => {
              self.saveOrder(self.order).then(response => {
                const res = response.data.data
                const state = response.data.state
                const message = response.data.msg
                if (state) {
                  self.$toast('订单取消成功!', 'sucess')
                  setTimeout(function(){
                    self.get()
                  }, 1000)
                } else {
                  self.$toast(message, 'error')
                }
              })
            });
          }
        }]
      })
    }
  },
  created() {
    if(this.isEdit) {
      this.get()
    }
  }
}
</script>
<style type="text/css">
@import '../../styles/order-confirm.css'
</style>
