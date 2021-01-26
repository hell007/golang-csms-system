<template>
  <!-- <page-header :title="title" /> -->

  <div class="page-main">
    <div class="page-scroller"></div>
  </div>
</template>
<script lang="ts">
import { defineComponent, onMounted } from 'vue';
import { theme } from '@/theme';

export default defineComponent({
  name: 'unit-test',
  components: {},
  setup(_props, _ctx) {
    const primary = theme.primary;
    const title = '模板';

    // const stateObj: {
    //   phone: string;
    //   time: any;
    // } = {
    //   phone: '13888845678',
    //   time: new Date()
    // };

    //const state = reactive(stateObj);

    const onTest = () => {
      console.log('成功');
    };

    onMounted(() => {
      //console.log(state);
    });

    return {
      primary,
      title,
      onTest
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
