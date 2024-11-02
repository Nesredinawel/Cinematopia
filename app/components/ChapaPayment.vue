<template>
    <div>
      <!-- Payment Form Container -->
      <div id="chapa-inline-form"></div>
  
      <!-- Payment Button to Trigger Form -->
      <button @click="initializePayment" class="pay-button">Pay Now</button>
    </div>
  </template>
  
  <script>
  export default {
    name: 'ChapaPayment',
    data() {
      return {
        chapa: null,
      };
    },
    methods: {
      initializePayment() {
        // Configure and initialize Chapa checkout
        this.chapa = new ChapaCheckout({
          publicKey: process.env.CHAPA_PUBLIC_KEY,  // Load from environment variable
          amount: '100',                                     // Specify the amount in ETB
          currency: 'ETB',
          availablePaymentMethods: ['telebirr', 'cbebirr', 'ebirr', 'mpesa', 'chapa'],
          customizations: {
            buttonText: 'Pay Now',
            styles: `
              .chapa-pay-button { 
                  background-color: #4CAF50; 
                  color: white;
              }
            `,
          },
          callbackUrl: `${window.location.origin}/callback`, // Define your callback URL
          returnUrl: `${window.location.origin}/success`,    // Define your return URL
          onSuccessfulPayment: this.handleSuccess,           // Callback for success
          onPaymentFailure: this.handleFailure,              // Callback for failure
          onClose: this.handleClose,                         // Callback when the popup is closed
        });
  
        // Render the payment form
        this.chapa.initialize('chapa-inline-form');
      },
      handleSuccess() {
        alert("Payment was successful!");
        // Optionally, add additional success handling code here
      },
      handleFailure() {
        alert("Payment failed. Please try again.");
        // Optionally, add additional failure handling code here
      },
      handleClose() {
        console.log("Payment popup closed");
        // Optionally, add any cleanup or close handling here
      }
    },
  };
  </script>
  
  <style scoped>
  .pay-button {
    padding: 10px 20px;
    background-color: #4CAF50;
    color: white;
    border: none;
    cursor: pointer;
    font-size: 16px;
  }
  </style>
  