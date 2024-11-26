<script>
    import { onMount } from 'svelte';

    let imageUrl = ""; // URL of the selected image
    let boxes = []; // Holds bounding box data from the backend
    let selectedFile = null; // Tracks the uploaded file
    let originalImageWidth = 1280; // Replace with the actual width of the backend image
    let originalImageHeight = 720; // Replace with the actual height of the backend image

    // Function to scale bounding boxes based on displayed image dimensions
    function scaleBoundingBoxes(boxes, originalWidth, originalHeight, renderedWidth, renderedHeight) {
        const xScale = renderedWidth / originalWidth;
        const yScale = renderedHeight / originalHeight;

        return boxes.map(({ x_min, y_min, x_max, y_max, confidence, class_name }) => ({
            x_min: x_min * xScale,
            y_min: y_min * yScale,
            x_max: x_max * xScale,
            y_max: y_max * yScale,
            confidence,
            class_name,
        }));
    }

    // Function to call the backend and process the image
    async function detectWeeds() {
        if (!selectedFile) {
            alert("Please upload an image.");
            return;
        }

        const formData = new FormData();
        formData.append("file", selectedFile);

        try {
            const response = await fetch("http://127.0.0.1:8000/detect-weed/", {
                method: "POST",
                body: formData,
            });

            if (response.ok) {
                const data = await response.json();

                console.log("Raw Bounding Boxes from Backend:", data.boxes);

                const renderedImage = document.querySelector(".image");
                const renderedWidth = renderedImage.clientWidth;
                const renderedHeight = renderedImage.clientHeight;

                // Scale bounding boxes to the displayed image dimensions
                boxes = scaleBoundingBoxes(
                    (data.boxes || []).filter(
                        ({ x_min, y_min, x_max, y_max }) =>
                            x_min >= 0 &&
                            y_min >= 0 &&
                            x_max <= originalImageWidth &&
                            y_max <= originalImageHeight
                    ),
                    originalImageWidth,
                    originalImageHeight,
                    renderedWidth,
                    renderedHeight
                );

                console.log("Filtered and Scaled Bounding Boxes:", boxes);
            } else {
                alert("Failed to detect weeds. Please try again.");
            }
        } catch (error) {
            console.error("Error:", error);
            alert("An error occurred. Please try again.");
        }
    }
</script>


<style>
    .container {
        position: relative;
        display: inline-block;
        margin-top: 20px;
    }

    .image {
        max-width: 100%;
        height: auto;
        border-radius: 10px;
        box-shadow: 0px 4px 10px rgba(0, 0, 0, 0.1);
    }

    .bounding-box {
        position: absolute;
        border: 2px solid red;
        pointer-events: none;
    }

    .upload-section {
        display: flex;
        flex-direction: column;
        align-items: center;
        margin: 20px 0;
    }

    .file-input {
        padding: 10px;
        font-size: 16px;
        margin-bottom: 10px;
        cursor: pointer;
    }

    .button {
        padding: 10px 20px;
        background-color: #007bff;
        color: white;
        border: none;
        border-radius: 5px;
        font-size: 16px;
        cursor: pointer;
        transition: background-color 0.3s;
    }

    .button:hover {
        background-color: #0056b3;
    }

    .result-container {
        text-align: center;
        margin-top: 20px;
    }

    .result-details {
        margin-top: 10px;
        font-size: 16px;
        color: #333;
    }
</style>



<div class="upload-section">
    <input
        class="file-input"
        type="file"
        accept="image/*"
        on:change="{(e) => {
            selectedFile = e.target.files[0];
            imageUrl = URL.createObjectURL(selectedFile);
        }}"
    />
    <button class="button" on:click="{detectWeeds}">Detect Weeds</button>
</div>

<div class="container" style="width: 100%; max-width: 500px;">
    {#if imageUrl}
        <img
            src="{imageUrl}"
            alt="Uploaded Image"
            class="image"
        />

        {#each boxes as { x_min, y_min, x_max, y_max }}
            <div
                class="bounding-box"
                style="top: {y_min}px; left: {x_min}px; width: {x_max - x_min}px; height: {y_max - y_min}px;"
            ></div>
        {/each}
    {/if}
</div>

<div class="result-container">
    {#if boxes.length > 0}
        <div class="result-details">Detected Weeds: {boxes.length}</div>
    {:else}
        <div class="result-details">No weeds detected.</div>
    {/if}
</div>