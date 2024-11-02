<template>
  <main class="py-20 container ">
    <section>
      <h1 class="text-3xl font-bold text-center mb-8">Movies List</h1>
      
      <div v-if="!error" class="grid grid-cols-2 md:grid-cols-2 lg:grid-cols- gap-6">
        <!-- Movie Card -->
        <div v-for="movie in movies" :key="movie.id" class="flex flex-col shadow-lg rounded-lg overflow-hidden">
          <!-- Movie Image -->
          <img 
            v-if="movie.images && movie.images.length"
            :src="movie.images[0].url" 
            alt="Movie image"
            class="h-64 object-cover w-full"
          />
          
          <!-- Movie Details -->
          <div class="flex flex-col p-4 flex-1">
            <h2 class="text-xl font-semibold text-gray-900 mb-1">{{ movie.title }}</h2>
            <p class="text-gray-700 text-sm mb-2">{{ movie.releaseYear }}</p>
            
            <p class="text-sm text-gray-800 mb-2">
              <span class="font-semibold">Genre:</span> {{ movie.genre }}
            </p>
            <p class="text-sm text-orange-600 mb-2">
              <span class="font-semibold">Rating:</span> {{ movie.rating }}
            </p>
            
            <!-- Schedule (if available) -->
            <div v-if="movie.schedules && movie.schedules.length" class="mt-2">
              <h3 class="text-sm font-semibold text-gray-600">Available Schedules:</h3>
              <ul class="text-sm text-gray-700 list-disc pl-4">
                <li v-for="schedule in movie.schedules" :key="schedule.id">
                  {{ schedule.cinema_location }} - {{ schedule.show_time }}
                </li>
              </ul>
            </div>

            <!-- View Details Link -->
            <NuxtLink
              :to="`/movies/${movie.id}`"
              class="mt-4 inline-block text-white bg-orange-600 rounded-md px-4 py-2 text-center text-base lg:text-lg"
            >
              View Details
            </NuxtLink>
          </div>
        </div>
      </div>

      <p v-else class="text-xl text-center text-gray-600">Oops, something went wrong. Please try again later or check your connection.</p>
    </section>
  </main>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      movies: [],
    };
  },
  methods: {
    async fetchMoviesData() {
      try {
        // Fetch movies data with embedded images and schedules
        const response = await axios.get('http://localhost:8081/movies');

        // Fetch all related images, schedules, and directors separately
        const imagesResponse = await axios.get('http://localhost:8081/images');
        const schedulesResponse = await axios.get('http://localhost:8081/schedules');
        const directorsResponse = await axios.get('http://localhost:8081/directors');

        const images = imagesResponse.data;
        const schedules = schedulesResponse.data;
        const directors = directorsResponse.data;

        // Map over each movie to attach images, schedules, and director by movieId,
        // but only include movies that have both images and schedules
        this.movies = response.data
          .map(movie => {
            // Attach images and schedules related to the movie
            movie.images = images.filter(image => image.movie_id === movie.id);
            movie.schedules = schedules.filter(schedule => schedule.movie_id === movie.id);

            // Attach the director related to the movie
            movie.director = directors.find(director => director.id === movie.director_id);

            return movie;
          })
          .filter(movie => movie.images.length > 0 && movie.schedules.length > 0);
      } catch (error) {
        console.error('Error fetching movies or related data:', error);
      }
    },
  },
  mounted() {
    this.fetchMoviesData();
  },
};
</script>

<style scoped>
.container {
  max-width: 900px;
}
</style>
