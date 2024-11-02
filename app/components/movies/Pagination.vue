<!-- Pagination.vue -->
<template>
    <nav aria-label="Pagination">
      <ul class="flex justify-center space-x-2">
        <li>
          <button
            @click="prevPage"
            :disabled="currentPage === 1"
            class="px-4 py-2 text-white bg-blue-500 rounded-md"
          >
            Previous
          </button>
        </li>
        <li v-for="page in totalPages" :key="page">
          <button
            @click="changePage(page)"
            :class="{
              'bg-blue-500 text-white': currentPage === page,
              'bg-white text-black': currentPage !== page
            }"
            class="px-4 py-2 rounded-md"
          >
            {{ page }}
          </button>
        </li>
        <li>
          <button
            @click="nextPage"
            :disabled="currentPage === totalPages"
            class="px-4 py-2 text-white bg-blue-500 rounded-md"
          >
            Next
          </button>
        </li>
      </ul>
    </nav>
  </template>
  
  <script setup lang="ts">
  import { ref, defineProps } from 'vue';
  
  const props = defineProps<{
    totalPages: number;
    currentPage: number;
    onPageChange: (page: number) => void;
  }>();
  
  const changePage = (page: number) => {
    props.onPageChange(page);
  };
  
  const prevPage = () => {
    if (props.currentPage > 1) {
      props.onPageChange(props.currentPage - 1);
    }
  };
  
  const nextPage = () => {
    if (props.currentPage < props.totalPages) {
      props.onPageChange(props.currentPage + 1);
    }
  };
  </script>
  
  <style scoped>
  /* Add any additional styles here */
  </style>
  