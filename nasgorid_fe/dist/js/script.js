// document.getElementById('starters').addEventListener('click', function () {
//     document.getElementById('starterSection').classList.remove('hidden');
//     document.getElementById('lunchSection').classList.add('hidden');
// });

// // Fungsi untuk menangani klik pada tombol Lunch
// document.getElementById('lunch').addEventListener('click', function () {
//     document.getElementById('lunchSection').classList.remove('hidden');
//     document.getElementById('starterSection').classList.add('hidden');
// });

// window.addEventListener('scroll', () => {
//     const header = document.querySelector('header');
//     const fixedNav = header.offsetTop;

//     if (window.pageYOffset > fixedNav) {
//         header.classList.add('navbar-fixed');
//     } else {
//         header.classList.remove('navbar-fixed');
//     }
// });

// const hamburger = document.querySelector('#hamburger');
// const navMenu = document.querySelector('#nav-menu');

// hamburger.addEventListener('click', () => {
//     hamburger.classList.toggle('hamburger-active');
//     navMenu.classList.toggle('hidden');
// });

// Kode JavaScript ini bisa dimasukkan ke dalam file dist/js/script.js
const cartButton = document.getElementById("cart-button");
const cartSidebar = document.getElementById("cart-sidebar");
const closeCartButton = document.getElementById("close-cart");

cartButton.addEventListener("click", () => {
    cartSidebar.classList.toggle("translate-x-full");
});

closeCartButton.addEventListener("click", () => {
    cartSidebar.classList.add("translate-x-full");
});
