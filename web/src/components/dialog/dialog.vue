<template>
  <div 
    v-if="visible"
    class="dialog dialog-bottom dialog-closable dialog-show">
    <div class="dialog-main">
      <div class="dialog-header">
        <button type="button" 
        	class="dialog-action dialog-action-cancel" 
        	@click="dialogCancle">
          {{cancleText}}
        </button>
        <button type="button" 
        	class="dialog-action dialog-action-confirm" 
        	@click="dialogSubmit">
          {{submitText}}
        </button>
      </div>
      <div class="dialog-body">
        <slot></slot>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'dialog',
  props: {
    visible: {
      type: Boolean,
      default: false
    },
    cancleText: {
    	type: String,
      default: '取消'
    },
    submitText: {
    	type: String,
      default: '确定'
    },
    cancle: {
      type: Function,
      default: null
    },
    submit: {
      type: Function,
      default: null
    }
  },
  data() {
    return {
    }
  },
  methods: {
    dialogCancle() {
      this.$emit('cancle');
    },
    dialogSubmit() {
      this.$emit('submit');
    }
  }
}
</script>
