// nodeList interface
export interface nodeListElem {
	uuid: string;
	name: string;
	location: string;
	sensors: Array<sensorsElem>;
}

export interface sensorsElem {
	uuid: string;
	name: string;
	value_list: Array<value_list_elem>;
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