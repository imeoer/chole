<style scoped>
.toolbar {
  overflow: auto;
  margin-bottom: 20px;
  .start {
    float: left;
  }
  .edit {
    float: right;
  }
}
</style>

<template>
  <div class="toolbar box" v-show="!show">
    <circle-button class="start" icon="play-circle">启动</circle-button>
    <circle-button class="edit green" icon="circle-o" v-link="'/edit'">编辑规则</circle-button>
  </div>
</template>

<script>
import CircleButton from './common/CircleButton.vue'

export default {
  components: {
    CircleButton
  },
  vuex: {
    getters: {
    }
  }
}
</script>
