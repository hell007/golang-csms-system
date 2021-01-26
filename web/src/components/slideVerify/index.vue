<template>
  <div ref="vPanel" class="verify-panel">
    <div ref="vSlide" class="verify-slide" />
    <div ref="vTips" class="verify-tips">{{ tips }}</div>
    <div ref="vBar" class="verify-bar">
      <span v-if="isSuccess" class="icon icon-checked" />
      <span v-else class="icon icon-bar" />
    </div>
    <div ref="vText" class="verify-message">
      <div v-if="loading" class="verify-loading">
        <span />
        <span />
        <span />
        <span />
        <span />
      </div>
      <span v-else>{{ message }}</span>
    </div>
  </div>
</template>
<script lang="ts">
import { defineComponent, reactive, ref, toRefs, watch, onMounted, onUnmounted } from 'vue';
import { theme } from '@/theme';

export default defineComponent({
  name: 'slide-verify',
  props: {
    tips: {
      type: String,
      default: '请按住滑块，拖动到最右边'
    },
    message: {
      type: String,
      default: '验证通过'
    },
    onSucess: {
      type: Function,
      default: null
    }
  },
  setup(_props, ctx) {
    const primary = theme.primary;

    //dom获取
    const vPanel = ref<any>(null);
    const vSlide = ref<any>(null);
    const vBar = ref<any>(null);
    const vTips = ref<any>(null);
    const vText = ref<any>(null);

    const stateObj: {
      x1: number;
      x2: number;
      max: number;
      isMousedown: boolean;
      loading: boolean;
      isSuccess: boolean;
    } = {
      x1: 0,
      x2: 0,
      max: 320 - 44 - 10,
      isMousedown: false,
      loading: false,
      isSuccess: false
    };

    let state = reactive(stateObj);

    //鼠标按下
    const mousedown = (e: any) => {
      if (state.isSuccess) {
        return;
      }
      state.x1 = e.x || e.targetTouches[0].clientX;
      state.isMousedown = true;
    };

    //鼠标移动
    const mousemove = (e: any) => {
      if (!state.isMousedown || state.isSuccess) {
        return;
      }
      e.preventDefault();
      e.stopPropagation();
      state.x2 = e.x || e.targetTouches[0].clientX;
      let diff = state.x2 - state.x1;
      if (diff < 0) {
        diff = 0;
      }
      if (diff >= state.max - 80 && diff < state.max) {
        state.loading = true;
        vTips.value.style.opacity = '0';
        vText.value.style.opacity = '1';
      }
      if (diff >= state.max) {
        diff = state.max;
        state.loading = false;
        state.isSuccess = true;
        ctx.emit('onSucess');
      }
      const transitionDuration = '0s';
      const transform = `translateX(${diff}px)`;
      setStyle(transitionDuration, transform);
    };

    //鼠标离开
    const mouseleave = () => {
      if (state.isSuccess || state.isMousedown) {
        return;
      }
      vSlide.value.style.opacity = '0';
      const transitionDuration = '.3s';
      const transform = `translateX(0)`;
      setStyle(transitionDuration, transform);
    };

    //鼠标松起
    const mouseup = () => {
      if (state.isSuccess) {
        return;
      }
      state.isMousedown = false;
      const transitionDuration = '.3s';
      const transform = `translateX(0)`;
      setStyle(transitionDuration, transform);
    };

    const setStyle = (transitionDuration: string, transform: string) => {
      // 滑块
      vSlide.value.style.transitionDuration = transitionDuration;
      vSlide.value.style.webkitTransitionDuration = transitionDuration;
      vSlide.value.style.transform = transform;
      vSlide.value.style.webkitTransform = transform;
      // 按钮
      vBar.value.style.transitionDuration = transitionDuration;
      vBar.value.style.webkitTransitionDuration = transitionDuration;
      vBar.value.style.transform = transform;
      vBar.value.style.webkitTransform = transform;
    };

    //重置
    const onReset = () => {
      state.isSuccess = false;
      state.isMousedown = false;
      const transitionDuration = '0s';
      const transform = `translateX(0)`;
      setStyle(transitionDuration, transform);
      vTips.value.style.opacity = '1';
      vText.value.style.opacity = '0';
      vSlide.value.style.opacity = '0';
    };

    //监听
    watch(
      () => state.isMousedown,
      () => {
        state.loading = false;
        vTips.value.style.opacity = '1';
        vText.value.style.opacity = '0';
      }
    );

    //初始化
    const init = () => {
      const slideW = vSlide.value.clientWidth;
      const barW = vBar.value.clientWidth;
      state.max = slideW - barW;
    };

    onMounted(() => {
      init();

      vBar.value.addEventListener('mouseleave', mouseleave);
      vBar.value.addEventListener('touchend', mouseleave, { passive: true });
      vBar.value.addEventListener('touchstart', mousedown, { passive: true });
      vBar.value.addEventListener('touchstart', mousedown, { passive: true });

      document.body.addEventListener('mousemove', mousemove, {
        passive: false
      });
      document.body.addEventListener('touchmove', mousemove, {
        passive: false
      });
      document.body.addEventListener('mouseup', mouseup);
      document.body.addEventListener('touchend', mouseup, { passive: true });
      vPanel.value.addEventListener('reset', onReset);
    });

    //卸载
    onUnmounted(() => {
      document.body.removeEventListener('mousemove', mousemove);
      document.body.removeEventListener('touchmove', mousemove);
      document.body.removeEventListener('mouseup', mouseup);
      document.body.removeEventListener('touchend', mouseup);
    });

    return {
      ...toRefs(state),
      primary,
      vPanel,
      vSlide,
      vBar,
      vTips,
      vText
    };
  }
});
</script>
<style scoped lang="scss">
.verify {
  &-panel {
    overflow: hidden;
    position: relative;
    max-width: 100%;
    min-height: 44px;
    margin: 10px;
    box-sizing: border-box;
    -webkit-user-select: none;
    border: 1px solid $color-border-gray;
    background-color: $color-gray;
    border-radius: 2px;
    background: linear-gradient(313deg, rgba(65, 209, 102, 1) 0%, rgba(90, 232, 118, 1) 100%);
    //background: $color-primary;

    &::after {
      position: absolute;
      bottom: -100%;
      left: -100%;
      z-index: 2;
      content: '';
      display: block;
      width: 100%;
      height: 100%;
      transition: all 0.2s;
      background: rgba(255, 255, 255, 0.3);
      transform: rotate(60deg);
      animation: wave 1.5s infinite;
      animation-delay: 1s;
    }
  }

  &-slide {
    position: absolute;
    left: 0px;
    top: 0;
    right: 0;
    bottom: 0;
    z-index: 2;
    width: 100%;
    height: 100%;
    transition: all 0.1s linear, transform 0.3s ease;
    background: #f7f7f7;
  }

  &-tips {
    position: absolute;
    left: 0px;
    top: 0;
    right: 0;
    bottom: 0;
    z-index: 3;
    width: 100%;
    height: 100%;
    @include flex-row();
    justify-content: center;
    align-items: center;
    color: $color-7d;
    font-size: 14px;
    opacity: 1;
    transition: opacity 0.1s linear;
    pointer-events: none;
  }

  &-bar {
    position: absolute;
    left: -1px;
    top: -1px;
    z-index: 4;
    @include flex-row();
    justify-content: center;
    align-items: center;
    width: 44px;
    height: 100%;
    cursor: pointer;
    transition: transform 0.3s ease;

    .icon {
      width: 100%;
      height: 100%;
      background-position: center;
      background-color: #fff;
    }

    .icon-bar {
      background-size: 100%;
      background-image: url('./images/bar.png');
    }

    .icon-checked {
      background-size: 50%;
      background-image: url('./images/checked.png');
    }
  }

  &-message {
    position: absolute;
    left: -1px;
    top: -1px;
    right: 0;
    bottom: 0;
    z-index: 5;
    width: calc(100% + 2px);
    height: calc(100% + 2px);
    color: #fff;
    @include flex-row();
    align-items: center;
    justify-content: center;
    font-size: 14px;
    font-weight: bold;
    opacity: 0;
    transition: opacity 0.1s linear;
    pointer-events: none;

    .icon {
      width: 20px;
      height: 20px;
      margin-right: 8px;
    }
  }

  &-loading {
    min-width: 80px;
    height: 8px;
    margin-top: -8px;
    text-align: center;

    span {
      display: inline-block;
      width: 8px;
      height: 100%;
      margin-right: 5px;
      background: #fff;
      animation: loading 1.04s ease infinite;

      &:nth-child(1) {
        animation-delay: 0.13s;
      }

      &:nth-child(2) {
        animation-delay: 0.26s;
      }

      &:nth-child(3) {
        animation-delay: 0.39s;
      }

      &:nth-child(4) {
        animation-delay: 0.52s;
      }

      &:nth-child(5) {
        animation-delay: 0.65s;
      }

      &:last-child {
        margin-right: 0px;
      }
    }
  }
}

@keyframes wave {
  0% {
    bottom: -100%;
    left: -100%;
    transition: all 0.2s;
  }
  100% {
    left: 110%;
    bottom: 110%;
    transition: all 0.2s;
  }
}

@keyframes loading {
  0% {
    opacity: 1;
  }
  100% {
    opacity: 0;
  }
}
</style>
