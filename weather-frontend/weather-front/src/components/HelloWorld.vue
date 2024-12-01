<template>
  <div>
    <h1>{{ message }}</h1>
  </div>
</template>

<script>
export default {
  data() {
    return {
      message: ''
    };
  },
  mounted() {
    this.fetchMessage();
  },
  methods: {
    async fetchMessage() {
      try {
        const response = await fetch('http://localhost:8080/api/hello');
        const data = await response.json();
        this.message = data.message; // Set the message from the response
        console.log('test');
      } catch (error) {
        console.error('Error fetching message:', error);
      }
    }
  }
};
</script>

<style scoped>
/* You can add styles here */
</style>
