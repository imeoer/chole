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
.box {
  background-color: white;
  border: 2px solid #eee;
  border-radius: 5px;
  padding: 15px 20px;
}
.show-transition {
  transition: all .3s ease;
  opacity: 1;
}
.show-enter, .show-leave {
  opacity: 0;
  height: 0;
  padding: 0;
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
#wrapper {
  width: 550px;
}
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
