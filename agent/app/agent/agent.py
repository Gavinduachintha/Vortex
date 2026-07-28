from app.llm.chat import chat_with_model
from app.embeddings.titan import make_embeddings


def ask_question(question: str):

    answer = chat_with_model(question)

    return answer


def create_embedding(text: str):

    return make_embeddings(text)