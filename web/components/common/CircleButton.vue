<style scoped>
button {
  @mixin hover;
  @mixin font-main;
  background-color: transparent;
  border: 0;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 0;
  .icon {
    @mixin shadow;
    background-color: #a20000;
    color: white;
    border: 0;
    width: 35px;
    height: 35px;
    display: flex;
    align-items: center;
    justify-content: center;
    border-radius: 50%;
    .fa {
      font-size: 18px;
    }
  }
  .text {
    margin-left: 10px;
    font-weight: bold;
  }
}
</style>

<template>
  <button class="button">
    <div class="icon">
      <i class="fa fa-{{icon}}" ></i>
    </div>
    <span class="text">
      <slot></slot>
    </span>
  </button>
</template>

<script>
export default {
  props: {
    icon: {
      type: String,
      required: true,
    },
  },
}
</script>
