<style scoped>
.line {
  display: flex;
  align-items: center;
  justify-content: space-between;
}
.line1 {
  .name {
    font-size: 22px;
    font-weight: bold;
    color: #aaa;
  }
  .status {
    background-color: green;
    color: white;
    width: 35px;
    height: 35px;
    border-radius: 50%;
    text-align: center;
    line-height: 35px;
    font-size: 18px;
  }
}
.line2 {
  font-size: 22px;
  font-weight: bold;
  margin: 20px 0;
  .arrow {
    border-bottom: 2px dotted #ddd;
    width: 100%;
    margin: 0 20px;
  }
}
.line3 {
  color: #aaa;
  .params {
    span {
      margin-left: 10px;
    }
  }
}
</style>

<template>
  <li class="item box">
    <div class="line line1">
      <div class="name">SSH</div>
      <i class="status fa fa-check"></i>
    </div>
    <div class="line line2">
      <div class="out">web.chole.io</div>
      <div class="arrow"></div>
      <div class="in">localhost:80</div>
    </div>
    <div class="line line3">
      <span class="status">状态：连接中...</span>
      <span class="params">
        <span class="speed">速度：213 kB/s</span>
        <span class="flow">流量：23 MB</span>
        <span class="eonnections">连接数：23</span>
      </span>
    </div>
  </li>
</template>

<script>
export default {
}
</script>
