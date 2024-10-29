document.addEventListener("DOMContentLoaded", () => {
    const userId = "63c9bca1c2579a4385f3e1b2"; // Ganti dengan ID user yang valid dari MongoDB
    const baseUrl = "http://localhost:8080"; // Pastikan URL ini sesuai dengan URL backend Anda

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
