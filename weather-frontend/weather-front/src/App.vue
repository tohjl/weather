<template>
  <div id="app">
    <WeatherForecast />
  </div>
</template>

<script>
import WeatherForecast from "./components/WeatherForecast.vue";

export default {
  components: {
    WeatherForecast,
  },
};
</script>

<style scoped>
/* Make sure the background is not black by default */
#app {
  background-color: #fff; /* Set a default white background */
}
</style>
