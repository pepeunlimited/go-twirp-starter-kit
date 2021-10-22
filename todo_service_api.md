##### CreateTodo
```
$ curl -H "Content-Type: application/json" \
-X POST "localhost:8080/twirp/pepeunlimited.twirpkit.services.TodoService/CreateTodo" \
-d '{"title": "HelloWorld"}'
```
##### GetTodo
```
$ curl -H "Content-Type: application/json" \
-X POST "localhost:8080/twirp/pepeunlimited.twirpkit.services.TodoService/GetTodo" \
-d '{"todo_id": 1}'
```
##### DeleteTodo
```
$ curl -H "Content-Type: application/json" \
-X POST "localhost:8080/twirp/pepeunlimited.twirpkit.services.TodoService/DeleteTodo" \
-d '{"todo_id": 1}'
```
##### UpdateTodo
```
$ curl -H "Content-Type: application/json" \
-X POST "localhost:8080/twirp/pepeunlimited.twirpkit.services.TodoService/UpdateTodo" \
-d '{ "operation": { "id": 4844229853851485532, "is_done": true }, "field_mask": "isDone" }'
```