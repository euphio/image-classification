
import grpc
from concurrent import futures
import image_recognition_pb2
import image_recognition_pb2_grpc
from PIL import Image
import io
import torch
import torchvision.transforms as transforms
from torchvision import models
from imagenet_classes import imagenet_classes

model = None


def load_model():
    global model
    model = models.mobilenet_v2(pretrained=True)
    model.eval()


def classify_image(image_bytes):
    transform = transforms.Compose([
        transforms.Resize((224, 224)),
        transforms.ToTensor(),
        transforms.Normalize([0.485, 0.456, 0.406], [0.229, 0.224, 0.225])
    ])
    image = Image.open(io.BytesIO(image_bytes)).convert("RGB")
    image = transform(image).unsqueeze(0)

    with torch.no_grad():
        outputs = model(image)
        _, predicted = outputs.max(1)
        predicted_class = predicted.item()

    class_label = imagenet_classes[predicted_class]
    return [class_label]


class ImageLabelServicer(image_recognition_pb2_grpc.ImageLabelServiceServicer):
    def GetLabels(self, request, context):
        image_bytes = request.image_data
        labels = classify_image(image_bytes)
        return image_recognition_pb2.ImageResponse(labels=labels)


def serve():
    load_model()
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    image_recognition_pb2_grpc.add_ImageLabelServiceServicer_to_server(
        ImageLabelServicer(), server)
    server.add_insecure_port("[::]:50051")
    print("Python gRPC service running on port 50051...")
    server.start()
    server.wait_for_termination()


if __name__ == "__main__":
    serve()
