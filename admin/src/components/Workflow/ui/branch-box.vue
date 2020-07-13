<template>
  <div class="branch-box">
    <AddBranch @addCondition="addCondition" />
    <slot>branch-box</slot>
  </div>
</template>
<script>
import AddBranch from './add-branch'
export default {
  name: 'branch-box',
  components: {
    AddBranch
  },
  methods: {
    addCondition () {
      this.$emit('addCondition')
    },
    delNode () {
      this.$emit('delNode')
    }
  }
}
</script>
<style lang="scss" scoped>
.branch {

  &-box {
    display: flex;
    overflow: visible;
    min-height: 180px;
    height: auto;
    border-bottom: 2px solid #ccc;
    border-top: 2px solid #ccc;
    position: relative;
    margin-top: 15px;
  }
}  
</style>