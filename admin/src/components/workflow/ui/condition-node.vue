<template>
  <div class="condition-node">
    <ConditionNodeBox
      :node="node"
      @addnode="addnode"
      @delConditionNode="delConditionNode"
      @addConditionFactor="addConditionFactor"
      @updateCondition="updateCondition"/>
  </div>
</template>
<script>
import ConditionNodeBox from './condition-node-box'
export default {
  name: 'condition-node',
  components: {
    ConditionNodeBox
  },
  props: {
    // value: {
    //   type: Number,
    //   default: 1
    // },
    // text: {
    //   type: String,
    //   default: '请设置条件'
    // },
    node: {
      type: Object,
      default: undefined
    }
  },
  methods: {
    addnode (node) {
      this.$emit('addnode', node)
    },
    delConditionNode () {
      this.$emit('delConditionNode')
    },
    addConditionFactor (node) {
      this.$emit('addConditionFactor', node)
    },
    updateCondition (node) {
      this.$emit('updateCondition', node)
    }
  },
  created() {
    //console.log('condition-node', this.node)
  }
}
</script>
<style lang="scss" scoped>
.condition-node {
  min-height: 220px;
  
  &,
  &-box {
    display: inline-flex;
    -webkit-box-orient: vertical;
    -webkit-box-direction: normal;
    flex-direction: column;
    -webkit-box-flex: 1;
  }

  &-box {
    padding-top: 30px;
    padding-right: 50px;
    padding-left: 50px;
    -webkit-box-pack: center;
    justify-content: center;
    -webkit-box-align: center;
    align-items: center;
    flex-grow: 1;
    position: relative;

    &::before {
      content: "";
      position: absolute;
      top: 0;
      left: 0;
      right: 0;
      bottom: 0;
      margin: auto;
      width: 2px;
      height: 100%;
      background-color: #cacaca;
    }
  }
}
</style>