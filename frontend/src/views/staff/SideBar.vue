<script setup>
</script>

<template>
  <div id="side-bar">
    <el-col :span="24">

      <el-menu
          active-text-color="#ffd04b"
          background-color="#38258b"
          class="el-menu-vertical-demo"
          default-active="2"
          text-color="#fff"
          @open="handleOpen"
          @close="handleClose"
      >
        <h2>Management</h2>
        <el-menu-item-group title="Alumni">
          <el-menu-item index="1-3">Information</el-menu-item>
          <el-menu-item index="1-3">Publication</el-menu-item>
        </el-menu-item-group>
        <el-menu-item-group title="Content">
          <el-menu-item index="1-1">Events</el-menu-item>
          <el-menu-item index="1-2">Announcements</el-menu-item>
          <el-menu-item index="1-3">Advertisements</el-menu-item>
        </el-menu-item-group>
        <el-menu-item-group title="Staff">
          <el-menu-item index="1-3">Staff</el-menu-item>
        </el-menu-item-group>
      </el-menu>

    </el-col>
  </div>
</template>

<style scoped>
h2 {
  padding-top: 20px;
  margin-left: 20px;
  margin-bottom: 20px;
  color: #f8f8f8;
  font-size: large;
}

#side-bar {
  height: 100vh;
  background-color: #38258b;
  overflow-y: auto;
  border-right: 60px;
}
</style>
