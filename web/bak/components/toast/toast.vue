<template>
  <div 
    v-if="visible" 
    :class="showContent ?'fadein':'fadeout'">
    <div class="dialog dialog-center msg dialog-show toast">
      <div class="dialog-main">
        <div class="dialog-body msg-body">
          <i v-if="icon=='success' " class="toast-icon toast-icon-success"></i>
          <i v-else class="toast-icon toast-icon-error"></i>
          <div class="msg-message">{{text}}</div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.fadein {
  animation: animate-in 0.25s;
}

.fadeout {
  animation: animate-out 0.25s;
  opacity: 0;
}

@keyframes animate-in {
  0% {
    opacity: 0;
  }
  100% {
    opacity: 1;
  }
}

@keyframes animate-out {
  0% {
    opacity: 1;
  }
  100% {
    opacity: 0;
  }
}
</style>
