<template>
  <router-view />
</template>

<script setup>
</script>

<style>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

body {
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, 'Helvetica Neue', Arial, sans-serif;
  background-color: #f7f8fa;
  font-size: 14px;
  color: #333;
}

.page-container {
  padding-bottom: 50px;
  min-height: 100vh;
}

.page-content {
  padding: 12px;
}

.section-title {
  font-size: 16px;
  font-weight: bold;
  margin: 16px 0 12px;
  padding-left: 4px;
  border-left: 3px solid #1989fa;
}

.van-card {
  margin-bottom: 12px;
}

.van-tag {
  margin-right: 4px;
}

.empty-container {
  padding: 40px 0;
  text-align: center;
  color: #999;
}
</style>
