document.addEventListener('DOMContentLoaded', function() {
    document.getElementById('loginForm').addEventListener('submit', function(event) {
      event.preventDefault();
  
      const formData = {
        email: document.getElementById('email').value,
        password: document.getElementById('password').value
      };
  
      fetch('http://localhost:5500/login', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(formData)
      })
      .then(response => {
        if (!response.ok) {
          throw new Error('Login failed');
        }
        return response.json();
      })
      .then(data => {
        // Assuming the backend responds with a status 200 and sets the token as a cookie
        const messageElement = document.getElementById('message');
        if (messageElement) {
          messageElement.textContent = 'Login successful. Redirecting...';
          window.location.href = '../app/dashboard.html'; // Redirect upon successful login
        }
      })
      .catch(error => {
        console.error('Error:', error);
        const messageElement = document.getElementById('message');
        if (messageElement) {
          messageElement.textContent = 'Login failed. Please try again.';
        }
      });
    });
  });
  