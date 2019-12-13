import { getPermission } from '@/utils/storage'

let auth = function(value) {
  let b = false;
  let p = getPermission();

  if (!p) {
    return false;
  }

  let permissions = getPermission();

  for (let i = 0; i < permissions.length; i++) {

    if (permissions[i].code === value) {
      b = true;
      break;
    }
  }

  return b;
}

let has = function(permission){
  if(!permissions[permission]){
    return false;
  }
  return true;
}

export default {
  bind() {
  },
  inserted(el, binding) {
    if (!auth(binding.value)) {
      el.parentNode.removeChild(el);
    }
  },
  update(el) {
    //console.log('update');
  },
  componentUpdated(el) {
    //console.log('componentUpdated');
  },
  unbind(el) {
    //console.log('unbind');
  }
}