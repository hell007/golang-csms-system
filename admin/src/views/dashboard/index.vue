<template>
  <el-card class="xa-block" shadow="hover">
    <div class="xa-blcok">
      <el-row class="xa-block" :gutter="20">
        <el-col :span="6">
          <el-card shadow="hover">
            <div class="home-card">
              <div class="home-card-icon">
                <img
                  src="../../assets/images/index/u1.png"
                />
              </div>
              <div class="home-card-text">
                <h5>今日商品支付订单数</h5>
                <b>200</b>
              </div>
            </div>
          </el-card>
        </el-col>
        <el-col :span="6">
          <el-card shadow="hover">
            <div class="home-card">
              <div class="home-card-icon">
                <img
                  src="../../assets/images/index/u2.png"
                />
              </div>
              <div class="home-card-text">
                <h5>今日支付总额</h5>
                <b>¥5000.00</b>
              </div>
            </div>
          </el-card>
        </el-col>
        <el-col :span="6">
          <el-card shadow="hover">
            <div class="home-card">
              <div class="home-card-icon">
                <img
                  src="../../assets/images/index/u3.png"
                />
              </div>
              <div class="home-card-text">
                <h5>今日商品客单价</h5>
                <b>¥50.00</b>
              </div>
            </div>
          </el-card>
        </el-col>
        <el-col :span="6">
          <el-card shadow="hover">
            <div class="home-card">
              <div class="home-card-icon">
                <img
                  src="../../assets/images/index/u4.png"
                />
              </div>
              <div class="home-card-text">
                <h5>今日业务办理数</h5>
                <b>500</b>
              </div>
            </div>
          </el-card>
        </el-col>
      </el-row>
      <el-row class="xa-block" :gutter="20">
        <el-col :span="6">
          <el-card shadow="hover">
            <div class="home-card">
              <div class="home-card-icon">
                <img
                  src="../../assets/images/index/u5.png"
                />
              </div>
              <div class="home-card-text">
                <h5>今日总访客数</h5>
                <b>200</b>
              </div>
            </div>
          </el-card>
        </el-col>
        <el-col :span="6">
          <el-card shadow="hover">
            <div class="home-card">
              <div class="home-card-icon">
                <img
                  src="../../assets/images/index/u6.png"
                />
              </div>
              <div class="home-card-text">
                <h5>今日总流量</h5>
                <b>50000</b>
              </div>
            </div>
          </el-card>
        </el-col>
        <el-col :span="6">
          <el-card shadow="hover">
            <div class="home-card">
              <div class="home-card-icon">
                <img
                  src="../../assets/images/index/u7.png"
                />
              </div>
              <div class="home-card-text">
                <h5>店铺总数</h5>
                <b>50</b>
              </div>
            </div>
          </el-card>
        </el-col>
        <el-col :span="6">
          <el-card shadow="hover">
            <div class="home-card">
              <div class="home-card-icon">
                <img
                  src="../../assets/images/index/u8.png"
                />
              </div>
              <div class="home-card-text">
                <h5>本月店铺活跃</h5>
                <b>50</b>
              </div>
            </div>
          </el-card>
        </el-col>
      </el-row>
    </div>

    <div class="home-panel" hidden>
      <div class="home-panel-hd">
        <h5>待处理事务</h5>
      </div>
      <div class="home-panel-bd">
        <ul class="home-todo xa-clearfix">
          <li class="home-todo-item">
            <a href="#">待付款订单</a>
            <span>
              (
              <em>10</em>
              )
            </span>
          </li>
          <li class="home-todo-item">
            <a href="#">待处理订单</a>
            <span>
              (
              <em>10</em>
              )
            </span>
          </li>
          <li class="home-todo-item">
            <a href="#">广告位即将到期</a>
            <span>
              (
              <em>10</em>
              )
            </span>
          </li>
          <li class="home-todo-item">
            <a href="#">待发货订单</a>
            <span>
              (
              <em>10</em>
              )
            </span>
          </li>
          <li class="home-todo-item">
            <a href="#">待处理售后申请</a>
            <span>
              (
              <em>10</em>
              )
            </span>
          </li>
        </ul>
      </div>
    </div>

    <div class="home-panel">
      <div class="home-panel-hd">
        <h5>运营快捷入口</h5>
      </div>
      <div class="home-panel-bd">
        <div class="home-menu">
          <a href="#/goods/add" class="home-menu-item">
            <div class="home-menu-icon">
              <img
                src="../../assets/images/index/m1.png"
              />
            </div>
            <span>商品发布</span>
          </a>
          <a href="#/goods/list" class="home-menu-item">
            <div class="home-menu-icon">
              <img
                src="../../assets/images/index/m2.png"
              />
            </div>
            <span>商品列表</span>
          </a>
          <a href="#/order/list" class="home-menu-item">
            <div class="home-menu-icon">
              <img
                src="../../assets/images/index/m3.png"
              />
            </div>
            <span>订购单列表</span>
          </a>
          <a href="#/member/list" class="home-menu-item">
            <div class="home-menu-icon">
              <img
                src="../../assets/images/index/m7.png"
              />
            </div>
            <span>会员管理</span>
          </a>
          <a href="#/permission/list" class="home-menu-item">
            <div class="home-menu-icon">
              <img
                src="../../assets/images/index/m4.png"
              />
            </div>
            <span>菜单管理</span>
          </a>
          <a href="#/user/list" class="home-menu-item">
            <div class="home-menu-icon">
              <img
                src="../../assets/images/index/m6.png"
              />
            </div>
            <span>管理员管理</span>
          </a>
        </div>
      </div>
    </div>

    <div class="home-panel">
      <div class="home-panel-hd">
        <h5>订单统计</h5>
      </div>
      <div class="home-panel-bd">
        <div class="home-chart">
          <div class="home-chart-label">
            <div class="home-chart-item">
              <h6>本月订单总数</h6>
              <b>10000</b>
              <dl class="up">
                <dt>10%</dt>
                <dd>同比上周</dd>
              </dl>
            </div>
            <div class="home-chart-item">
              <h6>本周订单数量</h6>
              <b>10000</b>
              <dl class="down">
                <dt>10%</dt>
                <dd>同比上周</dd>
              </dl>
            </div>
          </div>
          <div class="home-chart-line">
            <div class="xa-clearfix">
              <div class="xa-pull-left">近一周订单统计</div>
              <div class="xa-pull-right">
                <span class="btn">今日</span>
                <span
                  :class="`btn ${active==1 ? 'on' : ''}`"
                  @click="changeChart(1)"
                >
                  本周
                </span>
                <span
                  :class="`btn ${active==2 ? 'on' : ''}`"
                  @click="changeChart(2)"
                >
                  本月
                </span>
                <el-date-picker
                  size="small"
                  style="width:260px;"
                  v-model="date"
                  type="daterange"
                  range-separator="至"
                  start-placeholder="开始日期"
                  end-placeholder="结束日期"
                ></el-date-picker>
              </div>
            </div>
            <ve-line :data="chartData"></ve-line>
          </div>
        </div>
      </div>
    </div>

  </el-card>
</template>
<script>
import {
  mapGetters
} from 'vuex'

import VeLine from 'v-charts/lib/line.common'

export default {
  name: 'dashboard',
  components: {
    VeLine
  },
  data() {
    return {
      date: '',
      active: 1,
      chartData: {},
      chart: {
        weeks: {
          columns: ['日期', '订单量'],
          rows: [
            { 日期: '07-01周六', 订单量: 123 },
            { 日期: '07-02周日', 订单量: 1223 },
            { 日期: '07-03周一', 订单量: 2123 },
            { 日期: '07-04周二', 订单量: 4123 },
            { 日期: '07-05周三', 订单量: 3123 },
            { 日期: '07-06周四', 订单量: 7123 },
            { 日期: '07-07周五', 订单量: 123 }
          ]
        },
        months: {
          columns: ['日期', '订单量'],
          rows: [
            { 日期: '07-01', 订单量: 123 },
            { 日期: '07-02', 订单量: 1223 },
            { 日期: '07-03', 订单量: 2123 },
            { 日期: '07-04', 订单量: 4123 },
            { 日期: '07-05', 订单量: 3123 },
            { 日期: '07-06', 订单量: 7123 },
            { 日期: '07-07', 订单量: 123 },
            { 日期: '07-08', 订单量: 3123 },
            { 日期: '07-09', 订单量: 7123 }
          ]
        }
      }
    }
  },
  methods: {
    changeChart(type) {
      var app = this
      switch (type) {
        case 1:
          app.chartData = app.chart.weeks
          app.active = 1
          break
        case 2:
          app.chartData = app.chart.months
          app.active = 2
          break
      }
    }
  },
  created() {
    this.chartData = this.chart.weeks
  }
}
</script>
<style rel="stylesheet/scss" lang="scss">
@import '../../styles/home.scss';
</style>

