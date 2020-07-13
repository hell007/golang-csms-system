<template>
  <div class="end-node">
    <div class="end-node-text">
      流程结束
    </div>
  </div>
</template>
<script>
export default {
  name: 'end-node'
}
</script>
<style lang="scss" scoped>
@import "../../../styles/_global.scss";

.end-node {
  position:relative;

  &:before {
    content: "";
    position: absolute;
    top: -12px;
    left: 50%;
    transform: translateX(-50%);
    width: 0;
    height: 4px;
    border-style: solid;
    border-width: 8px 6px 4px;
    border-color: $color-primary transparent transparent;
    background: #f5f5f7;
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
