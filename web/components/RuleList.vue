<style scoped>
  .list {
    .item {
      margin-bottom: 20px;
    }
  }
</style>

<template>
  <search></search>
  <ul class="list" v-if="status">
    <rule class="item" v-for="rule in list" :data="rule" track-by="id" transition="show"></rule>
  </ul>
</template>

<script>
import CircleButton from './common/CircleButton.vue'
import Rule from './Rule.vue'
import Search from './Search.vue'

export default {
  components: {
    CircleButton,
    Rule,
    Search
  },
  vuex: {
    getters: {
      status: (state) => state.app.status,
      list: (state) => state.rule.list
    }
  }
}
</script>
