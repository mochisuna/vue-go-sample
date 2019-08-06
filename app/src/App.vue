<template>
  <div id="app">
    <img alt="Vue logo" src="./assets/logo.png" />
    <Monster msg="Monster" />
  </div>
</template>

<script>
import Monster from "./components/Monster.vue";

export default {
  name: "app",
  components: {
    Monster
  }
};
</script>

<style>
#app {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
