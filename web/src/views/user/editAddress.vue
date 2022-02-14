<template>
  <page-header :title="title" />

  <div class="page-main">
    <div class="page-scroller">
      <section class="address-form">
        <van-field
          v-model="form.consignee"
          required
          clearable
          label="收货人"
          input-align="right"
          placeholder="请输入"
        />

        <van-field
          v-model="form.phone"
          required
          clearable
          label="手机号"
          input-align="right"
          placeholder="请填写收货人手机号"
        />

        <van-field
          v-model="form.zone"
          required
          readonly
          is-link
          label="所在地区"
          input-align="right"
          placeholder="请选择"
          @click="onCity"
        />

        <van-field
          v-model="form.address"
          required
          clearable
          label="详细地址"
          input-align="right"
          placeholder="街道、楼牌号等"
        />

        <van-field
          v-model="form.tag"
          label="标签"
          input-align="right"
        >
          <template #input>
            <van-tag
              round
              size="large"
              :color="primary"
            >家</van-tag>
            <van-tag
              round
              size="large"
            >学校</van-tag>
            <van-tag
              round
              size="large"
            >公司</van-tag>
          </template>
        </van-field>

        <van-field
          v-model="form.state"
          label="默认地址"
          input-align="right"
        >
          <template #input>
            <van-switch
              v-model="form.checked"
              :size="20"
              :active-color="primary"
              inactive-color="#dcdee0"
            />
          </template>
        </van-field>
        <p class="address-tips">
          提醒：每次下单时会使用该地址，实际下单地址需要您在提交订单时确认！
        </p>
      </section>
    </div>
  </div>

  <!-- 城市 -->
  <van-popup
    v-model:show="pop.show"
    round
    position="bottom"
  >
    <van-cascader
      v-model="pop.val"
      title="请选择所在地区"
      :options="pop.list"
      :active-color="primary"
      @close="pop.show = false"
      @change="onChange"
      @finish="onFinish"
    />
  </van-popup>

  <div class="page-footer">
    <div class="address-bar">
      <van-button
        block
        round
        :color="primary"
      >
        保存
      </van-button>
    </div>
  </div>
</template>
<script lang="ts">
import { defineComponent, reactive, toRefs, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import { useStore } from 'vuex';
import { theme } from '@/theme';
import { getApp } from '@/hooks';
import { URIS } from '@/config';
import { isEmpty } from '@/utils/index';

export default defineComponent({
  name: 'user-edit-address',
  components: {},
  setup(_props, _ctx) {
    const app = getApp();
    const store = useStore();
    const router = useRouter();
    const primary = theme.primary;
    const title = '收货地址编辑';

    const stateObj: {
      form: any;
      pop: any;
    } = {
      form: {
        consignee: '',
        phone: '',
        province: '',
        city: '',
        district: '',
        address: '',
        state: 2,
        zone: '',
        tag: '家',
        checked: false
      },
      pop: {
        show: false,
        finished: false,
        level: 3,
        list: []
      }
    };
    const state = reactive(stateObj);

    // 从store中获取单个收货地址
    const getAddress = async () => {
      app.$toast.loading({
        message: '加载中...',
        forbidClick: true,
        overlay: true,
        loadingType: 'spinner'
      });
      let query = {
        aid: app.$route.query.aid || null
      };
      try {
        let res = await store.dispatch('user/getUserAddress', query);
        if (isEmpty(res)) {
          state.form = Object.assign(state.form, res);
          state.form.zone = state.form.province + state.form.city + state.form.district;
          state.form.checked = state.form.state == 1 ? true : false;
        }
      } catch (err) {
        console.log(err);
      }
    };

    // 城市级联组件的使用发现问题：
    //城市级别需要自己设置；
    //城市完成一次选择需要自己设置；
    const onCity = () => {
      state.pop.show = true;
    };

    const getCity = async (item: any) => {
      let aid = item?.value || 0;
      let index = item?.tabIndex;
      let selected = item?.selectedOptions;
      //console.log(aid, index, selected);
      let query = {
        aid: aid
      };
      try {
        let res = await store.dispatch('user/getCityList', query);
        const ok = res.data.state;
        const message = res.data.msg;
        if (ok) {
          if (aid == 0) {
            let list = res.data.data;
            for (const item of list) {
              item.children = [];
            }
            state.pop.list = list;
          } else {
            selected[index].children = res.data.data;
          }
        } else {
          app.$toast(message);
        }
      } catch (err) {
        console.log(err);
      }
    };

    const onChange = (item: any) => {
      let index = item?.tabIndex;
      //判断finished依据
      state.pop.finished = false;
      if (index == state.pop.level - 1) {
        state.pop.finished = true;
        return;
      }
      getCity(item);
    };

    const onFinish = (item: any) => {
      if (state.pop.finished) {
        state.pop.show = false;
        state.form.zone = item.selectedOptions.map((it: any) => it.text).join('/');
        state.form.province = item?.selectedOptions[0].value;
        state.form.city = item?.selectedOptions[1].value;
        state.form.district = item?.selectedOptions[2].value;
      }
    };

    const onSave = () => {
      let form = {};

      app.$http
        .post(URIS.user.saveAddress, form)
        .then((res: any) => {
          const ok = res.data.state;
          const message = res.data.msg;
          if (ok) {
            app.$toast.success(message);
            router.push({
              path: '/user/address'
            });
          } else {
            app.$toast.fail(message);
          }
        })
        .catch((err: any) => {
          console.log(err);
          app.$toast.fail(err);
        });
    };

    onMounted(() => {
      getAddress();
      getCity(null);
    });

    return {
      ...toRefs(state),
      primary,
      title,
      onCity,
      onChange,
      onFinish,
      onSave
    };
  }
});
</script>
<style scoped lang="scss">
.address {
  &-form {
    background-color: $color-white;
    border-radius: 10px;
    padding: 10px;
    margin: 10px;

    /deep/ .van-tag {
      min-width: 50px;
      margin-left: 10px;
      text-align: center;
      justify-content: center;
    }
  }

  &-tips {
    padding: 12px 12px 0;
    font-size: 12px;
    color: $color-7d;
  }

  &-bar {
    margin: 10px;
  }
}
</style>
