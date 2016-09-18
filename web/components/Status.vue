<style scoped>
.status {
  background-color: #eee;
  border: 2px solid #ddd;
  font-weight: bold;
  margin-bottom: 20px;
  text-align: center;
  display: flex;
  align-items: center;
  justify-content: center;
  .fa {
    font-size: 18px;
    margin-right: 5px;
  }
}
</style>

<template>
  <div class="status box">
    <i class="fa fa-exclamation-circle"></i>
    启动失败：端口80已经被占用
  </div>
</template>

<script>
export default {
}
</script>
