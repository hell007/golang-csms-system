import vue from 'vue'

import component from './message.vue'

let loading

const MessageConstructor = vue.extend(component)

const defaults = {
  visible: false,
  text: '',
  button: []
};

//这里关闭的时候返回promise
MessageConstructor.prototype.close = function() {
  var vm = this;
  loading = null;
  var promise = new Promise(function(resolve, reject) {
    if (vm.$el && vm.$el.parentNode) {
      vm.$el.parentNode.removeChild(vm.$el);
    }
    resolve();
    vm.$destroy();
    vm.visible = false;
  })
  return promise
};

function show(options) {
  options = Object.assign({}, defaults, options)

  let parent = document.body
  
  //没有关闭不允许新开
  if (loading) {
    return loading
  }

  const Dom = new MessageConstructor({
    el: document.createElement('div'),
    data: options
  })

  parent.appendChild(Dom.$el)

  vue.nextTick(() => {
    Dom.visible = true;
  });

  loading = Dom

  return Dom;
}

function registryMessage() {
  vue.prototype.$confirm = show
}

export default registryMessage


// let confirm = this.$confirm({
//   text: '',
//   button: [{
//     text: '取消',
//     ontap() {
//       confirm.close()
//     }
//   }, {
//     text: '确定',
//     ontap() {
//       confirm.close().then((res) => {
//         console.log('22222')
//       });
//     }
//   }]
// });