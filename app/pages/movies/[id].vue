<script setup lang="ts">
import { type Movie } from "../../../types/types.js";
import { ref } from 'vue';
import BuyTicketModal from '../../components/movies/BuyTicketModal.vue'; // Adjust the path as necessary
import Rating from '../../components/movies/Rating.vue';
const { id } = useRoute().params;
const { data, error } = await useFetch<Movie>(`https://api.mockfly.dev/mocks/65b3a960-b1ec-48d0-a968-76888c733843/movies/${id}`);

if (error.value) {
  throw createError({
    statusCode: error.value?.statusCode,
    statusMessage: error.value?.statusMessage,
  });
}

// Manage the modal visibility state
const showModal = ref(false);

// Define the close function
const closeModal = () => {
  showModal.value = false; // This will close the modal
};

// Handle rating
const handleRating = (rating: number) => {
  console.log(`User rated the movie: ${rating}`); // Replace with your logic
};


</script>

<template>
  <div class="container py-20 space-y-8">
    <div class="flex items-center justify-between gap-10">
      <!-- Header -->
      <div class="text-center">
        <h2 class="text-5xl mb-4 font-semibold text-gray-900">{{ data?.original_title }}</h2>
        <p class="text-lg text-gray-600 italic">{{ data?.release_date }} | {{ data?.original_language.toUpperCase() }}</p>
        <p class="mt-2 text-sm text-gray-500">Rating: {{ data?.vote_average }} ({{ data?.vote_count }} votes) | Popularity: {{ data?.popularity }}</p>
        <hr class="mt-6 border-gray-300" />
      </div>

      <!-- Image -->
      <NuxtImg
        :src="data?.poster_path"
        densities="x1"
        sizes="xs:100vw sm:100vw md:100vw lg:100vw"
        class="w-[60%] max-h-[400px] object-fill rounded-md shadow-md mb-12"
        alt="Poster for {{ data?.original_title }}"
      />
    </div>

    <div class="flex gap-9">
      <!-- Overview -->
      <div class="text-left">
        <h3 class="text-3xl font-semibold mb-4 text-gray-800">Overview</h3>
        <p class="text-lg text-white leading-relaxed">{{ data?.overview }}</p>
      </div>

      <!-- Buy Ticket Button -->
      <div class="md:ml-6 flex flex-col px-24 md:items-center">
        <button 
          class="w-40 py-3 bg-orange-600 text-white text-lg font-semibold rounded-md hover:bg-orange-700 transition"
          @click="showModal = true" <!-- Open modal on click -->
        >
          Buy Ticket
        </button>
        <p class="mt-4 text-sm text-black">Hurry! Limited seats available.</p>
      </div>
    </div>
    <!-- rating movie compo -->
    <Rating @onRate="handleRating" />
    <!-- Cast Section -->
    <div v-if="data?.casts.length" class="space-y-6">
      <h3 class="text-3xl font-semibold mb-4 text-gray-800">Cast</h3>
      <div class="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-8 gap-x-4 gap-y-8">
        <div v-for="cast in data?.casts" :key="cast.id" class="flex flex-col items-center shadow-lg p-4 rounded-md bg-slate-300">
          <NuxtImg
            :src="cast.profile_path"
            alt="NO Avatar"
            class="object-cover rounded-t-md mb-4"
          />
          <p class="text-sm font-semibold text-gray-900">{{ cast.name }}</p>
          <p class="text-xs text-gray-600">as {{ cast.character }}</p>
          <p class="text-xs text-gray-500 mt-1">Popularity: {{ cast.popularity }}</p>
        </div>
      </div>
    </div>
    
    <!-- Buy Ticket Modal -->
    <BuyTicketModal
      v-if="showModal"
      :movieName="data?.original_title"
      :isVisible="showModal"
      :onClose="closeModal" 
    />
  </div>
</template>

<style scoped>

</style>
