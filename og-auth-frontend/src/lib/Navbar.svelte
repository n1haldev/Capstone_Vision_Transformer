<script>
    export let activePage = '';
  </script>
  
  <style>
    nav {
      background-color: #333;
      color: white;
      padding: 1rem;
      display: flex;
      justify-content: space-between;
      align-items: center;
    }

    .logo img {
      width: 5%;
      height: 5%;
      border-radius: 100%;
    }
  
    .nav-links {
      display: flex;
      gap: 1rem;
    }
  
    .nav-link {
      color: white;
      border-radius: 100%;
      text-decoration: none;
      padding: 0.5rem 1rem;
      border-radius: 4px;
    }
  
    .nav-link:hover {
      background-color: #555;
    }
  
    .nav-link.active {
      background-color: #007BFF;
    }
  </style>
  
  <nav>
    <div class="logo">
      <a href="/">
        <img src="./logo.png" alt="">
      </a>
    </div>
    <div class="nav-links">
      <a href="/" class="nav-link {activePage === 'home' ? 'active' : ''}" >Home</a>
      <a href="/contact" class="nav-link {activePage === 'contact' ? 'active' : ''}">Contact</a>
      <a href="/choice" class="nav-link {activePage === 'choice' ? 'active' : ''}">Choice</a>
      <a href="/login" class="nav-link {activePage === 'login' ? 'active' : ''}">Login</a>
      <a href="/register" class="nav-link {activePage === 'register' ? 'active' : ''}">Register</a>
    </div>
  </nav>