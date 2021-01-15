<template>
  <div v-if="visible" 
  	class="dialog dialog-center msg dialog-show">
    <div class="dialog-main">
      <div class="dialog-body msg-body">
        <div class="msg-message">{{text}}</div>
      </div>
      <div class="dialog-footer msg-footer">
        <button type="button" 
        	@click="button[0].ontap"
        	class="msg-btn msg-btn-cancel">{{button[0].text}}</button>
        <button type="button" 
        	@click="button[1].ontap"
        	class="msg-btn msg-btn-confirm">{{button[1].text}}</button>
      </div>
    </div>
  </div>
</template>
<script>
export default {
  data() {
    return {
      visible: true,
      text: '',
      button: []
    }
  },
  methods: {
    onClosed() {
      this.close();
    }
  }
}    
</script>