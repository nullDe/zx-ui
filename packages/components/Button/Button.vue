<template>
  <button style="background-color: chocolate;color: #fff;">
    这是一个按钮
    <slot></slot>
  </button>
</template>

<script setup lang="ts">
  defineOptions({ name: 'zxButton' })
</script>

<style scoped>

</style>