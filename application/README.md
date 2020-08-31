# Backend of ToIoT

# Table of contents
- [Registration](#Registration)
- [Logic](#Logic)
- [Health Check](#Health-Check)


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
- GET /sink
- GET /sink/:id
- POST /sink
- DELETE /sink/:id

- GET /node
- GET /node/select
    - query param : key(uuid)
- POST /node
- DELETE /node

- GET /sensor
- POST /sensor
- DELETE /sensor

- GET /registerInfo

# Logic
Manage sensor data pipeline, event processing for sensor data

## Restful API
- GET /logiccore
- POST /logiccore
- DELETE /logiccore
- GET /websocket
    - websocket handler

# Health Check
Manage sensor status

## Status
1. Red
- When not running for "count * tick-interval" time
2. Yellow
- State changes intermittently
3. Green
- When it works well for "count * tick-interval" time

## Restful API
- GET /health-check
    - websocket handler