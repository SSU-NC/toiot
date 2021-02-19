import { type } from 'jquery';
import React, { Component } from 'react';
import Select from 'react-select';
import { logicElem } from '../../ElemInterface/LcElementsInterface';
import '../LogicCore.css';

interface InputActionCardProps {
	handleInputActionCardChange: (value: logicElem) => void;
	handleRemoveInputActionCardClick: () => void;
	index: number;
}

interface InputActionCardState {
	elem: string;
	arg: {
		text: string;
		elem: string;
		value: number;
		sleep: number;
	};
}
interface actionOptionsElem {
	label: string;
	value: string;
}

/*
InputActionCard
- Get input of action element
*/
class InputActionCard extends Component<
	InputActionCardProps,
	InputActionCardState
> {
	state: InputActionCardState = {
		elem: '',
		arg: { text: '', elem: '', value: 0, sleep: 0 },
	};

	// Handle action change (select alarm or email)
	handleActionChange = async (e: any) => {
		// Change this state and then..
		if (e.value === 'motor' || e.value === 'switch') {
			await this.setState({
				arg: {
					text: this.state.arg.text, 
					elem: e.value,
					value: this.state.arg.value,
					sleep: this.state.arg.sleep
				}
			});
		}
		else {
			await this.setState({
				elem: e.value,
			});
			// change parent's state
		}
		this.props.handleInputActionCardChange(this.state);
	};

	// Handle text change by typing
	handleTextChange = async (e: React.ChangeEvent<HTMLInputElement>) => {
		if (e.target.id === 'alarm_msg' || e.target.id === 'email') {
			await this.setState({
				arg: { 
					text: e.target.value, 
					elem: this.state.arg.elem,
					value: this.state.arg.value,
					sleep: this.state.arg.sleep
				},
			});
		}	
		else if (e.target.id === 'actuator_value') {
			await this.setState({
				arg: { 
					text: this.state.arg.text, 
					elem: this.state.arg.elem,
					value: parseInt(e.target.value),
					sleep: this.state.arg.sleep
				},
			});
		}
		else {
			await this.setState({
				arg: { 
					text: this.state.arg.text, 
					elem: this.state.arg.elem,
					value: this.state.arg.value,
					sleep: parseInt(e.target.value)
				},
			});
		}
		this.props.handleInputActionCardChange(this.state);
	};

	render() {
		let actionOptions: Array<actionOptionsElem> = [
			{ label: 'alarm', value: 'alarm' },
			{ label: 'email', value: 'email' },
			{ label: 'actuator', value: 'actuator'},
		];

		let actuatorOptios: Array<actionOptionsElem> = [
			{ label: 'motor', value: 'motor'},
			{ label: 'switch', value: 'switch'},
		]
		return (
			<div className="card form-group">
				<div className="card-body row">
					<div className=" col-2 right-divider">
						<span style={{ fontSize: '18pt', fontWeight: 500 }}>
							action #{this.props.index}
						</span>
						<button
							className="btn btn-sm float-right"
							type="button"
							id="button-addon2"
							onClick={this.props.handleRemoveInputActionCardClick}
						>
							<svg
								width="1em"
								height="1em"
								viewBox="0 0 16 16"
								className="bi bi-trash-fill"
								fill="currentColor"
								xmlns="http://www.w3.org/2000/svg"
							>
								<path
									fill-rule="evenodd"
									d="M2.5 1a1 1 0 0 0-1 1v1a1 1 0 0 0 1 1H3v9a2 2 0 0 0 2 2h6a2 2 0 0 0 2-2V4h.5a1 1 0 0 0 1-1V2a1 1 0 0 0-1-1H10a1 1 0 0 0-1-1H7a1 1 0 0 0-1 1H2.5zm3 4a.5.5 0 0 1 .5.5v7a.5.5 0 0 1-1 0v-7a.5.5 0 0 1 .5-.5zM8 5a.5.5 0 0 1 .5.5v7a.5.5 0 0 1-1 0v-7A.5.5 0 0 1 8 5zm3 .5a.5.5 0 0 0-1 0v7a.5.5 0 0 0 1 0v-7z"
								/>
							</svg>
						</button>
					</div>
					<div className="col-3">
						<Select
							options={actionOptions}
							name="action"
							classNamePrefix="select"
							onChange={this.handleActionChange}
						/>
					</div>

					<div className="col-1"></div>
					<div className="col-5">
						{this.state.elem === 'alarm' ? ( // If user select alarm
							<div>
								<span>Alarm MSG</span>
								<input
									type="text"
									className="form-control"
									name="alarm_msg"
									value={this.state.arg.text}
									placeholder="Enter alarm msg which you want to get alert"
									onChange={this.handleTextChange}
								/>
							</div>	
						) : this.state.elem === 'email' ? ( // If user select email
							<div>
								<span>Email address</span>
								<input
									type="email"
									className="form-control"
									id="email"
									value={this.state.arg.text}
									aria-describedby="emailHelp"
									placeholder="toiot@example.com"
									onChange={this.handleTextChange}
								/>
								<small id="emailHelp" className="form-text text-muted">
									We'll send message to this e-mail.
								</small>
							</div>
						) : this.state.elem === 'actuator' ? (
							<div>
								<Select
									options={actuatorOptios}
									name="action"
									classNamePrefix="select"
									onChange={this.handleActionChange}          // e.value === motor || e.value === switch ? arg의 elem update 
								/>
								<div className="col-1"></div>
								<div className="row">
									<div className="col-5">
										<span>value</span>
										<input
											type="number"
											className="form-control"
											id="actuator_value"
											value={this.state.arg.value}
											placeholder="Enter "
											onChange={this.handleTextChange}
										/>
										
									</div>
								</div>
								<div className="col-1"></div>
								<div className="row">
									<div className="col-5">
										<span>sleep</span>
										<input
											type="number"
											className="form-control"
											id="actuator_sleep"
											value={this.state.arg.sleep}
											placeholder="Enter "
											onChange={this.handleTextChange}
										/>                                          								
									</div>
								</div>
							</div>
								
						) : (
							<div></div>
						)}
					</div>
				</div>
			</div>
		);
	}
}

export default InputActionCard;
