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
ul {
  list-style: none;
  padding: 0;
  margin: 0;
}
a {
  color: #34495e;
  text-decoration: none;
  &:hover {
    opacity: 0.9;
  }
}
input, textarea {
  outline: none;
}
.box {
  background-color: white;
  border: 1px solid #ccc;
  padding: 20px 25px;
  box-shadow: 0 2px 3px rgba(0,0,0,0.05);
}
.show-transition {
  transition: opacity .3s ease;
  opacity: 1;
}
.show-enter, .show-leave {
  opacity: 0;
}
html, body {
  @mixin font-main;
  width: 100%;
  height: 100%;
  margin: 0;
  padding: 0;
  display: flex;
  justify-content: center;
  font-size: 12px;
  color: #34495e;
  background-color: #f1f1f1;
}
#wrapper {
  width: 800px;
  margin-top: 50px;
}
</style>

<template>
  <div id="wrapper">
    <title></title>
    <!-- <status></status> -->
    <router-view></router-view>
  </div>
</template>

<script>
import store from 'store'

import Title from './Title.vue'

export default {
  store,
  components: {
    Title
  }
}
</script>
