<!-- Rating.vue -->
<template>
    <div class="flex items-center space-x-2">
      <span class="text-lg font-semibold text-orange-600">Rate this movie:</span>
      <div class="flex items-center">
        <template v-for="star in stars" :key="star.id">
          <span
            @click="setRating(star.value)"
            class="cursor-pointer text-2xl"
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              class="h-6 w-6"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
              :class="star.value <= currentRating ? 'text-yellow-500' : 'text-gray-300'"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M12 17.27L18.18 21l-1.64-7.03L22 9.24l-7.19-.61L12 2 9.19 8.63 2 9.24l5.46 4.73L5.82 21z"
              />
            </svg>
          </span>
        </template>
      </div>
    </div>
  </template>
  
  <script setup lang="ts">
  import { ref, defineProps } from 'vue';
  
  // Props
  const props = defineProps<{
    onRate: (rating: number) => void; // Function to handle rating submission
  }>();
  
  // Rating logic
  const currentRating = ref(0);
  const stars = Array.from({ length: 5 }, (_, i) => ({
    id: i + 1,
    value: i + 1,
  }));
  
  const setRating = (rating: number) => {
    currentRating.value = rating;
    props.onRate(rating); // Call the rating function passed from parent
  };
  </script>
  
  <style scoped>
  /* Add any additional styles if necessary */
  </style>
  