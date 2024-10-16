document.getElementById("registerForm").addEventListener("submit", function(event) {
    event.preventDefault(); // Mencegah form dari pengiriman default

    const formData = new FormData(this); // Mengambil data dari form
    const data = Object.fromEntries(formData); // Mengonversi FormData ke objek biasa

    // Mengirim request ke endpoint register
    fetch("http://localhost:8080/register", { // Ganti dengan URL backend Anda
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
        // Anda bisa menampilkan pesan atau melakukan navigasi
    })
    .catch((error) => {
        console.error('Error:', error); // Menangani error
    });
});
