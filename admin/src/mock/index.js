import Mock from 'mockjs'
import UnitAPI from './unit'

Mock.setup({
  timeout: '350-600'
})

// 拦截需要mock处理的假数据 --》 url 请求

Mock.mock(/\/unit\/list\.*/, 'get', UnitAPI.getUnitList)
Mock.mock(/\/unit\/detail\.*/, 'get', UnitAPI.getUnit)
Mock.mock(/\/unit\/save\.*/, 'post', UnitAPI.saveUnit)
Mock.mock(/\/unit\/delete\.*/, 'post', UnitAPI.deleteUnit)
Mock.mock(/\/unit\/batchdelete\.*/, 'post', UnitAPI.batchDeleteUnit)

export default Mock
