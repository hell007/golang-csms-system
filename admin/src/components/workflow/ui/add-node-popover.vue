<template>
  <div class="add-node-popover"
    v-if="show">
    <div class="add-node-popover-arrow"></div>
    <div class="add-node-popover-header">
      <i class="el-icon-close close"
        @click="visible = false"></i>
    </div>
    <ul class="add-node-popover-body">
      <li
        v-for="(n, index) in nodes"
        :key="index"
        :class="'add-node-popover-item ' + n.type"
        @click="click(n)">
        <div class="item-wrapper">
          <i :class="n.icon"></i>
        </div>
        <p>{{ n.label }}</p>
      </li>
    </ul>
  </div>
</template>
<script>
export default {
  name: 'add-node-popover',
  props: {
    show: {
      type: Boolean,
      default: false
    },
    node: {
      type: Object,
      default: undefined
    }
  },
  data: () => ({
    visible: false,
    nodes: [
      {type: 'approver', label: '审批人', icon: 'el-icon-user-solid'},
      {type: 'notifier', label: '抄送人', icon: 'el-icon-position'},
      {type: 'route', label: '条件分支', icon: 'xa-icon xa-icon-tree'}
    ]
  }),
  watch: {
    'show' (val) {
      this.visible = val
    },
    visible (val) {
      this.$emit('update:show', val)
    }
  },
  methods: {
    click (item) {
      this.visible = false
      switch (item.type) {
        case 'approver':
          this.addApprover(item)
          break
        case 'route':
          this.addRoute(item)
          break
        case 'notifier':
          this.addNotifier(item)
          break
        default:
          this.$Message.error('暂时不支持')
      }
    },
    addApprover (item) {
      var node = {
        name: '审批人',
        prevId: this.node.nodeId,
        nodeId: '' + new Date().getTime(),
        type: 'approver'
      }
      this.$emit('addnode', node)
    },
    addNotifier (item) {
      var node = {
        name: '抄送人',
        prevId: this.node.nodeId,
        nodeId: '' + new Date().getTime(),
        type: 'notifier'
      }
      this.$emit('addnode', node)
    },
    addRoute (item) {
      var nodeId = new Date().getTime()
      var node = {
        type: 'route',
        prevId: this.node.nodeId,
        nodeId: '' + nodeId,
        conditionNodes: [{
          sort: 1, 
          name: '条件1', 
          type: 'condition', 
          prevId: '' + nodeId, 
          nodeId: '' + (nodeId + 10) 
        }, { 
          sort: 2,
          name: '条件2', 
          type: 'condition', 
          prevId: '' + (nodeId + 10), 
          nodeId: '' + (nodeId + 20) }
        ]
      }
      this.$emit('addnode', node)
    }
  }
}
</script>
<style lang="scss" scoped>
 .add-node {
  
  &-popover {
    position: absolute; 
    top: 18px; 
    left: 60%;
    z-index:1000;
    background-color: #fff;
    background-clip: padding-box;
    border-radius: 4px;
    box-shadow: 0 2px 8px rgba(0, 0, 0, 0.15);

    &-header {
      position: relative;
      margin-bottom: 5px;
      height: 15px;

      .close {
        cursor: pointer;
        position: absolute;
        top: 5px;
        right: 5px;
        color: rgba(0, 0, 0, 0.25);
        font-size:18px;
      }
    }

    &-body {
      display: flex;
      padding:0 15px 15px;
    }

    &-arrow {
      position: absolute;
      left: -4px;
      top: 12px;
      border-top-color: transparent;
      border-right-color: transparent;
      border-bottom-color: #fff;
      border-left-color: #fff;
      box-shadow: -3px 3px 7px rgba(0, 0, 0, 0.07);
      display: block;
      width: 8.48528137px;
      height: 8.48528137px;
      background: transparent;
      border-style: solid;
      border-width: 4px;
      transform: rotate(45deg);
    }

    &-item {
      margin-right: 10px;
      cursor: pointer;
      text-align: center;
      flex: none;
      color: #191f25!important;

      &.approver .item-wrapper {
        color: #ff943e;
      }
      
      &.notifier .item-wrapper {
        color: #3296fa;
      }

      .item-wrapper {
        user-select: none;
        display: inline-block;
        width: 80px;
        height: 80px;
        margin-bottom: 5px;
        background: #fff;
        border: 1px solid #e2e2e2;
        border-radius: 50%;
        transition: all .3s cubic-bezier(.645,.045,.355,1);
      }

      i {
        font-size: 35px;
        line-height: 80px;
      }

    }
  }  
}
</style> 