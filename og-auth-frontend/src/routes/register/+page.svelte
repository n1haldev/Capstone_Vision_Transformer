<script>
  let email = '';
  let password = '';
  let confirmPassword = '';
  let name = '';
  let isRedirecting = false;

  const handleSubmit = async () => {
    if (password !== confirmPassword) {
      alert('Passwords do not match');
      return;
    }

    const response = await fetch('http://localhost:8000/register/', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({ email, name, password })
    });

    const result = await response.json();
    console.log(result);

    // Show redirect message and initiate animation
    isRedirecting = true;

    // Wait for animation to complete
    setTimeout(() => {
      // Redirect to the login page
      window.location.href = 'http://localhost:5173/login';
    }, 4000); // Adjust this to match the duration of your animation
  };
</script>

<style>
  body {
    background: url('/capstone_login.jpg') no-repeat center center fixed;
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
    position: relative;
    overflow: hidden;
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

  .redirect-message {
    font-size: 1.2rem;
    color: #007bff;
    margin-top: 1rem;
    opacity: 0;
    animation: fadeIn 1s forwards, swoosh 1s 1s forwards;
  }

  @keyframes fadeIn {
    to {
      opacity: 1;
    }
  }

  @keyframes swoosh {
    to {
      transform: translateX(100%);
    }
  }
</style>

<div class="form-container">
  <form on:submit|preventDefault={handleSubmit}>
    <div class="form-group">
      <input type="email" bind:value={email} placeholder="Email" class="form-input" required />
    </div>

    <div class="form-group">
      <input type="text" bind:value={name} placeholder="Name" class="form-input" required />
    </div>

    <div class="form-group">
      <input type="password" bind:value={password} placeholder="Password" class="form-input" required />
    </div>

    <div class="form-group">
      <input type="password" bind:value={confirmPassword} placeholder="Confirm Password" class="form-input" required />
    </div>

    <div class="form-group">
      <button type="submit" class="form-button">Register</button>
    </div>
  </form>
  {#if isRedirecting}
    <div class="redirect-message">I am being redirected to the login page...</div>
  {/if}
</div>
