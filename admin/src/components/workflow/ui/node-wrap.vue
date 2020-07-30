<template>
  <div class="node-wrap">
    <NodeWrapBox
      :node="node"
      @delNode="delNode"/>
    <AddNodeBtnBox
      :node="node"
      @addnode="addnode"/>
  </div>
</template>
<script>
import NodeWrapBox from './node-wrap-box'
import AddNodeBtnBox from './add-node-btn-box'
export default {
  name: 'node-wrap',
  components: {
    NodeWrapBox,
    AddNodeBtnBox
  },
  props: {
    node: {
      type: Object,
      default: undefined
    }
  },
  methods: {
    addnode (node) {
      this.$emit('addnode', node)
    },
    delNode () {
      this.$emit('delNode')
    }
  }
}
</script>
<style lang="scss" scoped>
.node {

  &-wrap {
    flex-direction: column;
    justify-content: flex-start;
    align-items: center;
    flex-wrap: wrap;
    padding: 0 50px;
    position: relative;
    display: inline-flex;
    width: 100%;
  }
} 
</style>