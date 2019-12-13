<template>
  <div v-show="!loading">
    <el-pagination 
      class="xa-text-right"
      background
      @size-change="handleSizeChange" 
      @current-change="handleCurrentChange" 
      :current-page.sync="listQuery.pageNum" 
      :page-sizes="[10,20,30,40,50]" 
      :page-size="listQuery.pageSize" 
      :total="total" 
      layout="total, sizes, prev, pager, next, jumper">
      </el-pagination>
  </div>
</template>
<script>
/**
 * 分页组件
 */
export default {
  name: 'Pagination',
  props: {
    loading: {
      type: Boolean,
      default: false
    },
    listQuery: {
      type: Object,
      default: {}
    },
    total: {
      type: Number,
      default: 0
    },
    getList: {
      type: Function,
      default: null
    }
  },
  methods: {
    //页数跳转
    handleSizeChange(val) {
      this.listQuery.pageSize = val
      this.getList()
    },
    //当前页
    handleCurrentChange(val) {
      this.listQuery.pageNum = val
      this.getList()
    }
  }
}
</script>
