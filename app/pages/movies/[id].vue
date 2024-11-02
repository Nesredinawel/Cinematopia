<template>
  <main class="py-20 p-20">
    <section v-if="movie">
     
      
      <div class="flex flex-col md:flex-row md:gap-6">
        <!-- Movie Image -->
        <img 
          v-if="movie.images && movie.images.length" 
          :src="movie.images[0].url" 
          alt="Movie image" 
          class="h-64 md:h-96 object-fill w-full md:w-1/2 rounded-lg shadow-lg"
        />
        
        <!-- Movie Details -->
        <div class="flex flex-col p-4 md:p-6 flex-1">
          <h1 class="text-3xl font-bold  mb-8">{{ movie.title }}</h1>
          <p class="text-sm text-gray-800 mb-4"><span class="font-semibold">Description:</span> {{ movie.description }}</p>
          <p class="text-gray-800 mb-2 font-semibold">Release Year: <span class="font-semibold">{{ movie.release_year }}</span></p>
          <p class="text-sm text-gray-800 mb-2"><span class="font-semibold">Genre:</span> {{ movie.genre }}</p>
          <p class="text-sm text-orange-600 mb-2"><span class="font-semibold">Rating:</span> {{ movie.rating }}</p>
          
        

          <!-- Available Schedules -->
          <div v-if="movie.schedules && movie.schedules.length" class="mt-4">
            <h3 class="text-sm font-semibold text-gray-800">Available Schedules:</h3>
            <ul class="text-sm text-white list-disc pl-4 mb-4">
              <li v-for="schedule in movie.schedules" :key="schedule.id" class=" p-5">
                {{ schedule.cinema_location }} - {{ schedule.show_time }}
                <p class="text-2xl text-cyan-200 mb-2"><span class="font-medium">Price:</span> <span class="font-bold">{{ schedule.Price }}</span> <span class="font-extralight">Birr</span></p>
              </li>
             
            </ul>
            
          </div>

          <!-- Book Now Button -->
          <button 
            @click="openBookingModal"
            class="mt-4 inline-block text-white bg-orange-600 rounded-md px-4 py-2 text-center text-base lg:text-lg"
          >
            Book Now
          </button>
        </div>
      </div>
    </section>

    <p v-else-if="loading" class="text-xl text-center text-gray-600">Loading movie details...</p>
    <p v-else class="text-xl text-center text-gray-600">Oops, something went wrong. Please try again later.</p>

    <!-- Booking Modal -->
<div v-if="isBookingModalOpen" class="fixed inset-0 bg-gray-800 bg-opacity-50 flex items-center justify-center z-50">
  <div 
    class="bg-white rounded-lg shadow-lg p-6 max-w-md md:max-w-lg w-full h-auto max-h-screen overflow-y-auto relative" 
    style="padding: 2rem"
  >
    <button @click="closeBookingModal" class="absolute top-2 right-2 text-gray-600">X</button>
    <h3 class="text-2xl font-bold mb-4">Choose Seats</h3>
    
    <div v-if="selectedSchedule" class="mb-6">
      <h4 class="text-md font-semibold mb-2">Select Seats for {{ selectedSchedule.cinema_location }} at {{ selectedSchedule.show_time }}</h4>
      <div class="grid grid-cols-5 gap-2">
        <button 
          v-for="seat in generateSeats(selectedSchedule.available_seats)" 
          :key="seat" 
          @click="toggleSeat(seat)" 
          :class="{
            'bg-gray-200': !selectedSeats.includes(seat) && !bookedSeats.includes(seat),
            'bg-orange-600 text-white': selectedSeats.includes(seat),
            'bg-red-600 text-white cursor-not-allowed': bookedSeats.includes(seat),
          }" 
          :disabled="bookedSeats.includes(seat)" 
          class="p-2 rounded-md text-sm w-full text-center"
        >
          {{ seat }}
        </button>
      </div>
    </div>

    <button 
      @click="confirmBooking" 
      class="mt-4 text-white bg-orange-600 rounded-md px-4 py-2 w-full text-center"
    >
      Confirm Booking
    </button>
  </div>
</div>

  </main>
</template>

<script>
import axios from 'axios';
import QRCode from 'qrcode';

export default {
  data() {
    return {
      movie: null,
      loading: true,
      error: false,
      allMovies: [],
      allImages: [],
      allSchedules: [],
      isBookingModalOpen: false,
      selectedSeats: [],
      selectedSchedule: null,
      bookedSeats: [],
    };
  },
  methods: {
    async fetchAllData() {
      try {
        const moviesResponse = await axios.get('http://localhost:8081/movies');
        const imagesResponse = await axios.get('http://localhost:8081/images');
        const schedulesResponse = await axios.get('http://localhost:8081/schedules');

        this.allMovies = moviesResponse.data;
        this.allImages = imagesResponse.data;
        this.allSchedules = schedulesResponse.data;

        this.setMovie();
      } catch (error) {
        console.error('Error fetching data:', error);
        this.error = true;
      } finally {
        this.loading = false;
      }
    },
    setMovie() {
      const { id } = this.$route.params;
      this.movie = this.allMovies.find(movie => movie.id === parseInt(id));
      
      if (this.movie) {
        this.movie.images = this.allImages.filter(image => image.movie_id === this.movie.id);
        this.movie.schedules = this.allSchedules.filter(schedule => schedule.movie_id === this.movie.id);
      } else {
        this.error = true;
      }
    },
    async openBookingModal() {
  if (this.movie.schedules && this.movie.schedules.length) {
    this.selectedSchedule = this.movie.schedules[0];
    await this.fetchBookedSeats(this.selectedSchedule.id); // Fetch booked seats
  }
  this.isBookingModalOpen = true;
},
async fetchBookedSeats(scheduleId) {
  try {
    // Fetch all seats from the backend
    const response = await axios.get('http://localhost:8081/seats');
    
    // Filter seats based on the current schedule_id and availability
    this.bookedSeats = response.data
      .filter(seat => seat.schedule_id === scheduleId && !seat.is_available) // Select only booked seats for the schedule
      .map(seat => seat.seat_number); // Extract the seat numbers

  } catch (error) {
    console.error('Error fetching booked seats:', error);
  }
}


,
    closeBookingModal() {
      this.isBookingModalOpen = false;
      this.selectedSeats = [];
    },
    generateSeats(availableSeats) {
      const seats = [];
      const rows = Math.ceil(availableSeats / 10);
      for (let i = 0; i < rows; i++) {
        const rowLabel = String.fromCharCode(65 + i);
        for (let j = 1; j <= 10; j++) {
          const seatLabel = `${rowLabel}${j}`;
          seats.push(seatLabel);
          if (seats.length === availableSeats) break;
        }
      }
      return seats;
    },
    toggleSeat(seat) {
      if (this.selectedSeats.includes(seat)) {
        this.selectedSeats = this.selectedSeats.filter(s => s !== seat);
      } else {
        this.selectedSeats.push(seat);
      }
    },


  async confirmBooking() {
    const email = localStorage.getItem('email'); 
    if (!email) {
      alert('User is not logged in. Please log in to book tickets.');
      return;
    }

    try {
      const usersResponse = await axios.get('http://localhost:8081/users');
      const users = usersResponse.data;
      const user = users.find(user => user.email === email);
      if (!user) {
        alert('User not found. Please check your email and try again.');
        return;
      }

      const userId = user.id;
      const quantity = this.selectedSeats.length;
      const totalPrice = quantity * this.selectedSchedule.Price;
      const qrCodeData = `User ID: ${userId}, Schedule ID: ${this.selectedSchedule.id}`;
      let qrCodeUrl;

      try {
        qrCodeUrl = await QRCode.toDataURL(qrCodeData); // Generate QR code URL
      } catch (error) {
        console.error('Error generating QR code:', error);
        alert('Could not generate QR code. Please try again.');
        return;
      }

      const ticketData = {
        user_id: userId,
        schedule_id: this.selectedSchedule.id,
        quantity: quantity,
        total_price: totalPrice,
        payment_status: 'pending',
        qr_code_url: qrCodeUrl,
      };

      const ticketResponse = await axios.post('http://localhost:8081/tickets', ticketData);
      console.log('Ticket successfully booked:', ticketResponse.data);

      const bookedSeats = this.selectedSeats.map(seatLabel => {
        const row = seatLabel.charAt(0);
        const seatColumn = parseInt(seatLabel.slice(1), 10);
        return {
          schedule_id: this.selectedSchedule.id,
          seat_number: seatLabel,
          row: row.charCodeAt(0) - 64,
          seat_column: seatColumn,
          is_available: false,
        };
      });

      for (const seat of bookedSeats) {
        await axios.post('http://localhost:8081/seats', seat);
      }

      alert(`You have successfully booked the following seats: ${this.selectedSeats.join(', ')}`);
      this.closeBookingModal();
      this.selectedSeats = [];
    } catch (error) {
      console.error('Error booking ticket:', error);
      alert('There was an issue completing your booking. Please try again.');
    }
  },

  },
  mounted() {
    this.fetchAllData();
  },
};
</script>

<style scoped>
.container {
  max-width: 900px;
}
</style>
