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
 * @param phone
 */
export function formatPhone(phone: string) {
  if (!isEmpty(phone)) {
    return phone.replace(/(?=(\d{4})+$)/g, ' ');
  }
}

/**
 * 隐私号码
 * @param phone
 */
export function privatePhone(phone: string) {
  if (!isEmpty(phone)) {
    return phone.replace(/^(\d{3})\d{4}(\d+)/, '$1****$2');
  }
}

/**
 * 获取树
 * @param  nodes
 */
export function getTree(nodes: any[]) {
  // 删除 所有 children,以防止多次调用
  nodes.forEach(function(item) {
    delete item.children;
  });

  // 将数据存储为 以 id 为 KEY 的 map 索引数据列
  let map: any = {};
  nodes.forEach(function(item) {
    map[item.id] = item;
  });

  let val: any = [];
  nodes.forEach(function(item) {
    // 以当前遍历项，的pid,去map对象中找到索引的id
    let parent = map[item.pid];
    // 好绕啊，如果找到索引，那么说明此项不在顶级当中,那么需要把此项添加到，他对应的父级中
    if (parent) {
      (parent.children || (parent.children = [])).push(item);
    } else {
      //如果没有在map中找到对应的索引ID,那么直接把 当前的item添加到 val结果集中，作为顶级
      val.push(item);
    }
  });
  return val;
}

/**
 * html2Text
 * @param val
 */
export function html2Text(val: any) {
  const div = document.createElement('div');
  div.innerHTML = val;
  return div.textContent || div.innerText;
}

/**
 * 将传入的数字以3位添加逗号返回 1,234元
 * @param  num
 */
export function formatThousandsls(num: number) {
  if (!isEmpty(num)) {
    return num.toString().replace(/^-?\d+/g, (m) => m.replace(/(?=(?!\b)(\d{3})+$)/g, ','));
  }
}

/**
 * 数字格式化
 * @param num
 * @param digits
 */
export function nFormatter(num: number, digits: number) {
  const si = [
    {
      value: 1e18,
      symbol: 'E'
    },
    {
      value: 1e15,
      symbol: 'P'
    },
    {
      value: 1e12,
      symbol: 'T'
    },
    {
      value: 1e9,
      symbol: 'G'
    },
    {
      value: 1e6,
      symbol: 'M'
    },
    {
      value: 1e3,
      symbol: 'k'
    }
  ];

  for (let item of si) {
    if (num >= item.value) {
      return (
        (num / item.value + 0.1).toFixed(digits).replace(/\.0+$|(\.[0-9]*[1-9])0+$/, '$1') +
        item.symbol
      );
    }
  }
  return num.toString();
}
