# from fastapi import FastAPI, UploadFile, File
# import torch
# from PIL import Image, ImageDraw
# import io
# import base64
# import sys

# # Add the path to the models directory
# sys.path.append('/Users/naren/Capstone_Vision_Transformer/yolov5_model/model')

# app = FastAPI()

# # Load YOLOv5 model using torch.hub
# model = torch.hub.load('ultralytics/yolov5', 'custom', path='/Users/naren/Capstone_Vision_Transformer/yolov5_model/check_backend/best.pt')
# model.eval()  # Set the model to evaluation mode

# @app.post("/detect-weed/")
# async def detect_weed(file: UploadFile = File(...)):
#     # Read the uploaded image file
#     image_data = await file.read()
#     image = Image.open(io.BytesIO(image_data)).convert("RGB")

#     # Perform inference directly on the image
#     results = model(image)

#     # Flag to check if any weeds are detected
#     weed_detected = False

#     # Process the results
#     draw = ImageDraw.Draw(image)
#     for *box, confidence, class_id in results.xyxy[0]:  # YOLOv5 format
#         if confidence > 0.5:  # Adjust threshold as needed
#             x_min, y_min, x_max, y_max = map(int, box)
#             draw.rectangle([x_min, y_min, x_max, y_max], outline="red", width=2)
#             draw.text((x_min, y_min), "Weed", fill="red")  # Add label
#             weed_detected = True

#     # Convert the processed image to a returnable format
#     img_buffer = io.BytesIO()
#     image.save(img_buffer, format="JPEG")
#     img_str = base64.b64encode(img_buffer.getvalue()).decode("utf-8")

#     # Return the response based on detection
#     if weed_detected:
#         print("Detected")
#         return {"message": "Weed Detected", "image": img_str}
#     else:
#         print("Not Detected")
#         return {"message": "No Weed Detected"}
from fastapi import FastAPI, UploadFile, File
import torch
from PIL import Image, ImageDraw
import io
import os
import sys

# Add the path to the models directory
sys.path.append('/Users/naren/Capstone_Vision_Transformer/yolov5_model/model')

app = FastAPI()

# Load YOLOv5 model using torch.hub
model = torch.hub.load('ultralytics/yolov5', 'custom', path='/Users/naren/Capstone_Vision_Transformer/yolov5_model/check_backend/best.pt')
model.eval()  # Set the model to evaluation mode

# Directory to save the detected weed images
output_dir = "/Users/naren/Capstone_Vision_Transformer/returned_weed"
os.makedirs(output_dir, exist_ok=True)  # Ensure the directory exists

@app.post("/detect-weed/")
async def detect_weed(file: UploadFile = File(...)):
    # Read the uploaded image file
    image_data = await file.read()
    image = Image.open(io.BytesIO(image_data)).convert("RGB")

    # Perform inference directly on the image
    results = model(image)

    # Flag to check if any weeds are detected
    weed_detected = False

    # Process the results
    draw = ImageDraw.Draw(image)
    for *box, confidence, class_id in results.xyxy[0]:  # YOLOv5 format
        if confidence > 0.5:  # Adjust threshold as needed
            x_min, y_min, x_max, y_max = map(int, box)
            draw.rectangle([x_min, y_min, x_max, y_max], outline="red", width=2)
            draw.text((x_min, y_min), "Weed", fill="red")  # Add label
            weed_detected = True

    # Save the processed image if weed is detected
    if weed_detected:
        output_path = os.path.join(output_dir, "weed_detected.jpg")
        image.save(output_path, format="JPEG")
        print("Detected and saved to:", output_path)
        return {"message": "Weed Detected", "image_path": output_path}
    else:
        print("Not Detected")
        return {"message": "No Weed Detected"}
