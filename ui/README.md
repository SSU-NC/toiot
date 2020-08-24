# User Interface by using react

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

```
import React from 'react';
export const KIBANA_URL = 'http://<KIBANA_IP>:<KIBANA_PORT>';
export const SENSOR_URL = 'http://<DB_IP>:<DB_PORT>/sensor';
export const NODE_URL = 'http://<DB_IP>:<DB_PORT>/node';
```

## 2. Docker

If you use docker, follow this solution.

### 2.1. Download Docker

1. Download Docker image  
   For example...
   ```
   docker pull iamhge/pdk-ui:0.1.1
   ```
2. Make docker-compose.yml
   For example...

   ```
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

3. Execute instruction
   ```
   docker-compose up
   ```

###
