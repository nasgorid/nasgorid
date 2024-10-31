document.addEventListener("DOMContentLoaded", () => {
    const baseUrl = "http://localhost:8081"; // Sesuaikan dengan URL backend Anda
    const userId = "id"; // Ganti dengan ID user yang ingin ditampilkan

    async function fetchUserData() {
        try {
            const response = await fetch(`${baseUrl}/users/${userId}`, {
                method: "GET",
                headers: {
                    "Content-Type": "application/json"
                }
            });

            if (!response.ok) {
                throw new Error("Failed to fetch user data");
            }

            const user = await response.json();

            // Update konten profil di halaman profile dengan data dari backend
            document.querySelector(".content__title h1").textContent = user.name || "Nama tidak tersedia";
            document.querySelector(".content__title span").textContent = user.location || "Lokasi tidak tersedia";
            document.querySelector(".content__description p:nth-child(1)").textContent = `Email : ${user.email || "Email tidak tersedia"}`;
            document.querySelector(".content__description p:nth-child(2)").textContent = `Password : ${"*".repeat(user.password ? user.password.length : 4)}`;
            document.querySelector(".content__description p:nth-child(3)").textContent = `UMKM Name : ${user.umkm_name || "UMKM tidak tersedia"}`;
        } catch (error) {
            console.error("Error fetching user data:", error);
            alert("Terjadi kesalahan saat mengambil data pengguna.");
        }
    }

    fetchUserData(); // Panggil fungsi untuk mengambil data pengguna
});
