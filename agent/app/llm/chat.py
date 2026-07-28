from app.core.bedrock import get_chat_model

llm = get_chat_model()


def chat_with_model(prompt: str) -> str:
    response = llm.invoke(prompt)
    return response.content