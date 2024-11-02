<template>
  <div class="min-h-screen bg-gray-100 flex items-center justify-center">
    <div class="bg-white shadow-md rounded px-8 pt-6 pb-8 mb-4 max-w-md w-full">
      <h2 class="text-2xl font-bold mb-6 text-center">Add Image</h2>
      <form @submit.prevent="submitForm">
        <!-- Movie Selection Dropdown -->
        <div class="mb-4">
          <label class="block text-gray-700 text-sm font-bold mb-2" for="movieId">Select Movie</label>
          <select v-model="selectedMovieId" id="movieId" required
                  class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline">
            <option disabled value="">Please select a movie</option>
            <option v-for="movie in movies" :key="movie.id" :value="movie.id">
              {{ movie.title }} (ID: {{ movie.id }})
            </option>
          </select>
        </div>

        <!-- Image URL Field -->
        <div class="mb-4">
          <label class="block text-gray-700 text-sm font-bold mb-2" for="imageUrl">Image URL</label>
          <input v-model="imageUrl" id="imageUrl" type="text" required
                 class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                 placeholder="Enter Image URL"/>
        </div>

        <!-- Is Featured Checkbox -->
        <div class="mb-4 flex items-center">
          <input v-model="isFeatured" id="isFeatured" type="checkbox" class="mr-2 leading-tight"/>
          <label class="text-gray-700 text-sm font-bold" for="isFeatured">Featured Image</label>
        </div>

        <!-- Submit Button -->
        <div class="flex items-center justify-between">
          <button type="submit"
                  class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline">
            Add Image
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      selectedMovieId: '',
      imageUrl: '',
      isFeatured: false,
      movies: [], // Array to hold the fetched movies
    }
  },
  methods: {
    async fetchMovies() {
  try {
    const response = await axios.get('http://localhost:8081/movies');
    console.log("Fetched movies:", response.data); // Add this line for debugging
    this.movies = response.data; // Assuming the response is an array of movie objects
  } catch (error) {
    console.error("Error fetching movies:", error);
  }
},

    async submitForm() {
      try {
        const response = await axios.post('http://localhost:8081/images', {
          movie_id: this.selectedMovieId,
          url: this.imageUrl,
          is_featured: this.isFeatured,
        });
        alert('Image added successfully');
        // Clear the form after submission
        this.selectedMovieId = '';
        this.imageUrl = '';
        this.isFeatured = false;
      } catch (error) {
        console.error("Error adding image:", error);
        alert('Failed to add image');
      }
    }
  },
  mounted() {
    this.fetchMovies(); // Fetch movies when the component is mounted
  }
};
</script>

<style scoped>
/* Add any additional styling here if needed */
</style>
