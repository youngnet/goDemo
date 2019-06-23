<template>
  <div>
    hello world
  </div>
</template>
<script>
import { getUser } from "../../api";
export default {
  async created() {
    getUser()
    // let res = await getUser();
    // console.log("TCL: create -> res", res);
  },
  data() {
    return {};
  }
};
</script>
<style scoped lang='less'>
div {
  color: red;
}
</style>
