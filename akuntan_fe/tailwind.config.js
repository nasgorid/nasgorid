/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['index.html','login.html','dashboard.html'],
  theme: {
    container: {
      center: true,
      padding: '16px'
    },
    extend: {
      fontFamily: {
        'nunito': ["Nunito", "sans-serif"],
        'poppins': ["Poppins", "sans-serif"],
        'cursive': ['Comic Neue', 'cursive'],
      },
      zIndex: {
        '100': '100',
      },
      width: {
        '80': '80%',
      },
      screens:{
        '2xl': '1320px',
      }
    },
  },
  plugins: [],
};
