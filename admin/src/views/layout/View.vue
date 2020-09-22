<template>
  <!-- vue -->
  <div>
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
  </div>
</template>

<script>
export default {
  name: 'el-view',
  components: {},
  data() {
    return {
    }
  },
  methods: {
  },
  //实时计算属性
  computed: {
    key() {
      return this.$route.name !== undefined ? this.$route.name + +new Date() : this.$route + +new Date()
    }
  },
  created() {
  }
}
</script>
