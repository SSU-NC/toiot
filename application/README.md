# Backend of ToIoT

# Table of contents
- [Installation](#Installation)
- [Run](#Run)
- [Registration](#Registration)
- [Logic](#Logic)
- [Health Check](#Health-Check)

# Installation
```bash
$ go get github.com/KumKeeHyun/PDK
```
or
```bash
// toiot/application
$ go get github.com/KumKeeHyun/PDK/application

// toiot/logic-core
$ go get github.com/KumKeeHyun/PDK/logic-core

// toiot/health-check
$ go get github.com/KumKeeHyun/PDK/health-check
```

## docker
```bash
// latest tag is 0.3

// toiot/application
$ docker pull kbzjung359/pdk-app

// toiot/logic-core
$ docker pull kbzjung359/pdk-logic

// toiot/health-check
$ docker pull kbzjung359/pdk-health
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
|LOGIC_SERVER|127.0.0.1:8082|Logic microservice address|
|DB_DRIVER|mysql|DBMS type|
|DB_SERVER|localhost:3306|DBMS address|
|DB_USER|pdk|User name|
|DB_PASS|pdk1234|User password|
|DB_DATABASE|pdk-db|mysql database name|

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
|MONGO_ADDR|127.0.0.1|MongoDB address for logic repository|
|MONGO_PORT|27017|MongoDB port|


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

# Registration
Manage sensor registration and deletion

## Structure
The sensors are divided into three concepts and managed

### Sensor
A sensor used for measurement
```go
type Sensor struct {
    Name string
    ValueList []struct {
        Name string
    }
}
```

### Node
A board installed in a specific place to form a sensor network
```go
type Node struct {
    Name, Group string
    Location struct {
        Lat float64
        Lon float64
    }
    SinkID uint
    Sensors []Sensor
}
```

### Sink
A board that router, sink of sensor network
```go
type Sink struct {
    Name, Location, IP string
    Nodes []Node
}
```

## UUID
Unique number assigned to each sensor and node. uuid is used for sensor data collection, data enrichment.
- [https://github.com/rs/xid]

## Restful API
|Method|Path|Return/Request|
|---|------|------|
|GET|/sink|list sinks|
|GET|/sink/:id|sink that id is :id|
|POST|/sink|regist sink|
|DELETE|/sink/:id|delete sink|
|GET|/node|list nodes|
|GET|/node/select|list nodes that uuid is in query|
|POST|/node|regist node|
|DELETE|/node|delete node|
|GET|/sensor|list sensors|
|POST|/sensor|regist sensor|
|DELETE|/sensor|delete sensor|
|GET|/registerInfo|list node, sensor's metadata for logic-core|

# Logic
Manage sensor data pipeline, event processing for sensor data

## Structure
```go
type Ring struct {
	SensorUUID string
	LogicName string
	Logic []struct {
		Element string
		Argument map[string]interface{} 
	}
}
```

## Restful API
|Method|Path|Explain|
|---|------|------|
|GET|/logiccore|list logics|
|POST|/logiccore|regist logic|
|DELETE|/logiccore|delete logic|
|GET|/websocket|websocket connect for alert|

# Health Check
Manage sensor status

## Structure
```go
type States struct {
    Timestamp string
    State []struct{
        NodeUUID string
        State bool
    }
}
```

## Status
1. Red
- When not running for "count * tick-interval" time
2. Yellow
- State changes intermittently
3. Green
- When it works well for "count * tick-interval" time

## Restful API
|Method|Path|Explain|
|---|------|------|
|GET|/health-check|websocket connect for state data|