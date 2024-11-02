<template>
  <div>
    <MovieDetail :movie="selectedMovie" />
  </div>
</template>

<script setup lang="ts">
import MovieDetail from './[id].vue';

const selectedMovie = {
  id: 18502,
  original_title: "The Stuff",
  overview: "Amalgamated Dairies hires David Rutherford...",
  poster_path: "https://image.tmdb.org/t/p/original/bT0PfxX8YrZHGaaXau9RLN.jpg",
  release_date: "Fri, 06/14/1985",
  vote_average: 5.9,
  popularity: 9.94,
  casts: [
    {
      id: "11ca6d2d-3417-4e5c-bdd4-342ea0825db7",
      name: "Brooke Adams",
      profile_path: "https://image.tmdb.org/t/p/original/edvZi720lk3gkYVGuvMyz3R2nh.jpg",
      character: "Stuff Commercial Spokeswoman"
    }
  ]
};
</script>
