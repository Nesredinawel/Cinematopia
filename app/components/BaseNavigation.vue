<script setup lang="ts"></script>

<template>
  <header class="border-b border-b-red-600 border-opacity-35 shadow-2xl">
    <nav class="flex container items-center ">
      <NuxtLink to="/" class="flex  items-center">
        <NuxtImg width="100px" src="/cinemalogo.svg" alt="" />
        <span class="text-3xl font-bold">Cinematopia</span>
      </NuxtLink>
      <ul class="flex gap-6 ml-auto text-xl font-bold capitalize">
        <li>
          <NuxtLink to="/">Home</NuxtLink>
        </li>
        <li>
            <NuxtLink to="/search/browse" class="hover:text-gray-400">Browse Movies</NuxtLink>
      
        </li>
        <li>
            <NuxtLink to="/profile/profile" class="hover:text-gray-400">Profile</NuxtLink>
          
        </li>
        <li>
            <NuxtLink to="/auth/login" class="hover:text-gray-400">Login</NuxtLink>
        </li>
      </ul>
    </nav>
  </header>
</template>