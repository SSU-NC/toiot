import React, { Component } from 'react';
import {
	sensorListElem,
	sensorOptionsElem,
	nodeListElem,
} from '../ElemInterface/ElementsInterface';
import { logicElem } from '../ElemInterface/LcElementsInterface';
import './LogicCore.css';
import InputSensorCard from './InputCards/InputSensorCard';
import InputValueCard from './InputCards/InputValueCard';
//import InputGroupCard from './InputCards/InputGroupCard';
import InputTimeCard from './InputCards/InputTimeCard';
import InputActionCard from './InputCards/InputActionCard';
import { LOGICCORE_URL, SENSOR_URL, NODE_URL } from '../defineUrl';
import { Link } from 'react-router-dom';

interface RegisterLogicState {
	sensorList: Array<sensorListElem>;
	nodeList: Array<nodeListElem>;

	logic_name: string;
	sensor_info: sensorOptionsElem;
	selected_value: Array<logicElem>;
	selected_time: logicElem;
	selected_action: Array<logicElem>;
	selected_group: logicElem;
	nameValid: boolean;
	sensorValid: boolean;
	actionValid: boolean;
}

/* 
RegisterLogic
- Linked by register logic button
- register logic
*/
class RegisterLogic extends Component<{}, RegisterLogicState> {
	state: RegisterLogicState = {
		sensorList: [],
		nodeList: [],

		logic_name: '',
		sensor_info: {
			id: 0,
			value: '',
			label: '',
			sensor_values: [],
		},
		selected_value: [],
		selected_group: { elem: 'empty', arg: { group: [] } },
		selected_time: { elem: 'empty', arg: { range: [] } },
		selected_action: [],

		nameValid: false,
		sensorValid: false,
		actionValid: false,
	};
	componentDidMount() {
		this.getsensorList();
		this.getnodeList();
	}

	// Get sensor list from backend
	getsensorList() {
		var url = SENSOR_URL;

		fetch(url)
			.then((res) => res.json())
			.then((data) => {
				this.setState({ sensorList: data });
			})
			.catch((error) => console.error('Error:', error));
	}

	// Get node list from backend
	getnodeList() {
		var url = NODE_URL;

		fetch(url)
			.then((res) => res.json())
			.then((data) => this.setState({ nodeList: data }))
			.catch((error) => console.error('Error:', error));
	}

	// Handle node name change by typing
	handleLogicNameChange = (e: React.ChangeEvent<HTMLInputElement>) => {
		// name valid check : user should enter logic name
		if (e.target.value.length > 0) {
			this.setState({
				logic_name: e.target.value,
				nameValid: true,
			});
		} else {
			this.setState({
				logic_name: e.target.value,
				nameValid: false,
			});
		}
	};

	// handle sensor card change
	handleSensorCardChange = (sensor_info: sensorOptionsElem) => {
		// sensor valid check : user should select sensor
		if (sensor_info !== null) {
			this.setState({
				sensor_info,
				sensorValid: true,
			});
		} else {
			this.setState({
				sensor_info,
				sensorValid: false,
			});
		}
	};

	// handle group card change
	handleGroupCardChange = (selected_group: logicElem) => {
		this.setState({
			//selected_group :{ logic: "group", group : selectedGroups.map((selectedGroup: groupOptionsElem)=>(selectedGroup.value)),},
			selected_group,
		});
	};

	// handle value card change
	handleValueCardChange = (idx: number) => (selectedValue: logicElem) => {
		// Value card is updated dynamic. It can be added or removed freely.
		// so find changing field by using received idx and change state.
		const new_selected_value = this.state.selected_value.map(
			(value: logicElem, sidx: number) => {
				if (idx !== sidx) return value;
				return selectedValue;
			}
		);
		this.setState({ selected_value: new_selected_value });
	};

	// Handle time card change
	handleTimeCardChange = (selected_time: logicElem) => {
		this.setState({
			selected_time,
		});
	};

	// Handle action card change
	handleActionCardChange = (idx: number) => (selectedAction: logicElem) => {
		// Action card is updated dynamic. It can be added or removed freely.
		// so find changing field by using received idx and change state.
		const new_selected_action = this.state.selected_action.map(
			(action: logicElem, sidx: number) => {
				if (idx !== sidx) return action;
				return selectedAction;
			}
		);

		// action valid check : User should register more than a action
		if (
			new_selected_action !== null &&
			new_selected_action !== [] &&
			!new_selected_action.some((value) => value.elem === 'empty')
		) {
			this.setState({
				selected_action: new_selected_action,
				actionValid: true,
			});
		} else {
			this.setState({
				selected_action: new_selected_action,
				actionValid: false,
			});
		}
	};

	// handle add value card button click event
	handleAddValueCardClick = () => {
		this.setState({
			selected_value: [
				...this.state.selected_value,
				{ elem: 'empty', arg: { value: '', range: [{ min: 0, max: 255 }] } },
			],
		});
	};

	// handle add action card button click event
	handleAddActionCardClick = () => {
		this.setState({
			selected_action: [
				...this.state.selected_action,
				{ elem: 'empty', arg: { text: '' } },
			],
		});
	};

	// handle remove value card button click event
	handleRemoveValueCardClick = (idx: number) => () => {
		this.setState({
			selected_value: this.state.selected_value.filter(
				(s: any, sidx: number) => idx !== sidx
			),
		});
	};

	// handle remove action card button click event
	handleRemoveActionCardClick = (idx: number) => () => {
		this.setState({
			selected_action: this.state.selected_action.filter(
				(s: any, sidx: number) => idx !== sidx
			),
		});
	};

	// selected sensor에 따른 value_list에서, selected_value를 제외한 list 추출
	// this didn't work I have to implement this function..
	// getUnselectedValueList() {
	// 	var valueOptions = this.state.sensor_info.value_list;

	// 	for (let selected_value of this.state.selected_value) {
	// 		if (selected_value.elem !== 'value') continue;
	// 		valueOptions = valueOptions.filter(
	// 			(value) =>
	// 				!value.value_name.includes((selected_value.arg as lcValueArg).value)
	// 		);
	// 	}
	// 	this.setState({
	// 		valueOptions,
	// 	});
	// }

	handleSubmit = (e: React.MouseEvent<HTMLButtonElement>) => {
		e.preventDefault();
		let elems: Array<logicElem> = [
			this.state.selected_group,
			this.state.selected_time,
		];
		elems = elems.concat(
			this.state.selected_value,
			this.state.selected_action
		);

		// Filter elem: 'empty' field
		elems = elems.filter(function (logic) {
			return logic.elem !== 'empty';
		});

		var request_msg = {
			sensor_id: this.state.sensor_info.id,
			logic_name: this.state.logic_name,
			elems: elems,
		};
		var url = LOGICCORE_URL;

		// Valid check (unvalid -> alert)
		if (!this.state.nameValid) {
			alert('Please enter logic name.');
			return;
		}
		if (!this.state.sensorValid) {
			alert('Please select a sensor.');
			return;
		}
		if (!this.state.actionValid) {
			alert('Please set more than a action.');
			return;
		}

		// Check whether user really want to submit
		var submitValid: boolean;
		submitValid = window.confirm('Are you sure to register this logic?');
		if (!submitValid) {
			return;
		}

		fetch(url, {
			method: 'POST',
			body: JSON.stringify(request_msg),
			headers: {
				'Content-Type': 'application/json',
			},
		})
			.then((res) => res.json())
			.then((response) => console.log('Success:', JSON.stringify(response)))
			.catch((error) => console.error('Error:', error))
			.then(() => window.location.reload(false));
	};

	render() {
		return (
			<div>
				<form>
					<h2>Register Logic (Rule)</h2>
					<p style={{ fontSize: '14pt' }}>
						A logic is registered at Logic Core. A logic consists of
						elements(sensor, group, time, value, action). This is a kind of rule
						chain. You can build logic about a sensor, and that logic's elements
						would be 'rule' for action event(e.g. sending email, alert massage).
					</p>
					<p style={{ fontSize: '14pt' }}>
						For example, if you want to get email when 'sensor 1's 'value 1's
						scope is `0 ~ 15`, you should build logic just like you want. Select
						sensor 'sensor 1', select value 'value 1', and set action to send an
						email.
					</p>
					<br />
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
					<br />
					<br />
					<h5>Build Logic</h5>
					<InputSensorCard
						sensorList={this.state.sensorList}
						handleInputSensorCardChange={this.handleSensorCardChange}
					/>
					{/* <InputGroupCard
						nodeList={this.state.nodeList}
						handleInputGroupCardChange={this.handleGroupCardChange}
					/> */}
					<InputTimeCard
						handleInputTimeCardChange={this.handleTimeCardChange}
					/>
					{this.state.selected_value.map((d: any, idx: number) => (
						<InputValueCard
							valueList={this.state.sensor_info.sensor_values}
							handleRemoveInputValueCardClick={this.handleRemoveValueCardClick(
								idx
							)}
							handleInputValueCardChange={this.handleValueCardChange(idx)}
							index={idx}
						/>
					))}
					{this.state.selected_action.map((d: any, idx: number) => (
						<InputActionCard
							handleInputActionCardChange={this.handleActionCardChange(idx)}
							handleRemoveInputActionCardClick={this.handleRemoveActionCardClick(
								idx
							)}
							index={idx}
						/>
					))}
					<button
						type="button"
						className="btn margin-right"
						style={{ background: 'pink' }}
						onClick={this.handleAddValueCardClick}
					>
						Add value
					</button>
					<button
						type="button"
						className="btn"
						style={{ background: 'pink' }}
						onClick={this.handleAddActionCardClick}
					>
						Add action
					</button>
					<p></p>
					<Link to="/logicCore">
						<button type="button" className="btn btn-default float-right">
							Cancel
						</button>
					</Link>
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

export default RegisterLogic;
