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
    logic: "value";
	value: string;
    range: Array<numRange>;
}
export interface lcTime {
    logic: "group";
	range: Array<timeRange>;
}

export interface lcGroup {
    logic: "group";
    group: Array<string>;
}
export interface lcAction {
    logic: string; 
    text: string;
}

export interface lcAlarm {
    logic: "alarm"; 
	msg: string;
}
export interface lcEmail {
    logic: "email"; 
    email: string;
}
