import React, { Component } from 'react';
import Select from 'react-select';
import { sensorListElem, sensorOptionsElem, groupOptionsElem, nodeListElem } from './ElementsInterface';
import { timeRange, lcValue, lcAction, LogicCorePost, lcGroup, lcTime } from './LcElementsInterface'
import './LogicCore.css';
import SensorCard from './LogicCoreComponents/SensorCard';
import ValueCard from './LogicCoreComponents/ValueCard';
import GroupCard from './LogicCoreComponents/GroupCard';
import TimeCard from './LogicCoreComponents/TimeCard';
import ActionCard from './LogicCoreComponents/ActionCard';
import { on } from 'process';

interface LogicCoreProps{ 
	sensorList: Array<sensorListElem>;
	nodeList: Array<nodeListElem>;
}

interface LogicCoreState{
	sensor_info: sensorOptionsElem;
	selected_value: Array<lcValue>;
	selected_time: lcTime;
	selected_action: Array<lcAction>;
	selected_group: lcGroup;

	submit_msg: LogicCorePost;
}

class LogicCore extends Component<LogicCoreProps, LogicCoreState> {
	state: LogicCoreState = {
		sensor_info: {
			uuid: '',
			value: '',
			label: '',
			value_list:[]
		},

		selected_value: [],
		selected_group: {logic:"empty", group: []},
		selected_time: {logic:"empty", range: []},
		selected_action: [],

		submit_msg: {
			sensor_uuid: '',
			logic: []
		},

	}
	handleSensorCardChange = (sensor_info: sensorOptionsElem) => {
		this.setState({
			sensor_info,
		});
	};
	handleGroupCardChange = (selectedGroups: Array<groupOptionsElem>) => {
		this.setState({
			selected_group :{ logic: "group", group : selectedGroups.map((selectedGroup: groupOptionsElem)=>(selectedGroup.value)),},
			
		});
		
	};
	handleValueCardChange = (idx: number) => (selectedValue: lcValue) => {
		const new_selected_value = this.state.selected_value.map((value: lcValue, sidx: number) => {
            if (idx !== sidx) return value;
            return selectedValue;
        });
		this.setState({ selected_value: new_selected_value });
		
	};
	handleTimeCardChange = (selected_time: lcTime) => {
	 this.setState({
			selected_time,
		});
	
	};
	handleActionCardChange = (idx: number) => (selectedAction: lcAction) => {
		const new_selected_action = this.state.selected_action.map((action: lcAction, sidx: number) => {
            if (idx !== sidx) return action;
            return selectedAction;
        });
		this.setState({ selected_action: new_selected_action });
		
	};

	handleAddValueCardClick = () => {
       this.setState({
            selected_value: [...this.state.selected_value, {logic: "empty", value: "", range:[{min:0,max:255}]}]
		 });
	}
	handleAddActionCardClick = () => {
		 this.setState({
            selected_action: [...this.state.selected_action, {logic: "empty", text:""}]
		 });
	}
	handleRemoveValueCardClick = (idx: number) => () => {
        this.setState({
           selected_value: this.state.selected_value.filter((s: any, sidx:number) => idx !== sidx)
		});
	
	};
	handleRemoveActionCardClick = (idx: number) => () => {
        this.setState({
            selected_action: this.state.selected_action.filter((s: any, sidx:number) => idx !== sidx)
		});
	
	};

    handleSubmit = (e: React.MouseEvent<HTMLButtonElement>) => {
		//e.preventDefault();
		let logic_array : Array< lcValue | lcTime | lcGroup | lcAction>= [ 
			this.state.selected_group, 
			this.state.selected_time,
		];
		logic_array = logic_array.concat(this.state.selected_action, this.state.selected_value);
		
        this.setState({
			//selected_group :{ logic: "group", group : selectedGroups.map((selectedGroup: groupOptionsElem)=>(selectedGroup.value)),},
			submit_msg:  {sensor_uuid: this.state.sensor_info.uuid , logic: logic_array},
		});

        // var url = ;
        // var data : LogicCorePost = this.state.submit_msg;

        // fetch(url, {
        // method: 'POST', // or 'PUT'
        // body: JSON.stringify(data),
        // headers:{
        //     'Content-Type': 'application/json'
        // }
        // }).then(res => res.json())
        // .then(response => console.log('Success:', JSON.stringify(response)))
        // .catch(error => console.error('Error:', error));
    }

	render() {

		return (
			<div>
				<form>
					<SensorCard sensorList={this.props.sensorList} handleSensorCardChange={this.handleSensorCardChange}/>
					<GroupCard nodeList={this.props.nodeList} handleGroupCardChange={this.handleGroupCardChange}/>
					<TimeCard handleTimeCardChange={this.handleTimeCardChange}/>
					{this.state.selected_value.map((d: any , idx: number) => (
						<ValueCard valueList={this.state.sensor_info.value_list} handleRemoveValueCardClick={this.handleRemoveValueCardClick(idx)} handleValueCardChange={this.handleValueCardChange(idx)}/>
					))}
					{this.state.selected_action.map((d: any , idx: number) => (
					<ActionCard handleActionCardChange={this.handleActionCardChange(idx)} handleRemoveActionCardClick={this.handleRemoveActionCardClick(idx)}/>
					))}
					<button type="button" className="btn margin-right" style={{background:'pink'}} onClick={this.handleAddValueCardClick}>Add value</button>
					<button type="button" className="btn" style={{background:'pink'}} onClick={this.handleAddActionCardClick}>Add action</button>
					<p></p>
					<button
						//type="submit"
						type="button"
						className="btn float-right"
						
						style={{ background: 'pink' }}
						onClick={this.handleSubmit}
					>
						Submit
					</button>
				</form>
				
			</div>
		);
	}
}

export default LogicCore;
