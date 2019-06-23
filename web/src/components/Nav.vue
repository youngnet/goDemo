<template>
  <div class="navList">
    <router-link to="/">Home</router-link>
    <router-link to="/about">About</router-link>
  </div>
</template>
<script>
export default {
  data() {
    return {};
  }
};
</script>
<style scoped lang='less'>
.navList {
  display: flex;
  align-items: center;
  justify-content: space-around;
}
</style>
