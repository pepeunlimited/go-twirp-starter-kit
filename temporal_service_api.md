##### CreateWithdraw
```
$ curl --max-time 2.0 -H "Content-Type: application/json" \
-X POST "localhost:8080/twirp/pepeunlimited.twirpkit.services.TemporalService/CreateWithdraw" \
-d '{ "amount": 100, "user_id": 100 }'
```
##### GetWithdrawal (oneOf query)
```
$ curl -H "Content-Type: application/json" \
-X POST "localhost:8080/twirp/pepeunlimited.twirpkit.services.TemporalService/GetWithdrawal" \
-d '{"query": { "user_id": 100, "workflow_id": "8965f79e-2a2e-4611-bfc9-375d1082afd4", "workflow_run_id": "ba8ecadf-005b-40ee-a12a-81ddb8b064ab" }, "field_mask": "" }'
```
##### UpdateWithdrawal
```
$ curl -H "Content-Type: application/json" \
-X POST "localhost:8080/twirp/pepeunlimited.twirpkit.services.TemporalService/UpdateWithdrawal" \
-d '{ "operation": { "user_id": 100, "id": "8965f79e-2a2e-4611-bfc9-375d1082afd4", "workflow": { "status": 4 } }, "field_mask": "workflow"}'
```
###### UpdateWithdrawal.CancelWorkflow
```
$ curl -H "Content-Type: application/json" \
-X POST "localhost:8080/twirp/pepeunlimited.twirpkit.services.TemporalService/UpdateWithdrawal" \
-d '{ "operation": { "user_id": 100, "id": "8965f79e-2a2e-4611-bfc9-375d1082afd4", "workflow": { "status": 4 } }, "field_mask": "workflow"}'
```
###### UpdateWithdrawal.RerunWorkflow
```
$ curl -H "Content-Type: application/json" \
-X POST "localhost:8080/twirp/pepeunlimited.twirpkit.services.TemporalService/UpdateWithdrawal" \
-d '{ "operation": { "user_id": 100, "id": "8965f79e-2a2e-4611-bfc9-375d1082afd4", "workflow": { "status": 1 } }, "field_mask": "workflow"}'
```