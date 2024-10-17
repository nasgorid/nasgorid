document.getElementById("loginForm").addEventListener("submit", function(event) {
    event.preventDefault(); // Mencegah form dari pengiriman default

    const formData = new FormData(this); // Mengambil data dari form
    const data = Object.fromEntries(formData); // Mengonversi FormData ke objek biasa

    // Mengirim request ke endpoint login
    fetch("http://localhost:8081/login", { // Ganti dengan URL backend Anda
        method: "POST",
        headers: {
            "Content-Type": "application/json", // Tentukan tipe konten
        },
        body: JSON.stringify(data), // Mengonversi data objek ke JSON
    })
    .then(response => {
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }
        return response.json(); // Mengembalikan response sebagai JSON
    })
    .then(data => {
        console.log('Success:', data); // Proses sukses
        alert('Login berhasil!'); // Menampilkan pemberitahuan
        // Anda bisa melakukan navigasi setelah login berhasil
    })
    .catch((error) => {
        console.error('Error:', error); // Menangani error
        alert('Login gagal. Silakan coba lagi.'); // Menampilkan pesan error
    });
});
