<template>
  <div class="add-node-btn-box">
    <AddNodeBtn
      :node="node"
      @addnode="addnode"/>
  </div>
</template>
<script>
import AddNodeBtn from './add-node-btn'
export default {
  name: 'add-node-btn-box',
  components: {
    AddNodeBtn
  },
  props: {
    node: {
      type: Object,
      default: undefined
    }
  },
  data: () => ({
  }),
  methods: {
    addnode (node) {
      this.$emit('addnode', node)
    }
  }
}
</script>
<style lang="scss" scoped>
@import "../../../styles/_global.scss";

 .add-node-btn {

  &-box {
    width: 240px;
    display: inline-flex;
    flex-shrink: 0;
    position: relative;
    
    &::before {
      content: "";
      position: absolute;
      top: 0;
      left: 0;
      right: 0;
      bottom: 0;
      z-index: -1;
      margin: auto;
      width: 2px;
      height: 100%;
      background-color: $color-primary;
    }
  }
} 
</style>