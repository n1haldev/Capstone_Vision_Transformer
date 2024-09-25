<script>
    import Navbar from "$lib/Navbar.svelte";
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

    .contacts {
        width: 10;
        height: 10;
        margin: 1rem;
    }

    .contact-list {
        background-color: white;
        margin: 1rem auto;
        display: flex;
        justify-content: center;
        border-radius: 10%;
        padding: 1rem;
        width: 20%;
    }

    @media (max-width: 800px) {
        .contact-list {
            width: 50%;
        }
    }

    @media (max-width: 600px) {
        .contact-list {
            width: 80%;
        }
    }

</style>

<body>
    <Navbar activePage="contact" />
    <div class="hero">
        <h1>Contact Us</h1>
        <p>Here are some ways you can reach out to us</p>
        <br>
        <div class="contact-list">
            <a href="https://www.linkedin.com/in/nihal-tangaturu-murthy-601919164/" target="_blank" rel="noopener noreferrer" >
                <img src="https://img.icons8.com/fluent/48/000000/linkedin.png" alt="LinkedIn" class="contacts">
            </a>
            <a href="mailto:nihaltm2002@gmail.com" target="_blank" rel="noopener noreferrer">
                <img src="https://img.icons8.com/fluent/48/000000/gmail.png" alt="Gmail" class="contacts" />
            </a>
            <a href="https://github.com/n1haldev" target="_blank" rel="noopener noreferrer">
                <img src="https://img.icons8.com/fluent/48/000000/github.png" alt="GitHub" class="contacts" />
            </a>
        </div>
    </div>
</body>