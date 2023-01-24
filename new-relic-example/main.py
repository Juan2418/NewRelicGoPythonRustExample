from fastapi import FastAPI, Header
from pydantic import BaseSettings
from functools import lru_cache
import uuid
import httpx
import newrelic.agent


class Settings(BaseSettings):
    GO_SERVICE_BASE_URL: str = "http://localhost"

    class Config:
        env_file = ".env"


@lru_cache()
def get_settings():
    return Settings()


settings = get_settings()
newrelic.agent.initialize("newrelic.ini")
app = FastAPI()


@app.post("/")
async def requestMessage(x_request_id: str | None = Header(default=None)):
    request_id = x_request_id or str(uuid.uuid4())
    return {"message": "Hello World", "requestId": request_id}


@app.get("/trace")
async def trace(x_request_id: str = Header(default=None)):
    request_id = x_request_id or str(uuid.uuid4())
    GO_SERVICE_BASE_URL = settings.GO_SERVICE_BASE_URL
    return_message = ""

    with httpx.Client() as client:
        headers_map = {"x-request-id": request_id}
        go_service_response = client.post(
            f"{GO_SERVICE_BASE_URL}/reference", headers=headers_map
        )
        response_body = go_service_response.json()
        return_message = response_body["message"]

    return {"message": return_message, "requestId": request_id}
