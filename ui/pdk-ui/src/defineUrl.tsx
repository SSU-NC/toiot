import React from 'react';

declare global {
	namespace NodeJS {
		interface ProcessEnv {
			REACT_APP_KIBANA_IP: string;
			REACT_APP_KIBANA_PORT: string;
			REACT_APP_DB_IP: string;
			REACT_APP_DB_PORT: string;
		}
	}
}

export const KIBANA_URL = 'http://'
	.concat(process.env.REACT_APP_KIBANA_IP)
	.concat(':')
	.concat(process.env.REACT_APP_KIBANA_PORT);

export const KIBANA_VISUALIZE_URL = 'http://'
	.concat(process.env.REACT_APP_KIBANA_IP)
	.concat(':')
	.concat(process.env.REACT_APP_KIBANA_PORT)	
	.concat('/app/kibana#/visualize');
export const KIBANA_DASHBOARDS_URL= 'http://'
    .concat(process.env.REACT_APP_KIBANA_IP)
    .concat(':')
    .concat(process.env.REACT_APP_KIBANA_PORT)	
    .concat('/app/kibana#/dashboards');
export const SENSOR_URL = 'http://'
	.concat(process.env.REACT_APP_DB_IP)
	.concat(':')
	.concat(process.env.REACT_APP_DB_PORT)
	.concat('/sensor');

export const NODE_URL = 'http://'
	.concat(process.env.REACT_APP_DB_IP)
	.concat(':')
	.concat(process.env.REACT_APP_DB_PORT)
	.concat('/node');
