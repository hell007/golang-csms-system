<template>
  <el-card class="xa-block" shadow="hover">
    <!-- 筛选项-->
    <el-form
      class="p-filter" inline size="small"
      inline
      @submit.native.prevent
      label-suffix="：">
      <div class="p-filter__input">
        <el-form-item label="商品信息">
          <el-input style="width:300px;"
            clearable
            placeholder="请输入名称/编码"
            @keyup.enter.native="handleFilter"
            v-model="listQuery.name"></el-input>
        </el-form-item>
        <el-form-item label="商品分类">
          <el-select 
            v-model="listQuery.categoryId"
            clearable
            placeholder="请选择">
            <el-option
              v-for="item in options.categorys"
              :key="`category-${item.id}`"
              :label="item.categoryName"
              :value="item.id"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="上架">
          <el-radio-group v-model="listQuery.isOnSale">
            <el-radio-button
              v-for="item, index in options.sales"
              :key="`sales-${item.value}`"
              :label="item.value">{{item.label}}</el-radio-button>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="主推">
          <el-radio-group v-model="listQuery.isFirst">
            <el-radio-button
              v-for="item, index in options.firsts"
              :key="`first-${item.value}`"
              :label="item.value">{{item.label}}</el-radio-button>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="热门">
          <el-radio-group v-model="listQuery.isHot">
            <el-radio-button
              v-for="item, index in options.hots"
              :key="`hot-${item.value}`"
              :label="item.value">{{item.label}}</el-radio-button>
          </el-radio-group>
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
          icon="el-icon-close"
          v-waves
          @click="handleStatus">
          下架
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
          icon="el-icon-plus">
          <router-link to="/goods/add">
            新增
          </router-link>
        </el-button>
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
          label="商品信息"
          width="200"
          align="center">
          <template slot-scope="scope">
            <div class="p-info-block">
              <div class="p-info-block__pic">
                <div class="p-avatar" 
                  style="width:40px;height:40px;">
                  <img
                    class="p-avatar__img"
                    v-if="scope.row.small"
                    :src="links + scope.row.small" />
                </div>
              </div>
              <div class="p-info-block__content">
                <div>{{scope.row.goodsName}}</div>
                <div class="xa-text-secondary xa-text-left">
                  编号：{{scope.row.goodsSn}}
                </div>
              </div>
            </div>
          </template>
        </el-table-column>
        <el-table-column
          prop="categoryName"
          label="分类"
          width="100"
          align="center">
        </el-table-column>
        <el-table-column
          prop="goodsPrice"
          label="价格"
          align="center">
        </el-table-column>
        <el-table-column
          prop="isOnSale"
          label="上/下架"
          align="center">
          <template slot-scope="scope">
            <span v-if="scope.row.isOnSale==1">已上架</span>
            <span v-if="scope.row.isOnSale==2">已下架</span>
          </template>
        </el-table-column>
        <el-table-column
          prop="isFirst"
          label="主推"
          align="center">
          <template slot-scope="scope">
            <span v-if="scope.row.isFirst==1">是</span>
            <span v-if="scope.row.isFirst==2">否</span>
          </template>
        </el-table-column>
        <el-table-column
          prop="isHot"
          label="热销"
          align="center">
          <template slot-scope="scope">
            <span v-if="scope.row.isHot==1">是</span>
            <span v-if="scope.row.isHot==2">否</span>
          </template>
        </el-table-column>
        <el-table-column
          prop="createTime"
          label="创建时间"
          width="180"
          align="center">
          <template slot-scope="scope">
            {{scope.row.createTime | parseTime}}
          </template>
        </el-table-column>
        <el-table-column
          prop="updateTime"
          label="更新时间"
          width="180"
          align="center">
          <template slot-scope="scope">
            {{scope.row.updateTime | parseTime}}
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
                    @click="handleEdit(scope.row.id)">
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
import waves from '@/directive/waves/index'
import Pagination from '../components/pagination'
import {fetchGet, fetchPost} from '@/api'
import {URIS} from '@/config'

export default {
  name: 'goods',
  directives: {
    waves
  },
  components: {
    Pagination
  },
  data() {
    return {
      list: [],
      total: null,
      loading: true,
      tip: false,
      links: URIS.Goods,
      listQuery: {
        pageNumber: 1,
        pageSize: 10,
        name: '',
        categoryId: 0,
        isOnSale: 0,
        isFirst: 0,
        isHot: 0,
      },
      mulSelection: [],
      tableKey: 0,
      options: {
        categorys: [],
        sales: [
          {value: 0, label: '全部'},
          {value: 1, label: '上架'},
          {value: 2, label: '下架'}
        ],
        firsts: [
          {value: 0, label: '全部'},
          {value: 1, label: '是'},
          {value: 2, label: '否'}
        ],
        hots: [
          {value: 0, label: '全部'},
          {value: 1, label: '是'},
          {value: 2, label: '否'}
        ]
      }
    }
  },
  computed: {},
  methods: {
    //获取list数据，传入筛选对象
    getList() {
      const self = this
      self.loading = true

      fetchGet('/goods/product/list', self.listQuery).then(response => {
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
      }).catch(ex => {
        self.$notify({
          title: '请求错误',
          message: ex,
          type: 'error'
        })
      })
    },
    getCategorys() {
      const self = this
      fetchGet('/goods/category/list', {pid:0}).then(response => {
        const status = response.data.state
        const res = response.data.data
        const message = response.data.msg
        if (status) {
          self.options.categorys = [{id:0, categoryName:'全部'}, ...res]
        }
      }).catch(ex => {
        self.$notify({
          title: '请求错误',
          message: ex,
          type: 'error'
        })
      })
    },
    delete(ids) {
      fetchGet('/goods/product/delete', {
        id: ids
      }).then(response => {
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
      }).catch(ex => {
        self.$notify({
          title: '请求错误',
          message: ex,
          type: 'error'
        })
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
        categoryId: 0,
        isOnSale: 0,
        isFirst: 0,
        isHot: 0,
      }
    },
    //筛选状态
    filterStatus(value, row, column) {
      const property = column['property'];
      return row[property] == value;
    },
    //编辑
    handleEdit(id) {
      this.$router.push('/goods/form?id='+id);
    },
    //删除
    handleDelete(row) {
      const self = this
      let ids = []
      ids.push(row.id)

      self.$confirm(`确定要删除商品【${row.goodsName}】?`, '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(function(action) {
        self.delete(ids)
      })
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

      self.$confirm(`确定要批量删除商品?`, '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
      .then(function(action) {  
        let ids = []
        self.mulSelection.map(function(row) {
          ids.push(row.id)
        })
        self.delete(ids)
      })
    },
    handleStatus() {
      const self = this 
      if (self.mulSelection.length === 0) return

      self.$confirm(`确定要批量下架商品?`, '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
      .then(function(action) {  
        let ids = []
        self.mulSelection.map(function(row) {
          ids.push(row.id)
        })

        fetchGet('/goods/product/close', {id: ids}).then(response => {
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
        }).catch(ex => {
          self.$notify({
            title: '请求错误',
            message: ex,
            type: 'error'
          })
        })
      })
    }
  },
  created() {
    this.getList()
    this.getCategorys()
  }
}
</script>
