<template>
  <div class="col-box">
    <div
      v-if="pos == 0"
      class="top-left-cover-line"/>
    <div
      v-if="pos == 0"
      class="bottom-left-cover-line"/>
    <div
      v-if="pos == (total-1)"
      class="top-right-cover-line"/>
    <div
      v-if="pos == (total-1)"
      class="bottom-right-cover-line"/>
    <Node
      v-for="(item, index) in items"
      :key="index"
      :node="item"
      @addnode="addnode"
      @delNode="delNode(item)"
      @delConditionNode="delConditionNode"
      @addConditionFactor="addConditionFactor"/>
  </div>
</template>
<script>
import { iteratorData, addNewNode, delNode } from './process'
export default {
  name: 'col-box',
  props: {
    node: {
      type: Object,
      default: undefined
    },
    total: {
      type: Number,
      default: 0
    },
    pos: {
      type: Number,
      default: 0
    }
  },
  data: () => ({
    items: [],
    node1: null
  }),
  watch: {
    node: {
      handler (val) {
        // console.log(val)
        if (val) {
          this.getData(val)
          this.node1 = val
        }
      },
      deep: true
    }
  },
  mounted () {
    if (this.node) {
      this.getData(this.node)
      this.node1 = this.node
    }
  },
  methods: {
    getData (data) {
      this.items = []
      iteratorData(this.items, data)
    },
    addnode (node) {
      // console.log(this.node1)
      addNewNode(node, this.node1, this.items)
    },
    delNode (node) {
      delNode(node, this.node1, this.items)
    },
    delConditionNode () {
      this.$emit('delConditionNode')
    },
    addConditionFactor (node) {
      this.$emit('addConditionFactor', node)
    }
  }
}
</script>
<style lang="scss" scoped>
.col-box {
  display: inline-flex;
  -webkit-box-orient: vertical;
  -webkit-box-direction: normal;
  flex-direction: column;
  -webkit-box-align: center;
  align-items: center;
  position: relative;
  background: #f5f5f7;

  &::before {
    content: "";
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    z-index: 0;
    margin: auto;
    width: 2px;
    height: 100%;
    background-color: #cacaca;
  }
}

.bottom-left-cover-line, 
.bottom-right-cover-line {
  position: absolute;
  height: 3px;
  width: 50%;
  background-color: #f5f5f7;
  bottom: -2px;
}

.bottom-right-cover-line {
  right: -1px;
}

.bottom-left-cover-line {
  left: -1px;
}

.top-left-cover-line, 
.top-right-cover-line {
  position: absolute;
  height: 3px;
  width: 50%;
  background-color: #f5f5f7;
  top: -2px;
}

.top-left-cover-line {
  left: -1px;
}

.top-right-cover-line {
  right: -1px;
}
</style>
