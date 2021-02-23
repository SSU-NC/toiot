import { type } from 'jquery';
import React, { Component } from 'react';
import Select from 'react-select';
import { control, logicElem } from '../../ElemInterface/LcElementsInterface';
import '../LogicCore.css';

interface InputActionCardProps {
	handleInputActionCardChange: (value: logicElem) => void;
	//handleRemoveInputActionCardClick: () => void;
	//index: number;
}

interface InputActionCardState {
	elem: string;
	arg: {
		aid: number;
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
		elem: 'actuator',
		arg: {
			   aid: 0, 
               motion:[{ value: 0, sleep: 0 }] 
            },
	};

	// Handle action change (select alarm or email)
	handleActionChange = async (e: any) => {
		// Change this state and then..
		if (e.value === 'motor') {
			await this.setState({
				arg: {
					aid: 1,
					motion: this.state.arg.motion,
				}
			});
			console.log(this.state.arg.aid);
		}
		else if (e.value === 'switch') {
			await this.setState({
				arg: {
					aid: 2,
					motion: this.state.arg.motion,
				}
			});
		}
		this.props.handleInputActionCardChange(this.state);
	};

	// Handle text change by typing
	handleTextChange = async (e: React.ChangeEvent<HTMLInputElement>) => {
		await this.setState({
			arg: { 
				aid: this.state.arg.aid,
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
                return { ...motionElem, sleep: parseInt(e.target.value) };
                
			}
		);
		
		await this.setState({
			arg: { 
				aid: this.state.arg.aid,
				motion: new_motion_elem	
			}
		});

		this.props.handleInputActionCardChange(this.state);
	};
	
	handleAddClick = async () => {
		await this.setState({
			arg: {
				aid: this.state.arg.aid,
				motion: [...this.state.arg.motion, {value: 0, sleep: 0}],
			},
		});
		this.props.handleInputActionCardChange(this.state);
	};

	handleRemoveClick = (idx: number) => async () => {
		await this.setState({
			arg: {
				aid: this.state.arg.aid,
				motion: this.state.arg.motion.filter(
					(s: any, sidx: number) => idx !== sidx
				),
			},
		});
		this.props.handleInputActionCardChange(this.state)
	};

	render() {		
		let actuatorOptios: Array<actionOptionsElem> = [
			{ label: 'motor', value: 'motor'},
			{ label: 'switch', value: 'switch'},
		]
		return (
			<div>							
				<div className="col">
					<Select
						options={actuatorOptios}
						name="action"
						classNamePrefix="select"
						onChange={this.handleActionChange}         
					/>
				</div>
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
				<div className="col-5">
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
		);
	}
}

export default InputActionCard;
