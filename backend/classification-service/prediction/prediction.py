import pandas as pd
import numpy as np
import librosa
import io
from tensorflow.keras.models import load_model
from sklearn.preprocessing import LabelEncoder

le = LabelEncoder()
df = pd.read_csv('backend/classification-service/prediction/models/nyanya.csv', sep="\t")
y = df['label'].values
y_encoded = le.fit_transform(y)

loaded_model = load_model('backend/classification-service/prediction/models/audio_classi.keras')

def predict_genre(audio_bytes):
    features = extract_features(audio_bytes)
    features = features.reshape(1, -1) 
    prediction = loaded_model.predict(features)
    predicted_label = le.inverse_transform([np.argmax(prediction)])
    return predicted_label[0]

def extract_features(audio_bytes):
    y, sr = librosa.load(io.BytesIO(audio_bytes), duration=30, sr=None)  
    mfccs = np.mean(librosa.feature.mfcc(y=y, sr=sr, n_mfcc=13).T, axis=0)
    chroma = np.mean(librosa.feature.chroma_stft(y=y, sr=sr).T, axis=0)
    mel = np.mean(librosa.feature.melspectrogram(y=y, sr=sr).T, axis=0)
    contrast = np.mean(librosa.feature.spectral_contrast(y=y, sr=sr).T, axis=0)
    tonnetz = np.mean(librosa.feature.tonnetz(y=y, sr=sr).T, axis=0)
    
    return np.hstack([mfccs, chroma, mel, contrast, tonnetz])
