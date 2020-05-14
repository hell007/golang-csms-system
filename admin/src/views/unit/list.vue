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
          <el-button 
            size="small"
            v-waves
            v-auth="p">
            按钮权限
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
          <router-link to="/unit1/form">新增</router-link>
        </el-button>
        <el-button 
          size="small"
          type="primary"
          icon="el-icon-download"
          v-waves
          @click="handleDownload"  >
          导出
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
          prop="Role.roleName"
          label="角色"
          width="100"
          align="center">
        </el-table-column>
        <el-table-column
          prop="mobile"
          label="手机号"
          width="140"
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
          <template slot-scope="scope">
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
/**
vue 生命周期8个阶段：
beforeCreate（创建前）,
created（创建后）,
beforeMount(载入前),
mounted（载入后）,
beforeUpdate（更新前）,
updated（更新后）,
beforeDestroy（销毁前）,
destroyed（销毁后）
 */
import waves from '@/directive/waves/index' // 水波纹指令
import auth from '@/directive/auth/index' 
import Pagination from '../components/pagination'
import {
  mapState,
  mapMutations,
  mapActions,
  mapGetters
} from 'vuex'

import {
  noop,
  logger,
  formatJsonTime
} from '@/utils'

export default {
  name: 'userlist',
  directives: {
    waves,
    auth
  },
  components: {
    Pagination
  },
  data() {
    return {
      p:'permission:list', //测试无意义
      list: [],
      total: 0,
      loading: true,
      tip: false,
      listQuery: {
        pageNumber: 1,
        pageSize: 10,
        name: '',//字符串不能为null,否则影响查询
        status: null
      },
      mulSelection: [],
      tableKey: 0,
      options: {
        status: [
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
  computed: {},
  watch: {},
  methods: {
    ...mapActions(['getUserList', 'deleteUser', 'closeUser', 'batchDeleteUser']),
    //获取list数据，传入筛选对象
    getList() {
      const self = this
      self.loading = true
      self.getUserList(self.listQuery).then(response => {
        logger('info', response)
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
      this.listQuery.pageNum = 1
      this.getList()
    },
    //重置
    handleRest() {
      this.listQuery.name = ''
      this.listQuery.status = null
    },
    //筛选状态
    filterStatus(value, row, column) {
      const property = column['property'];
      //console.log(property,value,row)
      return row[property] == value;
    },
    //编辑
    handleEdit(id) {
      //路由跳转
      // 字符串
      this.$router.push('/unit1/form?id='+id);
      // 对象
      //this.$router.push({ path: '/user/form'})
      // 命名的路由
      //this.$router.push({ name: '/user/form', params: { id: id }})
    },
    //删除
    handleDelete(row) {
      const self = this
      self.deleteUser(row).then(response => {
        logger('info',response)
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
          this.$notify({
            title: '失败',
            message: message,
            type: 'error'
          })
        }
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

      self.batchDeleteUser(self.mulSelection).then(response => {
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
          this.$notify({
            title: '失败',
            message: message,
            type: 'error'
          })
        }
      })
    },
    handleStatus() {
      const self = this 
      if (self.mulSelection.length === 0) return

      self.closeUser(self.mulSelection).then(response => {
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
          self.getList() // 刷新
        } else {
          this.$notify({
            title: '失败',
            message: message,
            type: 'error'
          })
        }
      })
    },
    //导出xlsx
    handleDownload() {
      // require.ensure([], () => {
      //   //引入组件
      //   const {
      //     export_json_to_excel
      //   } = require('@/vendor/Export2Excel')
      //   const tHeader = ['序号', '姓名', '手机号', '状态', '创建时间']
      //   const filterVal = ['id', 'name', 'mobile', 'status', 'createTime']
      //   //const data = formatJsonTime(filterVal, 'createTime', this.list)
      //   //导出export_json_to_excel('头格式','数据','文件命名')
      //   export_json_to_excel(tHeader, data, '用户数据')
      // })
    }
  },
  beforeCreate: noop,
  created() {
    this.getList()
  },
  beforeMount: noop,
  mounted: noop,
  beforeDestory: noop,
  destoryed: noop
}
</script>
