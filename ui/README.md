# User Interface by using react

## Setting IP, PORT

### Solution 1 : environment variable

Add `/PDK/ui/pdk-ui/.env.development` file for using kibana, connect with backend.
For example..

```
REACT_APP_DB_IP=123.123.123.123
REACT_APP_DB_PORT=8080
REACT_APP_KIBANA_IP=456.456.456.456
REACT_APP_KIBANA_PORT=5601
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
