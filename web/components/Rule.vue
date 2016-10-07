<style scoped>
.item {
  position: relative;
  padding: 0;
  &:hover {
    cursor: pointer;
    background-color: #f9f9f9;
  }
  .wrap {
    padding: 20px;
    position: relative;
    z-index: 1;
  }
  .line {
    display: flex;
    align-items: center;
    justify-content: space-between;
  }
  .line1 {
    .name-wrap {
      .name {
        font-size: 22px;
        font-weight: bold;
        color: #bbb;
        text-transform: uppercase;
      }
      .out {
        font-size: 16px;
        font-weight: bold;
      }
    }
    .status {
      background-color: green;
      color: white;
      width: 30px;
      height: 30px;
      border-radius: 50%;
      text-align: center;
      line-height: 30px;
      font-size: 22px;
      padding: 4px;
      margin: 5px 0;
      &.grey {
        background-color: grey;
      }
    }
  }
  .line2 {
    font-size: 28px;
    font-weight: bold;
    margin: 20px 0;
    .arrow {
      border-bottom: 2px dotted #ddd;
      width: 100%;
      margin: 0 20px;
    }
  }
  .line3 {
    color: green;
    font-size: 18px;
    .bold {
      font-weight: bold;
    }
    .params {
      span {
        margin-left: 20px;
      }
    }
  }
}
</style>

<template>
  <li class="item box" @click="toggle">
    <div class="wrap">
      <div class="line line1">
        <span class="name-wrap">
          <div class="name">{{data.name}}</div>
          <a class="out" v-show="!expand" href="http://localhost:{{data.out}}/" target="_blank">localhost:{{data.out}}</a>
        </span>
        <i :class="['status', 'fa', data.status ? 'fa-check' : 'fa-close', { 'grey': !data.status }]"></i>
      </div>
      <div class="line line2" v-show="expand">
        <a class="out" href="http://localhost:{{data.out}}/" target="_blank">localhost:{{data.out}}</a>
        <div class="arrow"></div>
        <div class="in">{{data.in}}</div>
      </div>
      <div class="line line3" v-show="expand">
        <span class="status"><span class="bold">状态：</span>{{data.status ? '已连接' : '未连接'}}</span>
        <span class="params">
          <span class="speed"><span class="bold">速度：</span>{{speed | size}}/s</span>
          <span class="flow"><span class="bold">流量：</span>{{data.flow | size}}</span>
          <span class="eonnections"><span class="bold">连接数：</span>{{data.connections}}</span>
        </span>
      </div>
    </div>
    <chart v-show="expand" :speed="speed"></chart>
  </li>
</template>

<script>
import Vue from 'vue'
import sprintfjs from 'sprintf-js'
import Chart from './Chart.vue'

Vue.filter('size', (value) => {
  const attrs = ['B', 'kB', 'MB', 'GB', 'TB', 'PB', 'EB', 'ZB', 'YB']
  let i = 0
  while (value >= 1000) {
    value /= 1000
    i++
  }
  if (value === 0) {
    return '0'
  }
  return sprintfjs.sprintf('%.4g %s', value, attrs[i])
})

export default {
  components: {
    Chart
  },
  data: () => {
    return {
      expand: false,
      speed: 0
    }
  },
  props: {
    data: Object
  },
  methods: {
    toggle() {
      this.expand = !this.expand
    }
  },
  ready: function() {
    let lastFlow = this.data.flow
    setInterval(() => {
      const flow = this.data.flow
      this.speed = flow - lastFlow
      console.log(this.speed)
      lastFlow = flow
    }, 1000)
  }
}
</script>
