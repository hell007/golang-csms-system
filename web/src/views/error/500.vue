<template>
<div class="page">
  <div class="page-main error-main">
    <div class="page-scroller">
      <section class="err-panel">
        <div class="err-level err-500" 
          v-if="state=='500'">错误</div>
        <div class="err-level err-404"
          v-if="state=='404'">错误</div>
        <div class="err-text">
          <p>服务器累了呐</p>
          <span>The Server Is Not Available</span>
        </div>
        <div class="err-back" onclick="history.go(-1)">返回</div>
      </section>
    </div>
  </div>
</div>
</template>
<script>

export default {
  name: 'error',
  components: {
  },
  data() {
    return {
      state: '500'
    }
  },
  methods: {
    get() {},
  },
  created() {}
}
</script>
<style type="text/css">
@import '../../styles/error.css'
</style>
