from app.core.bedrock import get_embedding_model

embedding_model = get_embedding_model()

def make_embeddings(text: str):
    return embedding_model.embed_query(text)