# User Interface by using react

## Setting IP, PORT

### Solution 1 : environment variable

Add `/PDK/ui/pdk-ui/.env.development` file for using kibana, connect with backend.
For example..

```
DB_IP=123.123.123.123
DB_PORT=8080
KIBANA_IP=456.456.456.456
KIBANA_PORT=5601
```

### Solution 2 : export variable

If you cannot solve the problem, do solution 2.
Modify `/PDK/ui/pdk-ui/src/defineUrl.tsx`.
For example..

```
import React from 'react';
export const KIBANA_URL = 'http://<KIBANA_IP>:<KIBANA_PORT>';
export const SENSOR_URL = 'http://<DB_IP>:<DB_PORT>/sensor';
export const NODE_URL = 'http://<DB_IP>:<DB_PORT>/node';
```
