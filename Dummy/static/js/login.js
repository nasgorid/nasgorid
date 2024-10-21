const loginButton = document.getElementById('login-btn');
loginButton.addEventListener('click', async () => {
    const email = document.getElementById('email').value;
    const password = document.getElementById('password').value;

    try {
        const response = await fetch('http://localhost:8081/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ email: email, password: password })
        });

        if (response.ok) {
            const data = await response.json();
            alert(data.message);
            window.location.href = 'dashboard.html'; // Redirect on success
        } else {
            const errorData = await response.json();
            alert(errorData.message || 'Login failed');
        }
    } catch (error) {
        console.error('Error:', error);
        alert('Failed to fetch');
    }
});
    