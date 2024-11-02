<script setup lang="ts">
import MovieList from '~/components/movies/MovieList.vue';


</script>



<template>
    <main>
      <section class="">
         <div class="font-Montserrat ">
       <div class="container flex flex-col lg:flex-row items-center py-20 gap-10">
         <div class="flex-1 order-2 lg:order-1 text-center lg:text-left">
          <h1 class="text-4xl lg:text-6xl font-extrabold mb-6 text-balance">
            “Book Your Cinematic Experience: Simple, Fast!!”
          </h1>
          <p class="text-xl lg:text-2xl mb-8 text-balance">
            Catch the Latest Blockbusters: Seamless Ticket Booking for an Unforgettable Movie Night!
          </p>
         
          <NuxtLink  class="px-4 py-2 text-white self-start bg-orange-600 rounded-md text-lg cursor-pointer hover:opacity-80" to="/search/browse" >Browse Movies</NuxtLink>
         
        </div>

        <div class="flex-1 order-1 lg:order-2 rounded-lg" >
          <NuxtImg
          class="rounded-lg"
            sizes="xs:100vw sm:667px"
            src="/images/cinema_img.jpg"
            format="webp"
            densities="x1"
            alt=""
          />
        </div>

       </div>
    
    </div>
      </section>

      <section class="py-20 container">
         <h2 class="text-3xl lg:text-5xl mb-2"> Discover Your Movies</h2>
         <p class="text-lg lg:text-xl mb-8">Check out the latest movies!</p>
    

<!-- movie lists -->
 <MovieList/>

      </section>
    </main>
  </template>
  