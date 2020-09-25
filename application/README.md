# Backend of ToIoT

# Table of contents
- [Installation](#Installation)
- [Run](#Run)

# Installation
```bash
$ go get github.com/KumKeeHyun/toiot
```
or
```bash
// toiot/application
$ go get github.com/KumKeeHyun/toiot/application

// toiot/logic-core
$ go get github.com/KumKeeHyun/toiot/logic-core

// not yet
// toiot/health-check
$ go get github.com/KumKeeHyun/toiot/health-check
```

## docker
```bash
// latest tag is 0.3

// toiot/application
$ docker pull kbzjung359/toiot-app

// toiot/logic-core
$ docker pull kbzjung359/toiot-logic

// not yet
// toiot/health-check
$ docker pull kbzjung359/toiot-health
```

or

```bash
// In the directory where the Dockerfile are located
$ docker build -t (image name) .
```

# Run
```bash
// toiot/application, toiot/logic-core, toiot/health-check
$ go run main.go
```

or
```bash
// In the directory where the docker-compose.yml are located
$ docker-compose up
```


## Environment variable
The setting value of each microservice is set by reading environment variables.
### toiot/application (Registration)
|Key|Example|Explain|
|------|------|------|
|APP_SERVER|127.0.0.1:8081|Registration microservice address|
|DB_DRIVER|mysql|DBMS type|
|DB_SERVER|localhost:3306|DBMS address|
|DB_USER|pdk|User name|
|DB_PASS|pdk1234|User password|
|DB_DATABASE|pdk-db|mysql database name|
|TOPIC_NAME|sensor-data|default kafka topic for registration info|
|TOPIC_PARTITIONS|1|topic's partitions|
|TOPIC_RePLICATIONS|1|topic's replications|

- Currently DBMS only supports mysql

### toiot/logic-core (Logic)
|Key|Example|Explain|
|------|------|------|
|LOGIC_SERVER|127.0.0.1:8082|Logic microservice address|
|APP_SERVER|127.0.0.1:8081|Registration microservice server address|
|KAFKA_BROKER|localhost:9092|Kafka Cluster address|
|KAFKA_GROUP|logic-core|Kafka consumer group id|
|KAFKA_TOPIC|sensor-data|Kafka topic to steam sensor data|
|KAFKA_BUFSIZE|500|channel size for stream process in goroutines|
|ELASTIC_SERVER|localhost:9200|Elasticsearch cluster address|
|ELASTIC_RETRY|3|Rest API request retry when the requests fail|
|ELASTIC_BUFSIZE|500|channel size of stream process in goroutines|
|ELASTIC_BATCHTICKER|5|Bulk Indexing interval second|
|ELASTIC_BATCHSIZE|400|Bulk Indexing documents size|


### toiot/health-check (Health Check)
|Key|Example|Explain|
|------|------|------|
|HEALTH_SERVER|127.0.0.1:8083|Health microservice address|
|LOGIC_SERVER|127.0.0.1:8081|Logic microservice address|
|KAFKA_BROKER|localhost:9092|Kafka Cluster address|
|KAFKA_GROUP|health-check|Kafka consumer group id|
|KAFKA_TOPIC|health-data|Kafka topic to steam sensor data|
|KAFKA_BUFSIZE|100|channel size for stream process in goroutines|
|STATUS_COUNT|5|count to judge health state|
|STATUS_TICK|30| interval second to subtract|
|STATUS_DROP|12|drop metadata in repository interval hour when red state is maintained|
