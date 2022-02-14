<template>
  <!-- <page-header :title="title" /> -->

  <div class="page-main">
    <div class="page-scroller">
      <section class="login-form">
        <div class="login-logo">和云店</div>

        <van-field
          v-model="form.mobile"
          required
          clearable
          label="手机号："
          input-align="left"
          placeholder="请输入手机号"
        />

        <!-- <van-field
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
              class="login-confirm"
            >
              <span v-if="modal.msgVisible" class="login-count">
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
        </van-field> -->

        <!-- <van-field
          v-model="form.code"
          required
          clearable
          label="验证码："
          label-width="60px"
          input-align="left"
          placeholder="请输入"
        >
          <template #button>
            <span style="margin-right:10px;">
              <img :src="captcha" style="width:80px;" />
            </span>
            <van-button size="small" round :color="primary" @click="getCaptcha">看不清</van-button>
          </template>
        </van-field> -->

        <van-field
          v-model="form.password"
          required
          clearable
          label="密码："
          input-align="left"
          placeholder="请输入密码"
        />

        <br />
        <slideVerify tips="请按住滑块拖动验证" message="验证通过" @onSucess="onVerifySucess" />
      </section>

      <div class="login-bar">
        <van-button round block :color="primary" @click="onSubmit">登录</van-button>
      </div>

      <div class="login-register" @click="onRegister">未注册，去注册！</div>
    </div>
  </div>
</template>
<script lang="ts">
import { defineComponent, reactive, toRefs, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import { getApp } from '@/hooks';
import { theme } from '@/theme';
import * as utils from '@/utils/index';
import { set } from '@/utils/storage';
import { CSMSKEY, URIS } from '@/config';
import slideVerify from '@/components/slideVerify/index.vue';

export default defineComponent({
  name: 'user-login',
  components: {
    slideVerify
  },
  setup(_props, _ctx) {
    const app = getApp();
    const router = useRouter();
    const primary = theme.primary;
    const title = '登录';

    const stateObj: {
      captcha: string;
      form: any;
      modal: any;
    } = {
      captcha: '',
      form: {
        mobile: '13888888888',
        code: '',
        password: 'admin123'
      },
      modal: {
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

    const getCaptcha = () => {
      let time = new Date().getTime();
      state.captcha = URIS.base + URIS.user.captcha + '?t=' + time;
    };

    const onSubmit = () => {
      if (!utils.isPhone(state.form.mobile)) {
        return app.$toast('请正确输入手机号！');
      }

      if (utils.isEmpty(state.form.password)) {
        return app.$toast('请正确输入必填信息！');
      }

      app.$http.post(URIS.user.login, state.form).then((res:any) => {
        const data = res.data.data;
        const ok = res.data.state;
        const message = res.data.msg;
        if (ok) {
          set(CSMSKEY.token, data);
          router.push('/user/profile');
        } else {
          app.$toast(message);
        }
      });
    };

    const onVerifySucess = () => {
      console.log('通过');
    };

    const onRegister = () => {
      router.push('/user/register');
    };

    onMounted(() => {
      getCaptcha();
    });

    return {
      ...toRefs(state),
      primary,
      title,
      onSend,
      onMsgFinish,
      getCaptcha,
      onVerifySucess,
      onSubmit,
      onRegister
    };
  }
});
</script>
<style scoped lang="scss">
.login {
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
    margin: 40px 10px;
  }

  &-register {
    margin: 0 20px;
    text-align: center;
    color: $color-c8;
  }
}
</style>
