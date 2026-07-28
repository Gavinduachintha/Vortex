from fastapi import APIRouter, Request, Header
from app.memory.processor import process_github_event
router = APIRouter()

@router.post("/github/webhook")
async def github_webhook(
    request: Request,
    x_github_event: str=Header(None)):
    payload = await request.json()
    print("github event", x_github_event)

    if x_github_event == "push":
        await process_github_event(payload)

    
    