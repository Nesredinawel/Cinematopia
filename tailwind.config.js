/** @type {import('tailwindcss').Config} */
export default {
  content: ["./inedx.html","./src/**/*.{vue,js,ts}"],
  theme: {
    extend: {
      colors:{
        "bg-dodgeroll-gold":"#F79F1A",
        "apple-green":"#046E1B",
        "dire-wolf":"#292727"
      }
    },
    fontFamily:{
      Montserrat: "Montserrat, sans-serif"
    },
    container:{
      center:true,
      padding:"2rem",
    },
  },
  
  plugins: [],
}

