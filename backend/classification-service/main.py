import grpc
from concurrent import futures
from protos import classification_pb2
from protos import classification_pb2_grpc
from prediction import prediction
import os
from dotenv import load_dotenv

MAX_MESSAGE_LENGTH = 1024*1024*20

def get_genre_enum(genre_str):
     genre_mapping = {
        "unknown": classification_pb2.Genres.UNKNOWN,
        "blues": classification_pb2.Genres.blues,
        "classical": classification_pb2.Genres.classical,
        "country": classification_pb2.Genres.country,
        "disco": classification_pb2.Genres.disco,
        "hiphop": classification_pb2.Genres.hiphop,
        "jazz": classification_pb2.Genres.jazz,
        "metal": classification_pb2.Genres.metal,
        "pop": classification_pb2.Genres.pop,
        "reggae": classification_pb2.Genres.reggae,
        "rock": classification_pb2.Genres.rock,
     }
     return genre_mapping.get(genre_str.lower(), classification_pb2.Genres.UNKNOWN)


class ClassificationService(classification_pb2_grpc.ClassificationServiceServicer):
     def UploadAudio(self, request, context):
          print("яздесь")
          filename = request.filename
          file_data = request.file_data

          label = prediction.predict_genre(file_data)

          return classification_pb2.Genre(genre=get_genre_enum(label))

def serve():
     server = grpc.server(futures.ThreadPoolExecutor(max_workers=10), options = [
        ('grpc.max_send_message_length', MAX_MESSAGE_LENGTH),
        ('grpc.max_receive_message_length', MAX_MESSAGE_LENGTH)
    ])
     classification_pb2_grpc.add_ClassificationServiceServicer_to_server(ClassificationService(), server)
     server.add_insecure_port(f"[::]:{os.getenv('CLASSIFICATION_SERVER')}")
     server.start()
     print(f"Server is running on port {os.getenv('CLASSIFICATION_SERVER')}...")
     server.wait_for_termination()

if __name__ == '__main__':
     load_dotenv()
     serve()
   