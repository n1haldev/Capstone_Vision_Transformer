<script>
  import { goto } from '$app/navigation'; // Correct import for navigation

let choice = '';

const handleSelection = (/** @type {string} */ selected) => { // Explicitly typing 'selected' as a string
  choice = selected;

  if (selected === 'plant') {
    // Navigate to the Plant Species Detection page
    goto('/plant-detection'); // Redirects to /plant-detection route
  } else if (selected === 'weed') {
    // Navigate to the Weed Detection and Classification page
    goto('/weed-detection'); // Redirects to /weed-detection route
  }
};

</script>
  
  <style>
    .option {
      margin: 1rem;
      padding: 1rem;
      border: 1px solid #ccc;
      cursor: pointer;
      text-align: center;
      background-color: #f5f5f5;
    }
    .option:hover {
      background-color: #ddd;
    }
  </style>
  
  <!-- svelte-ignore a11y-click-events-have-key-events -->
  <!-- svelte-ignore a11y-no-static-element-interactions -->
  <div class="option" on:click={() => handleSelection('plant')}>Plant Species Detection</div>
  <!-- svelte-ignore a11y-click-events-have-key-events -->
  <div class="option" on:click={() => handleSelection('weed')}>Weed Detection and Classification</div>
  