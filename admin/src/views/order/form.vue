<template>
  <el-card class="xa-block" shadow="hover">
    <div class="xa-content xa-block">
      <el-row>
        <el-col :span="18">
          <el-steps :active="2">
            <el-step
              title="步骤1"
              description="下单"></el-step>
            <el-step
              title="步骤2"
              description="支付"></el-step>
            <el-step
              title="步骤3"
              description="发货"></el-step>
            <el-step
              title="步骤4"
              description="售后"></el-step>
          </el-steps>
        </el-col>
        <el-col :span="6" class="xa-block">
          <div class="xa-block--align-right">
            <el-button 
              icon="el-icon-back"
              size="small" 
              v-waves
              type="primary"
              @click="handleCancle">返回</el-button>
            <el-button 
              v-if="order.orderState == 4"
              icon="el-icon-delete"
              size="small" 
              v-waves
              type="danger">删除订单</el-button>
          </div>
        </el-col>
      </el-row>

      <el-form class="p-form p-form--condensed"
        ref="postForm"
        :rules="formRules"
        :model="order"
        label-suffix="："
        label-width="120px">
        <div class="p-form__section">
          <div class="p-form__s-head">订单信息</div>
          <el-row class="xa-block" :gutter="10">
            <el-col :span="8">
              <div class="p-form__info">
                <el-form-item label="支付流水">
                  {{order.ordersn}}
                </el-form-item>
              </div>
              <div class="p-form__info">
                <el-form-item label="交易时间">
                  {{order.createTime | parseTime}}
                </el-form-item>
              </div>
              <div class="p-form__info">
                <el-form-item label="支付时间">
                  {{order.payTime | parseTime}}
                </el-form-item>
              </div>
              <div class="p-form__info">
                <el-form-item label="发货时间">
                  {{order.shipTime | parseTime}}
                </el-form-item>
              </div>
              <div class="p-form__info">
                <el-form-item label="收货时间">
                  {{order.confirmTime | parseTime}}
                </el-form-item>
              </div>
            </el-col>
            <el-col :span="8">
              <el-form-item label="订单状态">
                <span
                  class="xa-text-primary"
                  v-for="item in options.states"
                  :key="`state-${item.value}`">
                  {{order.orderState == item.value ?  item.label : ''}}
                </span>
              </el-form-item>
              <el-form-item label="下单人">
                {{member.name}}
              </el-form-item>
              <el-form-item label="电话">
               {{member.mobile}}
              </el-form-item>
              <el-form-item label="订单备注">
                <div class="p-form__display">
                  <span class="xa-text-primary">{{order.remark}}</span>
                </div>
              </el-form-item>
            </el-col>
            <el-col :span="8">
              <el-form-item label="收货人">
                {{order.consignee}}
              </el-form-item>
              <el-form-item label="收货人电话">
                {{order.phone}} {{order.tell}}
              </el-form-item>
              <el-form-item label="收货人地址">
                <div class="p-form__display">
                  {{order.province}}{{order.city}}{{order.district}}{{order.address}}
                </div>
              </el-form-item>
              <el-form-item label="物流公司">
                {{order.shipName}}
                <el-button
                  v-if="order.orderState == 2"
                  type="primary"
                  size="small"
                  icon="el-icon-s-tools"
                  v-waves
                  @click="handleExpress">选择</el-button>
              </el-form-item>
              <el-form-item label="物流单号">
                {{order.shipNo}}
              </el-form-item>
            </el-col>
          </el-row>
        </div>
        
        <!-- 商品信息 -->
        <hr class="xa-separator" />
        <div class="p-form__section">
          <div class="p-form__s-head">商品信息</div>
          <div class="xa-block">
            <el-table
              class="p-table"
              :data="list"
              border
              ref="table-product"
              size="small">
              <el-table-column 
                label="商品信息" 
                header-align="center">
                <template slot-scope="scope">
                  <div class="p-info-block">
                    <div class="p-info-block__pic">
                      <div
                        class="p-avatar"
                        style="width: 60px; height: 60px;">
                        <img
                          class="p-avatar__img"
                          v-if="scope.row.small"
                          :src="links + scope.row.small"/>
                      </div>
                    </div>
                    <div class="p-info-block__main">
                      <div>编号：{{scope.row.goodsSn}}</div>
                      <div>名称：{{scope.row.goodsName}}</div>
                    </div>
                  </div>
                </template>
              </el-table-column>
              <el-table-column 
                label="价格/规格" 
                header-align="center">
                <template slot-scope="scope">
                  <div>价格：&yen;{{scope.row.goodsPrice}}</div>
                  <div>等级：{{scope.row.goodsSku}}</div>
                  <div>规格：{{scope.row.unit}}</div>
                </template>
              </el-table-column>
              <el-table-column
                label="数量(扎)"
                prop="goodsNum"
                width="160"
                align="center">
              </el-table-column>
              <el-table-column 
                label="小计(元)" 
                width="180" 
                align="center">
                <template slot-scope="scope">
                  &yen; {{scope.row.goodsPrice * scope.row.goodsNum}}
                </template>
              </el-table-column>
            </el-table>
          </div>
          <div class="p-product-info xa-text-right xa-block">
            <div class="p-product-info__item">
              <span class="xa-text-secondary">商品总价</span>
              <span class="p-product-info__num">
                &yen; {{order.totalPrice}}
              </span>
            </div>
            <div class="p-product-info__item">
              <span class="xa-text-secondary">运费(快递)</span>
              <span class="p-product-info__num">
                &yen; {{order.shipPrice}}
              </span>
            </div>
            <div class="p-product-info__item">
              <span class="xa-text-secondary">订单总价</span>
              <span class="p-product-info__num">
                &yen; {{order.orderPrice}}
              </span>
            </div>
            <div class="p-product-info__item">
              <span>实际支付</span>
              <span class="xa-text-danger p-product-info__num">
                &yen; {{order.payPrice}}
              </span>
            </div>
          </div>
        </div>

        <hr class="xa-separator" />

        <div class="p-form__section" hidden>
          <div class="p-form__s-head">操作记录</div>
          <div class="p-form__rich">
            <el-table
              class="p-table"
              :data="records"
              border
              ref="table-records"
              size="mini">
              <el-table-column
                prop=""
                label="操作者"
                width="100"
                align="center">
              </el-table-column>
              <el-table-column
                prop=""
                label="操作时间"
                width="200"
                align="center">
              </el-table-column>
              <el-table-column
                prop=""
                label="订单状态"
                width="100"
                align="center">
              </el-table-column>
              <el-table-column
                prop=""
                label="备注"
                align="center"> 
              </el-table-column>
            </el-table>
          </div>
        </div>
      </el-form>
    </div>
    
    <!-- 发货 -->
    <el-dialog
      :title="dialog.title"
      :visible.sync="dialog.visible"
      class="xa-dialog xa-dialog--sm">
      <el-form
        class="p-form"
        ref="postForm"
        :rules="formRules"
        :model="order"
        label-width="120px"
        label-suffix="："
        size="small">
        <div class="p-form__section">
          <div class="p-form__s-head">订单信息</div>
          <el-form-item label="订单号">
            {{order.ordersn}}
          </el-form-item>
          <el-form-item label="商品信息">
            <ul>
              <li class="xa-block" 
                v-for="item,index in list" >
                <div class="p-info-block">
                  <div class="xa-info-block__pic">
                    <div class="p-avatar" style="width:40px;height:40px;">
                      <img
                        class="p-avatar__img"
                        v-if="item.small"
                        :src="links + item.small"/>
                    </div>
                  </div>
                  <div class="p-info-block__main">
                    <div>名称：{{item.goodsName}}</div>
                    <div class="xa-text-gray">等级/数量：{{item.goodsSku}} x {{item.goodsNum}}</div>
                    <div class="xa-text-gray">
                      <span>价格：&yen;{{item.goodsPrice}}</span>&nbsp;&nbsp;
                      <span>小计：&yen;{{item.goodsPrice * item.goodsNum}}</span>
                    </div>
                  </div>
                </div>
              </li>
            </ul>
          </el-form-item>
          <el-form-item label="收货人信息">
            <div class="p-form__display">
              <div>姓名：{{order.consignee}} &nbsp;&nbsp; 手机：{{order.phone}} {{order.tell}}</div>
              <div class="xa-text-secondary">
                地址：{{order.province + order.city + order.district + order.address}}
              </div>
            </div>
          </el-form-item>
          <el-form-item label="备注">
            <div class="p-form__display">
              {{order.remark}}
            </div>
          </el-form-item>
        </div>
        <hr class="xa-separator" />
        <div class="p-form__section">
          <div class="p-form__s-head">物流服务商</div>
          <el-form-item 
            prop="shipName"
            label="物流服务商">
            <el-select v-model="order.shipName" >
              <el-option 
              v-for="item in options.ships"
              :key="`ship-${item.label}`"
              :value="item.label"
              :lable="item.label">
              </el-option>
            </el-select>
          </el-form-item>
          <el-form-item 
            prop="shipNo"
            label="物流单号">
            <el-input 
              v-model="order.shipNo" 
              placeholder="请输入">
            </el-input>
          </el-form-item>
          <el-form-item 
            prop="shipNote"
            label="发货备注">
            <el-input v-model="order.shipNote" type="textarea"></el-input>
          </el-form-item>
        </div>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button
          type="primary"
          :loading="dialog.processing"
          @click="handleSubmit">确定</el-button>
         <el-button 
          @click="dialog.visible = false">取消</el-button> 
      </span>
    </el-dialog>  
  </el-card>
</template>

<script>
import {
  mapActions
} from 'vuex'

import {
  URIS
} from '@/api/config'

import waves from '@/directive/waves/index' // 水波纹指令

export default {
  name: 'orderform',
  directives: {
    waves
  },
  components: {},
  data() {
    return {
      links: URIS.Goods,
      list: [],
      member: {},
      records: [],
      order: {
        shipName: '',
        shipNo: '',
        shipNote: ''
      },
      dialog: {
        visible: false,
        title: '发货',
        processing: false,
      },
      options: {
        states: [
          {value: 0, label: '全部'},
          {value: 1, label: '待支付'},
          {value: 2, label: '待发货'},
          {value: 3, label: '已发货'},
          {value: 4, label: '已取消'},
          {value: 5, label: '已完成'}
        ],
        ships: [
          {value: 1, label: '顺丰物流'},
          {value: 2, label: '京东物流'}
        ]
      },
      //验证规则
      formRules: {
        shipName: [{
          required: true,
          message: "请选择物流服务商",
          trigger: 'blur'
        }],
        shipNo: [{
          required: true,
          message: "请输入物流单号",
          trigger: 'blur'
        }, {
          min: 2,
          max: 20,
          message: "必须输入2-20个字符",
          trigger: 'blur'
        }]
      }
    }
  },
  computed: {
    isEdit() {
      this.order.ordersn = this.$route.query.id ? this.$route.query.id : null
      return this.order.ordersn
    }
  },
  methods: {
    ...mapActions(['getOrder', 'saveOrder']),
    getItem() {
      const self = this
      self.getOrder(self.order.ordersn).then(response => {
        const status = response.data.state
        const res = response.data.data
        const message = response.data.msg
        if (status) {
          self.order = res.order
          self.list = res.list
          self.member = res.member
          self.records = res.records
        } else {
          self.$notify({
            title: '失败',
            message: message,
            type: 'error'
          })
        } 
      })
    },
    handleExpress() {
      var dia = this.dialog
      dia.title = '发货'
      dia.visible = true
    },
    handleSubmit() {
      const self = this
      self.$refs.postForm.validate(valid => {
        if (valid) {
          let order = self.order
          self.dialog.processing = true
          self.$confirm(`订单号【${order.ordersn}】确定要发货?`, '提示', {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: 'warning'
          })
          .then(function(action) {
            self.saveOrder(order).then(response => {
              const status = response.data.state
              const res = response.data.data
              const message = response.data.msg
              if (status) {
                self.$notify({
                  title: '成功',
                  message: message,
                  type: 'success',
                  duration: 2000
                })
              } else {
                self.$notify({
                  title: '失败',
                  message: message,
                  type: 'error'
                })
              } 
              self.dialog.processing = false
            })
          })
          .catch(function(action) {
            self.dialog.processing = false
          })

        } else {
          self.$alert('请正确输入！', '提示', {
            confirmButtonText: '确定'
          });
          return false
        }
      })
    },
    handleCancle() {
      this.$router.push('/order/list')
    }
  },
  created() {
    if (this.isEdit) {
      this.getItem();
    }
  }
}
</script>