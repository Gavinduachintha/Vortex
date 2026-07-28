from fastapi import FastAPI
from pydantic import BaseModel
from app.agent.agent import ask_question
from app.embeddings.titan import make_embeddings
app = FastAPI(
    title="Vortex AI Agent",
    version="0.1.0",
)

class QuestionRequest(BaseModel):
    question: str

class AnswerResponse(BaseModel):
    answer: str

@app.get("/")
async def root():
    return {"message": "Welcome to the Vortex AI Agent API!"}

@app.post("/ask", response_model=AnswerResponse)
def ask_endpoint(request: QuestionRequest):
    answer =ask_question(request.question)
    return AnswerResponse(answer=answer)

@app.post("/make_embeddings")
def make_embeddings_endpoint(request: QuestionRequest):
    embeddings = make_embeddings(request.question)
    return {"embeddings": embeddings}