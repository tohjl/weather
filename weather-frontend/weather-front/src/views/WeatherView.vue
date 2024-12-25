<template>
  <div class="weather-container">
    <!-- Your weather forecast template here -->
  </div>
</template>
<script setup lang="ts">
</script>
<script >
import WeatherForecast from '../components/WeatherForecast.vue'; // Import your component

export default {
  name: 'WeatherView',
  components: {
    // eslint-disable-next-line vue/no-unused-components
    WeatherForecast,
  },
};
</script>
