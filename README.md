
# cadence-study


## Instructions

Install cadence-cli:

```
go install go.uber.org/cadence@latest
```

Start dependencies:
```
docker-compose up -d
```

Create domain:
```
cadence --domain samples-domain domain register
```

Start worker:
```
go run cmd/worker/main.go
```

Trigger workflow:
```
go run cmd/workflow/main.go
```

When you trigger the workflow, you will see on the `worker` output the `TaskToken` (it is a huge hexadecimal).

Trigger task completion:
```
go run cmd/taskcompleter/main.go 00 <PASTE task-token HERE>
```

Note: `00` is the response code that will be passed as response for of the activity.