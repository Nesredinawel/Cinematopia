<template>
  <div class="min-h-screen bg-gray-100 flex items-center justify-center">
    <div class="bg-white shadow-md rounded px-8 pt-6 pb-8 mb-4 max-w-md w-full">
      <h2 class="text-2xl font-bold mb-6 text-center">Add Movie</h2>
      <form @submit.prevent="submitForm">
        <!-- Title Field -->
        <div class="mb-4">
          <label class="block text-gray-700 text-sm font-bold mb-2" for="title">Title</label>
          <input v-model="title" id="title" type="text" required
                 class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                 placeholder="Movie Title"/>
        </div>

        <!-- Description Field -->
        <div class="mb-4">
          <label class="block text-gray-700 text-sm font-bold mb-2" for="description">Description</label>
          <textarea v-model="description" id="description" required
                    class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                    placeholder="Movie Description"></textarea>
        </div>

        <!-- Duration Field -->
        <div class="mb-4">
          <label class="block text-gray-700 text-sm font-bold mb-2" for="duration">Duration (minutes)</label>
          <input v-model="duration" id="duration" type="number" required
                 class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                 placeholder="Duration"/>
        </div>

        <!-- Release Year Field -->
        <div class="mb-4">
          <label class="block text-gray-700 text-sm font-bold mb-2" for="releaseYear">Release Year</label>
          <input v-model="releaseYear" id="releaseYear" type="number" required
                 class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                 placeholder="Release Year"/>
        </div>

        <!-- Genre Field -->
        <div class="mb-4">
          <label class="block text-gray-700 text-sm font-bold mb-2" for="genre">Genre</label>
          <input v-model="genre" id="genre" type="text" required
                 class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                 placeholder="Genre"/>
        </div>

        <!-- Rating Field -->
        <div class="mb-4">
          <label class="block text-gray-700 text-sm font-bold mb-2" for="rating">Rating (out of 10)</label>
          <input v-model="rating" id="rating" type="number" step="0.1" max="10" min="0"
                 class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                 placeholder="Rating"/>
        </div>

        <!-- Director ID Dropdown -->
        <div class="mb-4">
          <label class="block text-gray-700 text-sm font-bold mb-2" for="directorId">Director</label>
          <select v-model="selectedDirector" id="directorId" class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline">
            <option disabled value="">Please select a director</option>
            <option v-for="director in directors" :key="director.id" :value="director.id">
              {{ director.name }}
            </option>
          </select>
        </div>

        <!-- Button to Add New Director -->
        <div class="mb-4">
          <button type="button" @click="addNewDirector"
                  class="bg-green-500 hover:bg-green-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline">
            Add New Director
          </button>
        </div>

        <!-- Submit Button -->
        <div class="flex items-center justify-between">
          <button type="submit"
                  class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline">
            Add Movie
          </button>
        </div>
      </form>
    </div>

    <!-- Modal for Adding New Director -->
    <div v-if="showDirectorModal" class="fixed z-50 inset-0 flex items-center justify-center bg-gray-600 bg-opacity-50">
      <div class="bg-white rounded-lg p-6">
        <h3 class="text-xl font-bold mb-4">Add New Director</h3>
        <form @submit.prevent="submitNewDirector">
          <div class="mb-4">
            <label class="block text-gray-700 text-sm font-bold mb-2" for="newDirectorName">Director Name</label>
            <input v-model="newDirectorName" id="newDirectorName" type="text" required
                   class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                   placeholder="Enter new director's name"/>
          </div>
          <div class="mb-4">
            <label class="block text-gray-700 text-sm font-bold mb-2" for="newDirectorBio">Director Bio</label>
            <textarea v-model="newDirectorBio" id="newDirectorBio" required
                      class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                      placeholder="Enter new director's bio"></textarea>
          </div>
          <div class="flex justify-end">
            <button type="submit" class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded">
              Save Director
            </button>
            <button type="button" @click="showDirectorModal = false" class="ml-2 bg-gray-500 hover:bg-gray-700 text-white font-bold py-2 px-4 rounded">
              Cancel
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      title: '',
      description: '',
      duration: '',
      releaseYear: '',
      genre: '',
      rating: '',
      directors: [],
      selectedDirector: '',
      showDirectorModal: false,
      newDirectorName: '',
      newDirectorBio: '' // Added new data property for the director's bio
    }
  },
  methods: {
    async fetchDirectors() {
      try {
        const response = await axios.get('http://localhost:8081/directors');
        this.directors = response.data;
      } catch (error) {
        console.error("Error fetching directors:", error);
      }
    },
    addNewDirector() {
      this.showDirectorModal = true;
    },
    async submitNewDirector() {
      try {
        const response = await axios.post('http://localhost:8081/directors', {
          name: this.newDirectorName,
          bio: this.newDirectorBio // Include the bio in the request
        });

        // After adding the new director, fetch the updated list of directors
        this.fetchDirectors();

        // Close the modal and clear the new director's name and bio
        this.showDirectorModal = false;
        this.newDirectorName = '';
        this.newDirectorBio = ''; // Clear the bio field
        alert('Director added successfully');
      } catch (error) {
        console.error("Error adding new director:", error);
        alert('Failed to add director');
      }
    },
    async submitForm() {
      if (!this.selectedDirector) {
        alert('Please select a director before submitting the form');
        return;
      }

      try {
        const response = await axios.post('http://localhost:8081/movies', {
          title: this.title,
          description: this.description,
          duration: this.duration,
          release_year: this.releaseYear,
          genre: this.genre,
          rating: this.rating,
          director_id: this.selectedDirector
        });
        alert('Movie added successfully');
      } catch (error) {
        console.error("Error adding movie:", error);
        alert('Failed to add movie');
      }
    }
  },
  mounted() {
    this.fetchDirectors();
  }

  
};
</script>
