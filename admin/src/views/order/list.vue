<template>
  <el-card class="xa-block" shadow="hover">
    <!-- 筛选项-->
    <el-form
      class="p-filter" inline size="small"
      inline
      @submit.native.prevent
      label-suffix="：">
      <div class="p-filter__input">
        <el-form-item label="订单号">
          <el-input style="width:300px;"
            clearable
            placeholder="请输入订单号"
            @keyup.enter.native="handleFilter"
            v-model="listQuery.name"></el-input>
        </el-form-item>
        <el-form-item label="订单状态">
          <el-select 
            v-model="listQuery.orderState"
            clearable
            placeholder="请选择">
            <el-option
              v-for="item in options.states"
              :key="`states-${item.value}`"
              :label="item.label"
              :value="item.value"></el-option>
          </el-select>
        </el-form-item>
      </div>
      <div class="p-filter__action">
        <el-form-item>
          <el-button
            size="small"
            type="primary"
            icon="el-icon-search"
            v-waves
            @click="handleFilter">
            查询 
          </el-button>
          <el-button 
            size="small"
            v-waves
            @click="handleRest">
            重置
          </el-button>
        </el-form-item>
      </div>
    </el-form>
    
    <!-- 操作 -->
    <div class="xa-block xa-clearfix">  
      <div class="xa-pull-left">
        <span v-if="tip">
          <i class="el-icon-warning"></i>
          已选择<b>{{ mulSelection.length }}</b>项
        </span>
        <el-button 
          size="small" 
          type="success" 
          v-waves
          icon="el-icon-close"
          @click="handleStatus">
          取消
        </el-button>
        <el-button 
          size="small" 
          type="danger"
          icon="el-icon-delete"
          v-waves
          @click="handleBatchDelete">
          删除
        </el-button>
      </div>
      <div class="xa-pull-right">
        <el-button 
          size="small" 
          type="primary"
          icon="el-icon-download"
          v-waves
          @click="">
          导出报表
        </el-button>
      </div>
    </div>
    
    <!-- 数据 -->
    <div class="xa-block"
      v-loading="loading"
      element-loading-text="正在玩命加载中...">
      <el-table
        class="p-table"
        ref="listTable"
        :key="tableKey"
        :data="list"
        @selection-change="handleSelectionChange"
        tooltip-effect="dark"
        border
        fit
        highlight-current-row>
        <el-table-column
          fixed
          type="selection"
          width="48"
          align="center">
        </el-table-column>
        <el-table-column
          fixed
          prop="ordersn"
          label="订单号"
          width="180"
          align="center">
        </el-table-column>
        <el-table-column
          label="订单信息"
          min-width="230"
          align="center">
          <template slot-scope="scope">
            <div class="xa-text-left">
              <div class="xa-text-primary">订单状态：
                <span
                  v-for="item in options.states"
                  :key="`state-${item.value}`">
                  {{scope.row.orderState == item.value ?  item.label : ''}}
                </span>
              </div>
              <div>下单时间：{{scope.row.createTime | parseTime}}</div>          
              <div>付款时间：{{scope.row.payTime | parseTime}}</div>
              <div>发货时间：{{scope.row.shipTime | parseTime}}</div>
              <div>确认时间：{{scope.row.confirmTime | parseTime}}</div>
            </div>
          </template>
        </el-table-column>
        <el-table-column
          label="物流信息"
          min-width="260"
          align="center">
          <template slot-scope="scope">
            <div class="xa-text-left">
              <div>物流公司：{{scope.row.shipName}}</div>          
              <div>物流单号：{{scope.row.shipNo}}</div>
              <div>物流备注：{{scope.row.shipNote}}</div>
            </div>
          </template>
        </el-table-column>
        <el-table-column
          label="收货人信息"
          width="200"
          align="center">
          <template slot-scope="scope">
            <div class="xa-text-left">
              <div>收货人：{{scope.row.consignee}}</div>
              <div>联系方式：{{scope.row.phone}} {{scope.row.tell}}</div>
              <div>收货地址：{{scope.row.province}}{{scope.row.city}}{{scope.row.district}}{{scope.row.address}}</div>
            </div>
          </template>
        </el-table-column> 
        <el-table-column
          label="交易金额(元)"
          width="180"
          align="center">
          <template slot-scope="scope">
            <div class="xa-text-left">
              <div>商品总价：{{scope.row.totalPrice}}</div>
              <div>运费：{{scope.row.shipPrice}}</div>
              <div>订单总价：{{scope.row.orderPrice}}</div>
              <div class="xa-text-danger">支付总价：{{scope.row.payPrice}}</div>
              <div>支付方式：{{scope.row.payName}}</div>
            </div>
          </template>
        </el-table-column>
        <el-table-column
          label="操作"
          width="65"
          align="center"
          header-align="center">
          <template slot-scope="scope" v-if="scope.row.roleId != 1">
            <el-dropdown trigger="click">
              <span class="el-dropdown-link">
                <el-button
                  type="text"
                  size="mini"
                  icon="el-icon-more">
                </el-button>
              </span>
              <el-dropdown-menu slot="dropdown">
                <el-dropdown-item>
                  <i class="el-icon-edit" 
                    @click="handleEdit(scope.row.ordersn)">
                    编辑
                  </i>
                </el-dropdown-item>
                <el-dropdown-item>
                  <i
                    class="el-icon-delete"
                    @click="handleDelete(scope.row)">
                    删除
                  </i>
                </el-dropdown-item>
              </el-dropdown-menu>
            </el-dropdown>
          </template>
        </el-table-column>
      </el-table>
    </div>
    
    <!--分页-->
    <div class="xa-block">
      <Pagination 
        :getList="getList" 
        :listLoading="loading" 
        :listQuery="listQuery" 
        :total="total">
      </Pagination>
    </div>
  
  </el-card>
</template>

<script>
import waves from '@/directive/waves/index' // 水波纹指令
import Pagination from '../components/pagination'
import {
  mapActions
} from 'vuex'

export default {
  name: 'order',
  directives: {
    waves
  },
  components: {
    Pagination
  },
  data() {
    return {
      list: null,
      total: null,
      loading: true,
      tip: false,
      listQuery: {
        pageNumber: 1,
        pageSize: 10,
        name: '',
        orderState: 0
      },
      mulSelection: [],
      tableKey: 0,
      options: {
        states: [
          {value: 0, label: '全部'},
          {value: 1, label: '待支付'},
          {value: 2, label: '待发货'},
          {value: 3, label: '已发货'},
          {value: 4, label: '已取消'},
          {value: 5, label: '已完成'}
        ]
      }
    }
  },
  computed: {},
  methods: {
    ...mapActions(['getOrderList', 'deleteOrder', 'closeOrder']),
    //获取list数据，传入筛选对象
    getList() {
      const self = this
      self.loading = true
      self.getOrderList(self.listQuery).then(response => {
        const status = response.data.state
        const res = response.data.data
        const message = response.data.msg
        if (status) {
          self.list = res.rows
          self.total = res.count
        } else {
          self.$notify({
            title: '失败',
            message: message,
            type: 'error'
          })
        } 
        self.loading = false
      })
    },
    //查询
    handleFilter() {
      this.getList()
    },
    //重置
    handleRest() {
      this.listQuery = {
        pageNumber: 1,
        pageSize: 10,
        name: '',
        orderState: 0
      }
    },
    //编辑
    handleEdit(ordersn) {
      this.$router.push('/order/form?id='+ordersn);
    },
    //删除
    handleDelete(row) {
      const self = this
      let rows = []
      rows.push(row)

      self.$confirm(`确定要删除订单【${row.ordersn}】 ?`, '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
      .then(function(action) {
        self.deleteOrder(rows).then(response => {
          const status = response.data.state
          const message = response.data.msg
          if (status) {
            self.$notify({
              title: '成功',
              message: message,
              type: 'success',
              duration: 2000
            })
          } else {
            this.$notify({
              title: '失败',
              message: message,
              type: 'error'
            })
          }
        })
      })
      .catch(function(action) {})
    },
    //获取选中的row
    handleSelectionChange(rows) {
      this.tip = true
      this.mulSelection = rows
    },
    //批量删除
    handleBatchDelete() {
      const self = this 
      if (self.mulSelection.length === 0) return

      self.$confirm('确定要批量删除订单?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
      .then(function(action) {
        self.deleteOrder(self.mulSelection).then(response => {
          const status = response.data.state
          const message = response.data.msg
          if (status) {
            self.$notify({
              title: '成功',
              message: message,
              type: 'success',
              duration: 2000
            })
          } else {
            this.$notify({
              title: '失败',
              message: message,
              type: 'error'
            })
          }
        })
      })
      .catch(function(action) {})  
    },
    handleStatus() {
      const self = this 
      if (self.mulSelection.length === 0) return

      self.$confirm(`确定要取消订单?`, '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
      .then(function(action) {  
        self.closeOrder(self.mulSelection).then(response => {
          const status = response.data.state
          const message = response.data.msg
          if (status) {
            self.$notify({
              title: '成功',
              message: message,
              type: 'success',
              duration: 2000
            })
            self.getList() // 刷新
          } else {
            this.$notify({
              title: '失败',
              message: message,
              type: 'error'
            })
          }
        })
      })
      .catch(function(action) {}) 
    }
  },
  created() {
    this.getList()
  }
}
</script>
