<template>
  <div class="weather-container">
    <div class="glass-header">
      <div class="brand">
        <h1>WeatherScope</h1>
      </div>

      <!-- Enhanced Location Selector -->
      <div class="location-selector">
        <div class="autocomplete-wrapper">
          <div class="search-icon">
            <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <circle cx="11" cy="11" r="8"></circle>
              <line x1="21" y1="21" x2="16.65" y2="16.65"></line>
            </svg>
          </div>
          <input
            type="text"
            v-model="searchQuery"
            @focus="isOpen = true"
            @input="isOpen = true"
            placeholder="Enter city name..."
            class="location-input"
          />
          <transition name="fade">
            <div v-show="isOpen && filteredLocations.length > 0" class="autocomplete-dropdown">
              <div
                v-for="location in filteredLocations"
                :key="location"
                @click="selectLocation(location)"
                class="autocomplete-item"
                :class="{ 'active': location === selectedLocation }"
              >
                <span class="location-icon">
                  <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                    <path d="M21 10c0 7-9 13-9 13s-9-6-9-13a9 9 0 0 1 18 0z"></path>
                    <circle cx="12" cy="10" r="3"></circle>
                  </svg>
                </span>
                {{ location }}
              </div>
            </div>
          </transition>
        </div>
      </div>
    </div>

    <!-- Loading State -->
    <div v-if="loading" class="loading">
      <div class="loading-spinner"></div>
      <p>Getting your forecast ready...</p>
    </div>

    <!-- Error State -->
    <div v-if="error" class="error">
      <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
        <circle cx="12" cy="12" r="10"></circle>
        <line x1="12" y1="8" x2="12" y2="12"></line>
        <line x1="12" y1="16" x2="12.01" y2="16"></line>
      </svg>
      {{ error }}
    </div>

    <!-- Weather Cards -->
    <div v-if="filteredForecasts.length" class="forecast-list">
      <div
        v-for="forecast in filteredForecasts"
        :key="`${forecast.location.location_name}-${forecast.date}`"
        class="forecast-item"
      >
        <div class="forecast-header">
          <div class="location-info">
            <h2>{{ forecast.location.location_name }}</h2>
            <span class="date">{{ formatDate(forecast.date) }}</span>
          </div>
          <div class="weather-icon">
            <!-- Dynamic weather icon based on conditions -->
            <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <circle cx="12" cy="12" r="5"></circle>
              <line x1="12" y1="1" x2="12" y2="3"></line>
              <line x1="12" y1="21" x2="12" y2="23"></line>
              <line x1="4.22" y1="4.22" x2="5.64" y2="5.64"></line>
              <line x1="18.36" y1="18.36" x2="19.78" y2="19.78"></line>
              <line x1="1" y1="12" x2="3" y2="12"></line>
              <line x1="21" y1="12" x2="23" y2="12"></line>
              <line x1="4.22" y1="19.78" x2="5.64" y2="18.36"></line>
              <line x1="18.36" y1="5.64" x2="19.78" y2="4.22"></line>
            </svg>
          </div>
        </div>

        <div class="temp-range">
          <div class="temp min">
            <span class="label">Low</span>
            <span class="value">{{ forecast.min_temp }}°</span>
          </div>
          <div class="temp-graph">
            <div class="temp-bar" :style="{ '--min': forecast.min_temp, '--max': forecast.max_temp }"></div>
          </div>
          <div class="temp max">
            <span class="label">High</span>
            <span class="value">{{ forecast.max_temp }}°</span>
          </div>
        </div>

        <div class="forecast-details">
          <div class="forecast-time">
            <div class="time-block morning">
              <div class="time-icon">
                <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                  <path d="M17 18a5 5 0 0 0-10 0"></path>
                  <line x1="12" y1="2" x2="12" y2="9"></line>
                  <line x1="4.22" y1="10.22" x2="5.64" y2="11.64"></line>
                  <line x1="1" y1="18" x2="3" y2="18"></line>
                  <line x1="21" y1="18" x2="23" y2="18"></line>
                  <line x1="18.36" y1="11.64" x2="19.78" y2="10.22"></line>
                  <line x1="23" y1="22" x2="1" y2="22"></line>
                  <polyline points="8 6 12 2 16 6"></polyline>
                </svg>
              </div>
              <h3>Morning</h3>
              <p>{{ forecast.morning_forecast }}</p>
            </div>
            <div class="time-block afternoon">
              <div class="time-icon">
                <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                  <circle cx="12" cy="12" r="5"></circle>
                  <line x1="12" y1="1" x2="12" y2="3"></line>
                  <line x1="12" y1="21" x2="12" y2="23"></line>
                  <line x1="4.22" y1="4.22" x2="5.64" y2="5.64"></line>
                  <line x1="18.36" y1="18.36" x2="19.78" y2="19.78"></line>
                  <line x1="1" y1="12" x2="3" y2="12"></line>
                  <line x1="21" y1="12" x2="23" y2="12"></line>
                  <line x1="4.22" y1="19.78" x2="5.64" y2="18.36"></line>
                  <line x1="18.36" y1="5.64" x2="19.78" y2="4.22"></line>
                </svg>
              </div>
              <h3>Afternoon</h3>
              <p>{{ forecast.afternoon_forecast }}</p>
            </div>
            <div class="time-block night">
              <div class="time-icon">
                <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                  <path d="M21 12.79A9 9 0 1 1 11.21 3 7 7 0 0 0 21 12.79z"></path>
                </svg>
              </div>
              <h3>Night</h3>
              <p>{{ forecast.night_forecast }}</p>
            </div>
          </div>

          <div class="summary">
            <div class="summary-icon">
              <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <line x1="8" y1="6" x2="21" y2="6"></line>
                <line x1="8" y1="12" x2="21" y2="12"></line>
                <line x1="8" y1="18" x2="21" y2="18"></line>
                <line x1="3" y1="6" x2="3.01" y2="6"></line>
                <line x1="3" y1="12" x2="3.01" y2="12"></line>
                <line x1="3" y1="18" x2="3.01" y2="18"></line>
              </svg>
            </div>
            <p>{{ forecast.summary_forecast }}</p>
            <span class="summary-time">Updated {{ forecast.summary_when }}</span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import axios from "axios";
import { ref, computed, onMounted, onBeforeUnmount } from "vue";

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
const selectedLocation = ref('Kuala Lumpur');
const searchQuery = ref('');
const isOpen = ref(false);

const formatDate = (dateString: string) => {
  return new Date(dateString).toLocaleDateString('en-US', {
    weekday: 'long',
    year: 'numeric',
    month: 'long',
    day: 'numeric'
  });
};

const uniqueLocations = computed(() => {
  const locations = forecasts.value.map(forecast => forecast.location.location_name);
  return [...new Set(locations)].sort();
});

const filteredLocations = computed(() => {
  const query = searchQuery.value.toLowerCase().trim();
  if (!query) return uniqueLocations.value;
  return uniqueLocations.value.filter(location =>
    location.toLowerCase().includes(query)
  );
});

const selectLocation = (location: string) => {
  selectedLocation.value = location;
  searchQuery.value = location;
  isOpen.value = false;
};

const handleClickOutside = (event: MouseEvent) => {
  const target = event.target as HTMLElement;
  if (!target.closest('.autocomplete-wrapper')) {
    isOpen.value = false;
  }
};

onMounted(() => {
  document.addEventListener('click', handleClickOutside);
  fetchForecasts();
});

onBeforeUnmount(() => {
  document.removeEventListener('click', handleClickOutside);
});

const fetchForecasts = () => {
  loading.value = true;
  error.value = null;

  axios
    .get("http://localhost:8080/api/weather")
    .then((response) => {
      forecasts.value = response.data;
      if (!selectedLocation.value && uniqueLocations.value.length > 0) {
        const firstLocation = uniqueLocations.value[0];
        selectedLocation.value = firstLocation;
        searchQuery.value = firstLocation;
      }
    })
    .catch((err) => {
      error.value = "Failed to fetch weather data. Please try again later.";
      console.error(err);
    })
    .finally(() => {
      loading.value = false;
    });
};

const filteredForecasts = computed(() => {
  if (!selectedLocation.value) return [];

  const locationForecasts = forecasts.value.filter(
    (forecast) => forecast.location.location_name === selectedLocation.value
  );

  const uniqueForecasts = new Map();
  locationForecasts.forEach(forecast => {
    uniqueForecasts.set(forecast.date, forecast);
  });

  return Array.from(uniqueForecasts.values()).sort((a, b) =>
    new Date(a.date).getTime() - new Date(b.date).getTime()
  );
});
</script>

