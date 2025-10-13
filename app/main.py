from fastapi import FastAPI

app = FastAPI()

@app.get('/')
def root():
    return "chirp-core-service"