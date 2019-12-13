<template>
<div class="main-screen">
  <div class="main-screen__main">
    <div class="main-screen__inner">
      <div class="top-bar container">
        <!-- <div class="top-bar__left"> -->
        <a href="javascript:;" class="logo">
          system
        </a>
        <!-- </div> -->
        <!-- <div class="top-bar__right">
          <div class="top-bar__btn-wrap">
            <a href="javascript:;" class="top-bar__btn">商户入驻</a>
          </div>
        </div> -->
      </div>
      <div class="platform container">
        <div class="platform__info">
          <!-- <div class="platform__info-main">
            <div class="platform__name">全渠道运营中心</div>
            <div>共享千万用户流量，线上线下协同销售<br>营销玩法丰富多彩，把握商机精密经营</div>
          </div> -->
        </div>

        <!-- 登录密码 -->
        <div v-bind:class="login">
          <div class="login__type">登录</div>
          <el-form label-width="0px"
            ref="postForm"
            :rules="formRules"
            :model="form">
            <el-tabs v-model="logintype">
              <el-tab-pane label="手机登录" name="login-1">
                <div class="login__form" v-if="logintype=='login-1'">
                  <div class="login__item" prop="mobile">
                    <fieldset class="login__fieldset">
                      <legend>手机号</legend>
                      <label class="login__block">
                        <span class="login__icon">
                          <i class="xa-icon xa-icon-phone"></i>
                        </span>
                        <input
                          v-model="form.mobile"
                          class="login__input"
                          type="text"
                          placeholder="输入手机号"
                          autofocus
                        />
                      </label>
                    </fieldset>
                  </div>
                  <div class="login__item login__item--verify" prop="code">
                    <fieldset class="login__fieldset" style="width:191px;">
                      <legend>验证码</legend>
                      <label class="login__block">
                        <span class="login__icon">
                          <i class="xa-icon xa-icon-verify"></i>
                        </span>
                        <input
                          v-model="form.code"
                          class="login__input"
                          type="text"
                          placeholder="输入验证码"
                        />
                      </label>
                    </fieldset>
                    <button type="button" class="login__btn-verify"
                      @click="handleSendCode">
                      获取验证码
                    </button>
                  </div>
                  <div class="login__actions">
                    <div class="login__action-bar">
                      <button type="button" class="login__action"
                        @click="handleLogin">登录</button>
                    </div>
                  </div>
                </div>
              </el-tab-pane>
              <el-tab-pane label="账号密码登录" name="login-2">
                <div class="login__form" v-if="logintype=='login-2'">
                  <div class="login__item" prop="username">
                    <fieldset class="login__fieldset">
                      <legend>账号</legend>
                      <label class="login__block">
                        <span class="login__icon">
                          <i class="xa-icon xa-icon-account"></i>
                        </span>
                        <input
                          v-model="form.username"
                          class="login__input"
                          type="text"
                          placeholder="输入账号"
                          autofocus
                        />
                      </label>
                    </fieldset>
                  </div>
                  <div class="login__item" prop="password">
                    <fieldset class="login__fieldset">
                      <legend>密码</legend>
                      <label class="login__block">
                        <span class="login__icon">
                          <i class="xa-icon xa-icon-password"></i>
                        </span>
                        <input
                          v-model="form.password"
                          class="login__input"
                          type="password"
                          placeholder="输入密码"
                        />
                      </label>
                    </fieldset>
                  </div>
                  <div class="login__actions">
                    <div class="login__action-bar">
                      <button type="button" class="login__action"
                        @click="handleLogin">登录</button>
                    </div>
                    <div class="login__tip">
                      <a href="javascript:;"
                        @click="handleGoRecovery">忘记密码？</a>
                    </div>
                  </div>
                </div>
              </el-tab-pane>
            </el-tabs>
          </el-form>
        </div>

        <!-- 重设密码 -->
        <div v-bind:class="recovery">
          <div class="login__type">重设密码</div>
            <div class="login__item">
              <fieldset class="login__fieldset">
                <legend>手机号</legend>
                <label class="login__block">
                  <span class="login__icon">
                    <i class="xa-icon xa-icon-phone"></i>
                  </span>
                  <input
                    class="login__input"
                    type="text"
                    placeholder="请输入注册时的手机号"
                    required
                  />
                </label>
              </fieldset>
            </div>
            <div class="login__item login__item--verify">
              <fieldset class="login__fieldset">
                <legend>验证码</legend>
                <label class="login__block">
                  <span class="login__icon">
                    <i class="xa-icon xa-icon-verify"></i>
                  </span>
                  <input
                    class="login__input"
                    type="text"
                    placeholder="输入验证码"
                    required
                  />
                </label>
              </fieldset>
              <button type="button" class="login__btn-verify">
                获取验证码
              </button>
            </div>

            <div class="login__item">
              <fieldset class="login__fieldset">
                <legend>密码</legend>
                <label class="login__block">
                  <span class="login__icon">
                    <i class="xa-icon xa-icon-password"></i>
                  </span>
                  <input
                    class="login__input"
                    type="password"
                    placeholder="请输入新密码，6-10位字母和数字"
                    required
                  />
                </label>
              </fieldset>
            </div>

            <div class="login__item">
              <fieldset class="login__fieldset">
                <legend>密码</legend>
                <label class="login__block">
                  <span class="login__icon">
                    <i class="xa-icon xa-icon-password"></i>
                  </span>
                  <input
                    class="login__input"
                    type="password"
                    placeholder="请再次输入新密码"
                    required
                  />
                </label>
              </fieldset>
            </div>

            <div class="login__actions">
              <div class="login__action-bar">
                <button type="submit" class="login__action">
                  确认修改
                </button>
              </div>
              <div class="login__tip">
                <a href="javascript:;"
                  @click="handleGoLogin">去登录</a>
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
  mapState,
  mapMutations,
  mapActions,
  mapGetters
} from 'vuex'

import * as validate from '@/utils/validate'

export default {
  components: {},
  name: 'login',
  data() {
    return {
      login: ['login'],
      recovery:['login','login--recovery','login--hide'],
      logintype: 'login-2',
      form: {
        phone: '',
        code: '',
        username: '曹操',
        password: '123456',
      },
      //验证规则
      formRules: {
        mobile: [{
          required: true,
          message: "请输入手机号",
          trigger: 'blur'
        }, {
          validator: validate.validateMobile,
          trigger: 'blur'
        }],
        code: [{
          required: true,
          message: "请输入验证码",
          trigger: 'blur'
        }, {
          min: 6,
          max: 6,
          message: "验证码长度为6位",
          trigger: 'blur'
        }],
        name: [{
          required: true,
          message: "请输入账号",
          trigger: 'blur'
        }, {
          min: 2,
          max: 20,
          message: "必须输入2-20个字符",
          trigger: 'blur'
        }],
        password: [{
          required: true,
          message: "请输入密码",
          trigger: 'blur'
        }]
      }
    }
  },
  methods: {
    ...mapActions(['LoginByName']),
    handleSendCode() {
      console.log('handleSendCode')
    },
    handleLogin() {
      const self = this
      this.$refs.postForm.validate(valid => {
        if (valid) {
          //dispatch 调用store里面的action,相当于 ...mapActions(['L...'])
          //this.$store.dispatch('LoginByName', this.form).then((res) => {
          self.LoginByName(this.form).then((res) => {
            const state = res.data.state
            const msg = res.data.msg
            if (state) {
              this.$router.push('/');
            } else {
              self.$notify({
                title: '登录失败',
                message: msg,
                type: 'error'
              })
            }
          }).catch(ex => {
            self.$notify({
                title: '登录失败',
                message: ex,
                type: 'error'
              })
          })
        } else {
          self.$alert('请正确输入，并登录！', '提示', {
            confirmButtonText: '确定'
          });
        }
      })
    },
    handleGoLogin() {
      this.recovery.push('login--hide');
      this.login.pop('login--hide');
    },
    handleGoRecovery() {
      this.login.push('login--hide');
      this.recovery.pop('login--hide');
    }  
  },
  created() {
    
  }
}
</script>
<style rel="stylesheet/scss" lang="scss">
@import '../../styles/xa-base.scss';
@import '../../styles/login.scss';
</style>
