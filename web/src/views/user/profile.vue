<template>
  <page-header :title="title" />

  <div class="page-main">
    <div class="page-scroller">
      <!-- 头像 -->
      <section class="profile-header">
        <div class="profile-header-avatar">
          <img :src="require('../../assets/images/user/vatar.png')" />
        </div>
        <div class="profile-header-text">
          <div class="profile-header-name">
            <span>{{ user.name }}</span>
          </div>
          <p>ID：{{ user.mobile }}</p>
        </div>
      </section>

      <!-- 订单 -->
      <section class="profile-order">
        <div class="profile-order-hd">
          <h6 class="profile-order-title">我的订单</h6>
          <span class="profile-order-all" @click="getOrder">
            <span>全部订单</span>
            <van-icon name="arrow" size="12" />
          </span>
        </div>
        <ul class="profile-order-bd">
          <li
            class="profile-order-link"
            v-for="(item, index) in orderList"
            :key="index"
            @click="getOrder(item)"
          >
            <van-icon class-prefix="xa-icon" :name="item.icon" size="28" :color="primary" />
            <span>{{ item.name }}</span>
            <!-- <em>2</em> -->
          </li>
        </ul>
      </section>

      <!-- 卡片 -->
      <section class="profile-card">
        <van-cell
          :title="item.title"
          is-link
          :value="item.value"
          v-for="(item, index) in cardList"
          :key="index"
          @click="goPage(item)"
        >
          <template #icon>
            <van-icon class-prefix="xa-icon" :name="item.icon" size="20" />
          </template>
        </van-cell>
      </section>

      <div class="profile-out" @click="logout">退出登录</div>
    </div>
  </div>
</template>
<script lang="ts">
import { defineComponent, reactive, toRefs, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import { theme } from '@/theme';
import { getApp } from '@/hooks';
import { fetchGet } from '@/utils/api';
import { getToken, clear } from '@/utils/storage';
import { URIS } from '@/config';

export default defineComponent({
  name: 'user-profile',
  components: {},
  setup(_props, _ctx) {
    const app = getApp();
    const router = useRouter();
    const primary = theme.primary;
    const title = '个人中心';

    const stateObj: {
      user: any;
      orderList: any;
      cardList: any;
    } = {
      user: {},
      orderList: [
        {
          state: 1,
          name: '待付款',
          icon: 'pay-money'
        },
        {
          state: 2,
          name: '待发货',
          icon: 'deliver'
        },
        {
          state: 3,
          name: '待收货',
          icon: 'receive'
        },
        {
          state: 5,
          name: '已完成',
          icon: 'complete'
        }
      ],
      cardList: [
        {
          url: '/user/address',
          icon: 'address',
          title: '收货地址',
          value: '管理我的地址'
        },
        {
          url: '/user/info',
          icon: 'edit',
          title: '修改密码',
          value: '编辑'
        }
      ]
    };
    const state = reactive(stateObj);

    const getUser = async () => {

      fetchGet(URIS.user.profile, {})
        .then(function(res) {
          let data = res.data.data;
          const ok = res.data.state;
          const message = res.data.msg;
          if (ok) {
            state.user = data;
          } else {
            app.$toast(message);
          }
        })
        .catch((err) => {
          console.log(err);
        });
    };

    //订单
    const getOrder = (item: any) => {
      // console.log(item);
      router.push({
        path: '/order',
        query: { uid: state.user.token, state: item.state }
      });
    };

    //卡片
    const goPage = (item: any) => {
      router.push({
        path: item.url
      });
    };

    //退出
    const logout = () => {
      let query = { token: getToken() };
      fetchGet(URIS.user.profile, query)
        .then(function(res) {
          const ok = res.data.state;
          const message = res.data.msg;
          if (ok) {
            app.$toast.success(message);
            clear();
          } else {
            app.$toast.fail(message);
          }
        })
        .catch((err) => {
          console.log(err);
          app.$toast.fail(err);
        });
    };

    onMounted(() => {
      getUser();
    });

    return {
      ...toRefs(state),
      primary,
      title,
      getOrder,
      logout,
      goPage
    };
  }
});
</script>
<style scoped lang="scss">
.profile {
  &-header {
    @include flex-row();
    align-items: center;
    padding: 10px;

    &-avatar {
      width: 90px;
      height: 90px;
      overflow: hidden;
      margin-right: 20px;
      border-radius: 999px;
      border: 1px solid $color-primary;

      & img {
        border-radius: 999px;
      }
    }

    &-text {
      flex: 4;
      overflow: hidden;
      font-size: 12px;
      color: $color-7d;

      & p {
        margin-bottom: 0;
      }
    }

    &-name {
      font-size: 16px;
      color: $color-32;
    }
  }

  &-order {
    padding: 10px;
    margin: 10px;
    border-radius: 10px;
    background-color: $color-white;
    font-size: 14px;

    &-hd {
      @include flex-row();
      align-items: center;
      padding-bottom: 8px;
      border-bottom: 1px solid #e9e9e9;
    }

    &-title {
      flex: 3;
      font-size: 16px;
      color: $color-32;
      margin-bottom: 0;
    }

    &-all {
      flex: 1;
      @include flex-row();
      align-items: center;
      color: $color-7d;
      justify-content: flex-end;
    }

    &-bd {
      @include flex-row();
      align-items: center;
    }

    &-link {
      flex: 1;
      text-align: center;
      position: relative;
      padding: 10px 0;

      & span {
        display: block;
        color: $color-primary;
      }

      & em {
        position: absolute;
        right: 10px;
        top: 10px;
        display: block;
        width: 1rem;
        height: 1rem;
        border-radius: 999px;
        background-color: $color-primary;
        color: $color-white;
        font-size: 12px;
        font-style: normal;
      }
    }
  }

  &-card {
    padding: 5px 0;
    margin: 10px;
    border-radius: 10px;
    background-color: $color-white;
  }

  &-out {
    position: absolute;
    left: 0;
    bottom: 2rem;
    right: 0;
    padding: 10px;
    margin: 10px;
    border-radius: 6px;
    text-align: center;
    background-color: $color-primary;
    font-size: 16px;
    color: $color-white;
  }
}
</style>
