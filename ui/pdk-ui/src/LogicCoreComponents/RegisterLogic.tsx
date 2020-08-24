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
import InputGroupCard from './InputCards/InputGroupCard';
import InputTimeCard from './InputCards/InputTimeCard';
import InputActionCard from './InputCards/InputActionCard';
import { LOGICCORE_URL } from '../defineUrl';

interface RegisterLogicProps {
	sensorList: Array<sensorListElem>;
	nodeList: Array<nodeListElem>;
}

interface RegisterLogicState {
	logic_name: string;
	sensor_info: sensorOptionsElem;
	selected_value: Array<logicElem>;
	selected_time: logicElem;
	selected_action: Array<logicElem>;
	selected_group: logicElem;
}

class RegisterLogic extends Component<RegisterLogicProps, RegisterLogicState> {
	state: RegisterLogicState = {
		logic_name: '',
		sensor_info: {
			uuid: '',
			value: '',
			label: '',
			value_list: [],
		},
		selected_value: [],
		selected_group: { elem: 'empty', arg: { group: [] } },
		selected_time: { elem: 'empty', arg: { range: [] } },
		selected_action: [],
	};
	handleLogicNameChange = (e: React.ChangeEvent<HTMLInputElement>) => {
		this.setState({
			logic_name: e.target.value,
		});
	};

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
		const new_selected_value = this.state.selected_value.map(
			(value: logicElem, sidx: number) => {
				if (idx !== sidx) return value;
				return selectedValue;
			}
		);
		this.setState({ selected_value: new_selected_value });
	};
	handleTimeCardChange = (selected_time: logicElem) => {
		this.setState({
			selected_time,
		});
	};
	handleActionCardChange = (idx: number) => (selectedAction: logicElem) => {
		const new_selected_action = this.state.selected_action.map(
			(action: logicElem, sidx: number) => {
				if (idx !== sidx) return action;
				return selectedAction;
			}
		);
		this.setState({ selected_action: new_selected_action });
	};

	// handle add card button click event
	handleAddValueCardClick = () => {
		this.setState({
			selected_value: [
				...this.state.selected_value,
				{ elem: 'empty', arg: { value: '', range: [{ min: 0, max: 255 }] } },
			],
		});
	};
	handleAddActionCardClick = () => {
		this.setState({
			selected_action: [
				...this.state.selected_action,
				{ elem: 'empty', arg: { text: '' } },
			],
		});
	};

	// handle remove card button click event
	handleRemoveValueCardClick = (idx: number) => () => {
		this.setState({
			selected_value: this.state.selected_value.filter(
				(s: any, sidx: number) => idx !== sidx
			),
		});
	};
	handleRemoveActionCardClick = (idx: number) => () => {
		this.setState({
			selected_action: this.state.selected_action.filter(
				(s: any, sidx: number) => idx !== sidx
			),
		});
	};

	// selected sensor에 따른 value_list에서, selected_value를 제외한 list 추출
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
		let logic_array: Array<logicElem> = [
			this.state.selected_group,
			this.state.selected_time,
		];
		logic_array = logic_array.concat(
			this.state.selected_value,
			this.state.selected_action
		);

		// Filter elem: 'empty' field
		logic_array = logic_array.filter(function (logic) {
			return logic.elem !== 'empty';
		});

		var request_msg = {
			sensor_uuid: this.state.sensor_info.uuid,
			logic_name: this.state.logic_name,
			logic: logic_array,
		};
		var url = LOGICCORE_URL;

		fetch(url, {
			method: 'POST',
			body: JSON.stringify(request_msg),
			headers: {
				'Content-Type': 'application/json',
			},
		})
			.then((res) => res.json())
			.then((response) => console.log('Success:', JSON.stringify(response)))
			.catch((error) => console.error('Error:', error));
	};

	render() {
		return (
			<div>
				<form>
					<h3>Register Logic</h3>
					<p>A logic is registered at Logic Core</p>
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
						sensorList={this.props.sensorList}
						handleInputSensorCardChange={this.handleSensorCardChange}
					/>
					<InputGroupCard
						nodeList={this.props.nodeList}
						handleInputGroupCardChange={this.handleGroupCardChange}
					/>
					<InputTimeCard
						handleInputTimeCardChange={this.handleTimeCardChange}
					/>
					{this.state.selected_value.map((d: any, idx: number) => (
						<InputValueCard
							valueList={this.state.sensor_info.value_list}
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
