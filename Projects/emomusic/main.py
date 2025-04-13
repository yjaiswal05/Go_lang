import uvicorn
from fastapi import FastAPI
from pydantic import BaseModel
import httpx

app = FastAPI()

class EmotionRequest(BaseModel):
    user_id: str

class MusicResponse(BaseModel):
    emotion: str
    playlist_uri: str

@app.post("/play_music", response_model=MusicResponse)
async def play_music(request: EmotionRequest):
    async with httpx.AsyncClient() as client:
        # Get emotion from Emotion Detection service
        emotion_response = await client.post("http://emotion-service:8001/detect", json={"user_id": request.user_id})
        emotion = emotion_response.json()["emotion"]

        # Get playlist from Spotify service
        spotify_response = await client.post("http://spotify-service:8002/get_playlist", json={"emotion": emotion})
        playlist_uri = spotify_response.json()["playlist_uri"]

        # Play music using Spotify service
        await client.post("http://spotify-service:8002/play", json={"playlist_uri": playlist_uri})

    return MusicResponse(emotion=emotion, playlist_uri=playlist_uri)

if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=8000)
