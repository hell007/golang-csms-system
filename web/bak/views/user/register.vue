<template>
<div class="page login-page">
  <div class="page-main">
    <div class="page-scroller">
      <div class="login-logo">和云店</div>
      <div class="login-form reg-form">
        <div class="login-form-item">
          <span class="login-form-key">手机号：</span>
          <div class="login-form-val">
            <input
              class="login-form-input"
              type="text"
              placeholder="请输入手机号"
              v-model="form.mobile"
            />
          </div>
        </div>
        <div class="login-form-item">
          <span class="login-form-key">验证码：</span>
          <div class="login-form-val">
            <input
              class="login-form-input-small"
              type="number"
              placeholder="请输入验证码"
              v-model="form.code"
            />
            <div class="login-form-code">
              <span v-if="show" @click="sendMsg">{{state}}</span>
              <span v-if="!show">{{count}}s</span>
            </div>
          </div>
        </div>
        <div class="login-form-item">
          <span class="login-form-key">姓名：</span>
          <div class="login-form-val">
            <input
              class="login-form-input"
              type="text"
              placeholder="请输入真实姓名"
              v-model="form.name"
            />
          </div>
        </div>
        <div class="login-form-item">
          <span class="login-form-key">性别：</span>
          <div class="login-form-val">
            <label class="u-check-radio">
              <input class="u-check-native-control" 
                type="radio" 
                name="sex" 
                checked=""
                v-model="form.gender"
                value="男">
              <i class="u-check-indicator"></i>
              <span>男</span>
            </label>
            <label class="u-check-radio">
              <input class="u-check-native-control" 
                type="radio" 
                name="sex"
                v-model="form.gender"
                value="女">
              <i class="u-check-indicator"></i>
              <span>女</span>
            </label>
          </div>
        </div>
        <div class="login-form-item">
          <span class="login-form-key">密码：</span>
          <div class="login-form-val">
            <input
              class="login-form-input"
              type="password"
              placeholder="请输入密码"
              v-model="form.password"
            />
          </div>  
        </div>
        <div class="login-form-item">
          <span class="login-form-key">确认密码：</span>
          <div class="login-form-val">
            <input
              class="login-form-input"
              type="password"
              placeholder="确认密码"
              v-model="form.repeat"
            />
          </div>
        </div>
        <div class="login-form-tips">
          <label class="u-check-checkbox">
            <input class="u-check-native-control" 
              type="checkbox" 
              checked>
            <i class="u-check-indicator"></i>
          </label>
          <span>注册即视为同意</span>
          <a href="/agree" class="login-form-agree">《和小店注册协议》</a>
        </div>
        <div class="login-form-btn btn-green"
          @click="handleSubmit">注册</div>
        <a href="/login" class="login-form-go"
          style="margin-top:0;">已注册，去登录！</a> 
      </div>
      
    </div>
  </div>
</div>
</template>
<script>
import {
  mapActions
} from 'vuex'

export default {
  name: 'register',
  components: {
  },
  data() {
    return {
      show: true,
      state: '获取验证码',
      count: '',
      timer: null,
      form: {
        mobile: '13888888888',
        name: '范逸尘',
        gender: '男',
        code: '8888',
        password: '123456',
        repeat: '123456',
      }
    }
  },
  methods: {
    ...mapActions(['register']),
    countDown() {
      const TIME_COUNT = 10;
      if (!this.timer) {
        this.count = TIME_COUNT;
        this.show = false;
        this.timer = setInterval(() => {
          if (this.count > 1 && this.count <= TIME_COUNT) {
            this.count--;
          } else {
            this.show = true;
            clearInterval(this.timer);
            this.timer = null;
            this.state = "重新发送"
          }
        }, 1000)
      }
　　},
    sendMsg() {
      const self = this
      const reg = /^1[34578]\d{9}$/
      if(!reg.test(self.form.mobile)) {
        self.$toast('请输入手机号码！','error')
        return
      }
      this.countDown()
    },
    handleSubmit() {
      const self = this
      if(!self.form.mobile || !self.form.name || !self.form.code || !self.form.password) {
        self.$toast('请正确输入！','error')
        return
      }
      if(self.form.password !== self.form.repeat) {
        self.$toast('两次输入的密码不相同！','error')
        return
      }
      self.register(self.form).then(response => {
        const res = response.data.data
        const state = response.data.state
        const message = response.data.msg
        if (state) {
          self.$toast(message,'success')
          setTimeout(function(){
            self.$router.push('/login');
          },3000)
        } else {
          self.$toast(message,'error')
        }
      });
    },
  },
  created() {
    
  }
}
</script>
<style type="text/css">
@import '../../styles/login.css'
</style>
