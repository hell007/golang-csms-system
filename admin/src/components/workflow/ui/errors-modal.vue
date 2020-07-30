<template>
  <el-dialog
    :append-to-body="true"
    class="xa-dialog"
    title="错误"
    :visible.sync="visible">
    <div class="error-modal">
      <div class="error-modal-tips">
        <i class="el-icon-circle-close"></i>
        <span class="ant-modal-confirm-title">当前无法发布</span>
      </div>
      <p class="error-modal-desc">
        以下内容不完善，需进行修改
      </p>
      <ul class="error-modal-list">
        <li
          v-for="(item, index) in data"
          :key="index"
          class="error-modal-item">
          <div class="error-modal-item-label">
            流程设计
          </div>
          <div class="error-modal-item-val">
            {{ item.name }} 未选择
          </div>
        </li>
      </ul>
    </div>
    <span slot="footer" class="dialog-footer">
      <el-button 
        size="small"
        @click="visible = false">我知道了</el-button>
    </span>
  </el-dialog>
</template>
<script>
export default {
  name: 'errors-modal',
  props: {
    show: {
      type: Boolean,
      default: false
    },
    data: {
      type: Array,
      default: null
    }
  },
  data () {
    return {
      visible: false
    }
  },
  watch: {
    'show' (val) {
      this.visible = val
    },
    visible (val) {
      this.$emit('update:show', val)
    }
  }
}
</script>
<style lang="scss" scoped>
.error-modal {

  &-tips {
    font-size:16px;
    margin-bottom:10px;
    
    i {
      font-size:20px;
      margin-right:5px;
    }
  }

  &-desc {
    font-size:12px;
    color:$color-primary;
  }

  &-list {
    height: 200px;
    overflow-y: auto;
    margin-right: -25px;
    padding-right: 25px;
  }
  
  &-item {
    padding: 10px 20px;
    line-height: 21px;
    background: #f6f6f6;
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 8px;
    border-radius: 4px;

    &-label {
      flex: none;
      font-size: 15px;
      color: rgba(25,31,37,.56);
      padding-right: 10px;
    }

    &-val {
      text-align: right;
      flex: 1;
      font-size: 14px;
      color: #666;
    }
  }
}
</style>
