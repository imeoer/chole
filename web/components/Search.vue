<style scoped>
.wrap {
  display: flex;
  margin-top: 50px;
  margin-bottom: 20px;
  position: relative;
  display: none;
  input {
    width: 100%;
    padding: 15px 20px;
    border: 1px solid #ddd;
    font-size: 22px;
    &:focus {
      border: 1px solid #ccc;
      background-color: white;
    }
  }
  .fa {
    position: absolute;
    font-size: 24px;
    right: 20px;
    top: 17px;
    color: #ccc;
  }
}
</style>

<template>
  <div class="wrap">
    <input type="text" class="box" placeholder="搜索...">
    <i class="fa fa-search"></i>
  </div>
</template>

<script>
export default {
}
</script>
