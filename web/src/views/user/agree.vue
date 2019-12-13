<template>
<div class="page">
  <div class="page-header">
    <header class="header">
      <div class="header-left">
        <a href="javascript:history.go(-1);" 
          class="header-btn header-btn-back">
          返回
        </a>
      </div>
      <div class="header-title">
        注册协议
      </div>
    </header>
  </div>

  <div class="page-main">
    <div class="page-scroller">
      <section class="agree-card">
        <div class="agree-card-bd">
          <h6>一、购买流程</h6>
          <p>
            亲爱的客户，很抱歉您的商品出现这样的问题。请准备好商品全套（商品、附件、发票原件、包装）按以下地址寄往和小店售后部，不支持到付及平邮。如果收到商品后情况属实，我们会尽快为您办理。感谢您对我们的支持。
          </p>
        </div>
      </section>
    </div>
  </div>
</div>
</template>
<script>

export default {
  name: 'user-agree',
  components: {
  },
  data() {
    return {
      text: ''
    }
  },
  methods: {
    get() {
      
    },
  },
  created() {
    
  }
}
</script>
<style type="text/css">
.agree-card {
  padding: 0.55555556rem;
  margin: 0.55555556rem;
  -webkit-border-radius: 0.27777778rem;
  border-radius: 0.27777778rem;
  background-color: #fff;
  font-size: 0.83333333rem;
}

.agree-card-bd p {
  font-size: 0.77777778rem;
  color: #999;
  margin-bottom: 0.41666667rem;
}

</style>
