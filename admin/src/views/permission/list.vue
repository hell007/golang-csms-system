<template>
  <el-card class="xa-block" shadow="hover">

    <div class="xa-block xa-clearfix">  
      <div class="xa-pull-left">
        <el-button 
          @click="handleToggleTree(true)" 
          type="primary"
          size="small">
          全部展开
        </el-button>
        <el-button 
          @click="handleToggleTree(false)" 
          type="primary"
          size="small">
          全部收起
        </el-button>
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
          <router-link to="/permission/form">
            新增
          </router-link>
        </el-button>
      </div>
    </div>
    
    <!-- tree -->
    <el-tree
      class="p-tree"
      ref="tree"
      :data="list"
      :expand-on-click-node="false"
      default-expand-all
      show-checkbox
      node-key="id"
      v-loading="loading"
      element-loading-text="正在玩命加载中...">
      <div class="p-tree__content" slot-scope="item">
        <div class="p-tree__main">
          <span class="p-tree__actions">
            <el-button 
              size="small" 
              type="text"
              icon="el-icon-plus"
              v-if="item.data.level ==2 "
              @click="handleAdd()">新增</el-button>
            <el-button 
              size="small" 
              type="text"
              icon="el-icon-edit"
              @click="handleEdit(item.data.id)">编辑</el-button>
            <el-button 
              size="small" 
              type="text"
              icon="el-icon-delete"
              v-if="item.data.level ==3 "
              @click="handleDelete(item.data)">删除</el-button>
          </span>
        </div>
        <div class="p-tree__left">
          <span class="p-tree__name">
            {{ item.data.name }}
            <el-tooltip
              v-if="item.data.status== 2"
              effect="dark"
              placement="top">
              <div slot="content">停用</div>
              <i class="el-icon-info"></i>
            </el-tooltip>
          </span>
        </div>
      </div>
    </el-tree>
  
  </el-card>
</template>

<script>
import waves from '@/directive/waves/index'
import {fetchGet, fetchPost} from '@/api'
import {getTree} from '@/utils'

export default {
  name: 'permission',
  directives: {
    waves
  },
  components: {},
  data() {
    return {
      list: [],
      loading: true,
      tableKey: 0,
      options: {
        status: [
          {value: 0, label: '全部'},
          {value: 1, label: '启用'},
          {value: 2, label: '停用'}
        ]
      },
      checked: []
    }
  },
  methods: {
    getList() {
      const self = this
      self.loading = true
      fetchGet('/sys/permission/list').then(response => {
        const status = response.data.state
        const res = response.data.data.rows
        const message = response.data.msg
        if (status) {
          self.list = getTree(res)
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
      fetchGet('/sys/permission/delete', {
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
          self.getList()
        } else {
          this.$notify({
            title: '失败',
            message: message,
            type: 'error'
          })
        }
      }) 
    },
    handleToggleTree: function(expanded) {
      this.$refs.tree.store._getAllNodes().forEach(function(item) {
        item.expanded = expanded
      })
    },
    //编辑
    handleEdit(id) {
      this.$router.push('/permission/form?id='+id);
    },
    handleAdd() {
      this.$router.push('/permission/form');
    },
    //删除
    handleDelete(item) {
      const self = this
      let ids = []
      ids.push(item.id)

      self.$confirm(`确定要删除【${item.name}】?`, '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
      .then(function(action) {
        self.delete(ids)
      })
      .catch(function(action) {})
    },
    //批量删除
    handleBatchDelete() {
      const self = this 
      self.checked = self.$refs.tree.getCheckedKeys()
      if (self.checked.length === 0) return

      self.$confirm('确定要批量删除?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
      .then(function(action) {
        let ids = []
        self.checked.map(function(row) {
          ids.push(row)
        })
        self.delete(ids)
      })
      .catch(function(action) {
        self.$refs.tree.setCheckedKeys([])
      })
    },
    handleStatus() {
      const self = this 
      self.checked = self.$refs.tree.getCheckedKeys()
      if (self.checked.length === 0) return

      self.$confirm(`确定要批量停用?`, '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
      .then(function(action) { 
        let ids = []
        self.checked.map(function(row) {
          ids.push(row)
        })
        fetchGet('/sys/permission/close', {id: ids }).then(response => { 
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
