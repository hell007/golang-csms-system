import Mock from 'mockjs'
import { param2Obj } from '@/utils'

const List = []
const count = 100

const baseContent = '<p>我是测试数据</p>'
const image = 'https://wpimg.wallstcn.com/577965b9-bb9e-4e02-9f0c-095b41417191'

for (let i = 0; i < count; i++) {
  List.push(Mock.mock({
    id: '@increment',
    avatar: 'https://wpimg.wallstcn.com/577965b9-bb9e-4e02-9f0c-095b41417191',
    name: '@first',
    'sex|1': ['男', '女', '保密'],
    mobile: '@integer(13500000000, 19900000000)',
    'status|1': [1, 2],
    'isSuper|1' : [1, 2],
    'roleId|1': [1, 2],
    createTime: '@datetime',// +Mock.Random.date('T'),
    note: '请注意',
    area: [],
    'address|': ['南屏街','小西门'],
    ids: '530129201903260000',
    photos: [],
    content:baseContent,
  }))
}

export default {
  getUnitList: config => {

    //操作4-> 返回模拟数据 开发环境 （生产环境由后台服务返回）

    //console.log('34-> 调用mock里面的getUserlist',param2Obj(config.url))

    const { name, mobile, pageNum = 1, pageSize = 10, sort } = param2Obj(config.url)

    let mockList = List.filter(item => {
      if (mobile && item.mobile !== mobile) return false
      if (name && item.name.indexOf(name) < 0) return false
      return true
    })

    if (sort === '-id') {
      mockList = mockList.reverse()
    }

    const pageList = mockList.filter((item, index) => index < pageSize * pageNum && index >= pageSize * (pageNum - 1))

    //console.log(50,'mock user -> getUserList -> pageList=', pageList)

    return {
      status: true,
      data: {
        total: mockList.length,
        items: pageList
      }
    }
  },
  getUnit: (config) => {
    const { id } = param2Obj(config.url)
    for (const user of List) {
      if (user.id === +id) {
        user.status = 1
        user.sex = '男'
        user.mobile = '13888888888'
        return {
          status: true,
          data: user
        }
      }
    }
  },
  saveUnit: (config) => {
    const { id } = param2Obj(config.url)
    for (const user of List) {
      if (user.id === +id) {
        user.name = '曹操'
        user.status = 0
        user.sex = '男'
        user.note = '我是更改后的数据'
        return {
          status: true,
          data: user
        }
      }
    }
  },
  deleteUnit: (config) => {
    const { id } = param2Obj(config.url)
    for (const user of List) {
      if (user.id === +id) {
        return {
          status: true
        }
      }
    }
  },
  batchDeleteUnit: (config) => {
    const { ids } = param2Obj(config.url)
    return {
      status: true
    }
  }
}