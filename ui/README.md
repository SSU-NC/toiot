# User Interface by using react

# Table of contents

<!--ts-->

- [Introduce](#Introduce)
- [Setting](#Setting)
- [Execute](#Execute)
- [Structure](#Structure)
- [Logic Core](#Logic-Core)

<!--te-->

# Introduce

ToioT use typescript, html, css and react frame work.

# Setting

## 1. Setting IP, PORT

### 1.1. Solution 1 : environment variable

We recomment this solution then solution 2.  
Add `/PDK/ui/pdk-ui/.env.development` file for using kibana, connect with backend.  
For example..

```
REACT_APP_DB_IP=0.0.0.0
REACT_APP_DB_PORT=8080
REACT_APP_KIBANA_IP=0.0.0.0
REACT_APP_KIBANA_PORT=8080
REACT_APP_HEALTHCHECK_IP=0.0.0.0
REACT_APP_HEALTHCHECK_PORT=8080
REACT_APP_LOGICCORE_IP=0.0.0.0
REACT_APP_LOGICCORE_PORT=8080
REACT_APP_ALARM_IP=0.0.0.0
REACT_APP_ALARM_PORT=8080
```

### 1.2. Solution 2 : export variable

If you cannot solve the problem, do solution 2.
Modify `/PDK/ui/pdk-ui/src/defineUrl.tsx`.  
For example..

```typescript
import React from "react";
export const KIBANA_URL = "http://<KIBANA_IP>:<KIBANA_PORT>";
export const SENSOR_URL = "http://<DB_IP>:<DB_PORT>/sensor";
export const NODE_URL = "http://<DB_IP>:<DB_PORT>/node";
```

## 2. Docker

If you use docker, follow this solution.

### 2.1. Download Docker

1. Download Docker image  
   For example...
   ```shell
   docker pull iamhge/pdk-ui:0.1.1
   ```
2. Make docker-compose.yml
   For example...

   ```docker
   version: "3.2"

   services:
   pdk-ui:
       container_name: pdk-ui
       image: iamhge/pdk-ui:0.1.1
       ports:
       - "3001:3000"
       environment:
       - NODE_ENV=development
       - CHOIDAR_USEPOLLING=true
       - REACT_APP_DB_IP=0.0.0.0
       - REACT_APP_DB_PORT=8080
       - REACT_APP_KIBANA_IP=0.0.0.0
       - REACT_APP_KIBANA_PORT=8080
       - REACT_APP_HEALTHCHECK_IP=0.0.0.0
       - REACT_APP_HEALTHCHECK_PORT=8080
       - REACT_APP_LOGICCORE_IP=0.0.0.0
       - REACT_APP_LOGICCORE_PORT=8080
       - REACT_APP_ALARM_IP=0.0.0.0
       - REACT_APP_ALARM_PORT=8080
       stdin_open: true
       tty: true
   ```

3. Execute instruction. React will start.

   ```shell
   docker-compose up
   ```

4. Find your docker ip by type like next. You'll find the ip where the docker runs.
   ```shell
   ip addr | grep inet
   ```

# Execute

If you want to start react, execute this instruction at `/ui/pdk-ui`. After you execute instruction, enter `localhost://3000` in address.

```shell
npm install    // install dependency modules
npm run start  // start react
```

# Structure

All components are routed at `src/App.tsx`.  
`App.tsx` request nodeList, sensorList, sinkList and logicCore to backend, and pass on to lower components by props.

![components_structure](./img/components_structure.png)

## 1. Navigation bar

Navigation bar is implemented so that each page can be entered.  
Used Compoponent : `src/Navigation.tsx`  
![navigation_bar](./img/navigation_bar.png)

## 2. HOME

User can see Kibana Dashboard at here.  
Used Component : `src/Home.tsx`  
![navbar_home](./img/navbar_home.png)

## 3. MANAGEMENT

`MANAGEMENT` tab's components consist of `Sensor`, `Node`, and `Sink` tab.

### 3.1. components

Basically, components consist of sensor, node, and sink. User can manage each sensor, node, and sink at `MANAGEMENT` tab.

![navbar_management](./img/navbar_management.png)

#### 3.1.1. Sensor

Used Components : `src/components/SensorManagement.tsx`, `src/components/Register/RegisterSensor.tsx`, `src/components/Table/SensorTable`

##### SensorManagement.tsx

This component manage sensor tab.

##### RegisterSensor.tsx

User can register sensor by using this component.  
When user click `register sensor` button, a modal will show up.  
User enter sensor's informations( sensor name, sensor's values' name ).  
For example...  
![register_sensor_ex](./img/register_sensor_ex.png)

- sensor name
- sensor's values : Sensor can have more than two values, so user can register all values' name.

##### SensorTable.tsx

User can see sensors' informations.  
User can delete sensors by click wastebasket img.  
![sensor_table](./img/sensor_table.png)

#### 3.1.2. Node

Used Components : `src/components/NodeManagement.tsx`, `src/components/Register/RegisterNode.tsx`, `src/components/Table/NodeTable`

##### NodeManagement.tsx

This component manage node tab.

##### RegisterNode.tsx

User can register node by using this component.  
When user click `register node` button, a modal will show up.  
User enter node's informations( node name, group, location, selecting sensors... ).  
For example...  
![register_node_ex](./img/register_node_ex.png)

- node name
- group : The group of nodes is then used to group when registering logics. (e.g.location info(Seoul, Busan, ...))
- location : User should type latitude and longitude.
- sensors : User should select sensors which want to make belong to this node.
- sink : User should select sink which want to connect.

##### NodeTable.tsx

User can see nodes' informations.  
User can delete nodes by click wastebasket img.  
![node_table](./img/node_table.png)

### 3.1.4. Sink

Used Components : `src/components/SinkManagement.tsx`, `src/components/Register/RegisterSink.tsx`, `src/components/Table/SinkTable`

##### SinkManagement.tsx

This component manage sink tab.

##### RegisterSink.tsx

User can register sink by using this component.  
When user click `register sink` button, a modal will show up.  
User enter sink's informations( sink name, location, ip:port ).  
For example...  
![register_sink_ex](./img/register_sink_ex.png)

- sink name
- location
- ip:port : User should type sink's ip:port info.

##### SinkTable.tsx

User can see sinks' informations.  
User can delete sinks by click wastebasket img.  
![sink_table](./img/sink_table.png)

## 4. Service

![navbar_service](./img/navbar_service.png)

## 5. KIBANA

![navbar_kibana](./img/navbar_kibana.png)

### 5.1. Visualize

### 5.2. Dashboard

## 6. Alarm

Alert alarm should be service all time, so this function is implemented by 'web socket'.  
This service can be serviced by registering logic.  
Used Component : `src/components/AlertAlarm.tsx`

![alarm_service](./img/alarm_service.png)

# Logic Core

## 1. Logic

A logic

## 2.
