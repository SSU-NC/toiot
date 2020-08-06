import React, { Component } from 'react';
import Select from 'react-select';
import { sensorListElem, sensorOptionsElem, valueOptionsElem, groupOptionsElem, nodeListElem } from './ElementsInterface';
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
	values: Array<valueOptionsElem>;
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
		values: [], 
	}
	handleSensorCardChange = (sensor: any) => {
		this.setState({
			sensor,
		});
	};
	handleGroupCardChange =(group: any) => {
		this.setState({
			group,
		});
	};
	handleValueCardChange = (values: any) => {
		this.setState({
			values,
		});
	};
	render() {

		return (
			<div>
				<form>
					<SensorCard sensorList={this.props.sensorList} handleSensorCardChange={this.handleSensorCardChange}/>
					<ValueCard valueList={this.state.sensor.value_list} handleValueCardChange={this.handleValueCardChange}/>
					<GroupCard nodeList={this.props.nodeList} handleGroupCardChange={this.handleGroupCardChange}/>
					<TimeCard/>
					<ActionCard/>
					<button
						type="submit"
						className="btn"
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
