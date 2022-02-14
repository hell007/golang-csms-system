<!--
 * @Descripttion: 
 * @Author: zenghua.wang
 * @Date: 2021-02-23 21:12:37
 * @LastEditors: zenghua.wang
 * @LastEditTime: 2022-02-14 09:42:49
-->
<template>
  <!-- <page-header :title="title" /> -->

  <div class="page-main">
    <div class="page-scroller">
      <span @click="onAdd">增</span>
      <br />
      <span @click="onDel">删</span>
    </div>
  </div>
</template>
<script lang="ts">
import { defineComponent, onMounted } from 'vue';
import { useStore } from 'vuex';
import { theme } from '@/theme';
import { getApp } from '@/hooks';

export default defineComponent({
  name: 'unit-test',
  components: {},
  setup(_props, _ctx) {
    const primary = theme.primary;
    const title = '模板';
    const store = useStore();
    const app = getApp();

    
    console.log('app===', app);

    // const stateObj: {
    //   phone: string;
    //   time: any;
    // } = {
    //   phone: '13888845678',
    //   time: new Date()
    // };

    //const state = reactive(stateObj);

    const onAdd = async () => {
      let res = await store.dispatch('unit/login', {});
      console.log(50, res);
    };

    
    const onDel = () => {
      app.$http
        .get('/data-report/model/query', { params: {name: 'test'} })
        .then((res: any) => {
          console.log(res);
          // if (res.code === 200) {
          //   this.data = res.data.records
          //   this.page.total = res.data.total || res.data.records.length
          // }
          // this.loading = false
        })
        .catch((err: any) => {
          console.log(err);
          // this.loading = false
          // this.$message.error(err.msg)
        });
    };

    onMounted(() => {
      //console.log(state);
    });

    return {
      primary,
      title,
      onAdd,
      onDel
    };
  }
});
</script>
<style scoped lang="scss">
h2 {
  padding: 10px;
}

.demo {
  margin: 10px 0;
  padding: 10px;
  background-color: $color-gray;
}
</style>
