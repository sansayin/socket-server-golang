from fastapi import FastAPI
import uvicorn
from fastapi.responses import UJSONResponse, ORJSONResponse
from fastapi.staticfiles import StaticFiles

app = FastAPI()

app.mount("/", StaticFiles(directory="./"), name="/")

if __name__ == "__main__":
    uvicorn.run("fast-api:app", host="127.0.0.1", port=9988, log_config=None, workers=200)
