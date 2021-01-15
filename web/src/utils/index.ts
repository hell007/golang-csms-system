import moment from 'moment';
import lodash from 'lodash';

moment.locale('zh-cn');

/**
 * Lodash 中文文档 https://www.lodashjs.com/
 */
export { lodash };

/**
 * 格式化时间
 * @param datetime
 * @param type YYYY-MM-DD HH:mm:ss
 */
export function formatDate(datetime: any, type: string) {
  return moment(datetime).format(type);
}

/**
 * 获取年(years)、月(months)、日(days)、周(weeks)的上n个时间
 * @param number
 * @param mdate
 * @param type
 */
export function lastDate(number: number, mdate: any, type: string) {
  return moment()
    .subtract(number, mdate)
    .format(type);
}

/**
 * 获取当前 年(YYYY) 月(MM) 日(DD)
 * @param type
 */
export function getDate(type: string) {
  return parseInt(moment().format(type));
}

/**
 * 判断any是否为空
 * @param value
 */
export function isEmpty(value: any) {
  return lodash.isEmpty(value);
}

/**
 * 延迟
 * @param time
 */
export function delay(time: number) {
  return new Promise(function(resolve, _reject) {
    window.setTimeout(function() {
      resolve(time);
    }, time);
  });
}

/**
 * 将形如url?(page=1&sort=4)param转换为{page:'1',sort:'4'}
 * @param url
 */
export function param2Obj(url: string) {
  const search = url.split('?')[1];
  if (!search) return {};
  return JSON.parse(
    '{"' +
      decodeURIComponent(search)
        .replace(/"/g, '\\"')
        .replace(/&/g, '","')
        .replace(/=/g, '":"') +
      '"}'
  );
}

/**
 * 将形如{page:'1',sort:'4'}JSON转换为(page=1&sort=4)param
 * @param json
 */
export function obj2Param(json: any) {
  if (!json) return '';
  return Object.keys(json)
    .map((key) => {
      if (json[key] === undefined) return '';
      return encodeURIComponent(key) + '=' + encodeURIComponent(json[key]);
    })
    .join('&');
}

/**
 * 
 //通过正则表达式判断手机号码格式是否正确,根据电信，联通、移动手机号码规则可以到以下正则
  // 手机号码第一位是[1]开头，第二位[3,4,5,7,8]中的一位，第三位到第十一位则是[0-9]中的数字；
  //^1表示开头为1
  //[3|4|5|7|8] 表示3、4、5、7、8中的一位数值
  //[0-9]{9} 匹配包含0-9的数字
  * @param phone 
 */
export function isPhone(phone: string) {
  if (isEmpty(phone)) {
    return false;
  }
  let reg = /^1[3|4|5|7|8][0-9]{9}/;
  return reg.test(phone) ? true : false;
}

/**
 * 手机号码分段显示
 * @param {string} phone
 */
export function formatPhone(phone: string) {
  if (!isEmpty(phone)) {
    return phone.replace(/(?=(\d{4})+$)/g, ' ');
  }
}

/**
 * 隐私号码
 * @param {string} phone
 */
export function privatePhone(phone: string) {
  if (!isEmpty(phone)) {
    return phone.replace(/^(\d{3})\d{4}(\d+)/, '$1****$2');
  }
}

/**
 * 将传入的数字以3位添加逗号返回 1,234元
 * [toThousandslsFilter description]
 * @param  {number} num
 */
export function toThousandslsFilter(num: number) {
  if (!isEmpty(num)) {
    return num.toString().replace(/^-?\d+/g, m => m.replace(/(?=(?!\b)(\d{3})+$)/g, ','))
  }
}
