<template>
  <section style="width:100%;height:100%;">
    <transition name="fade" mode="out-in">
      <keep-alive>
        <router-view v-if="$route.meta.keepAlive" :key="`view-{key}`">
        </router-view>
      </keep-alive>
    </transition>
    <transition name="fade" mode="out-in">
      <router-view v-if="!$route.meta.keepAlive" :key="`view-{key}`">
      </router-view>
    </transition>
  </section>
</template>
<script>
export default {
  name: 'APP',
  components: {},
  data() {
    return {}
  },
  methods: {},
  //实时计算属性
  computed: {
    key() {
      return this.$route.name !== undefined ? this.$route.name + +new Date() : this.$route + +new Date()
    }
  },
  created: function() {}
}
</script>
<style type="text/css">
@import './styles/common.css'
</style>
