// nodeList interface
export interface nodeListElem {
	id: number;
	name: string;
	lat: number;
	lng: number;
	sink_id: number;
	sink: Array<sinkListElem>;
	sensors: Array<sensorListElem>;
}

export interface locationElem {
	lat: number;
	lng: number;
}

// sensorList interface
export interface sensorListElem {
	id: number;
	name: string;
	sensor_values: Array<value_list_elem>;
	nodes: Array<nodeListElem>;
}

export interface value_list_elem {
	sensor_id: number;
	value_name: string;
	index: number;
}

// sinkList interface
export interface sinkListElem {
	id: number;
	name: string;
	addr: string;
	topic_id: number;
	topic: topicListElem;
	nodes: Array<nodeListElem>;
}

// topic
export interface topicListElem {
	id: number;
	name: string;
	partitions: number;
	replications: number;
	sinks: Array<sinkListElem>;
	logic_services: Array<logic_services_elem>;
}

export interface logic_services_elem {
	id: number;
	addr: string;
	topic_id: number;
}

// used to <Select> options
export interface sensorOptionsElem {
	label: string;
	value: string; // 이때 value는 sensor내의 value값들이 아닌, select component를 사용하기 위해
	id: number;
	sensor_values: Array<value_list_elem>;
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
	n_id: number;
	state: number;
}

// alarm
export interface alarmElem {
	sensor_id: number;
	sensor_name: string;
	msg: string;
}
