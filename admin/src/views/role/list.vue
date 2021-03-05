<template>
  <el-card class="xa-block" shadow="hover">
    <!-- 筛选项-->
    <div class="xa-block p-filter">
      <el-form
        class="p-filter" inline size="small"
        @submit.native.prevent
        label-suffix="：">
        <div class="p-filter__input">
          <el-form-item label="角色名">
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
    </div>
    
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
        <!-- <el-button 
          size="small" 
          type="danger"
          icon="el-icon-delete"
          v-waves
          @click="handleBatchDelete">
          删除
        </el-button> -->
      </div>
      <div class="xa-pull-right">
        <el-button 
          size="small" 
          type="primary" 
          icon="el-icon-plus">
          <router-link to="/role/form">
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
        :default-sort="{prop:'',order:'descending'}"
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
          prop="rolename"
          label="角色名"
          width="200"
          align="center">
        </el-table-column>
        <el-table-column
          prop="status"
          label="状态"
          width="200"
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
              <div slot="content">该角色已下架，给予停用。</div>
              <i class="el-icon-info"></i>
            </el-tooltip>
          </template>
        </el-table-column>
        <el-table-column
          prop="rolenote"
          label="描述"
          align="left">
        </el-table-column>
        <el-table-column
          label="操作"
          width="150"
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
                <el-dropdown-item v-if="scope.row.pid !=0">
                  <i class="el-icon-setting" 
                    @click="handleMenu(scope.row.id)">
                    权限菜单
                  </i>
                </el-dropdown-item>
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
    
    <!-- 弹窗 -->
    <el-dialog
      title="权限菜单"
      :visible.sync="menu.visible"
      class="xa-dialog xa-dialog--sm">
      <el-tree
        ref="tree"
        node-key="id"
        :data="nodes"
        :props="defaultProps"
        :default-checked-keys="checked"
        default-expand-all
        show-checkbox
        highlight-current>
      </el-tree>
      <span slot="footer" class="xa-dialog-footer">
        <el-button 
          size="small"
          @click="handleCancel">取消</el-button>
        <el-button
          type="primary"
          size="small"
          @click="handleMenuSubmit"
          :loading="menu.processing">确定</el-button>
      </span>
    </el-dialog>

  </el-card>
</template>
<script>

import waves from '@/directive/waves/index'
import Pagination from '../components/pagination'

import {
  mapActions
} from 'vuex'

import {
  getTree,
  uuid
} from '@/utils'

import {
  setPermission
} from '@/utils/storage'

export default {
  name: 'rolelist',
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
        status: 1
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
      },
      // 权限菜单
      menu: {
        visible: false,
        processing: false
      },
      form: {
        rid: null,
        list: []
      },
      nodes: [],
      checked: [],  // 默认选择项
      defaultProps: {
        children: 'children',
        label: 'name'
      }
    }
  },
  computed: {},
  methods: {
    ...mapActions(['getRoleList', 'deleteRole', 'closeRole', 'getPermissions', 'getAccessList', 'saveAccess']),
    getList() {
      const self = this
      self.loading = true
      self.getRoleList(self.listQuery).then(response => {
        const status = response.data.state
        const res = response.data.data
        if (status) {
          self.list = res.rows
          self.total = res.count
        } else {
          self.$notify({
            title: '失败',
            message: response.data.message,
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
      this.listQuery.status = 1
    },
    //筛选状态
    filterStatus(value, row, column) {
      const property = column['property'];
      return row[property] == value;
    },
    //编辑
    handleEdit(id) {
      this.$router.push('/role/form?id='+id);
    },
    //删除
    handleDelete(row) {
      const self = this
      let rows = []
      rows.push(row)

      self.$confirm(`确定要删除角色【${row.rolename}】?`, '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
      .then(function(action) {
        self.deleteRole(rows).then(response => {
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
    // 停用
    handleStatus() {
      const self = this 
      if (self.mulSelection.length === 0) return

      self.$confirm(`确定要批量停用角色?`, '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
      .then(function(action) {
        self.closeRole(self.mulSelection).then(response => {
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
    },
    // 访问菜单
    handleAcesslist(rid) {
      const self = this
      self.getAccessList(rid).then(response=> {
        const status = response.data.state
        const rows = response.data.data.rows
        if (status) {
          self.checked = []
          for (const item of rows) {
            self.checked.push(item.id)
            self.checked.push(item.mid)
          }
          setTimeout(function(){
            self.menu.visible = true
          }, 200);
        }
      })
    },
    // 权限菜单
    handleMenu(id) {
      const self = this
      const query = {
        status: 1
      }
      self.form.rid = id
      self.getPermissions(query).then(response => {
        const status = response.data.state
        const res = response.data.data.rows
        const message = response.data.msg
        if (status) {
          self.nodes = getTree(res)
          self.handleAcesslist(id);
        } else {
          this.$notify({
            title: '失败',
            message: message,
            type: 'error'
          })
        }
      })
    },
    // 清空
    handleCancel() {
      this.checked = null
      this.menu.visible = false
    },
    // 权限菜单提交
    handleMenuSubmit() {
      const self = this
      self.checked = self.$refs.tree.getCheckedKeys()
      if(self.checked.length === 0 || !self.form.rid) return; 
      
      let form = self.form
      self.menu.processing = true

      for (let i = 0, len = self.checked.length; i < len; i++) {
        form.list.push({
          id: uuid(),
          rid: self.form.rid,
          mid: self.checked[i]
        })
      }

      self.saveAccess(form).then(res => {
        const status = res.data.state
        const message = res.data.msg
        if (status) {
          this.$notify({
            title: '成功',
            message: message,
            type: 'success'
          })
          self.menu.processing = false
        } else {
          this.$notify({
            title: '失败',
            message: message,
            type: 'error'
          })
        }
      })
    }
  },
  created() {
    this.getList()
  }
}
</script>
