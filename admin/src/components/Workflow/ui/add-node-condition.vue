<template>
  <el-dialog
    :append-to-body="true"
    class="xa-dialog"
    title="条件编辑"
    :visible.sync="visible">
    <div class="condition-tips"
      v-if="showTips">
      <i class="el-icon-warning-outline"></i>
      <span>当审批单同时满足以下条件时进入此流程</span>
      <em @click="showTips=false">我知道了</em>
    </div>
    <ConditionGroup
      v-for="(item, index) in conditions"
      :key="index"
      :data.sync="item"
      :pos="index"
      @del="delGroup(index)"/>
    <el-button 
      type="primary"
      size="small"
      icon="el-icon-plus"
      @click="addCondition">添加条件</el-button>
    <span slot="footer" class="dialog-footer">
      <el-button 
        type="info"
        @click="cancel">取 消</el-button>
      <el-button
        type="primary"
        @click="save">保 存</el-button>
    </span>
  </el-dialog>  
</template>
<script>
import ConditionGroup from './condition-group'
export default {
  name: 'add-node-condition',
  components: {
    ConditionGroup
  },
  props: {
    show: {
      type: Boolean,
      default: false
    },
    properties: {
      type: Object,
      default: null
    }
  },
  data () {
    return {
      visible: false,
      showTips: true,
      conditions: [],
      type1: 'dingtalk_actioner_range_condition'
    }
  },
  watch: {
    show (val) {
      this.visible = val
    },
    visible (val) {
      this.$emit('update:show', val)
    },
    properties() {
      this.conditions = this.properties.conditions[0]
    }
  },
  created () {
    console.log('==',this.properties)
  },
  methods: {
    ifValid (cond) {
      if (!cond.paramKey || !cond.paramLabel) {
        alert('参数的paramKey和paramLabel 值不能为空')
        return false
      }
      if (!cond.type || cond.type === '') {
        alert('条件类型不能为空')
        return false
      }

      return true
    },
    cancel () {
      this.$emit('update:show', false)
    },
    save () {
      if (this.conditions.length === 0) {
        alert('条件不能为空')
        return
      }
      var flag = this.conditions.every(cond => {
        return this.ifValid(cond)
      })
      
      if (flag) {
        this.visible = false
        this.properties.conditions[0] = this.conditions
        this.$emit('on-success', this.properties)
      }
    },
    addCondition () {
      this.conditions.push({
        pos: this.conditions.length + 1,
        type: this.type1
      })
    },
    delGroup (index) {
      this.conditions.splice(index, 1)
    }
  }
}
</script>
<style lang="scss" scoped>
.condition {

  &-tips {
    @include flex-row();
    padding: 9px 16px;
    margin-bottom: 10px;
    background: #e6f7ff;
    border: 1px solid #91d5ff;
    border-radius: 4px;
    
    i {
      color:#096dd9;
      margin-right:10px;
    }

    span {
      flex:3;
      font-size:12px;
    }

    em {
      text-align:right;
      font-size:12px;
      font-style:normal;
      color:#096dd9;
      cursor:pointer;
    }
  }
}  
</style>
