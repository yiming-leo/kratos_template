# Kratos Project Template
this template is the modified version for traditional Kratos layout, it has a better design and arch.

## Template Arch
- /api (controller layer)
- /cmd (startup)
- /configs (config.yaml for project)
- /internal (others)
  - /conf (config for server)
  - /data (mapper, client, entity layer)
    - /cluster (cluster function)
    - /entity (model repo)
    - /mapper (mapper function layer for service)
    - /other clients
  - /server (proxy definition)
  - /service (service layer)
- /third_party (do not touch)

## Components
1. gorm: used for easy-parse-orm for Mysql, PostgreSQL, SQLite, SQL Server, and TiDB
2. go-redis: v8

## Install Kratos and Generate Project
1. pls makesure your environment has golang & kratos cli

```bash
go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
```

2. use following cmd to generate a new project <b>(replace <your_kratos_name> to your project name)</b>:
```bash
kratos new <your_kratos_name> --repo-url=https://github.com/ingru-ai/kratos_template.git
```

## Create a service
```
# Add a proto template
kratos proto add api/server/your_server_name.proto
# Generate the proto code  
kratos proto client api/server/your_server_name.proto
# Generate the source code of service by proto file
kratos proto server api/server/your_server_name.proto -t internal/service

go generate ./...
go build -o ./bin/ ./...
./bin/server -conf ./configs
```
now you can enjoy coding!
## Development Regulation
1. DO NOT use Unified Result Set in gRPC server
2. the grocery above shows how to build a correct & normative api & server

## Generate other auxiliary files by Makefile
```
# Download and update dependencies
make init
# Generate API files (include: pb.go, http, grpc, validate, swagger) by proto file
make api
# Generate all files
make all
```

## Docker
```bash
# build
docker build -t <your-docker-image-name> .

# run
docker run --rm -p 8000:8000 -p 9000:9000 -v </path/to/your/configs>:/data/conf <your-docker-image-name>
```

# 