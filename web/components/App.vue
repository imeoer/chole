<style>
@define-mixin font-main {
  font-family: 'Source Sans Pro', 'Microsoft Yahei', 'Hiragino Sans GB', 'WenQuanYi Micro Hei', monospace;
}
@define-mixin hover {
  cursor: pointer;
  transition: .2s all ease;
  &:hover {
    opacity: 0.9;
  }
}
@define-mixin shadow {
  box-shadow: 0 2px 4px rgba(0, 0 , 0, .2);
}
.box {
  background-color: white;
  border: 2px solid #eee;
  border-radius: 5px;
  padding: 15px 20px;
}
html, body {
  @mixin font-main;
  width: 100%;
  height: 100%;
  margin: 0;
  padding: 0;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 12px;
  color: #34495e;
}
ul {
  list-style: none;
  padding: 0;
  margin: 0;
}
#wrapper {
  width: 500px;
}
.header {
  display: flex;
  justify-content: space-between;
  margin-bottom: 20px;
}
</style>

<template>
  <div id="wrapper">
    <title></title>
    <div class="header box">
      <circle-button icon="play-circle">启动</circle-button>
      <circle-button icon="circle-o">编辑规则</circle-button>
    </div>
    <!-- <status></status> -->
    <ul class="list">
      <item></item>
    </ul>
  </div>
</template>

<script>
import CircleButton from './common/CircleButton.vue'
import Title from './Title.vue'
import Item from './Item.vue'
import Status from './Status.vue'

export default {
  components: {
    Title,
    CircleButton,
    Item,
    Status,
  },
}
</script>
