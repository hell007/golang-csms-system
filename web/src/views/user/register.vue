<template>
  <!-- <page-header :title="title" /> -->

  <div class="page-main">
    <div class="page-scroller">
      <section class="register-form">
        <div class="register-logo">和云店</div>

        <van-field
          v-model="form.mobile"
          required
          clearable
          type="tel"
          maxlength="11"
          label="手机号："
          input-align="left"
          placeholder="请输入手机号"
        />

        <van-field
          v-model="form.code"
          required
          clearable
          label="验证码："
          input-align="left"
          placeholder="请输入验证码"
        >
          <template #button>
            <van-button
              size="small"
              round
              :color="primary"
              :disabled="modal.msgDisabled"
              @click="onSend"
              class="register-confirm"
            >
              <span v-if="modal.msgVisible" class="register-count">
                <van-count-down
                  millisecond
                  :time="30000"
                  :auto-start="true"
                  format="ss"
                  @finish="onMsgFinish"
                />
                秒
              </span>
              <span v-else>发送验证码</span>
            </van-button>
          </template>
        </van-field>

        <van-field
          v-model="form.name"
          required
          clearable
          label="姓名："
          input-align="left"
          placeholder="请输入真实姓名"
        />

        <van-field label="性别：" input-align="left" placeholder="请输入真实姓名">
          <template #input>
            <van-radio-group v-model="form.gender" direction="horizontal" class="register-radio">
              <van-radio name="1" :checked-color="primary">男</van-radio>
              <van-radio name="2" :checked-color="primary">女</van-radio>
            </van-radio-group>
          </template>
        </van-field>

        <van-field
          v-model="form.password"
          required
          clearable
          maxlength="32"
          type="password"
          label="密码："
          input-align="left"
          placeholder="请输入密码"
        />

        <van-field
          v-model="form.repeat"
          required
          clearable
          maxlength="32"
          type="password"
          label="确认密码："
          input-align="left"
          placeholder="请再次输入密码"
        />
      </section>

      <div class="register-form-tips">
        <van-checkbox
          v-model="form.checked"
          :checked-color="primary"
          icon-size="14px"
          shape="square"
        >
          注册即视为同意
        </van-checkbox>
        <span class="agree" @click="onAgree">《和小店注册协议》</span>
      </div>

      <div class="register-bar">
        <van-button round block :color="primary" @click="onSubmit">注册</van-button>
      </div>

      <div class="register-login" @click="onLogin">已注册，去登录！</div>
    </div>
  </div>

  <!-- dialog -->
  <page-dialog type="dialog" v-model:show="modal.show" :lock="true">
    <template #body>
      <section class="agree-card">
        <h6>一、购买流程</h6>
        <p>
          亲爱的客户，很抱歉您的商品出现这样的问题。请准备好商品全套（商品、附件、发票原件、包装）按以下地址寄往和小店售后部，不支持到付及平邮。如果收到商品后情况属实，我们会尽快为您办理。感谢您对我们的支持。
        </p>
      </section>
    </template>
    <template #footer>
      <van-button
        round
        :disabled="modal.agreeDisabled"
        :color="primary"
        @click="onConfirm"
        class="register-confirm"
      >
        <span v-if="modal.agreeVisible" class="register-count">
          <van-count-down
            millisecond
            :time="10000"
            :auto-start="true"
            format="ss"
            @finish="onAgreeFinish"
          />
          秒
        </span>
        <span v-else>已阅读协议</span>
      </van-button>
    </template>
  </page-dialog>
</template>
<script lang="ts">
import { defineComponent, reactive, toRefs, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import { getApp } from '@/hooks';
import { theme } from '@/theme';
import { fetchPost } from '@/utils/api';
import * as utils from '@/utils/index';
import { URIS } from '@/config';

export default defineComponent({
  name: 'user-register',
  components: {},
  setup(_props, _ctx) {
    const app = getApp();
    const router = useRouter();
    const primary = theme.primary;
    const title = '注册';

    const stateObj: {
      form: any;
      modal: any;
    } = {
      form: {
        mobile: '',
        name: '',
        gender: '1',
        code: '',
        password: '',
        repeat: '',
        checked: false
      },
      modal: {
        show: false,
        agreeVisible: true,
        agreeDisabled: true,
        msgVisible: false,
        msgDisabled: false
      }
    };

    const state = reactive(stateObj);

    //短信验证
    const onSend = () => {
      state.modal.msgDisabled = true;
      state.modal.msgVisible = true;
      app.$toast('短信验证码已发送！');
    };

    //短信倒计时
    const onMsgFinish = () => {
      state.modal.msgDisabled = false;
      state.modal.msgVisible = false;
    };

    //注册
    const onSubmit = () => {
      if (!utils.isPhone(state.form.mobile)) {
        return app.$toast('请正确输入手机号！');
      }

      if (
        utils.isEmpty(state.form.code) ||
        utils.isEmpty(state.form.name) ||
        utils.isEmpty(state.form.password)
      ) {
        return app.$toast('请正确输入必填信息！');
      }

      if (state.form.password !== state.form.repeat) {
        return app.$toast('两次输入的密码不相同！');
      }

      if (!state.form.checked) {
        return app.$toast('请勾选协议！');
      }

      fetchPost(URIS.user.register, state.form).then((res) => {
        const ok = res.data.state;
        const message = res.data.msg;
        if (ok) {
          app.$toast(message);
          setTimeout(function() {
            router.push('/user/login');
          }, 3000);
        } else {
          app.$toast(message);
        }
      });
    };

    //协议
    const onAgree = () => {
      state.modal.show = true;
    };

    //协议倒计时
    const onAgreeFinish = () => {
      state.modal.agreeDisabled = false;
      state.modal.agreeVisible = false;
    };

    //确认协议
    const onConfirm = () => {
      state.modal.show = false;
    };

    //去登录
    const onLogin = () => {
      router.push('/user/login');
    };

    onMounted(() => {});

    return {
      ...toRefs(state),
      primary,
      title,
      onSend,
      onMsgFinish,
      onSubmit,
      onConfirm,
      onAgree,
      onAgreeFinish,
      onLogin
    };
  }
});
</script>
<style scoped lang="scss">
.register {
  &-logo {
    margin: 40px auto;
    font-size: 40px;
    font-weight: bold;
    text-align: center;
  }

  &-form {
    margin: 10px;
    padding: 10px;
    border-radius: 10px;
    box-shadow: 2px 2px 4px rgba(0, 0, 0, 0.1);
    background: $color-white;

    &-tips {
      margin: 10px;
      @include flex-row();
      align-items: center;
      padding: 20px;
      font-size: 12px;
      color: $color-7d;

      /deep/ .van-checkbox__label {
        color: $color-96;
      }

      .agree {
        color: $color-red;
      }
    }
  }

  &-radio {
    @include flex-row();
    align-items: center;

    /deep/ .van-radio {
      margin-right: 15px;
    }
  }

  &-confirm {
    min-width: 100px;
  }

  &-count {
    @include flex-row();
    align-items: center;
    color: $color-white;

    /deep/ .van-count-down {
      color: $color-white;
    }
  }

  &-bar {
    margin: 10px;
  }

  &-login {
    margin: 30px 20px;
    text-align: center;
    color: $color-c8;
  }
}

.agree {
  &-card {
    padding: 10px;

    p {
      font-size: 12px;
      color: $color-32;
      margin-bottom: 10px;
    }
  }
}
</style>
