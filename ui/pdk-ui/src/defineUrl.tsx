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
	.concat(process.env.REACT_APP_KIBANA_PORT)
	.concat('/app/kibana#/dashboard/7adfa750-4c81-11e8-b3d7-01146121b73d?_g=(refreshInterval%3A(pause%3A!f%2Cvalue%3A900000)%2Ctime%3A(from%3Anow-24h%2Cto%3Anow),filters:!(),fullScreenMode:!f)')
export const KIBANA_URL_1 = 'http://'
	.concat(process.env.REACT_APP_KIBANA_IP)
	.concat(':')
	.concat(process.env.REACT_APP_KIBANA_PORT)

export const KIBANA_URL_2 = 'http://'
	.concat(process.env.REACT_APP_KIBANA_IP)
	.concat(':')
	.concat(process.env.REACT_APP_KIBANA_PORT)



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

export const SINK_URL = 'http://'
	.concat(process.env.REACT_APP_DB_IP)
	.concat(':')
	.concat(process.env.REACT_APP_DB_PORT)
	.concat('/sink');

export const LOGIC_URL = 'http://'
	.concat(process.env.REACT_APP_DB_IP)
	.concat(':')
	.concat(process.env.REACT_APP_DB_PORT)
	.concat('/logic/new');

