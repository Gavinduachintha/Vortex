 curl -X POST http://localhost:8000/github/webhook -H "Content-Type: application/json" -H "X-GitHub-Event: push" -d '{
  "repository": {
    "full_name": "gavi/vortex"
  },
  "commits": [
    {
      "message": "added websocket"
    }
  ]
}'
{"message":"Web hook received"}
