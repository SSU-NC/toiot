import React, { Component } from 'react';
import Select from 'react-select';
import { sensorListElem, sensorOptionsElem, groupOptionsElem, nodeListElem } from './ElementsInterface';
import { timeRange, lcValue, lcTime, lcGroup, lcAlarm, lcEmail, lcAction} from './LcElementsInterface'
import './LogicCore.css';
import SensorCard from './LogicCoreComponents/SensorCard';
import ValueCard from './LogicCoreComponents/ValueCard';
import GroupCard from './LogicCoreComponents/GroupCard';
import TimeCard from './LogicCoreComponents/TimeCard';
import ActionCard from './LogicCoreComponents/ActionCard';

interface LogicCoreProps{ 
	sensorList: Array<sensorListElem>;
	nodeList: Array<nodeListElem>;
}

interface LogicCoreState{
	sensor: sensorOptionsElem;
	group: Array<groupOptionsElem>;
	value: Array<lcValue>;
	time_range: Array<timeRange>;
	action: Array<lcAction>;
}

interface LogicCorePost {
	sensor_uuid: string; 
	logic: Array<lcValue | lcTime | lcGroup | lcAlarm | lcEmail>;
}

class LogicCore extends Component<LogicCoreProps, LogicCoreState> {
	state: LogicCoreState = {
		sensor: {
			uuid: '',
			value: '',
			label: '',
			value_list:[]
		},
		group: [],
		value: [],
		time_range: [],
		action: [],
	}
	handleSensorCardChange = (sensor: sensorOptionsElem) => {
		this.setState({
			sensor,
		});
	};
	handleGroupCardChange = (group: Array<groupOptionsElem>) => {
		this.setState({
			group,
		});
	};
	handleValueCardChange = (value: lcValue) => {
		this.setState({
			value: [value],
		});
	};
	handleTimeCardChange = (time_range: Array<timeRange>) => {
		this.setState({
			time_range,
		});
	};
	handleActionCardChange = (action: lcAction) => {
		this.setState({
			action: [action]
		});
	};
	render() {

		return (
			<div>
				<form>
					<SensorCard sensorList={this.props.sensorList} handleSensorCardChange={this.handleSensorCardChange}/>
					<ValueCard valueList={this.state.sensor.value_list} handleValueCardChange={this.handleValueCardChange}/>
					<GroupCard nodeList={this.props.nodeList} handleGroupCardChange={this.handleGroupCardChange}/>
					<TimeCard handleTimeCardChange={this.handleTimeCardChange}/>
					<ActionCard handleActionCardChange={this.handleActionCardChange}/>
					<button
						type="submit"
						className="btn float-right"
						data-dismiss="modal"
						style={{ background: 'pink' }}
					>
						Submit
					</button>
				</form>
			</div>
		);
	}
}

export default LogicCore;
