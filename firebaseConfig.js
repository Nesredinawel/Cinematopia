// firebaseConfig.js for Firebase v9+
import { initializeApp } from "firebase/app";
import { getAuth } from "firebase/auth";

const firebaseConfig = {
    apiKey: "AIzaSyD7ln41ozDiYcXLe_ItYjMI6kI_hHXloLU",
    authDomain: "cinema-schedule-5cd2c.firebaseapp.com",
    projectId: "cinema-schedule-5cd2c",
    storageBucket: "cinema-schedule-5cd2c.firebasestorage.app",
    messagingSenderId: "274898147078",
    appId: "1:274898147078:web:c3768b2d263dcd16a2f087",
    measurementId: "G-9G7TE9WNFS"
};

// Initialize Firebase
const app = initializeApp(firebaseConfig);
const auth = getAuth(app);

export { auth };
