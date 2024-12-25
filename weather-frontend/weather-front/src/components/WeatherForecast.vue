<template>
  <div class="weather-container">
    <h1>Weather Forecast</h1>
    <div v-if="loading" class="loading">Loading...</div>
    <div v-if="error" class="error">{{ error }}</div>
    <div v-if="filteredForecasts.length" class="forecast-list">
      <div
        v-for="(forecast, index) in filteredForecasts"
        :key="index"
        class="forecast-item"
      >
        <h2>{{ forecast.location.location_name }}</h2>
        <p><strong>Date:</strong> {{ forecast.date }}</p>
        <p><strong>Morning:</strong> {{ forecast.morning_forecast }}</p>
        <p><strong>Afternoon:</strong> {{ forecast.afternoon_forecast }}</p>
        <p><strong>Night:</strong> {{ forecast.night_forecast }}</p>
        <p><strong>Summary:</strong> {{ forecast.summary_forecast }} ({{ forecast.summary_when }})</p>
        <p><strong>Temperature:</strong> {{ forecast.min_temp }}°C - {{ forecast.max_temp }}°C</p>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import axios from "axios";
import { ref, computed, onMounted } from "vue";


// Define the structure of the forecast data
interface Location {
  location_name: string;
}

interface Forecast {
  location: Location;
  date: string;
  morning_forecast: string;
  afternoon_forecast: string;
  night_forecast: string;
  summary_forecast: string;
  summary_when: string;
  min_temp: number;
  max_temp: number;
}

const loading = ref(false);
const error = ref<string | null>(null);
const forecasts = ref<Forecast[]>([]);

const fetchForecasts = () => {
  loading.value = true;
  error.value = null; // Clear error before making a new request
  axios
    .get("http://localhost:8080/api/weather")
    .then((response) => {
      forecasts.value = response.data; // Assuming the API returns an array of forecast objects
    })
    .catch((err) => {
      error.value = "Failed to fetch weather data. Please try again later.";
      console.error(err);
    })
    .finally(() => {
      loading.value = false;
    });
};

// Use onMounted to call fetchForecasts when the component mounts
onMounted(fetchForecasts);

// Filter forecasts based on location
const filteredForecasts = computed(() => {
  return forecasts.value.filter(
    (forecast) => forecast.location.location_name === "Cheras"
  );
});
</script>

<style scoped>
.weather-container {
  font-family: Arial, sans-serif;
  max-width: 1000px;
  margin: 0 auto;
  padding: 20px;
  text-align: left;
  background-color: #f7f9fc;
  border-radius: 10px;
}

h1 {
  text-align: center;
  color: #333;
}

.loading {
  color: blue;
  text-align: center;
}

.error {
  color: red;
  text-align: center;
}

.forecast-list {
  display: flex; /* Display cards in a row */
  flex-wrap: wrap; /* Allow wrapping to the next line if needed */
  gap: 15px;
  justify-content: center; /* Center-align cards */
}

.forecast-item {
  border: 1px solid #ccc;
  padding: 15px;
  border-radius: 5px;
  background-color: #ffffff;
  box-shadow: 0px 2px 4px rgba(0, 0, 0, 0.1);
  width: 300px; /* Fixed width for each card */
  flex: 1 1 calc(33.333% - 30px); /* 3 cards per row, with spacing */
  box-sizing: border-box;
}

.forecast-item h2 {
  color: #2c3e50;
}

.forecast-item p {
  margin: 5px 0;
  color: #34495e;
}

.forecast-item p strong {
  color: #2980b9;
}
</style>
