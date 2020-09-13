// Define URLs

declare global {
	namespace NodeJS {
		interface ProcessEnv {
			REACT_APP_KIBANA_IP: string;
			REACT_APP_KIBANA_PORT: string;
			REACT_APP_DB_IP: string;
			REACT_APP_DB_PORT: string;
			REACT_APP_LOGICCORE_IP: string;
			REACT_APP_LOGICCORE_PORT: string;
			REACT_APP_HEALTHCHECK_IP: string;
			REACT_APP_HEALTHCHECK_PORT: string;
			REACT_APP_ALARM_IP: string;
			REACT_APP_ALARM_PORT: string;
		}
	}
}

export const KIBANA_URL = 'http://'
	.concat(process.env.REACT_APP_KIBANA_IP)
	.concat(':')
	.concat(process.env.REACT_APP_KIBANA_PORT)
	.concat(
		'/app/kibana#/dashboard/7adfa750-4c81-11e8-b3d7-01146121b73d?_g=(refreshInterval%3A(pause%3A!f%2Cvalue%3A900000)%2Ctime%3A(from%3Anow-24h%2Cto%3Anow),filters:!(),fullScreenMode:!f)'
	);
export const KIBANA_VISUALIZE_URL = 'http://'
	.concat(process.env.REACT_APP_KIBANA_IP)
	.concat(':')
	.concat(process.env.REACT_APP_KIBANA_PORT)
	.concat('/app/kibana#/visualize');
export const KIBANA_DASHBOARDS_URL = 'http://'
	.concat(process.env.REACT_APP_KIBANA_IP)
	.concat(':')
	.concat(process.env.REACT_APP_KIBANA_PORT)
	.concat('/app/kibana#/dashboards');

export const SENSOR_URL = 'http://'
	.concat(process.env.REACT_APP_DB_IP)
	.concat(':')
	.concat(process.env.REACT_APP_DB_PORT)
	.concat('/regist/sensor');
export const NODE_URL = 'http://'
	.concat(process.env.REACT_APP_DB_IP)
	.concat(':')
	.concat(process.env.REACT_APP_DB_PORT)
	.concat('/regist/node');
export const SINK_URL = 'http://'
	.concat(process.env.REACT_APP_DB_IP)
	.concat(':')
	.concat(process.env.REACT_APP_DB_PORT)
	.concat('/regist/sink');

export const LOGICCORE_URL = 'http://'
	.concat(process.env.REACT_APP_LOGICCORE_IP)
	.concat(':')
	.concat(process.env.REACT_APP_LOGICCORE_PORT)
	.concat('/regist/logic');

export const HEALTHCHECK_URL = 'ws://'
	.concat(process.env.REACT_APP_HEALTHCHECK_IP)
	.concat(':')
	.concat(process.env.REACT_APP_HEALTHCHECK_PORT)
	.concat('/health-check');

export const ALARM_URL = 'ws://'
	.concat(process.env.REACT_APP_ALARM_IP)
	.concat(':')
	.concat(process.env.REACT_APP_ALARM_PORT)
	.concat('/websocket');
