# User Interface of ToIoT

# Table of contents

<!--ts-->

- [Introduction](#Introdution)
- [Setting](#Setting)
  - [Setting IP, PORT](##1.-Setting-IP,-PORT)
- [Installation & Run](#Installation--Run)
  - [Git](##1.-Git)
  - [Docker](##2.-Docker)

<!--te-->

# Introduce
`ToioT` use typescript, html, css and react framework.

# Setting

## Setting IP, PORT

### Solution 1 : environment variable
We recomment this solution.  
Add `/toiot/ui/.env.development` file for using kibana, connect with backend.  
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
REACT_APP_KAKAO_MAP_KEY=abcdefg12345678
```

### 1.2. Solution 2 : export variable
If you cannot solve the problem, do solution 2.
Modify `/toiot/ui/src/defineUrl.tsx`.  
For example..

```typescript
import React from 'react';
export const KIBANA_URL = 'http://<KIBANA_IP>:<KIBANA_PORT>';
export const SENSOR_URL = 'http://<DB_IP>:<DB_PORT>/sensor';
export const NODE_URL = 'http://<DB_IP>:<DB_PORT>/node';
```

# Installation & Run

## Solution 1 : Git

If you want to start toiot ui at localhost or get source code, you can follow this.

1. Git clone

```shell
git clone https://github.com/SSU-NC/toiot.git
```

2. Setting IP/PORT as [Setting](#Setting)
3. Execute this instruction at `/ui`. After you execute instruction, enter `localhost://3000` at web address. (default port : 3000)

```shell
npm install    // install dependency modules
npm run start  // start react
```

## Solution 2 : Docker

If you use docker, follow this solution.

### 2.1. Download Docker

1. Download Docker image  
   For example...
   ```shell
   docker pull iamhge/toiot-ui:0.1.1
   ```
2. Make docker-compose.yml
   For example...

   ```docker
   version: "3.2"

   services:
   toiot-ui:
       container_name: toiot-ui
       image: iamhge/toiot-ui:0.1.1
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
       - REACT_APP_KAKAO_MAP_KEY=abcdefg12345678
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

5. Enter your docker ip and port at web address. ex) http://123.123.123.123:3001
