//缓存
let CSMSKEY = {
  expire: 1 * 60 * 60, //单位秒：设置为1小时
  token: 'csms-web-token'
};

//加密
let CRYPTO = {
  secret: 'csms-web-jie-secret'
};

//请求
let URIS = {
  base: 'http://127.0.0.1:7000/',
  user: {
    login: '/api/user/login',
    register: '/api/user/register',
    captcha: 'api/user/captcha',
    city: '/api/user/city',
    profile: '/api/user/profile',
    address: '/api/user/userAddress',
    delAddress: '/api/user/deleteAddress',
    saveAddress: '/api/user/saveAddress'
  },
  goods: 'http://127.0.0.1:7000/uploads/goods/',
  kindeditor: 'http://127.0.0.1:7000/goods/product/upload'
};

export { CSMSKEY, CRYPTO, URIS };
