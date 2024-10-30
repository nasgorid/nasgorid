document.addEventListener("DOMContentLoaded", () => {
    const userId = "63c9bca1c2579a4385f3e1b2"; // Ganti dengan ID user yang valid dari MongoDB
    const baseUrl = "http://localhost:8081"; // Pastikan URL ini sesuai dengan URL backend Anda

    async function fetchUserData() {
        try {
            const response = await fetch(`${baseUrl}/users/${userId}`);
            
            if (!response.ok) {
                throw new Error("Failed to fetch user data");
            }

            const user = await response.json();

            // Update konten profil di profile.html dengan data yang diambil dari backend
            document.querySelector(".content__title h1").textContent = user.name || "Nama tidak tersedia";
            document.querySelector(".content__title span").textContent = user.location || "Lokasi tidak tersedia";
            document.querySelector(".content__description p:nth-child(1)").textContent = `Email : ${user.email || "Email tidak tersedia"}`;
            document.querySelector(".content__description p:nth-child(2)").textContent = `Password : ${"*".repeat(user.password ? user.password.length : 4)}`;
            document.querySelector(".content__description p:nth-child(3)").textContent = `UMKM Name : ${user.umkmName || "UMKM tidak tersedia"}`;
        } catch (error) {
            console.error("Error fetching user data:", error);
        }
    }

    // Memanggil fungsi untuk mengambil data user
    fetchUserData();
});

// document.addEventListener('DOMContentLoaded', function () {
//     const loginButton = document.getElementById('login-btn');
//     const emailInput = document.getElementById('email');
//     const passwordInput = document.getElementById('password');

//     loginButton.addEventListener('click', function (event) {
//         event.preventDefault();
//         const email = emailInput.value;
//         const password = passwordInput.value;

//         if (!email || !password) {
//             alert('Please complete all fields');
//             return;
//         }

//         const data = {
//             email: email,
//             password: password,
//         };

//         fetch('http://localhost:8081/login', {
//             method: 'POST',
//             headers: {
//                 'Content-Type': 'application/json',
//             },
//             body: JSON.stringify(data),
//         })
//         .then(response => {
//             if (!response.ok) {
//                 throw new Error('Login failed');
//             }
//             return response.json();
//         })
//         .then(result => {
//             const token = result.token;
//             localStorage.setItem('authToken', token); // Simpan token ke localStorage
//             alert('Login successful!');
//             window.location.href = 'profile.html';
//         })
//         .catch(error => {
//             console.error('Error:', error);
//             alert('Login failed. Check your email and password.');
//         });
//     });
// });
