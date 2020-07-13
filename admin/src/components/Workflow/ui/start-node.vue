<template>
  <div class="start-node">
    <div class="start-node-text">
      流程开始
    </div>
  </div>
</template>
<script>
export default {
  name: 'start-node'
}
</script>
<style lang="scss" scoped>
@import "../../../styles/_global.scss";

.start-node {
  height:80px;
  position:relative;

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
    background-color:$color-primary;
  }

  &-text {
    width:90px;
    padding:12px 0;
    margin:0 auto;
    text-align: center;
    border-radius: 50% / 50%;
    font-size: 14px;
    color:#fff;
    background-color:#aaa;
  }
}	
</style>
