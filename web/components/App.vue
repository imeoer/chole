<style scoped>
  .app {
    color: red;
    .xx {
      color: blue;
    }
  }
</style>

<template>
  <div class="app">
    <component-a></component-a>
    <div class="xx">ABC</div>
  </div>
</template>

<script>
import ComponentA from './ComponentA.vue'

export default {
  components: {
    ComponentA,
  },
}
</script>