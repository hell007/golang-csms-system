<template>
  <page-header :title="title" />

  <div class="page-main">
    <div class="page-scroller">
      <section class="address-list">
        <div class="address-item" v-for="(item, index) in list" :key="index">
          <div class="address-item-bd">
            <div class="address-item-name">
              <span>{{ item.consignee }}</span>
              <em>{{ formatPhone(item.phone) }}</em>
            </div>
            <p>{{ item.province }}{{ item.city }}{{ item.district }}{{ item.address }}</p>
          </div>
          <div class="address-item-ft">
            <div class="address-item-radio">
              <van-radio
                v-if="item.state == 1"
                :name="item.state"
                :checked="item.state == 1 ? true : false"
                :checked-color="primary"
              >
                默认地址
              </van-radio>
              <!-- <van-radio v-else :name="item.state" :checked-color="primary">
                设置为默认
              </van-radio> -->
            </div>
            <div class="address-item-opt">
              <span @click="onEdit(item)">
                <van-icon name="edit" size="14" />
                编辑
              </span>
              <span @click="onDel(item)">
                <van-icon name="delete-o" size="14" />
                删除
              </span>
            </div>
          </div>
        </div>
      </section>
      <div class="no-entry" v-if="list.length == 0">暂时没有收货地址</div>
    </div>
  </div>

  <div class="page-footer">
    <div class="address-bar" @click="onAdd">
      <van-button block round :color="primary">新增收货地址</van-button>
    </div>
  </div>
</template>
<script lang="ts">
import { defineComponent, reactive, toRefs, onMounted } from 'vue';
import { useStore } from 'vuex';
import { useRouter } from 'vue-router';
import { getApp } from '@/hooks';
import { theme } from '@/theme';
import { getToken } from '@/utils/storage';
import { fetchGet } from '@/utils/api';
import { URIS } from '@/config';
import { formatPhone } from '@/utils/index';

export default defineComponent({
  name: 'user-address',
  components: {},
  setup(_props, _ctx) {
    const app = getApp();
    const store = useStore();
    const router = useRouter();
    const primary = theme.primary;
    const title = '收货地址管理';

    const stateObj: {
      list: any;
    } = {
      list: []
    };
    const state = reactive(stateObj);

    // 收货地址列表
    const loadData = async () => {
      app.$toast.loading({
        message: '加载中...',
        forbidClick: true,
        overlay: true,
        loadingType: 'spinner'
      });
      let query = { uid: getToken() };
      try {
        let res = await store.dispatch('user/getUserAddressList', query);
        let data = res.data.data;
        const ok = res.data.state;
        const message = res.data.msg;
        if (ok) {
          state.list = data;
        } else {
          app.$toast.fail(message);
        }
        app.$toast.clear();
      } catch (err) {
        console.log(err);
        app.$toast.fail(err);
      }
    };

    const onAdd = () => {
      router.push({
        path: '/user/editAddress'
      });
    };

    const onEdit = (row: any) => {
      router.push({
        path: '/user/editAddress',
        query: { aid: row.id }
      });
    };

    const onDel = (row: any) => {
      fetchGet(URIS.user.delAddress, { uid: row.id })
        .then(function(res) {
          const ok = res.data.state;
          const message = res.data.msg;
          if (ok) {
            app.$toast(message, 'success');
          } else {
            app.$toast(message, 'error');
          }
        })
        .catch((err) => {
          console.log(err);
        });
    };

    onMounted(() => {
      loadData();
    });

    return {
      ...toRefs(state),
      primary,
      title,
      formatPhone,
      onAdd,
      onEdit,
      onDel
    };
  }
});
</script>
<style scoped lang="scss">
.address {
  &-item {
    background-color: $color-white;
    border-radius: 8px;
    padding: 10px;
    margin: 10px;

    &-bd {
      padding-top: 10px;
      border-bottom: 1px solid $color-border-gray;

      p {
        font-size: 14px;
        color: $color-7d;
      }
    }

    &-name {
      color: $color-32;
      margin-bottom: 10px;

      span {
        font-size: 16px;
        margin-right: 10px;
      }

      em {
        font-size: 16px;
        font-style: normal;
        color: $color-primary;
      }
    }

    &-ft {
      @include flex-row();
      align-items: center;
      font-size: 14px;
      padding-top: 10px;
    }

    &-radio {
      flex: 1;
    }

    &-opt {
      flex: 2;
      justify-content: flex-end;
      @include flex-row();
      align-items: center;

      span {
        @include flex-row();
        align-items: center;
        margin-left: 20px;
      }
    }
  }

  &-bar {
    margin: 10px;
  }
}
</style>
