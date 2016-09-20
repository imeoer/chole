<style scoped>

</style>

<template>
  <ul class="list" v-show="!show" transition="show">
    <rule></rule>
  </ul>
</template>

<script>
import CircleButton from './common/CircleButton.vue'
import Rule from './Rule.vue'

export default {
  components: {
    CircleButton,
    Rule
  },
  vuex: {
    getters: {
    }
  }
}
</script>
