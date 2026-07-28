from app.embeddings.titan import make_embeddings
from app.config.db import get_connection


async def process_github_event(payload):
    memories = []
    repository = payload["repository"]["full_name"]
    for commit in payload.get("commits",[]):
        message = commit["message"]
        # author = commit["author"]["name"]
        memory_text = f"Repository: {repository} Commit:{message}"
        print(memory_text)
        vector = make_embeddings(memory_text)
        insert_into_db(
            repo_name=repository,
            commit_msg = message,
            embedding = vector
        )
        memories.append({
            "content": memory_text,
            "vector": vector
        })
    return memories


def insert_into_db(repo_name, commit_msg, embedding):
    conn = get_connection()
    with conn.cursor() as cursor:
        cursor.execute(
            "INSERT INTO memories (repo_name, commit_msg, embedding) VALUES (%s, %s, %s)",
            (repo_name, commit_msg, embedding)
        )
    conn.commit()
    conn.close()