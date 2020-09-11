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
 .add-node-btn {

  &-box {
    width: 240px;
    display: -webkit-inline-box;
    display: -ms-inline-flexbox;
    display: inline-flex;
    -ms-flex-negative: 0;
    flex-shrink: 0;
    -webkit-box-flex: 1;
    -ms-flex-positive: 1;
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
      background-color: #cacaca;
    }
  }
} 
</style>