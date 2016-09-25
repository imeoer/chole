<style scoped>
header {
  text-align: center;
  margin-bottom: 50px;
  .head {
    display: flex;
    justify-content: center;
    align-items: center;
    font-size: 28px;
    font-weight: bold;
    .fa {
      font-size: 38px;
      color: #a20000;
    }
    .title {
      margin-left: 15px;
    }
  }
  .version {
    font-size: 12px;
    color: #aaa;
    margin-top: 5px;
  }
}
</style>

<template>
  <header>
    <div class="head">
      <i class="fa fa-heartbeat"></i>
      <span class="title">内网映射控制</span>
    </div>
    <div class="version">RELEASE 2016-09-18</div>
  </header>
</template>

<script>
export default {
}
</script>
