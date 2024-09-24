<script>
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
    background: url('og-auth-frontend\static\capstone_login.jpg') no-repeat center center fixed;
    background-size: cover;
    font-family: Arial, sans-serif;
    margin: 0;
    padding: 0;
  }

  .form-container {
    max-width: 400px;
    margin: 2rem auto;
    padding: 2rem;
    background: #fff;
    border-radius: 8px;
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  }

  .form-group {
    margin-bottom: 1rem;
  }

  .form-input {
    width: 100%;
    padding: 0.5rem;
    border: 1px solid #ced4da;
    border-radius: 4px;
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

<div class="form-container">
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

    <div class="form-group">
      Haven't made an account yet?
    </div>

    <div class="form-group">
      <a href="http://localhost:5173/register">Register</a>
    </div>
  </form>
</div>
