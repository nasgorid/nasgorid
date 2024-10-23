document.getElementById('loginForm').addEventListener('submit', async function(event) {
    event.preventDefault(); // Mencegah form dari reload halaman

    // Mengambil nilai dari input form
    const email = document.getElementById('email').value;
    const password = document.getElementById('password').value;

    try {
        const response = await fetch('http://localhost:8081/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                email: email,
                password: password
            })
        });

        // Menangani response dari server
        const result = await response.json();
        
        if (response.ok) {
            alert("Login berhasil!");
            // Redirect ke halaman yang diinginkan setelah login
            window.location.href = "/dashboard.html";
        } else {
            alert("Login gagal: " + result.message);
        }

    } catch (error) {
        console.error("Error:", error);
        alert("Terjadi kesalahan saat login. Silakan coba lagi.");
    }
});