<script>
  import Navbar from "$lib/Navbar.svelte";
  import { goto } from '$app/navigation';
  let choice = "";

  const handleSelection = () => {
    if (choice === "plant") {
      goto("/plant-detection");
    } else if (choice === "weed") {
      goto("/weed-detection");
    }
  };
</script>

<style>
  body {
  
    background-size: cover;
    font-family: Arial, sans-serif;
    margin: 0;
    padding: 0;
  }

  .hero {
    text-align: center;
    padding: 2rem 2rem;
    background: rgba(0, 0, 0, 0.5);
    color: white;
  }

  .hero h1 {
    font-size: 3rem;
    margin-bottom: 0.5rem;
  }

  .hero p {
    font-size: 1.5rem;
  }

  .services {
    display: flex;
    justify-content: space-around;
    padding: 2rem;
    background: rgba(255, 255, 255, 0.8);
  }

  .service {
    width: 45%;
    padding: 1rem;
    background: white;
    border-radius: 8px;
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
    text-align: center;
  }

  .service img {
    width: 100px;
    height: 100px;
    border-radius: 50%;
    margin-bottom: 1rem;
  }

  .service h2 {
    font-size: 1.5rem;
    margin-bottom: 1rem;
  }

  .service:hover {
    background-color: lightblue;
    transform: scale(1.03);
    transition: transform 0.3s ease-in-out, box-shadow 0.3s ease-in-out;
  }

  .service p {
    font-size: 1rem;
    color: #555;
  }
</style>

<body>
  <Navbar activePage="choice" />
  <div class="hero">
    <h1>Welcome to PlantVision</h1>
    <p>Empowering Farmers and Academics with Advanced Plant Detection</p>
  </div>
  <div class="services">
    <div class="service" on:click = {() => goto('/weed-detection')}>
      <img src="./weed_detection_icon_resized.png" alt="Weed Detection">
      <h2>Weed Detection for Farmers</h2>
      <p>Utilize our advanced YOLOv5 model to detect weeds in your field and improve crop yield.</p>
    </div>
    <div class="service">
      <img src="./plant_species_identification.jpg" alt="Plant Species Identification">
      <h2>Plant Species Identification for Academics</h2>
      <p>Identify plant species accurately for your research and academic projects using our cutting-edge technology.</p>
    </div>
  </div>
</body>