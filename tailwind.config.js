/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./**/**/*.{html,js,go}"],
  theme: {
    extend: {},
  },
  plugins: [require("daisyui")],
  daisyui: {
    themes: ["nord", "night"],
  },
}

