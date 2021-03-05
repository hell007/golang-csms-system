Vue.directive('has', {
  bind: function(el, binding) {
    if (!Vue.prototype.$_has(binding.value)) {
      el.parentNode.removeChild(el);
    }
  }
});
//权限检查方法
Vue.prototype.$_has = function(value) {
  debugger
  let isExist = false;
  let buttonpermsStr = sessionStorage.getItem("buttenpremissions");
  if (buttonpermsStr == undefined || buttonpermsStr == null) {
    return false;
  }
  let buttonperms = JSON.parse(buttonpermsStr);
  for (let i = 0; i < buttonperms.length; i++) {
    if (buttonperms[i].perms.indexOf(value) > -1) {
      isExist = true;
      break;
    }
  }
  return isExist;
};