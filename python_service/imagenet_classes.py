import json
import os

with open(os.path.join(os.path.dirname(__file__), 'imagenet_classes.json'), 'r') as f:
    imagenet_classes = json.load(f)

# Optional: Convert keys to integers if the JSON keys are strings
imagenet_classes = {int(k): v for k, v in imagenet_classes.items()}