<template>
<div class="page login-page">
  <div class="page-main">
    <div class="page-scroller">
      <div class="login-logo">和云店</div>
      <div class="login-form">
        <div class="login-form-item">
          <input
            class="login-form-input"
            type="text"
            placeholder="请输入手机号"
            v-model="form.mobile"
          />
        </div>
        <div class="login-form-item">
          <input
            class="login-form-input"
            type="text"
            placeholder="请输入密码"
            v-model="form.password"
          />
        </div>
        <div class="login-form-item">
          <input
            class="login-form-input-small"
            type="text"
            placeholder="请输入验证码"
            v-model="form.code"
          />
          <div class="login-form-code">
            <img :src="captcha" />
          </div>
          <span class="login-form-switch"
            @click="getCaptcha">看不清</span>
        </div>
        <br />
        <div class="login-form-btn btn-primary"
          v-if="submit"
          @click="handleSubmit">登录</div>
        <a href="/register" class="login-form-go">还没账号？去注册！</a>  
      </div>
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
  name: 'login',
  components: {
  },
  data() {
    return {
      captcha: 'http://127.0.0.1:7000/api/user/captcha',
      submit: true,
      times: 1,
      form: {
        mobile: '13888888888',
        password: '123456',
        code: '',
      }
    }
  },
  methods: {
    ...mapActions(['login']),
    getCaptcha() {
      const self = this
      let time = new Date().getTime()
      self.captcha = 'http://127.0.0.1:7000/api/user/captcha?t='+ time
    },
    handleHide() {
      let times = storage.get('login-times')
      if(times >= 5) {
        this.submit = false
        return
      }
    },
    handleSubmit() {
      const self = this
      if(!self.form.mobile || !self.form.password || !self.form.code) {
        self.$toast('请正确输入！','error')
        return
      }
      self.login(self.form).then(response => {
        const res = response.data.data
        const state = response.data.state
        const message = response.data.msg
        if(state) {
          storage.set(CSMSKEY.Token, res)
          //self.$router.push('/profile')
          window.history.back(-1)
        } else {
          self.times++
          self.handleHide()
          storage.set('login-times', self.times)
          self.$toast(message,'error')
        }
      });
    },
  },
  created() {
    this.handleHide()
  }
}
</script>
<style type="text/css">
@import '../../styles/login.css'
</style>
