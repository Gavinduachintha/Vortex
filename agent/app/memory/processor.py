from app.embeddings.titan import make_embeddings


async def process_github_event(payload):
    repository = payload["repository"]["full_name"]
    for commit in payload.get("commits",[]):
        message = commit["message"]
        # author = commit["author"]["name"]
        memory_text = f"Repository: {repository} Commit:{message}"
        print(memory_text)
        vector = make_embeddings(memory_text)
        print (vector)
        return vector