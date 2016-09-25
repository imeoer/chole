<style>
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
    color: white;
    border: 0;
    width: 35px;
    height: 35px;
    display: flex;
    align-items: center;
    justify-content: center;
    border-radius: 50%;
    background-color: #d81919;
    .fa {
      font-size: 17px;
    }
  }
  .text {
    margin-left: 10px;
    font-weight: bold;
  }
  &.green {
    .icon {
      background-color: green;
    }
  }
  &.grey {
    .icon {
      background-color: grey;
    }
  }
  &.orange {
    .icon {
      background-color: #f97f17;
    }
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
      required: true
    }
  }
}
</script>
