# User Interface by using react

## Setting IP, PORT

<!--

Please add `/PDK/ui/pdk-ui/.env` file for using kibana, connect with backend
For example..

```
DB_IP=123.123.123.123
DB_PORT=8080
KIBANA_IP=456.456.456.456
KIBANA_PORT=5601
```
-->

Please add `/PDK/ui/pdk-ui/src/defineUrl.tsx` file for using kibana, connect with backend
For example..

```
import React from 'react';
export const KIBANA_URL = 'http://<KIBANA_IP>:<KIBANA_PORT>';
export const SENSOR_URL = 'http://<DB_IP>:<DB_PORT>/sensor';
export const NODE_URL = 'http://<DB_IP>:<DB_PORT>/node';
```
