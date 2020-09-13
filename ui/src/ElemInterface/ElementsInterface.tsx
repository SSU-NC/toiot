// nodeList interface
export interface nodeListElem {
	uuid: string;
	name: string;
	group: string;
	location: locationElem;
	sink_id: number;
	sensors: Array<sensorListElem>;
}

export interface locationElem {
	lon: number; // longitude
	lat: number; // latitude
}

// sensorList interface
export interface sensorListElem {
	uuid: string;
	name: string;
	value_list: Array<value_list_elem>;
}

export interface value_list_elem {
	sensor_uuid: string;
	value_name: string;
	index: number;
}

// sinkList interface
export interface sinkListElem {
	id: number;
	name: string;
	location: string;
	ip: string;
}

// used to <Select> options
export interface sensorOptionsElem {
	label: string;
	value: string; // 이때 value는 sensor내의 value값들이 아닌, select component를 사용하기 위해
	uuid: string;
	value_list: Array<value_list_elem>;
}

export interface groupOptionsElem {
	label: string;
	value: string;
}

export interface valueOptionsElem {
	label: string;
	value: string;
}

export interface sinkOptionsElem {
	label: string;
	value: string;
	id: number;
}

// node health check
export interface nodeHealthCheckElem {
	n_uuid: string;
	state: number;
}

// alarm
export interface alarmElem {
	sensor_uuid: string;
	sensor_name: string;
	msg: string;
}
