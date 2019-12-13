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
        修改密码
      </div>
    </header>
  </div>

  <div class="page-main">
    <div class="page-scroller">
      <section class="me-panel">
        <div class="me-panel-bd">
          <div class="me-panel-item">
            <div class="me-panel-key">姓名</div>
            <div class="me-panel-val">
              <input
                class="v-input"
                type="text"
                :readonly="listQuery.token ? true : false"
                v-model="user.name"
                placeholder="请输入姓名"
              />
            </div>
          </div>
          <div class="me-panel-item">
            <div class="me-panel-key">手机号</div>
            <div class="me-panel-val">
              <input
                class="v-input"
                type="number"
                :readonly="listQuery.token ? true : false"
                v-model="user.mobile"
                placeholder="请输入手机号"
              />
            </div>
          </div>
          <div class="me-panel-item">
            <div class="me-panel-key">验证码</div>
            <div class="me-panel-val">
              <input
                class="v-input"
                style="width:40px;"
                type="text"
                v-model="user.code"
                placeholder="请输入"/>
              <span style="width:100px;">
                <img :src="captcha" />
              </span>
              <span style="padding:0 10px;white-space:nowrap;" 
                @click="getCaptcha">看不清</span>
            </div>
          </div>
          <div class="me-panel-item">
            <div class="me-panel-key">新密码</div>
            <div class="me-panel-val">
              <input
                class="v-input"
                type="password"
                v-model="user.password"
                placeholder="6-20位数字或字母或数字和密码组合"
              />
            </div>
          </div>
          <div class="me-panel-item">
            <div class="me-panel-key">确认密码</div>
            <div class="me-panel-val">
              <input
                class="v-input"
                type="text"
                v-model="user.repeat"
                placeholder="请再次输入新密码"
              />
            </div>
          </div>
        </div>
        <div class="me-panel-btn"
          v-if="submit"
          @click="handleSubmit">提交</div>
      </section>
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
  name: 'user-info',
  components: {
  },
  data() {
    return {
      captcha: 'http://127.0.0.1:7000/api/user/captcha',
      listQuery: {
        token: storage.get(CSMSKEY.Token)
      },
      submit: true,
      user: {
        name: '',
        mobile: '',
        password: '',
        repeat: '',
        code: '',
      }
    }
  },
  methods: {
    ...mapActions(['getUser', 'saveUser']),
    getCaptcha() {
      const self = this
      let time = new Date().getTime()
      self.captcha = 'http://127.0.0.1:7000/api/user/captcha?t='+ time
    },
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
    handleSubmit() {
      const self = this
      if(!self.user.mobile || !self.user.password || !self.user.code) {
        self.$toast('请正确输入！','error')
        return
      }
      if(self.user.password !== self.user.repeat) {
        self.$toast('两次输入的密码不相同！','error')
        return
      }
      self.saveUser(self.user).then(response => {
        const res = response.data.data
        const state = response.data.state
        const message = response.data.msg
        if(state) {
          self.$toast(message,'success')
          setTimeout(function(){
            self.$router.push('/login')
          }, 2000);
        } else {
          if(res =='error') {
            self.submit = false
          }
          self.$toast(message,'error')
        }
      });
    },
  },
  created() {
    if(this.listQuery.token) {
      this.get()
    }
  }
}
</script>
<style type="text/css">
@import '../../styles/me-edit.css'
</style>
