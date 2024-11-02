<script setup lang="ts">
import { type MoviesResponse } from "../../../types/types.js";

const { data, error } = await useFetch<MoviesResponse>("https://api.mockfly.dev/mocks/65b3a960-b1ec-48d0-a968-76888c733843/movies");
</script>

<template>
  <main>
    <section class="py-20 container">
    

      <div v-if="!error" class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-x-4 gap-y-4 ">
        <div v-for="movie in data?.movies" :key="movie.id" class="flex flex-col shadow-lg rounded-md">
          <NuxtImg
            :src="movie.poster_path"
            sizes="xs:100vw sm:50vw lg:400px"
            format="webp"
            densities="x1"
            alt="Poster for {{ movie.original_title }}"
            class="rounded-t-md"
          />
          <div class="flex flex-col py-6 px-4 flex-1">
            <p class="text-xl lg:text-2xl font-semibold mb-2">{{ movie.original_title }}</p>
            <p class="text-sm text-gray-800 mb-4">{{ movie.release_date }}</p>
           
            <p class="text-sm text-orange-600 mb-2">
              <span class="font-semibold">Popularity:</span> {{ movie.popularity }}
            </p>
            <p class="text-sm text-orange-600 mb-2">
              <span class="font-semibold">Vote Average:</span> {{ movie.vote_average }} ({{ movie.vote_count }} votes)
            </p>
            <NuxtLink
              :to="`./movies/${movie.id}`"
              class="px-4 py-2 text-white self-start bg-orange-600 rounded-md text-base lg:text-lg cursor-pointer"
            >
              View
            </NuxtLink>
          </div>
        </div>
      </div>

      <p v-else class="text-xl">Oops, something went wrong. Please try again later or check your connection</p>
    </section>
  </main>
</template>
