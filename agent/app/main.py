from fastapi import FastAPI, Request

app = FastAPI()

@app.post("/webhook")
async def github_webhook(request: Request):
    payload = await request.json()
    # Process the GitHub webhook payload here
    print(payload)
    return {"status": "received"}