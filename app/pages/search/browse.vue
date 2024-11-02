<template>
  <div class="min-h-screen bg-gray-100 flex items-center justify-center bg-gradient-to-r from-red-700 to-indigo-600"">
    <div class="bg-white shadow-md rounded px-8 pt-6 pb-8 mb-4 w-full max-w-7xl"> <!-- Wider container -->
      <div class="grid grid-cols-1 md:grid-cols-2 gap-8"> <!-- Create two side-by-side sections -->

        <!-- Add Movie Form Section -->
        <div>
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

            <!-- Add New Director Button -->
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

        <!-- Add Images Form Section -->
         <div>
        <div>
          <h2 class="text-2xl font-bold mb-6 text-center">Add Images for Movie</h2>
          <form @submit.prevent="submitimageForm">
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

            <!-- Multiple Image URLs -->
            <div class="mb-4">
              <label class="block text-gray-700 text-sm font-bold mb-2">Image URLs</label>
              <div v-for="(image, index) in images" :key="index" class="mb-2 flex">
                <input v-model="image.url" type="text" required
                       class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                       placeholder="Enter Image URL"/>
                <button type="button" @click="removeImageField(index)"
                        class="ml-2 bg-red-500 hover:bg-red-700 text-white font-bold py-1 px-2 rounded">
                  X
                </button>
              </div>
            </div>

            <!-- Add New Image Button -->
            <div class="mb-4">
              <button type="button" @click="addImageField"
                      class="bg-green-500 hover:bg-green-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline">
                Add Another Image
              </button>
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
                Add Images
              </button>
            </div>
          </form>
        </div>
         <!-- Add Schedule Section -->
         <div>
          <h2 class="text-2xl font-bold mb-6 text-center">Add Schedule for Movie</h2>
          <form @submit.prevent="submitScheduleForm">
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

            <!-- Cinema Location Input -->
            <div class="mb-4">
              <label class="block text-gray-700 text-sm font-bold mb-2" for="cinemaLocation">Cinema Location</label>
              <input v-model="schedule.cinema_location" type="text" id="cinemaLocation" required
                     class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                     placeholder="Enter cinema location" />
            </div>

            <!-- Available Seats Input -->
            <div class="mb-4">
              <label class="block text-gray-700 text-sm font-bold mb-2" for="availableSeats">Available Seats</label>
              <input v-model="schedule.available_seats" type="number" id="availableSeats" required
                     class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                     placeholder="Enter available seats" />
            </div>
             <!-- Available Seats Input -->
             <div class="mb-4">
              <label class="block text-gray-700 text-sm font-bold mb-2" for="availableSeats">Price</label>
              <input v-model="schedule.Price" type="number" id="Price" required
                     class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                     placeholder="Enter available seats" />
            </div>

            <!-- Show Time Input -->
            <div class="mb-4">
              <label class="block text-gray-700 text-sm font-bold mb-2" for="showTime">Show Time</label>
              <input v-model="schedule.show_time" type="datetime-local" id="showTime" required
                     class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline" />
            </div>

            <!-- Submit Button -->
            <div class="flex items-center justify-between">
              <button type="submit"
                      class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline">
                Add Schedule
              </button>
            </div>
          </form>
        </div>
             <!-- add movie star form -->
         
  <div>
          <h2 class="text-2xl font-bold mb-6 text-center"> Add Star For The Movie</h2>
          <form @submit.prevent="submitMovieStarForm">

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
           

            
            <div class="mb-4">
              <label class="block text-gray-700 text-sm font-bold mb-2" for="movieId">Select Movie</label>
              <select v-model="selectedStar" id="starId" required
                      class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline">
                <option disabled value="">Please select a movie</option>
                <option v-for="star in stars" :key="star.id" :value="star.id">
                  {{ star.name }} (ID: {{ star.id }})
                </option>
              </select>
            </div>

            

            <!-- Add New Star Button -->
            <div class="mb-4">
              <button type="button" @click="addnewStar"
                      class="bg-green-500 hover:bg-green-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline">
                Add New Star
              </button>
            </div>

            <!-- Submit Button -->
            <div class="flex items-center justify-between">
              <button type="submit"
                      class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline">
                Add Star For The Movie
              </button>
            </div>
          </form>
        </div>
      </div>
    </div>
 
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
  <!-- Modal for Adding New star -->
  <div v-if="showStarModal" class="fixed z-50 inset-0 flex items-center justify-center bg-gray-600 bg-opacity-50">
    <div class="bg-white rounded-lg p-6">
      <h3 class="text-xl font-bold mb-4">Add New Star</h3>
      <form @submit.prevent="submitNewStar">
        <div class="mb-4">
          <label class="block text-gray-700 text-sm font-bold mb-2" for="newDirectorName">star Name</label>
          <input v-model="newStarName" id="newStarName" type="text" required
                 class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                 placeholder="Enter new star's name"/>
        </div>
        <div class="mb-4">
          <label class="block text-gray-700 text-sm font-bold mb-2" for="newDirectorBio">Star Bio</label>
          <textarea v-model="newStarBio" id="newStarBio" required
                    class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                    placeholder="Enter new star's bio"></textarea>
        </div>
        <div class="flex justify-end">
          <button type="submit" class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded">
            Save Director
          </button>
          <button type="button" @click="showStarModal = false" class="ml-2 bg-gray-500 hover:bg-gray-700 text-white font-bold py-2 px-4 rounded">
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
      stars:[],
      selectedDirector: '',
      selectedStar:'',
      showDirectorModal: false,
      showStarModal:false,
      newDirectorName: '',
      newDirectorBio: '', // Added new data property for the director's bio
      newStarName:'',
      newStarBio:'',
      selectedMovieId: '', // Stores the selected movie ID
      images: [{ url: '' }], // Stores multiple image URLs
      isFeatured: false, // Checkbox to indicate featured image
      movies: [] ,
      selectedMovieId: '', // Stores the selected movie ID
      schedule: { // Stores schedule details
        cinema_location: '',
        available_seats: '',
        Price: '',
        show_time: '',
        newStarName:'',
        newStarBio:'',
      },
     
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
    },
    
    async fetchMovies() {
      try {
        const response = await axios.get('http://localhost:8081/movies'); // Adjust the endpoint as needed
        this.movies = response.data; // Assuming the response is an array of movie objects
      } catch (error) {
        console.error("Error fetching movies:", error);
      }
    },
    addImageField() {
      // Add a new image URL field
      this.images.push({ url: '' });
    },
    removeImageField(index) {
      // Remove the image URL field at the given index
      if (this.images.length > 1) {
        this.images.splice(index, 1);
      }
    },
    async submitimageForm() {
      if (!this.selectedMovieId) {
        alert('Please select a movie before submitting the form');
        return;
      }

      try {
        // Iterate over the images array and submit each image to the backend
        const requests = this.images.map(image => {
          return axios.post('http://localhost:8081/images', {
            movie_id: this.selectedMovieId,
            url: image.url,
            is_featured: this.isFeatured,
          });
        });

        // Wait for all the requests to finish
        await Promise.all(requests);
        alert('Images added successfully');
        
        // Reset form
        this.selectedMovieId = '';
        this.images = [{ url: '' }];
        this.isFeatured = false;
      } catch (error) {
        console.error("Error adding images:", error);
        alert('Failed to add images');
      }
    },
    
    async fetchMovies() {
      try {
        const response = await axios.get('http://localhost:8081/movies'); // Adjust the endpoint as needed
        this.movies = response.data; // Assuming the response is an array of movie objects
      } catch (error) {
        console.error("Error fetching movies:", error);
      }
    },
    async submitScheduleForm() {
      if (!this.selectedMovieId) {
        alert('Please select a movie before submitting the form');
        return;
      }

      try {
        const scheduleData = {
          movie_id: this.selectedMovieId,
          cinema_location: this.schedule.cinema_location,
          available_seats: this.schedule.available_seats,
          show_time: this.schedule.show_time,
          Price: this.schedule.Price,
        };

        // Send schedule data to the backend
        await axios.post('http://localhost:8081/schedules', scheduleData);
        alert('Schedule added successfully');
        
        // Reset form
        this.selectedMovieId = '';
        this.schedule = {
          cinema_location: '',
          available_seats: '',
          Price: '',
          show_time: '',
        };
      } catch (error) {
        console.error("Error adding schedule:", error);
        alert('Failed to add schedule');
      }
    },
    async fetchstars() {
      try {
        const response = await axios.get('http://localhost:8081/stars');
        this.stars = response.data;
      } catch (error) {
        console.error("Error fetching stars:", error);
      }
    },
    addnewStar() {
      this.showStarModal = true;
    },
    async submitNewStar() {
      try {
        const response = await axios.post('http://localhost:8081/stars', {
          name: this.newStarName,
          bio: this.newStarBio // Include the bio in the request
        });

        // After adding the new director, fetch the updated list of directors
        this.fetchstars();

        // Close the modal and clear the new director's name and bio
        this.showStarModal = false;
        this.newStarName = '';
        this.newStarBio = ''; // Clear the bio field
        alert('star added successfully');
      } catch (error) {
        console.error("Error adding new star:", error);
        alert('Failed to add star');
      }
    },
    async submitMovieStarForm() {
      if (!this.selectedStar) {
        alert('Please select a star before submitting the form');
        return;
      }
      try {
        const moviestarData = {
          movie_id: this.selectedMovieId,
          star_id: this.selectedStar,
        };

        // Send schedule data to the backend
        await axios.post('http://localhost:8081/movie-stars', moviestarData);
        alert('Movie star added successfully');
        
        // Reset form
        this.selectedMovieId = '';
        this.selectedStar= '';
      } catch (error) {
        console.error("Error adding Star:", error);
        alert('Failed to add Star');
      }
  },
  },
  mounted() {
    this.fetchMovies(); // Fetch movies when the component is mounted
    this.fetchDirectors();
    this.fetchstars();
  }

  
};
</script>
