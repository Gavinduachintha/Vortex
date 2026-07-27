from app.core.bedrock import get_model

llm = get_model()
def chat_with_model(prompt: str) -> str:
    try:
        response = llm.invoke(prompt)
        return response.text
    except Exception as e:
        print(f"Error occurred while invoking the model: {e}")
        return "An error occurred while processing your request."