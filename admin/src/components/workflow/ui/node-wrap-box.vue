<template>
  <div class="node-wrap-box">
    <div class="node-panel">
      <div
        class="node-panel-title"
        :style="titlebg">
        <i :class="icon" 
          v-if="node.type != 'start'"></i>
        <span class="edit">{{ nodetype }}</span>
        <i class="el-icon-close node-panel-close"
          v-if="node.type != 'start'"
          @click="delNode"></i>
      </div>
      <div
        class="node-panel-content"
        @click="dialog = true">
        <div class="node-panel-text">
          {{ content }}
        </div>
        <i class="el-icon-arrow-right arrow"></i>
      </div>
    </div>
    <AddNodeApprover
      :show.sync="dialog"
      :title="`选择${nodetype}`"
      :properties="node.properties"
      @setProperties="setProperties"
    />
  </div>
</template>
<script>
import AddNodeApprover from './add-node-approver'
export default {
  name: 'node-wrap-box',
  components: {
    AddNodeApprover
  },
  props: {
    node: {
      type: Object,
      default: Object
    }
  },
  data: () => ({
    dialog: false,
    text: '请选择',
    titlebg: 'background: rgb(87, 106, 149);',
    nodetype: '发起人',
    content: '选择审批人',
    icon: ''
  }),
  mounted () {
    this.setText()
  },
  methods: {
    delNode () {
      this.$emit('delNode')
    },
    addApprover () {
      this.dialog = true
    },
    setProperties (properties) {
      this.node.properties = properties
      this.setText()
    },
    setText () {
      switch (this.node.type) {
        case 'start':
          this.content = '所有人'
          break
        case 'approver':
          this.nodetype = '审批人'
          this.titlebg = 'background: rgb(255, 148, 62);'
          this.icon = 'el-icon-user-solid'
          break
        case 'notifier':
          this.nodetype = '抄送人'
          this.titlebg = 'background: rgb(50, 150, 250)'
          this.icon = 'el-icon-position'
          break
        default:
      }
      if (this.node.properties && this.node.properties.actionerRules) {
        var acr = this.node.properties.actionerRules[0]
        switch (acr.type) {
          case 'target_management':
            this.content = '直接主管'
            this.icon = 'el-icon-user-solid'
            break
          case 'target_approval':
            this.content = ''
            for (var i = 0; i < acr.approvals.length; i++) {
              this.content += acr.approvals[i].userName + '、'
            }
            this.content = this.content.substr(0, this.content.length - 1)
            break
          case 'target_label':
            this.nodetype = acr.labelNames
            this.content = acr.labelNames
            this.icon = 'el-icon-user-solid'
            if (this.node.type === 'approver') this.content += acr.actType === 'and' ? '会签' : '或签'
            break
          default:
        }
      }
    }
  }
}
</script>
<style lang="scss" scoped>
@import "../../../styles/_global.scss";

.node {
  
  &-wrap {

    &-box {
      display: inline-flex;
      flex-direction: column;
      position: relative;
      width: 220px;
      border:1px solid #f5f5f7;
      min-height: 72px;
      flex-shrink: 0;
      background: #fff;
      border-radius: 4px;
      box-shadow: 0 2px 5px rgba(0,0,0,.045);
      cursor: pointer;

      &:hover {
        border-color:rgb(50, 150, 250);
      }

      &::before {
        content: "";
        position: absolute;
        top: -13px;
        left: 50%;
        transform: translateX(-50%);
        width: 0;
        height: 4px;
        border-style: solid;
        border-width: 8px 6px 4px;
        border-color: $color-primary transparent transparent;
        background: #f5f5f7;
      }
    }
  }

  &-panel {

    &-close {
      position: absolute;
      right: 10px;
      top: 50%;
      transform: translateY(-50%);
      width: 20px;
      height: 20px;
      font-size: 14px;
      color: #fff;
      border-radius: 50%;
      text-align: center;
      line-height: 20px;
    }

    &-title {
      position: relative;
      display: flex;
      align-items: center;
      padding-left: 16px;
      padding-right: 30px;
      width: 100%;
      height: 24px;
      line-height: 24px;
      font-size: 12px;
      color: #fff;
      text-align: left;
      background: #576a95;
      border-radius: 4px 4px 0 0;

      .edit {
        margin-left: 10px;
        line-height: 15px;
        overflow: hidden;
        white-space: nowrap;
        text-overflow: ellipsis;
        border-bottom: 1px dashed transparent;
        font-size: 0.7em;

        &::before {
          content: "";
          position: absolute;
          top: 0;
          left: 0;
          bottom: 0;
          right: 40px;
        }
      }
    }

    &-content {
      position: relative;
      font-size: 14px;
      padding: 16px;
      padding-right: 30px;

      .arrow {
        position: absolute;
        right: 10px;
        top: 50%;
        transform: translateY(-50%);
        width: 20px;
        height: 14px;
        font-size: 18px;
        color: #979797;
      }
    }

    &-text {
      @include ellipsis();
    }
  }
} 
</style>