<script>

import Navbar from "$lib/Navbar.svelte";

let email = '';
  let password = '';

  const handleSubmit = async () => {
    const response = await fetch('http://localhost:8000/login/', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({ email, password })
    });

    if (response.ok) {
      // Redirect to the choice page on successful login
      window.location.href = '/choice'; // New route for choosing between plant or weed detection
    } else {
      const result = await response.json();
      console.error(result);
      alert('Invalid email or password');
    }
  };
</script>

<style>
  body {
    background: url('./plant_bgi.jpg') no-repeat center center fixed;
    background-size: cover;
    font-family: Arial, sans-serif;
    margin: 0;
    padding: 0;
  }

  .form-container {
    max-width: 400px;
    margin: 2rem auto;
    padding: 2rem;
    background: rgba(50, 50, 50, 0.8); /* Slightly lighter black-gray with transparency */
    border-radius: 8px;
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
    position: relative;
    overflow: hidden;
  }

  .form-group {
    margin-bottom: 1rem;
    color: #fff;
  }

  .form-group a {
    color: #007bff;
    text-decoration: none;
  }

  .form-input {
    width: 95%;
    padding: 0.5rem;
    border: 1px solid #ced4da;
    border-radius: 4px;
    background-color: #fff;
  }

  .form-button {
    width: 100%;
    padding: 0.5rem;
    background-color: #007bff;
    color: white;
    border: none;
    border-radius: 4px;
    cursor: pointer;
  }

  .form-button:hover {
    background-color: #0056b3;
  }
</style>

<body>
  <Navbar activePage="login" />
  <div class="form-container">
    <h2 style="color: chartreuse; text-align: center;">Login</h2>
      <form on:submit|preventDefault={handleSubmit}>
        <div class="form-group">
          <input type="email" bind:value={email} placeholder="Email" class="form-input" required>
        </div>
    
        <div class="form-group">
          <input type="password" bind:value={password} placeholder="Password" class="form-input" required>
        </div>
    
        <div class="form-group">
          <button type="submit" class="form-button">Login</button>
        </div>
    
      </form>
      <div class="form-group">
        Haven't made an account yet?
        <a href="/register">Register</a>
      </div>
    </div>
</body>
