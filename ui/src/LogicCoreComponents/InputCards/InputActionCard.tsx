import { type } from 'jquery';
import React, { Component } from 'react';
import Select from 'react-select';
import { control, logicElem } from '../../ElemInterface/LcElementsInterface';
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
		motion: Array<control>;
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
		arg: { text: '', 
			   elem: '', 
			   motion:[{ value: 0, sleep: 0 }] },
	};

	// Handle action change (select alarm or email)
	handleActionChange = async (e: any) => {
		// Change this state and then..
		if (e.value === 'motor' || e.value === 'switch') {
			await this.setState({
				arg: {
					text: this.state.arg.text, 
					elem: e.value,
					motion: this.state.arg.motion,
				}
			});
			console.log(this.state.arg.elem);
		}
		else {
			await this.setState({
				elem: e.value,
			});
			// change parent's state
			console.log(this.state.elem);
		}
		this.props.handleInputActionCardChange(this.state);
	};

	// Handle text change by typing
	handleTextChange = async (e: React.ChangeEvent<HTMLInputElement>) => {
		await this.setState({
			arg: { 
				text: e.target.value, 
				elem: this.state.arg.elem,
				motion: this.state.arg.motion,
			},
		});
		this.props.handleInputActionCardChange(this.state);
	};

	handleControlChange = (idx: number) => async (e:any) => {
		const new_motion_elem = this.state.arg.motion.map(
			(motionElem: control, sidx: number) => {
				if (idx !== sidx) return motionElem;
				if (e.target.id === 'actuator_value') 
					return { ...motionElem, value: parseInt(e.target.value) };
				return { ...motionElem, sleep: parseInt(e.target.sleep) };
			}
		);
			
		await this.setState({
			arg: { 
				text: this.state.arg.text,
				elem: this.state.arg.elem,
				motion: new_motion_elem	
			}
		});

		this.props.handleInputActionCardChange(this.state);
	};
	
	handleAddClick = async () => {
		await this.setState({
			arg: {
				text: this.state.arg.text, 
				elem: this.state.arg.elem,
				motion: [...this.state.arg.motion, {value: 0, sleep: 0}],
			},
		});
		this.props.handleInputActionCardChange(this.state);
	};

	handleRemoveClick = (idx: number) => async () => {
		await this.setState({
			arg: {
				text: this.state.arg.text,
				elem: this.state.arg.elem,
				motion: this.state.arg.motion.filter(
					(s: any, sidx: number) => idx !== sidx
				),
			},
		});
		this.props.handleInputActionCardChange(this.state)
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
						{/*<div className="col-4">*/}
						{this.state.elem === 'alarm' ? ( // If user select alarm
							<div className="col-5">
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
							<div className="col-5">
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
							<div className="col-3">
								<Select
									options={actuatorOptios}
									name="action"
									classNamePrefix="select"
									onChange={this.handleActionChange}         
								/>
							{/*<div className="col-1"></div>*/}
							<div className="col">
								<button
									type="button"
									className="btn float-right"
									style={{ background: 'pink' }}
									onClick={this.handleAddClick}
								>
									Add scope
								</button>
							</div>
							<div className="col">
								{this.state.arg.motion.map((Control: control, idx: number) => (
									<div className="input-group mb-2">
										<span>value</span>
										<input
											type="number"
											className="form-control"
											id="actuator_value"
											value={Control.value}
											placeholder="Enter "
											onChange={this.handleControlChange(idx)}
										/>

										<div className="col-1"></div>
										<span>sleep</span>
										<input
											type="number"
											className="form-control"
											id="actuator_sleep"
											value={Control.sleep}
											placeholder="Enter "
											onChange={this.handleControlChange(idx)}
										/>

		<								button
											className="btn btn-sm"
											type="button"
											id="button-addon2"
											onClick={this.handleRemoveClick(idx)}
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
								))}
								
							</div>                                          								
							</div>
						) : (
							<div></div>
						)}
					</div>
				</div>
			// </div>
		);
	}
}

export default InputActionCard;
