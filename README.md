# go-twirp-starter-kit

A starting point project to create `todo`-service

## Go Directories  

### `/build`

### `/api`
``` $ brew install protobuf ```  
``` $ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest ```  
``` $ go install github.com/twitchtv/twirp/protoc-gen-twirp@latest ```  

### `/cmd`

#### `/worker`
- [`debug/SAGA_TASK_QUEUE`](http://localhost:8088/namespaces/default/task-queues/SAGA_TASK_QUEUE)

### `/deployments`
Deployments folder contains files for the `k8s` environment; `service`, `deployment` and `ingress`. `dev` folder have an additional `external-mysql-service.yaml` file which allow access from local cluster to external MySQL database.  

`docker-compose up` spin up the MySQL database for local development 

### `/init`

### `/internal`

#### `/ent`
Speed up implementing the database access using [`ent`](https://github.com/facebookincubator/ent).
Of course you can implement repositories with a raw sql statements,
but it is very time consuming and boring repeat x10 times same CRUD functions.

#### [`ent`](https://github.com/facebookincubator/ent)
``` $ cd $PROJECT_ROOT/internal/pkg/ ```  
``` $ go generate ./ent ```  

#### [`twirp`](https://github.com/twitchtv/twirp)

-   [`twriptest`](https://github.com/twitchtv/twirp/tree/master/internal/twirptest)
``` $ protoc --proto_path=$GOPATH/src:. --twirp_out=. --go_out=. todo.proto ```  

#### [`saga`](https://github.com/temporalio/temporal)
