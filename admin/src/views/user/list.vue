<template>
  <el-card class="xa-block" shadow="hover">
    <!-- 筛选项-->
    <el-form
      class="p-filter" inline size="small"
      @submit.native.prevent
      label-suffix="：">
      <div class="p-filter__input">
        <el-form-item label="姓名">
          <el-input
            placeholder="请输入"
            clearable
            @keyup.enter.native="handleFilter"
            v-model="listQuery.name"></el-input>
        </el-form-item>
        <el-form-item label="状态">
          <el-select
            v-model="listQuery.status"
            clearable
            placeholder="请选择">
            <el-option
              v-for="item in options.status"
              :key="item.value"
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
          icon="el-icon-close"
          v-waves
          @click="handleStatus">
          停用
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
          <router-link to="/user/form">
            新增
          </router-link>
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
        :default-sort="{prop:'createdTime',order:'descending'}"
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
          prop="username"
          label="姓名"
          width="100"
          align="center">
        </el-table-column>
        <el-table-column
          prop="roleName"
          label="角色"
          width="120"
          align="center">
        </el-table-column>
        <el-table-column
          prop="mobile"
          label="手机号"
          width="120"
          align="center">
        </el-table-column>
        <el-table-column
          prop="status"
          label="状态"
          width="100"
          align="center"
          :filters="options.filters"
          :filter-method="filterStatus">
          <template slot-scope="scope">
            <span v-if="scope.row.status==1">启用</span>
            <span v-if="scope.row.status==2">停用</span>
            <el-tooltip
              v-if="scope.row.status== 2"
              effect="dark"
              placement="top">
              <div slot="content">该用户已经离职，给予停用。</div>
              <i class="el-icon-info"></i>
            </el-tooltip>
          </template>
        </el-table-column>
        <el-table-column
          prop="email"
          label="邮箱"
          align="center">
        </el-table-column>
        <el-table-column
          prop="ip"
          label="IP"
          width="120"
          align="center">
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
          prop="loginTime"
          label="登录时间"
          width="180"
          align="center">
          <template slot-scope="scope">
            {{scope.row.loginTime | parseTime}}
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
import {
  mapState,
  mapGetters
} from 'vuex'

export default {
  name: 'user',
  directives: {
    waves
  },
  components: {
    Pagination
  },
  data() {
    return {
      list: [],
      total: 0,
      loading: true,
      tip: false,
      listQuery: {
        pageNumber: 1,
        pageSize: 10,
        name: '',//字符串不能为null,否则影响查询
        status: 0
      },
      mulSelection: [],
      tableKey: 0,
      options: {
        status: [
          {value: 0, label: '全部'},
          {value: 1, label: '启用'},
          {value: 2, label: '停用'}
        ],
        filters: [
          {text: '启用', value: 1},
          {text: '停用', value: 2}
        ]
      }
    }
  },
  computed: {
    ...mapState([
      'admin'
    ]),
    ...mapGetters([
      'test',
      'admins'
    ])
  },
  methods: {
    //获取list数据，传入筛选对象
    getList() {
      const self = this
      self.loading = true
      fetchGet('/sys/user/list', self.listQuery).then(response => {
        const status = response.data.state
        const res = response.data.data
        const message = response.data.msg
        if (status) {
          self.list = res.rows
          self.total = res.total
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
    delete(ids) {
      const self = this
      fetchGet('/sys/user/delete', {id: ids }).then(response => {
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
    },
    //查询
    handleFilter() {
      this.listQuery.pageNum = 1
      this.getList()
    },
    //重置
    handleRest() {
      this.listQuery.name = ''
      this.listQuery.status = 1
    },
    //筛选状态
    filterStatus(value, row, column) {
      const property = column['property'];
      //console.log(property,value,row)
      return row[property] == value;
    },
    //编辑
    handleEdit(id) {
      this.$router.push('/user/form?id='+id);
    },
    //删除
    handleDelete(row) {
      self.$confirm(`确定要删除管理员【${row.username}】?`, '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
      .then(function(action) {
        let ids = []
        ids.push(row.id)
        self.delete(ids)
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

      self.$confirm(`确定要批量删除?`, '提示', {
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

      self.$confirm(`确定要批量停用?`, '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
      .then(function(action) {  
        let ids = []
        self.mulSelection.map(function(row) {
          ids.push(row.id)
        })

        fetchGet('/sys/user/close', {id: ids }).then(response => {
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
      .catch(function(action) {})
    }
  },
  created() {
    this.getList()
  }
}
</script>
