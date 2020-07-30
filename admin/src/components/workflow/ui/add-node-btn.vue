<template>
  <div class="add-node-btn">
    <button
      class="btn"
      type="button"
      @click="click">
      <i class="el-icon-plus"></i>
    </button>
    <AddNodePopover
      :show.sync="visible"
      :node="node"
      @addnode="addnode"/>
  </div>
</template>
<script>
import AddNodePopover from './add-node-popover'
export default {
  name: 'add-node-btn',
  components: {
    AddNodePopover
  },
  props: {
    node: {
      type: Object,
      default: undefined
    }
  },
  data: () => ({
    visible: false
  }),
  methods: {
    click () {
      this.visible = true
    },
    addnode (node) {
      this.$emit('addnode', node)
    }
  }
}
</script>
<style lang="scss" scoped>
@import "../../../styles/_global.scss";

.add-node {

  &-btn {
    user-select: none;
    width: 240px;
    padding: 20px 0 32px;
    display: flex;
    justify-content: center;
    flex-shrink: 0;
    flex-grow: 1;

    .btn {
      outline: none;
      box-shadow: 0 2px 4px 0 rgba(0,0,0,.1);
      width: 30px;
      height: 30px;
      background: $color-primary;
      border-radius: 50%;
      position: relative;
      border: none;
      line-height: 30px;
      transition: all .3s cubic-bezier(.645,.045,.355,1);

      .el-icon-plus {
        color: #fff;
        font-size: 16px;
      }
    }
  }
} 
</style>
