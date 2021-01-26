<template>
  <!-- <page-header :title="title" /> -->

  <div class="page-main">
    <slideVerify />

    <div class="page-scroller" style="display:none">
      <h2>utils</h2>
      <section class="demo">
        <code>
          <b>引入</b>
          <p>import * as utils from '@/utils/index';</p>

          <b>使用</b>
          <p>
            utils.isEmpty() ....
          </p>
        </code>
      </section>

      <h2><a href="https://www.lodashjs.com/">lodash 一个工具类</a></h2>
      <section class="demo">
        <code>
          <b>引入</b>
          <p>import {lodash} from '@/utils/index';</p>

          <b>使用</b>
          <p>
            lodash.isEmpty() ....
          </p>
        </code>
      </section>
    </div>
  </div>
</template>
<script lang="ts">
import { defineComponent, reactive, onMounted } from 'vue';
import { theme } from '@/theme';
import * as utils from '@/utils/index';

export default defineComponent({
  name: 'unit-utils',
  components: {},
  setup(_props, _ctx) {
    const primary = theme.primary;
    const title = '模板';

    const stateObj: {
      phone: string;
      time: any;
    } = {
      phone: '13888845678',
      time: new Date()
    };

    const state = reactive(stateObj);

    onMounted(() => {
      //空判断
      console.log('----空判断----');
      console.log('isEmpty==', utils.isEmpty(state.phone));

      //时间
      console.log('----时间----');
      console.log('年==', utils.getDate('YYYY'));
      console.log('月==', utils.getDate('MM'));
      console.log('日==', utils.getDate('DD'));
      console.log('formatDate==', utils.formatDate(state.time, 'YYYY-MM-DD HH:mm:ss'));
      console.log('3天前==', utils.lastDate(3, 'days', 'YYYY-MM-DD'));
      console.log('去年==', utils.lastDate(1, 'years', 'YYYY-MM-DD'));
      console.log('上月==', utils.lastDate(3, 'months', 'YYYY-MM-DD'));
      console.log('上周==', utils.lastDate(1, 'weeks', 'YYYY-MM-DD'));

      //手机号码
      console.log('----手机号码----');
      console.log('isPhone==', utils.isPhone(state.phone));
      console.log('formatPhone==', utils.formatPhone(state.phone));
      console.log('privatePhone==', utils.privatePhone(state.phone));
    });

    return {
      primary,
      title
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
