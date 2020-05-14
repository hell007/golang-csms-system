<template>
  <el-card class="xa-block" shadow="hover">
    <div class="xa-block chart-btn">
      <el-button
        v-for="item,index in list"
        :key="index"
        @click="setCharts(item, index)"
        :type="active==index ? 'primary' : 'info'"
        size="small">{{item.name}}</el-button>
    </div>
    <div class="xa-block" ref="wrap">
      <svg 
        v-for="item,index in list"
        :key="index"
        v-if="active==index"
        :class="`container-${index}`"></svg>
    </div>
  </el-card>
</template>

<script>
import { Theme } from '@/config'
import * as d3 from '@/vendor/d3/index'
export default {
  name: 'charts-d3',
  components: {
  },
  data() {
    return {
      primary: Theme.primary,
      active: 0,
      list: [{
        name: '柱状图',
        callback: function(opt){
          d3.setBar(opt)
        },
        opt: {
          container: null,
          containerWidth: 1200,
          containerHeight: 500,
          margin: {
            top: 80,
            right: 20,
            bottom: 30,
            left: 40
          },
          bgs: ['#ccc', '#ddd'],
          colors: [
            '#19d4ae',
            '#5867c3',
            '#f7ba2a',
            '#13ce66',
            '#8acc6d',
            '#20a0ff',
            '#fa6e86'
          ],
          title: '本周市民外出率',
          data: [
            { letter: '一', frequency: 0.08167 },
            { letter: '二', frequency: 0.13492 },
            { letter: '三', frequency: 0.16782 },
            { letter: '四', frequency: 0.14253 },
            { letter: '五', frequency: 0.32702 },
            { letter: '六', frequency: 0.62288 },
            { letter: '日', frequency: 0.22288 }
          ]
        }
      }, {
        name: '组合柱状图',
        callback: function(opt) {
          d3.setGroupedBar(opt)
        },
        opt: {
          container: null,
          containerWidth: 1200,
          containerHeight: 600,
          margin: {
            top: 80,
            right: 80,
            bottom: 30,
            left: 60
          },
          title: '某某公司近几年各季度产生营收情况汇总',
          data: [
            { date: '2011', q1: 155, q2: 200, q3: 214, q4: 234 },
            { date: '2012', q1: 165, q2: 210, q3: 244, q4: 254 },
            { date: '2013', q1: 175, q2: 230, q3: 274, q4: 274 },
            { date: '2014', q1: 185, q2: 250, q3: 304, q4: 294 },
            { date: '2015', q1: 195, q2: 270, q3: 334, q4: 314 },
            { date: '2016', q1: 205, q2: 290, q3: 364, q4: null }
          ]
        }  
      }, {
        name: '堆栈柱状图',
        callback: function(opt) {
          d3.setStackedBar(opt)
        },
        opt: {
          container: null,
          containerWidth: 1200,
          containerHeight: 500,
          margin: {
            top: 80,
            right: 80,
            bottom: 30,
            left: 60
          },
          title: '某某公司近几年各季度产生营收情况汇总',
          data: [
            { date: '2011', q1: 155, q2: 200, q3: 214, q4: 234 },
            { date: '2012', q1: 165, q2: 210, q3: 244, q4: 254 },
            { date: '2013', q1: 175, q2: 230, q3: 274, q4: 274 },
            { date: '2014', q1: 185, q2: 250, q3: 304, q4: 294 },
            { date: '2015', q1: 195, q2: 270, q3: 334, q4: 314 },
            { date: '2016', q1: 205, q2: 290, q3: 364, q4: null }
          ]
        }  
      }, {
        name: '径向堆栈柱状图',
        callback: function(opt) {
          d3.setRadialStackedBar(opt)
        },
        opt: {
          container: null,
          containerWidth: 1200,
          containerHeight: 600,
          margin: {
            top: 80,
            right: 80,
            bottom: 30,
            left: 60
          },
          title: '云南省2018年各季度GDP汇总',
          data: [
            { city: 'A城市', q1: 3546, q2: 3132, q3: 2299, q4: 1337 },
            { city: 'B城市', q1: 199, q2: 275, q3: 275, q4: 299 },
            { city: 'C城市', q1: 175, q2: 235, q3: 238, q4: 268 },
            { city: 'D城市', q1: 154, q2: 200, q3: 214, q4: 234 },
            { city: 'E城市', q1: 123, q2: 127, q3: 168, q4: 139 },
            { city: 'F城市', q1: 137, q2: 153, q3: 177, q4: 172 },
            { city: 'G城市', q1: 148, q2: 186, q3: 198, q4: 207 },
            { city: 'H城市', q1: 155, q2: 200, q3: 214, q4: 234 },
            { city: 'I城市', q1: 165, q2: 210, q3: 244, q4: 254 },
            { city: 'J城市', q1: 175, q2: 230, q3: 274, q4: 274 },
            { city: 'K城市', q1: 185, q2: 250, q3: 304, q4: 294 },
            { city: 'L城市', q1: 195, q2: 270, q3: 334, q4: 314 },
            { city: 'M城市', q1: 205, q2: 290, q3: 364, q4: 330 },
            { city: 'N城市', q1: 546, q2: 988, q3: 1024, q4: 1254 },
            { city: 'O城市', q1: 3514, q2: 2541, q3: 1987, q4: 1752 },
            { city: 'P城市', q1: 3654, q2: 3787, q3: 3654, q4: 2415 },
            { city: 'Q城市', q1: 368, q2: 385, q3: 244, q4: 545 },
            { city: 'R城市', q1: 232, q2: 555, q3: 274, q4: 274 },
            { city: 'S城市', q1: 358, q2: 344, q3: 304, q4: 787 },
            { city: 'T城市', q1: 855, q2: 865, q3: 334, q4: 875 },
            { city: 'U城市', q1: 453, q2: 422, q3: 364, q4: 330 },
            { city: 'V城市', q1: 568, q2: 478, q3: 875, q4: 254 },
            { city: 'W城市', q1: 554, q2: 234, q3: 695, q4: 948 },
            { city: 'X城市', q1: 938, q2: 875, q3: 304, q4: 585 },
            { city: 'Y城市', q1: 247, q2: 757, q3: 578, q4: 857 },
            { city: 'Z城市', q1: 368, q2: 695, q3: 757, q4: 875 }
          ]
        }  
      }, {
        name: '面积图',
        callback: function(opt) {
          d3.setArea(opt)
        },
        opt: {
          container: null,
          containerWidth: 1200,
          containerHeight: 500,
          margin: {
            top: 80,
            right: 20,
            bottom: 130,
            left: 40
          },
          margin2: {
            top: 410,
            right: 40,
            bottom: 60,
            left: 40
          },
          title: '2019年昆明饭店每月接待顾客人数',
          data: [
            { day: '2017-01', quantity: 1240 },
            { day: '2017-02', quantity: 1905 },
            { day: '2017-03', quantity: 6232 },
            { day: '2017-04', quantity: 7545 },
            { day: '2017-05', quantity: 543 },
            { day: '2017-06', quantity: 443 },
            { day: '2017-07', quantity: 246 },
            { day: '2017-08', quantity: 5445 },
            { day: '2017-09', quantity: 1154 },
            { day: '2017-10', quantity: 448 },
            { day: '2017-11', quantity: 1545 },
            { day: '2017-12', quantity: 4585 },
            { day: '2018-01', quantity: 1520 },
            { day: '2018-02', quantity: 9015 },
            { day: '2018-03', quantity: 632 },
            { day: '2018-04', quantity: 745 },
            { day: '2018-05', quantity: 343 },
            { day: '2018-06', quantity: 6443 },
            { day: '2018-07', quantity: 546 },
            { day: '2018-08', quantity: 1545 },
            { day: '2018-09', quantity: 1354 },
            { day: '2018-10', quantity: 848 },
            { day: '2018-11', quantity: 2155 },
            { day: '2018-12', quantity: 4585 },
            { day: '2019-01', quantity: 1540 },
            { day: '2019-02', quantity: 905 },
            { day: '2019-03', quantity: 632 },
            { day: '2019-04', quantity: 745 },
            { day: '2019-05', quantity: 3543 },
            { day: '2019-06', quantity: 4443 },
            { day: '2019-07', quantity: 2546 },
            { day: '2019-08', quantity: 545 },
            { day: '2019-09', quantity: 154 },
            { day: '2019-10', quantity: 4848 },
            { day: '2019-11', quantity: 155 },
            { day: '2019-12', quantity: 4585 }
          ]
        }
      }, {
        name: '饼图',
        callback: function(opt) {
          d3.setPie(opt)
        },
        opt: {
          container: null,
          containerWidth: 1200,
          containerHeight: 600,
          margin: {
            top: 80,
            right: 20,
            bottom: 80,
            left: 60
          },
          title: '昆明市人口年龄结构',
          data: [
            { age: '<5', population: 2704659 },
            { age: '5-13', population: 4499890 },
            { age: '14-17', population: 2159981 },
            { age: '18-24', population: 3853788 },
            { age: '25-44', population: 14106543 },
            { age: '45-64', population: 8819342 },
            { age: '≥65', population: 612463 }
          ]
        }
      }, {
        name: '圆环图',
        callback: function(opt) {
          d3.setRing(opt)
        },
        opt: {
          container: null,
          containerWidth: 1200,
          containerHeight: 600,
          margin: {
            top: 80,
            right: 20,
            bottom: 80,
            left: 60
          },
          title: '昆明市人口年龄结构',
          data: [
            { age: '<5', population: 2704659 },
            { age: '5-13', population: 4499890 },
            { age: '14-17', population: 2159981 },
            { age: '18-24', population: 3853788 },
            { age: '25-44', population: 14106543 },
            { age: '45-64', population: 8819342 },
            { age: '≥65', population: 612463 }
          ]
        }
      }, {
        name: '线状图',
        callback: function(opt) {
          d3.setLine(opt)
        },
        opt: {
          container: null,
          containerWidth: 1200,
          containerHeight: 600,
          margin: {
            top: 80,
            right: 20,
            bottom: 80,
            left: 60
          },
          title: '昆明市昨天PM2.5及空气质量指数(AQI)',
          data: [
            { time: '00:00', pm25: 75 },
            { time: '01:00', pm25: 66 },
            { time: '02:00', pm25: 43 },
            { time: '03:00', pm25: 32 },
            { time: '04:00', pm25: 20 },
            { time: '05:00', pm25: 18 },
            { time: '06:00', pm25: 16 },
            { time: '07:00', pm25: 33 },
            { time: '08:00', pm25: 53 },
            { time: '09:00', pm25: 66 },
            { time: '10:00', pm25: 55 },
            { time: '11:00', pm25: 67 },
            { time: '12:00', pm25: 99 },
            { time: '13:00', pm25: 138 },
            { time: '14:00', pm25: 110 },
            { time: '15:00', pm25: 99 },
            { time: '16:00', pm25: 119 },
            { time: '17:00', pm25: 125 },
            { time: '18:00', pm25: 173 },
            { time: '19:00', pm25: 168 },
            { time: '20:00', pm25: 162 },
            { time: '21:00', pm25: 143 },
            { time: '22:00', pm25: 132 },
            { time: '23:00', pm25: 87 }
          ]
        }
      }, {
        name: '散点图',
        callback: function(opt) {
          d3.setPoints(opt)
        },
        opt: {
          container: null,
          containerWidth: 1200,
          containerHeight: 600,
          margin: {
            top: 80,
            right: 20,
            bottom: 80,
            left: 60
          },
          title: '激烈运动年限与健康指数之间的关系抽样检查',
          data: [
            [5, 66],
            [7, 55],
            [4, 99],
            [11, 78],
            [28, 65],
            [7, 88],
            [5, 56],
            [2, 60],
            [4, 57],
            [6, 98],
            [27, 33],
            [26, 77],
            [23, 95],
            [34, 87],
            [7, 68],
            [1, 68],
            [2, 60],
            [22, 84],
            [6, 96],
            [13, 87]
          ]
        }
      }, {
        name: '弦图',
        callback: function(opt) {
          d3.setChord(opt)
        },
        opt: {
          container: null,
          containerWidth: 1200,
          containerHeight: 600,
          margin: {
            top: 80,
            right: 20,
            bottom: 80,
            left: 60
          },
          title: '激烈运动年限与健康指数之间的关系抽样检查',
          data: {
            names: ['北京', '上海', '广州', '深圳'],
            matrix: [
              [11975, 5871, 8916, 2868],
              [1951, 10048, 2060, 6171],
              [8010, 16145, 8090, 8045],
              [1013, 990, 940, 6907]
            ]
          }
        }
      }, {
        name: '树状图',
        callback: function(opt) {
          d3.setDendrogram(opt)
        },
        opt: {
          container: null,
          containerWidth: 1200,
          containerHeight: 600,
          margin: {
            top: 80,
            right: 120,
            bottom: 40,
            left: 120
          },
          data: {
            name: '某企业人员组织',
            children: [{
              name: '财务部',
              children: [
                {
                  name: '财务总监',
                  children: [
                    { name: '刘备' }
                  ]
                },
                {
                  name: '财务职员',
                  children: [
                    { name: '关羽' },
                    { name: '张飞' },
                    { name: '赵云' },
                    { name: '黄忠' },
                    { name: '马超' }
                  ]
                }
              ]
            },
            {
              name: '开发部',
              children: [
                {
                  name: '开发总监',
                  children: [
                    { name: '曹操' }
                  ]
                },
                {
                  name: '开发人员',
                  children: [
                    { name: '曹洪' },
                    { name: '夏侯淳' },
                    { name: '夏侯渊' },
                    { name: '张辽' },
                    { name: '许褚' }
                  ]
                }
              ]
            },
            {
              name: '人事部',
              children: [
                {
                  name: '人事总监',
                  children: [
                    { name: '孙权' }
                  ]
                },
                {
                  name: '人事专员',
                  children: [
                    { name: '周瑜' },
                    { name: '黄盖' },
                    { name: '吕蒙' },
                    { name: '陆轩' }
                  ]
                }
              ]
            }]
          }
        }
      }, {
        name: '打包图',
        callback: function(opt) {
          d3.setPack(opt)
        },
        opt: {
          container: null,
          containerWidth: 1200,
          containerHeight: 600,
          margin: {
            top: 80,
            right: 80,
            bottom: 80,
            left: 120
          },
          data: {
            name: '中国一线城市',
            children: [{
              name: '北京',
              children: [{
                name: '东城区',
                value: 42
              },
              {
                name: '朝阳区',
                value: 71
              },
              {
                name: '海淀区',
                value: 31
              }]
            }, {
              name: '上海',
              children: [
                {
                  name: '黄浦区',
                  value: 20
                },
                {
                  name: '徐汇区',
                  value: 55
              }]
            }]
          }
        }
      }, {
        name: '雷达图',
        callback: function(opt) {
          d3.setRadar(opt)
        },
        opt: {
          container: null,
          containerWidth: 1200,
          containerHeight: 700,
          margin: {
            top: 80,
            right: 80,
            bottom: 60,
            left: 60
          },
          title: '六年级期中考试平均成绩(分), 满分100分',
          data: [{
            className: '甲班',
            chinese: 88,
            math: 92,
            physics: 90,
            chemistry: 88,
            english: 95
          },
          {
            className: '乙班',
            chinese: 67,
            math: 78,
            physics: 80,
            chemistry: 72,
            english: 74
          },
          {
            className: '丙班',
            chinese: 77,
            math: 83,
            physics: 68,
            chemistry: 69,
            english: 65
          },
          {
            className: '丁班',
            chinese: 72,
            math: 67,
            physics: 62,
            chemistry: 67,
            english: 68
          }]
        }  
      }, {
        name: '力导向图',
        callback: function(opt) {
          d3.setForce(opt)
        },
        opt: {
          container: null,
          containerWidth: 1200,
          containerHeight: 700,
          margin: {
            top: 60,
            right: 60,
            bottom: 60,
            left: 60
          },
          title: '人物关系图',
          data: {
             nodes: [
              { name: 'A人物' },
              { name: 'B人物' },
              { name: 'C人物' },
              { name: 'D人物' },
              { name: 'E人物' },
              { name: 'F人物' },
              { name: 'G人物' },
              { name: 'H人物' },
              { name: 'I人物' },
              { name: 'J人物' },
              { name: 'K人物' },
              { name: 'L人物' },
              { name: 'M人物' }
            ],
            edges: [
              // value越小关系越近
              { source: 0, target: 1, relation: '朋友', value: 3 },
              { source: 0, target: 2, relation: '朋友', value: 3 },
              { source: 0, target: 3, relation: '朋友', value: 6 },
              { source: 1, target: 2, relation: '朋友', value: 6 },
              { source: 1, target: 3, relation: '朋友', value: 7 },
              { source: 2, target: 3, relation: '朋友', value: 7 },
              { source: 0, target: 4, relation: '朋友', value: 3 },
              { source: 0, target: 5, relation: '朋友', value: 3 },
              { source: 4, target: 5, relation: '夫妻', value: 1 },
              { source: 0, target: 6, relation: '兄弟', value: 2 },
              { source: 4, target: 6, relation: '同学', value: 3 },
              { source: 5, target: 6, relation: '同学', value: 3 },
              { source: 4, target: 7, relation: '同学', value: 4 },
              { source: 5, target: 7, relation: '同学', value: 3 },
              { source: 6, target: 7, relation: '同学', value: 3 },
              { source: 4, target: 8, relation: '师生', value: 4 },
              { source: 5, target: 8, relation: '师生', value: 5 },
              { source: 6, target: 8, relation: '师生', value: 3 },
              { source: 7, target: 8, relation: '师生', value: 5 },
              { source: 8, target: 9, relation: '师生', value: 4 },
              { source: 3, target: 9, relation: '师生', value: 5 },
              { source: 2, target: 10, relation: '母子', value: 1 },
              { source: 10, target: 11, relation: '雇佣', value: 6 },
              { source: 10, target: 12, relation: '雇佣', value: 6 },
              { source: 11, target: 12, relation: '同事', value: 7 }
            ]
          }
        }  
      }, {
        name: '泰森多边形',
        callback: function(opt) {
          d3.setVoronoi(opt)
        },
        opt: {
          container: null,
          containerWidth: 1200,
          containerHeight: 1000,
          margin: {
            top: 80,
            right: 20,
            bottom: 30,
            left: 60
          },
          title: '泰森多边形/冯洛诺伊图(Voronoi diagram)',
          data: {}
        }  
      }, {
        name: '竖向业务合作伙伴图',
        callback: function(opt) {
          d3.setVerticalBP(opt)
        },
        opt: {
          container: null,
          containerWidth: 1200,
          containerHeight: 1000,
          margin: {
            top: 80,
            right: 80,
            bottom: 30,
            left: 60
          },
          title: '手机2019年全国各渠道出货量数据汇总',
          data: [['天猫', '上海', 15216, 25216],
            ['线下自营店', '上海', 11278, 13244],
            ['苏宁易购', '上海', 27, 24],
            ['网上自营店', '上海', 27648, 35411],
            ['线下代理商', '上海', 1551, 1545],
            ['京东', '上海', 22141, 25441],
            ['天猫', '广州', 15453, 15353],
            ['线下自营店', '广州', 24683, 24623],
            ['苏宁易购', '广州', 1862, 654],
            ['线下代理商', '广州', 16228, 13228],
            ['天猫', '北京', 15001, 18001],
            ['线下自营店', '北京', 15001, 1654],
            ['苏宁易购', '北京', 5001, 6541],
            ['网上自营店', '北京', 28648, 29648],
            ['线下代理商', '北京', 9648, 9648],
            ['天猫', '深圳', 3313, 541],
            ['线下自营店', '深圳', 22396, 24396],
            ['苏宁易购', '深圳', 3362, 3762],
            ['网上自营店', '深圳', 22396, 21396],
            ['线下代理商', '深圳', 2473, 2973],
            ['京东', '深圳', 16541, 11541],
            ['苏宁易购', '杭州', 3541, 3599],
            ['线下代理商', '杭州', 3541, 8741],
            ['京东', '杭州', 3654, 9874],
            ['天猫', '武汉', 1738, 110],
            ['线下自营店', '武汉', 12925, 13],
            ['苏宁易购', '武汉', 15413, 0],
            ['线下自营店', '重庆', 2166, 654],
            ['苏宁易购', '重庆', 2286, 3654],
            ['网上自营店', '重庆', 348, 3654],
            ['线下代理商', '重庆', 4244, 3654],
            ['京东', '重庆', 1536, 1654],
            ['线下自营店', '长沙', 351, 654],
            ['线下代理商', '长沙', 405, 541],
            ['线下自营店', '成都', 914, 654],
            ['苏宁易购', '成都', 127, 354],
            ['线下代理商', '成都', 1470, 654],
            ['京东', '成都', 516, 354],
            ['天猫', '东莞', 43, 0],
            ['线下自营店', '东莞', 667, 654],
            ['苏宁易购', '东莞', 172, 354],
            ['网上自营店', '东莞', 149, 541],
            ['线下代理商', '东莞', 1380, 3254],
            ['京东', '东莞', 791, 754],
            ['线下自营店', '苏州', 541, 687],
            ['线下代理商', '苏州', 654, 541],
            ['线下自营店', '南京', 1070, 654],
            ['线下代理商', '南京', 1171, 1541],
            ['京东', '南京', 33, 45],
            ['线下自营店', '佛山', 407, 654],
            ['苏宁易购', '佛山', 541, 874],
            ['线下代理商', '佛山', 457, 674],
            ['京东', '佛山', 541, 365],
            ['线下自营店', '天津', 557, 654],
            ['苏宁易购', '天津', 167, 541],
            ['网上自营店', '天津', 95, 100],
            ['线下代理商', '天津', 1090, 1321],
            ['京东', '天津', 676, 541],
            ['天猫', '合肥', 1195, 1654],
            ['线下自营店', '合肥', 5412, 6541],
            ['苏宁易购', '合肥', 212, 241],
            ['线下代理商', '合肥', 1509, 1654],
            ['天猫', '温州', 3899, 389],
            ['线下自营店', '温州', 147, 321],
            ['苏宁易购', '温州', 455, 541],
            ['网上自营店', '温州', 321, 254],
            ['线下代理商', '温州', 4100, 4512],
            ['天猫', '南宁', 123, 133],
            ['线下自营店', '南宁', 634, 654],
            ['苏宁易购', '南宁', 749, 541],
            ['网上自营店', '南宁', 119, 654],
            ['线下代理商', '南宁', 3705, 4574],
            ['京东', '南宁', 3456, 4000],
            ['线下自营店', '厦门', 828, 1201],
            ['苏宁易购', '厦门', 2808, 3541],
            ['网上自营店', '厦门', 1452, 2000],
            ['线下代理商', '厦门', 2625, 1541],
            ['京东', '厦门', 1920, 1234],
            ['线下自营店', '西安', 1146, 1541],
            ['苏宁易购', '西安', 212, 321],
            ['网上自营店', '西安', 223, 241],
            ['线下代理商', '西安', 1803, 2000],
            ['京东', '西安', 761, 465],
            ['线下自营店', '长春', 527, 654],
            ['苏宁易购', '长春', 90, 120],
            ['线下代理商', '长春', 930, 1241],
            ['京东', '长春', 395, 410],
            ['天猫', '哈尔滨', 7232, 8451],
            ['线下自营店', '哈尔滨', 1272, 2141],
            ['苏宁易购', '哈尔滨', 1896, 3541],
            ['网上自营店', '哈尔滨', 200, 1241],
            ['线下代理商', '哈尔滨', 10782, 15412],
            ['京东', '哈尔滨', 1911, 2000],
            ['线下自营店', '青岛', 495, 521],
            ['苏宁易购', '青岛', 432, 541],
            ['网上自营店', '青岛', 241, 320],
            ['线下代理商', '青岛', 1557, 1600],
            ['京东', '青岛', 24, 30],
            ['线下自营店', '沈阳', 460, 541],
            ['网上自营店', '沈阳', 88, 99],
            ['线下代理商', '沈阳', 956, 365],
            ['线下自营店', '济南', 232, 365],
            ['苏宁易购', '济南', 71, 99],
            ['线下代理商', '济南', 575, 654],
            ['京东', '济南', 368, 354]
          ]
        }  
      }],
    }
  },
  methods: {
    setCharts(item, index) {
      const self = this
      let containerClass = `.container-${index}`
      this.active = index
      setTimeout(function(){
        item.opt.container = document.querySelector(containerClass)
        item.opt.containerWidth = self.$refs.wrap.offsetWidth
        item.opt.primary = self.primary
        item.callback(item.opt)
      }, 200)
    }
  },
  created() {

  },
  mounted() {
    const self = this
    const index = 0
    const container = '.container-' + index
    self.active = index
    setTimeout(function(){
      self.list[index].opt.container = document.querySelector(container)
      self.list[index].opt.containerWidth = self.$refs.wrap.offsetWidth
      self.list[index].opt.primary = self.primary
      d3.setBar(self.list[index].opt)
    },200)
  }
}
</script>

<style scoped lang="scss">
@import "../../styles/_global.scss";
.chart {

  &-btn {
    .el-button {
      margin-bottom:10px;
    }
  }
} 
</style>
