<style scoped>
.toolbar {
  overflow: auto;
  margin-bottom: 60px;
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
    <circle-button class="start {{status ? 'green' : 'orange'}}" icon="rocket">{{status ? '运行中' : '未启动'}}</circle-button>
    <circle-button class="edit" icon="circle-o" v-link="'/edit'">编辑规则</circle-button>
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
      status: (state) => state.app.status
    }
  }
}
</script>
