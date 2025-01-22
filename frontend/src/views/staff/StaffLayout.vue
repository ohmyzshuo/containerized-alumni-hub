<script setup>
import Nav from '../../components/Nav.vue';
</script>
<template>
  <div id="app">
    <Nav/>
    <div class="content">
      <RouterView/>
    </div>

  </div>
</template>

<style>
#app {
  display: grid;
  justify-items: center;
}

.content {
  margin-top: 60px;
  width: 100%;
  max-width: 1600px;
}
</style>
