import vue from 'vue'

import component from './toast.vue'

const ToastConstructor = vue.extend(component)

// 定义弹出组件的函数 接收2个参数, 要显示的文本 和 显示时间
function showToast(text, icon, duration = 4000) {
  const Dom = new ToastConstructor({
    el: document.createElement('div'),
    data() {
      return {
        text: text,
        visible: false, // 是否显示组件
        icon: icon, // 图标
        showContent: true // 作用:在隐藏组件之前,显示隐藏动画
      }
    }
  })
  document.body.appendChild(Dom.$el)

  vue.nextTick(() => {
    Dom.visible = true;
  });

  // 提前 250ms 执行淡出动画(因为我们再css里面设置的隐藏动画持续是250ms)
  setTimeout(() => {
      Dom.showContent = false
    }, duration - 1250)
    // 过了 duration 时间后隐藏整个组件
  setTimeout(() => {
    Dom.visible = false
  }, duration)
}

// 注册为全局组件的函数
function registryToast() {
  // 将组件注册到 vue 的 原型链里去,
  // 这样就可以在所有 vue 的实例里面使用 this.$toast()
  vue.prototype.$toast = showToast
}

export default registryToast