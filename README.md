# GetJob
[![Build](https://github.com/quangdangfit/getjob/workflows/master/badge.svg)](https://github.com/quangdangfit/getjob/actions)

Backend service for *Human Resources System* based on Golang, Mysql

## How to run

### Required

- Postgres
- Redis

### Config
Simply run `make startup`, or run following commands step - by - step:
- Copy config file: `cp config/config.sample.yaml config/config.yaml`


- You should modify `config/config.yaml`

```yaml
database:
  host: localhost
  port: 5432
  name: getjob
  user: mysql
  password: 1234

redis:
  enable: true
  host: localhost
  port: 6397
  password:
  database: 0

cache:
  enable: true
  expiry_time: 3600
```

### Migration - Create Admin User
```shell script
$ make admin
```

### Run
```shell script
$ go run main.go 
```

### Document
* API document at: `http://localhost:8080/swagger/index.html`

### Structure
```shell
├── src
│   ├── api             # Handle request & response
│   ├── dbs             # Database Layer
│   ├── middleware      # Middlewares
│   │   ├── cache           # Cache middleware
│   │   ├── jwt             # JWT middleware
│   │   └── roles           # Authorization middleware
│   ├── migration       # Migration
│   ├── mocks           # Mocks
│   ├── models          # Models
│   ├── repositories    # Repository Layer
│   │   └── impl            # Implement repositories
│   ├── router          # Router api
│   ├── schema          # Schemas
│   ├── services        # Business Logic Layer
│   │   └── impl            # Implement services
├── cmd                 # Contains commands 
├── config              # Config files 
├── docs                # Swagger API document
├── pkg                 # Internal packages
│   ├── app                 # App packages
│   ├── errors              # Errors packages
│   ├── http                # HTTP packages
│   ├── jwt                 # JWT packages
│   └── utils               # Utils packages
├── scripts             # Scripts
├── test                # Test package
```
