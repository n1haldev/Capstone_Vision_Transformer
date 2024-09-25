<script>
  import Navbar from "$lib/Navbar.svelte";
  import { goto } from '$app/navigation';
  
</script>

<style>
  body {
    background: url('./plant_bgi.jpg') no-repeat center center fixed;
    background-size: cover;
    font-family: Arial, sans-serif;
    margin: 0;
    padding: 0;
  }

  .hero {
    text-align: center;
    padding: 2rem 2rem;
    background: rgba(0, 0, 0, 0.5);
    height: 80vh;
    color: white;
  }

  .hero h1 {
    font-size: 3rem;
    margin-bottom: 0.5rem;
  }

  .hero p {
    font-size: 1.5rem;
  }

  .hero button {
    padding: 0.5rem;
    margin: 1rem;
    background-color: aqua;
    color: black;
    border-radius: 10%;
    border-color: #555;
  }

  .hero button:hover {
    background-color: lightskyblue;
    transform: scale(1.1);
    transition: transform 0.5s ease-in-out, box-shadow 0.5s ease-in-out;
  }
</style>

<body>
  <Navbar activePage="home" />
  <div class="hero">
    <h1>Welcome to PlantVision</h1>
    <p>Empowering Farmers and Academics with Advanced Plant Detection</p>
    <br>
    <button on:click={() => goto('/register')}>Register</button>
    <button on:click={() => goto('/login')}>Login</button>
  </div>
  <!-- <div class="services">
    <div class="service">
      <img src="./weed_detection_icon_resized.png" alt="Weed Detection">
      <h2>Weed Detection for Farmers</h2>
      <p>Utilize our advanced YOLOv5 model to detect weeds in your field and improve crop yield.</p>
    </div>
    <div class="service">
      <img src="./plant_species_identification.jpg" alt="Plant Species Identification">
      <h2>Plant Species Identification for Academics</h2>
      <p>Identify plant species accurately for your research and academic projects using our cutting-edge technology.</p>
    </div>
  </div> -->
</body>