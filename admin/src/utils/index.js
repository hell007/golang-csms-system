/**
 * uitls
 * @time 2017-10-17
 * @author 1048523306@qq.com
 */
import UUID from 'uuid';
import moment from 'moment';

moment.locale('zh-cn');

// noop
export function noop() {}

// uuid
export function uuid() {
  return UUID.v1().toString().replace(/-/g, '');
}

// 格式化时间 
export function formatDate(datetime, type) {
  return moment(datetime).format(type);
}

// 获取年(years)、月(months)、日(days)、周(weeks)的上n个时间
export function lastDate(number, mdate, type) {
  return moment().subtract(number, mdate).format(type)
}

// 获取当前 年(YYYY) 月(MM) 日(DD) 
export function getDate(type) {
  return parseInt(moment().format(type))
}

// 延时请求模拟
export function delay(time) {
  return new Promise(function(resolve, reject) {
    window.setTimeout(function() {
      resolve(time)
    }, time)
  })
}

// 随机数
export function random(low, high) {
  if (arguments.length === 1) {
    high = low
    low = 0
  }
  return Math.floor(low + Math.random() * (high - low))
}

export function randomOne(arr) {
  return arr[random(0, arr.length)]
}

export function kebabCase(s) {
  return s.replace(/[A-Z]/g, function($0) {
    return '-' + $0.toLowerCase()
  })
}

export function forIn(obj, fn) {
  obj = obj || {}
  for (var key in obj) {
    if (obj.hasOwnProperty(key)) {
      fn.call(obj, obj[key], key, obj)
    }
  }
}

// 判断空
export function isEmpty(obj) {
  //typeof
  if (typeof obj == 'undefined' || obj == null || obj == '') {
    return true
  } else if (typeof obj == 'object' && Object.keys(obj).length == 0) {
    return true
  } else if (typeof obj == 'array' && obj.length == 0) {
    return true
  } else {
    return false
  }
}

//logger
export function logger(type, text) {
  if (Object.prototype.toString.call(text) === '[Object Object]') {
    text = JSON.stringify(text)
  }
  // return false;
  console && console[type] && console[type]('node front: ', text)
}

/* 将时间戳转换为datetime
 * [parseTime description]
 * @param  {String} time    [时间戳]
 * @param  {String} cFormat [转换格式]
 * @return {String}         [datetime]
 */
export function parseTime(time, cFormat) {
  if (arguments.length === 0) {
    return null
  }
  const format = cFormat || '{y}-{m}-{d} {h}:{i}:{s}'
  let date
  if (typeof time === 'object') {
    date = time
  } else {
    if (('' + time).length === 10) time = parseInt(time) * 1000
    date = new Date(time)
  }
  const formatObj = {
      y: date.getFullYear(),
      m: date.getMonth() + 1,
      d: date.getDate(),
      h: date.getHours(),
      i: date.getMinutes(),
      s: date.getSeconds(),
      a: date.getDay()
    }
    /**
     * str.replace(reg,function(result, key, index, str){...})
     * @param reg 正则匹配
     * @param callback
     * result 第一个参数表示匹配到的字符
     * key    第二个参数表示匹配时key
     * index  第三个参数表示匹配时的字符最小索引位置
     * str    第四个参数表示被匹配的字符串 str
     * 注意：根据不同的正则表达式，传入的参数表示不同 (见 getQueryObject 函数)
     */
  const timeStr = format.replace(/{(y|m|d|h|i|s|a)+}/g, (result, key) => {
    let value = formatObj[key]
    if (key === 'a') return ['一', '二', '三', '四', '五', '六', '日'][value - 1]
    if (result.length > 0 && value < 10) {
      value = '0' + value
    }
    return value || 0
  })
  return timeStr
}


// 时间格式(2014-02-02 14:10:00)转换成时间戳   
export function formatTimestamp(datetime) {
  var arr = datetime.replace(/:/g, "-").replace(/ /g, "-").split("-");
  var timestamp = new Date(Date.UTC(arr[0], arr[1] - 1, arr[2], arr[3] - 8, arr[4], arr[5]));
  return timestamp.getTime();
}


/**
 * 将时间转换为日期 或者是 多久前
 * [formatTime description]
 * @param  {String} time   [时间戳]
 * @param  {String} option [cFormat]
 * @return {月日时分|datetime} 
 */
export function formatTime(time, option) {
  time = +time * 1000
  const d = new Date(time)
  const now = Date.now()

  const diff = (now - d) / 1000

  if (diff < 30) {
    return '刚刚'
  } else if (diff < 3600) { // less 1 hour
    return Math.ceil(diff / 60) + '分钟前'
  } else if (diff < 3600 * 24) {
    return Math.ceil(diff / 3600) + '小时前'
  } else if (diff < 3600 * 24 * 2) {
    return '1天前'
  }
  if (option) {
    return parseTime(time, option)
  } else {
    return d.getMonth() + 1 + '月' + d.getDate() + '日' + d.getHours() + '时' + d.getMinutes() + '分'
  }
}


/**
 * 多少时间前
 * [timeAgo description]
 * @param  {Number} time 
 * @return {String}     
 */
export function timeAgo(time) {
  const between = Date.now() / 1000 - Number(time)
  if (between < 3600) {
    return pluralize(~~(between / 60), ' minute')
  } else if (between < 86400) {
    return pluralize(~~(between / 3600), ' hour')
  } else {
    return pluralize(~~(between / 86400), ' day')
  }

  function pluralize(time, label) {
    if (time === 1) {
      return time + label
    }
    return time + label + 's'
  }
}


/**
 * 数字 格式化
 * [nFormatter description]
 * @param  {[type]} num    [description]
 * @param  {[type]} digits [description]
 * @return {[type]}        [description]
 */
export function nFormatter(num, digits) {
  const si = [{
    value: 1E18,
    symbol: 'E'
  }, {
    value: 1E15,
    symbol: 'P'
  }, {
    value: 1E12,
    symbol: 'T'
  }, {
    value: 1E9,
    symbol: 'G'
  }, {
    value: 1E6,
    symbol: 'M'
  }, {
    value: 1E3,
    symbol: 'k'
  }]
  for (let i = 0; i < si.length; i++) {
    if (num >= si[i].value) {
      return (num / si[i].value + 0.1).toFixed(digits).replace(/\.0+$|(\.[0-9]*[1-9])0+$/, '$1') + si[i].symbol
    }
  }
  return num.toString()
}


/**
 * 根据url获取？后面的参数，以对象的形式返回
 * [getQueryObject description]
 * @param  {String} url 
 * @return {Object}     [{prame:value}]
 */
export function getQueryObject(url) {
  url = url == null ? window.location.href : url
  const search = url.substring(url.lastIndexOf('?') + 1)
  const obj = {}
  const reg = /([^?&=]+)=([^?&=]*)/g
    //注意这里的 $1 $2 与上文的区别
  search.replace(reg, (rs, $1, $2) => {
    const name = decodeURIComponent($1)
    let val = decodeURIComponent($2)
    val = String(val)
    obj[name] = val
    return rs
  })
  return obj
}


/**
 * 结果与getQueryObject相同
 * [param2Obj description]
 * @param  {String} url 
 * @return {jsonString}   
 */
export function param2Obj(url) {
  const search = url.split('?')[1]
  if (!search) return {}
  return JSON.parse('{"' + decodeURIComponent(search).replace(/"/g, '\\"').replace(/&/g, '","').replace(/=/g, '":"') + '"}')
}


/**
 * 将形如{page:'1',sort:'4'}JSON转换为(page=1&sort=4)param
 * [obj2Param description]
 * @param  {Object} json 
 * @return {String}    
 */
export function obj2Param(json) {
  if (!json) return ''
  return Object.keys(json).map(key => {
    if (json[key] === undefined) return ''
    return encodeURIComponent(key) + '=' +
      encodeURIComponent(json[key])
  }).join('&')
}


/**
 * [getByteLen description]
 * @param  {[type]} val [description]
 * @return {[type]}     [description]
 */
export function getByteLen(val) {
  let len = 0
  for (let i = 0; i < val.length; i++) {
    if (val[i].match(/[^\x00-\xff]/ig) != null) {
      len += 1
    } else {
      len += 0.5
    }
  }
  return Math.floor(len)
}


/**
 * [html2Text description]
 * @param  {[type]} val 
 * @return {[type]}     
 */
export function html2Text(val) {
  const div = document.createElement('div')
  div.innerHTML = val
  return div.textContent || div.innerText
}


/**
 * 将传入的数字以3位添加逗号返回 1,234元
 * [toThousandslsFilter description]
 * @param  {Number} num 
 * @return {Stirng}     
 */
export function toThousandslsFilter(num) {
  return (+num || 0).toString().replace(/^-?\d+/g, m => m.replace(/(?=(?!\b)(\d{3})+$)/g, ','))
}


/**
 * 将两个参数合并为一个object
 * [objectMerge description]
 * @param  {any} target 
 * @param  {any} source 
 * @return {Object}    
 * 注意,对于已知识obj的可以使用：obj = Object.assign(obj1,obj2)    
 */
export function objectMerge(target, source) {
  if (typeof target !== 'object') {
    target = {}
  }
  if (Array.isArray(source)) {
    return source.slice()
  }
  for (const property in source) {
    if (source.hasOwnProperty(property)) {
      const sourceProperty = source[property]
      if (typeof sourceProperty === 'object') {
        target[property] = objectMerge(target[property], sourceProperty)
        continue
      }
      target[property] = sourceProperty
    }
  }
  return target
}


/**
 * [toggleClass description]
 * @param  {[type]} element   
 * @param  {[type]} className 
 * @return {[type]}           
 */
export function toggleClass(element, className) {
  if (!element || !className) return
  let classString = element.className
  const nameIndex = classString.indexOf(className)
  if (nameIndex === -1) {
    classString += '' + className
  } else {
    classString = classString.substr(0, nameIndex) + classString.substr(nameIndex + className.length)
  }
  element.className = classString
}


/**
 * element ui datatimepicker
 * [pickerOptions description]
 * @type {Array}
 */
export const pickerOptions = [{
  text: '今天',
  onClick(picker) {
    const end = new Date()
    const start = new Date(new Date().toDateString())
    end.setTime(start.getTime())
    picker.$emit('pick', [start, end])
  }
}, {
  text: '最近一周',
  onClick(picker) {
    const end = new Date(new Date().toDateString())
    const start = new Date()
    start.setTime(end.getTime() - 3600 * 1000 * 24 * 7)
    picker.$emit('pick', [start, end])
  }
}, {
  text: '最近一个月',
  onClick(picker) {
    const end = new Date(new Date().toDateString())
    const start = new Date()
    start.setTime(start.getTime() - 3600 * 1000 * 24 * 30)
    picker.$emit('pick', [start, end])
  }
}, {
  text: '最近三个月',
  onClick(picker) {
    const end = new Date(new Date().toDateString())
    const start = new Date()
    start.setTime(start.getTime() - 3600 * 1000 * 24 * 90)
    picker.$emit('pick', [start, end])
  }
}]


/**
 * [getTime description]
 * @param  {[type]} type 
 * @return {[type]}      
 */
export function getTime(type) {
  if (type === 'start') {
    return new Date().getTime() - 3600 * 1000 * 24 * 90
  } else {
    return new Date(new Date().toDateString())
  }
}


/**uuid
 * [createUniqueString description]
 * @return {[type]} [description]
 */
export function createUniqueString() {
  const timestamp = +new Date() + ''
  const randomNum = parseInt((1 + Math.random()) * 65536) + ''
  return (+(randomNum + timestamp)).toString(32)
}

/**
 * 获取树
 * [getTree description]
 * @param  {[type]} nodes
 * @return {[type]} 
 */
export function getTree(nodes) {
  // 删除 所有 children,以防止多次调用
  nodes.forEach(function(item) {
    delete item.children;
  });

  // 将数据存储为 以 id 为 KEY 的 map 索引数据列
  var map = {};
  nodes.forEach(function(item) {
    map[item.id] = item;
  });

  var val = [];
  nodes.forEach(function(item) {
    // 以当前遍历项，的pid,去map对象中找到索引的id
    var parent = map[item.pid];
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


export function debounce(func, wait, immediate) {
  let timeout, args, context, timestamp, result

  const later = function() {
    // 据上一次触发时间间隔
    const last = +new Date() - timestamp

    // 上次被包装函数被调用时间间隔last小于设定时间间隔wait
    if (last < wait && last > 0) {
      timeout = setTimeout(later, wait - last)
    } else {
      timeout = null
        // 如果设定为immediate===true，因为开始边界已经调用过了此处无需调用
      if (!immediate) {
        result = func.apply(context, args)
        if (!timeout) context = args = null
      }
    }
  }

  return function(...args) {
    context = this
    timestamp = +new Date()
    const callNow = immediate && !timeout
      // 如果延时不存在，重新设定延时
    if (!timeout) timeout = setTimeout(later, wait)
    if (callNow) {
      result = func.apply(context, args)
      context = args = null
    }

    return result
  }
}


export function deepClone(source) {
  if (!source && typeof source !== 'object') {
    throw new Error('error arguments', 'shallowClone')
  }
  const targetObj = source.constructor === Array ? [] : {}
  for (const keys in source) {
    if (source.hasOwnProperty(keys)) {
      if (source[keys] && typeof source[keys] === 'object') {
        targetObj[keys] = source[keys].constructor === Array ? [] : {}
        targetObj[keys] = deepClone(source[keys])
      } else {
        targetObj[keys] = source[keys]
      }
    }
  }
  return targetObj
}


export function scrollTo(element, to, duration) {
  if (duration <= 0) return
  const difference = to - element.scrollTop
  const perTick = difference / duration * 10
  setTimeout(() => {
    console.log(new Date())
    element.scrollTop = element.scrollTop + perTick
    if (element.scrollTop === to) return
    scrollTo(element, to, duration - 10)
  }, 10)
}

// 去重
export function uniquelizeArray(a) {
  var r = [];
  for (var i = 0; i < a.length; i++) {
    var flag = true;
    var temp = a[i];
    for (var j = 0; j < r.length; j++) {
      if (temp === r[j]) {
        flag = false;
        break;
      }
    }
    if (flag) {
      r.push(temp);
    }
  }
  return r;
}

// 去重
export function unique(arr = []) {
  return arr.reduce((t, v) => t.includes(v) ? t : [...t, v], []);
}

// 交集
export function intersectArray(a, b) {
  return a.concat(b.filter(function(v) {
    return a.indexOf(v) === -1
  }))
}

// 并集
export function unionArray(a, b) {
  return a.filter(function(v) {
    return b.indexOf(v) > -1
  })
}

// 差集
export function minusArray(a, b) {
  return a.filter(function(v) {
    return b.indexOf(v) === -1
  })
}

// 差集
export function difference(arr = [], oarr = []) {
  return arr.reduce((t, v) => (!oarr.includes(v) && t.push(v), t), []);
}

// 数组最大值
export function max(arr = []) {
  return arr.reduce((t, v) => t > v ? t : v);
}

// 数组最小值
export function min(arr = []) {
  return arr.reduce((t, v) => t < v ? t : v);
}

// 数字千分化
export function thousandNum(num = 0) {
  const str = (+num).toString().split(".");
  const int = nums => nums.split("").reverse().reduceRight((t, v, i) => t + (i % 3 ? v : `${v},`), "").replace(/^,|,$/g, "");
  const dec = nums => nums.split("").reduce((t, v, i) => t + ((i + 1) % 3 ? v : `${v},`), "").replace(/^,|,$/g, "");
  return str.length > 1 ? `${int(str[0])}.${dec(str[1])}` : int(str[0]);
}