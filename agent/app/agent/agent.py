from app.llm.chat import chat_with_model

def ask_question(question: str) -> str:
    answer =  chat_with_model(question)
    return answer