<template>
  <div class="page me-page">
    <div class="page-header">
      <header class="header profie-header">
        <div class="header-left">
          <a href="javascript:history.go(-1);" 
            class="header-btn header-btn-back">
            返回</a>
        </div>
        <div class="header-title">
          个人中心
        </div>
      </header>
    </div>
    <div class="page-main">
      <div class="page-scroller">
        <section class="me-header">
          <div class="me-header-pic">
            <img src="uploads/images/pic1.jpg" />
          </div>
          <div class="me-header-text">
            <div class="me-header-name">
              <span>{{user.name}}</span>
            </div>
            <p>ID：{{user.mobile}}</p>
          </div>
        </section>
        <section class="me-order">
          <div class="me-order-hd">
            <h6 class="fl me-order-title">我的订单</h6>
            <a class="fr me-order-all" 
              :href="`/order?uid=${user.token}&state=0`">全部订单</a>
          </div>
          <div class="me-order-bd">
            <a class="me-order-link" 
              :href="`/order?uid=${user.token}&state=1`">
              <span>待付款</span>
              <!-- <em>2</em> -->
            </a>
            <a class="me-order-link" 
              :href="`/order?uid=${user.token}&state=3`">
              <span>待收货</span>
              <!-- <em>1</em> -->
            </a>
            <!-- <a class="me-order-link" 
              :href="`/order?uid=${user.token}&state=5`">
              <span>售后</span>
            </a> -->
            <a class="me-order-link" 
              :href="`/order?uid=${user.token}&state=5`">
              <span>已完成</span>
            </a>
          </div>
        </section>

        <section class="me-card">
          <div class="me-card-item arrow">
            <div class="me-card-key">收货地址</div>
            <a class="me-card-val" 
              :href="`/address?uid=${user.token}`" >管理我的地址</a>
          </div>
          <div class="me-card-item arrow">
            <div class="me-card-key">修改密码</div>
            <a class="me-card-val" href="/info">编辑</a>
          </div>
        </section>
        <div class="me-out" @click="handleLogout">退出登录</div>
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

export default {
  name: 'profile',
  data() {
    return {
      listQuery: {
        token: storage.get(CSMSKEY.Token)
      },
      user: {}
    }
  },
  methods: {
    ...mapActions(['getUser', 'logout']),
    get() {
      const self = this
      self.getUser(self.listQuery).then(response => {
        const res = response.data.data
        const state = response.data.state
        const message = response.data.msg
        if (state) {
          self.user = res
        } else {
          self.$toast(message, 'error')
        }
      })
    },
    handleLogout() {
      const self = this
      self.logout(self.listQuery).then(response => {
        const res = response.data.data
        const state = response.data.state
        const message = response.data.msg
        if (state) {
          storage.clear()
          self.$router.push('/login')
        } else {
          self.$toast(message, 'error')
        }
      })
    }
  },
  created() {
    this.get();
  }
}
</script>
<style type="text/css">
@import '../../styles/profile.css'
</style>
