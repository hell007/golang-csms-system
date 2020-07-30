<template>
  <button
    class="add-branch"
    @click="addCondition">
    添加条件
  </button>
</template>
<script>
export default {
  name: 'add-branch',
  methods: {
    addCondition () {
      this.$emit('addCondition')
    }
  }
}
</script>
<style lang="scss" scoped>
.add-branch {
  border: none;
  outline: none;
  user-select: none;
  justify-content: center;
  font-size: 12px;
  padding: 0 10px;
  height: 30px;
  line-height: 30px;
  border-radius: 15px;
  color: #3296fa;
  background: #fff;
  box-shadow: 0 2px 4px 0 rgba(0,0,0,.1);
  position: absolute;
  top: -16px;
  left: 50%;
  transform: translateX(-50%);
  transform-origin: center center;
  cursor: pointer;
  z-index: 1;
  display: inline-flex;
  align-items: center;
  transition: all .3s cubic-bezier(.645,.045,.355,1);
}  
</style>
