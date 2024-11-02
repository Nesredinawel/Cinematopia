<template>
  <div class="min-h-screen bg-gradient-to-r from-red-700 to-indigo-600 flex items-center justify-center">
    <div class="bg-white shadow-md rounded px-8 pt-6 pb-8 mb-4 max-w-md w-full">
      <h2 class="text-2xl font-bold mb-6 text-center">{{ isLogin ? "Login" : "Sign Up" }}</h2>
      <form @submit.prevent="submitForm">
        <div class="mb-4">
          <label class="block text-gray-700 text-sm font-bold mb-2" for="email">Email</label>
          <input v-model="email" id="email" type="email" required
                 class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                 placeholder="Email"/>
        </div>
        <div class="mb-4" v-if="!isLogin">
          <label class="block text-gray-700 text-sm font-bold mb-2" for="name">Name</label>
          <input v-model="name" id="name" type="text" required
                 class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                 placeholder="Name"/>
        </div>
        <div class="mb-6">
          <label class="block text-gray-700 text-sm font-bold mb-2" for="password">Password</label>
          <input v-model="password" id="password" type="password" required
                 class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                 placeholder="********"/>
        </div>
        <div class="flex items-center justify-between">
          <button type="submit"
                  class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline">
            {{ isLogin ? "Login" : "Sign Up" }}
          </button>
        </div>
        <p class="text-center mt-4 text-ssm text-blue-500 hover:text-blue-700 cursor-pointer" @click="toggleMode">
          {{ isLogin ? "Create an account" : "Already have an account? Login" }}
        </p>
      </form>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import { auth } from "../../../firebaseConfig"; // Import Firebase configuration
import { createUserWithEmailAndPassword,sendEmailVerification } from 'firebase/auth';

export default {
  data() {
    return {
      isLogin: true,
      name: '',
      email: '',
      password: ''
    }
  },
  methods: {
    toggleMode() {
      this.isLogin = !this.isLogin;
    },
    async submitForm() {
      if (this.isLogin) {
        await this.login();
      } else {
        await this.signup();
      }
    },
    async signup() {
      try {
        // Step 1: Create user with email and password
        const userCredential = await createUserWithEmailAndPassword(auth, this.email, this.password);
        const user = userCredential.user;

        // Step 2: Send verification email
        await sendEmailVerification(user);
        alert("Verification email sent! Please verify your email before logging in.");

        // Step 3: Wait for email verification
    const verificationInterval = setInterval(async () => {
      await user.reload();
      if (user.emailVerified) {
        clearInterval(verificationInterval);

        // Step 4: After verification, post data to your backend
        await axios.post("http://localhost:8081/signup", {
          name: this.name,
          email: this.email,
          password: this.password
        });

        // Store email in local storage
        localStorage.setItem("email", this.email);
        alert("Email verified and user created successfully!");
      }
    }, 3000); 
      } catch (error) {
        console.error(error);
        alert("Signup or verification failed");
      }
    }
,
    async login() {
      try {
        const response = await axios.post("http://localhost:8081/login", {
          email: this.email,
          password: this.password
        });
        const token = response.data.token;
        localStorage.setItem("token", token);
        localStorage.setItem("email", this.email);
        alert("Login successful");
      } catch (error) {
        console.error(error);
        alert("Login failed");
      }
    }
  }
};
</script>
