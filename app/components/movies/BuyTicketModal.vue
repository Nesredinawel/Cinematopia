<template>
    <div class="modal-overlay" v-if="isVisible" @click="closeModal">
      <div class="modal-content" @click.stop>
        <header class="modal-header">
          <h2 class="text-2xl font-semibold">{{ movieName }}</h2>
          <button class="close-button" @click="closeModal">✖️</button>
        </header>
        <div class="seat-chooser">
          <h3 class="text-xl mb-4">Choose Your Seats</h3>
          <div class="grid grid-cols-5 gap-4">
            <div
              v-for="seat in seats"
              :key="seat.id"
              :class="['seat', { 'selected': seat.selected, 'occupied': seat.occupied }]"
              @click="toggleSeat(seat)"
            >
              {{ seat.number }}
            </div>
          </div>
        </div>
        <button
          class="purchase-button mt-6"
          :disabled="!selectedSeats.length"
          @click="purchaseTickets"
        >
          Purchase Tickets
        </button>
        <p v-if="selectedSeats.length" class="mt-4">Selected Seats: {{ selectedSeats.join(', ') }}</p>
      </div>
    </div>
  </template>
  
  <script setup lang="ts">
  import { ref, defineProps } from 'vue';
  
  // BuyTicketModal.vue
const props = defineProps<{
  movieName: string;
  isVisible: boolean;
  onClose: () => void; // Ensure this is correctly defined
}>();

  
  // Seat selection logic
  const seats = ref(Array.from({ length: 20 }, (_, i) => ({
    id: i,
    number: (i + 1).toString(),
    occupied: Math.random() < 0.3, // Randomly occupy some seats
    selected: false,
  })));
  
  const selectedSeats = ref<string[]>([]);
  
  const toggleSeat = (seat: { id: number; number: string; occupied: boolean; selected: boolean }) => {
    if (!seat.occupied) {
      seat.selected = !seat.selected;
      if (seat.selected) {
        selectedSeats.value.push(seat.number);
      } else {
        selectedSeats.value = selectedSeats.value.filter(s => s !== seat.number);
      }
    }
  };
  
  const purchaseTickets = () => {
    alert(`Tickets purchased for seats: ${selectedSeats.value.join(', ')}`);
    closeModal();
  };
  
  const closeModal = () => {
    props.onClose();
  };
  </script>
  
  <style scoped>
  .modal-overlay {
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background: rgba(0, 0, 0, 0.5);
    display: flex;
    justify-content: center;
    align-items: center;
  }
  
  .modal-content {
    background: white;
    padding: 2rem;
    border-radius: 8px;
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
    width: 90%;
    max-width: 600px;
  }
  
  .modal-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
  }
  
  .close-button {
    background: none;
    border: none;
    font-size: 1.5rem;
    cursor: pointer;
  }
  
  .seat-chooser {
    margin-top: 1.5rem;
  }
  
  .seat {
    background: #f0f0f0;
    padding: 1rem;
    text-align: center;
    border-radius: 4px;
    cursor: pointer;
    user-select: none;
  }
  
  .seat.selected {
    background: #4caf50;
    color: white;
  }
  
  .seat.occupied {
    background: #f44336;
    color: white;
    cursor: not-allowed;
  }
  
  .purchase-button {
    background: #2196f3;
    color: white;
    padding: 0.75rem 1.5rem;
    border: none;
    border-radius: 4px;
    cursor: pointer;
    transition: background 0.3s;
  }
  
  .purchase-button:disabled {
    background: #9e9e9e;
    cursor: not-allowed;
  }
  </style>
  