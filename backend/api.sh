url="http://127.0.0.1:4400"
todoId=6ac07b617d5e8f7b9e89b899

# health
# echo -e "\ncheck health\n"
# curl -sS -X GET "${url}/healthz" | jq

# get all todos
echo -e "fetching all todos..\n"
curl -sS -X GET "${url}/api/todos" \
    | jq

# create todo
# timestamp=$(date +%s)
# echo -e "\nCreating todo...\n"
# curl -sS -X POST "${url}/api/todos" \
#   -H "Content-Type: application/json" \
#   -d "$(jq -n \
#     --arg title "todo ${timestamp}" \
#     --arg description "description for todo ${timestamp}" \
#     '{title: $title, description: $description}')" \
#   | jq

# mark done
# echo -e "completing todo.."
# curl -sS -X PATCH "${url}/api/todos/${todoId}/done" \
#     | jq

# delete todo
# echo -e "Deleting todo..\n"
# curl -sS -X PATCH "${url}/api/todos/${todoId}/remove" \
#     | jq
