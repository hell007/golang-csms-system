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
        <span v-if="order.orderState==1">待付款</span>
        <span v-if="order.orderState==2">待发货</span>
        <span v-if="order.orderState==3">已发货</span>
        <span v-if="order.orderState==4">已取消</span>
        <span v-if="order.orderState==5">已完成</span>
      </div>
    </header>
  </div>

  <div class="page-main">
    <div class="page-scroller">
      <section class="order-card">
        <div class="order-card-item">
          <div class="order-card-key">收货人</div>
          <div class="order-card-val">{{order.consignee}} {{order.phone}}</div>
        </div>
        <div class="order-card-item">
          <div class="order-card-key">收货地址</div>
          <div class="order-card-val">
            {{order.province}}{{order.city}}{{order.district}}
            {{order.address}}
          </div>
        </div>
        <div class="order-card-item">
          <div class="order-card-key">下单人</div>
          <div class="order-card-val">{{user.name}} {{user.mobile}}</div>
        </div>
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
          <div class="order-card-key">订单编号</div>
          <div class="order-card-val">
            <span>{{order.ordersn}}</span>
            <span
              class="v-copy js-copy"
              @click="handleClipboard"
              :data-clipboard-text="order.ordersn">
              复制
            </span>
          </div>
        </div>
        <div class="order-card-item">
          <div class="order-card-key">下单时间</div>
          <div class="order-card-val">
            {{order.createTime | parseTime}}
          </div>
        </div>
      </section>

      <section class="order-card">
        <div class="order-card-item">
          <div class="order-card-key">支付方式</div>
          <div class="order-card-val">{{order.payName}}</div>
        </div>
        <div class="order-card-item">
          <div class="order-card-key">支付时间</div>
          <div class="order-card-val">
            {{order.payTime | parseTime}}
          </div>
        </div>
      </section>

      <section class="order-card">
        <div class="order-card-item">
          <div class="order-card-key">配送方式</div>
          <div class="order-card-val">{{order.shipName}}</div>
        </div>
        <div class="order-card-item">
          <div class="order-card-key">物流单号</div>
          <div class="order-card-val">
            <span>{{order.shipNo}}</span>
            <span
              class="v-copy js-copy"
              @click="handleClipboard"
              :data-clipboard-text="order.shipNo">
              复制
            </span>
          </div>
        </div>
        <div class="order-card-item">
          <div class="order-card-key">发货时间</div>
          <div class="order-card-val">
            {{order.shipTime | parseTime}}
          </div>
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
        <div class="order-card-item">
          <div class="order-card-key">订单金额</div>
          <div class="order-card-val">￥{{order.orderPrice}}</div>
        </div>
        <div class="order-card-item">
          <div class="order-card-key">应付款</div>
          <div class="order-card-val">
            <span class="v-price">￥{{order.payPrice}}</span>
          </div>
        </div>
      </section>
    </div>
  </div>

  <div class="page-footer">
    <div class="bar order-bar">
      <div class="bar-right order-bar-right">
        <button class="btn btn-gray" type="button"
          @click="handleEdit">取消订单</button>
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

import Clipboard from 'clipboard'

export default {
  name: 'order-detail',
  components: {
  },
  data() {
    return {
      link: URIS.Goods,
      id: '',
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
    handleClipboard() {
      const self = this
      const clipboard = new Clipboard('.js-copy')
      clipboard.on('success', () => {
        self.$toast('复制成功！', 'error')
        clipboard.destroy()
      })
      clipboard.on('error', () => {
        self.$toast('复制失败！', 'error')
        clipboard.destroy()
      })
      clipboard.onClick(event)
    },
    handleEdit() {
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
@import '../../styles/order-detail.css'
</style>
