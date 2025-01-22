<template>
  <div class="app-container">
    <Nav/>
    <div class="content-container">
      <RouterView/>
    </div>
  </div>
</template>

<style>

.content-container {
  margin-top: 60px;
  position: relative;
  flex: 1;
  height: 100%;
  width: 100%;
}

.content-container :deep(> *) {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100%;
  width: 100%;
  position: relative;
  z-index: 2;
}
</style>

<script setup>
import Nav from "@/components/Nav.vue";
</script>
