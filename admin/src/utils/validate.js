/**
 * validate
 * @time 2017-10-17
 * @author 1048523306@qq.com
 */


/* 合法用户 */
export function validateUsername(rule, value, callback) {
  const valid_map = ['admin', 'editor'] //>0 true
  if (valid_map.indexOf(value.trim()) < 0) {
    callback(new Error('请输入正确的用户名'))
  } else {
    callback()
  }
}


/* 合法密码 */
export function validatePassword(rule, value, callback) {
  if (value.length < 6) {
    callback(new Error('密码不能小于6位'))
  } else {
    callback()
  }
}


/* 合法手机号码 */
export function validateMobile(rule, value, callback) {
  const reg = /^1[34578]\d{9}$/
  if (!reg.test(value)) {
    callback(new Error("输入的手机号码不正确"));
  } else {
    callback();
  }
}


/* 合法uri*/
export function validateURL(str) {
  const reg = /^(https?|ftp):\/\/([a-zA-Z0-9.-]+(:[a-zA-Z0-9.&%$-]+)*@)*((25[0-5]|2[0-4][0-9]|1[0-9]{2}|[1-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|1[0-9]{2}|[1-9]?[0-9])){3}|([a-zA-Z0-9-]+\.)*[a-zA-Z0-9-]+\.(com|edu|gov|int|mil|net|org|biz|arpa|info|name|pro|aero|coop|museum|[a-zA-Z]{2}))(:[0-9]+)*(\/($|[a-zA-Z0-9.,?'\\+&%$#=~_-]+))*$/
  return reg.test(textval)
}

/* 小写字母*/
export function validateLowerCase(str) {
  const reg = /^[a-z]+$/
  return reg.test(str)
}

/* 大写字母*/
export function validateUpperCase(str) {
  const reg = /^[A-Z]+$/
  return reg.test(str)
}

/* 大小写字母*/
export function validatAlphabets(str) {
  const reg = /^[A-Za-z]+$/
  return reg.test(str)
}