<style>
</style>

<template>
  <toolbar></toolbar>
  <rule-list></rule-list>
</template>

<script>
import store from 'store'

import Toolbar from './Toolbar.vue'
import RuleList from './RuleList.vue'

export default {
  store,
  components: {
    Toolbar,
    RuleList
  }
}
</script>
