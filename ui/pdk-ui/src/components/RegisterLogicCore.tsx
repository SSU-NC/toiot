import React, { Component } from 'react';
import { sensorListElem, sensorOptionsElem, nodeListElem } from './ElementsInterface';
import { logicElem, LogicCorePost } from './LcElementsInterface'
import './LogicCore.css';
import SensorCard from './LogicCoreComponents/SensorCard';
import ValueCard from './LogicCoreComponents/ValueCard';
import GroupCard from './LogicCoreComponents/GroupCard';
import TimeCard from './LogicCoreComponents/TimeCard';
import ActionCard from './LogicCoreComponents/ActionCard';
import { } from '../defineUrl';

interface LogicCoreProps{ 
	sensorList: Array<sensorListElem>;
	nodeList: Array<nodeListElem>;
}

interface LogicCoreState{
	logic_name: string;
	sensor_info: sensorOptionsElem;
	selected_value: Array<logicElem>;
	selected_time: logicElem;
	selected_action: Array<logicElem>;
	selected_group: logicElem;

	submit_msg: LogicCorePost;
}

class LogicCore extends Component<LogicCoreProps, LogicCoreState> {
	state: LogicCoreState = {
		logic_name: '',
		sensor_info: {
			uuid: '',
			value: '',
			label: '',
			value_list:[]
		},
		selected_value: [],
		selected_group: {elem:"empty", arg: {group: []}},
		selected_time: {elem:"empty", arg:{range: []}},
		selected_action: [],

		submit_msg: {
			sensor_uuid: '',
			logic_name: '',
			logic: []
		},
	}
	handleLogicNameChange  = (e: React.ChangeEvent<HTMLInputElement>) => {
        this.setState({
            logic_name: e.target.value
        });
    }

	// handle each card change
	handleSensorCardChange = (sensor_info: sensorOptionsElem) => {
		this.setState({
			sensor_info,
		});
	};
	handleGroupCardChange = (selected_group: logicElem) => {
		this.setState({
			//selected_group :{ logic: "group", group : selectedGroups.map((selectedGroup: groupOptionsElem)=>(selectedGroup.value)),},
			selected_group,
		});
	};
	handleValueCardChange = (idx: number) => (selectedValue: logicElem) => {
		const new_selected_value = this.state.selected_value.map((value: logicElem, sidx: number) => {
            if (idx !== sidx) return value;
            return selectedValue;
        });
		this.setState({ selected_value: new_selected_value });
		
	};
	handleTimeCardChange = (selected_time: logicElem) => {
	 this.setState({
			selected_time,
		});
	};
	handleActionCardChange = (idx: number) => (selectedAction: logicElem) => {
		const new_selected_action = this.state.selected_action.map((action: logicElem, sidx: number) => {
            if (idx !== sidx) return action;
            return selectedAction;
        });
		this.setState({ selected_action: new_selected_action });
	};

	// handle add card button click event
	handleAddValueCardClick = () => {
       this.setState({
            selected_value: [...this.state.selected_value, {elem: "empty", arg: {value: "", range:[{min:0,max:255}]}}]
		 });
	}
	handleAddActionCardClick = () => {
		 this.setState({
            selected_action: [...this.state.selected_action, {elem: "empty", arg:{text:""}}]
		 });
	}

	// handle remove card button click event
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
		let logic_array : Array< logicElem>= [ 
			this.state.selected_group, 
			this.state.selected_time,
		];
		logic_array = logic_array.concat(this.state.selected_action, this.state.selected_value);
		
        this.setState({
			//selected_group :{ logic: "group", group : selectedGroups.map((selectedGroup: groupOptionsElem)=>(selectedGroup.value)),},
			submit_msg:  {sensor_uuid: this.state.sensor_info.uuid, logic_name: this.state.logic_name , logic: logic_array},
		});

        var url = ''; // ??
        var data : LogicCorePost = this.state.submit_msg;

        fetch(url, {
        method: 'POST',
        body: JSON.stringify(data),
        headers:{
            'Content-Type': 'application/json'
        }
        }).then(res => res.json())
        .then(response => console.log('Success:', JSON.stringify(response)))
        .catch(error => console.error('Error:', error));
    }

	render() {

		return (
			<div>
				<form>
					<h3>Register Logic</h3>
					<p>A logic is registered at Logic Core</p>
					<br/>
					<div>
						<h5>Logic name</h5>
					<input
						type="text"
						className="form-control"
						name="logic_chain_name"
						placeholder="Enter Logic name"
						value={this.state.logic_name}
						onChange={this.handleLogicNameChange}
					/>
					</div>
					<br/><br/>
					<h5>Build Logic</h5>
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
