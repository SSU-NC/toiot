// LogicCore(lc) interface
export interface numRange {
	min: number;
	max: number;
}
export interface timeRange {
	start: string;
	end: string;
}

export interface lcValue {
    logic: string;
	value: string;
    range: Array<numRange>;
}
export interface lcTime {
    logic: string;
	range: Array<timeRange>;
}
export interface lcGroup {
    logic: string;
    group: Array<string>;
}
export interface lcAction {
    logic: string; 
    text: string;
}

export interface LogicCorePost {
	sensor_uuid: string; 
	logic: Array< lcValue | lcTime | lcGroup | lcAction >;
}
