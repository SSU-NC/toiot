import React, { Component } from 'react';
import Select from 'react-select';
import { lcAction } from '../LcElementsInterface';
import '../LogicCore.css';

interface ActionCardProps { 
	handleActionCardChange: (value: lcAction) => void;
}
interface ActionCardState { 
	logic: string;
	text: string;
}
interface actionOptionsElem {
	label: string;
	value: string;
}

class ActionCard extends Component< ActionCardProps, ActionCardState > {
	state: ActionCardState = {
		logic: '',
		text: ''
	}
	handleActionChange = async(e: any) => {
		await this.setState({
			logic: e.value,
		})
		this.props.handleActionCardChange(this.state);
	}

	handleTextChange = async(e: React.ChangeEvent<HTMLInputElement>) => {
		await this.setState({
			text: e.target.value,
		})
		this.props.handleActionCardChange(this.state);
	}

    render() {
		let actionOptions : Array<actionOptionsElem> = [
			{label: 'alarm', value: 'alarm'}, {label:'email', value: 'email'}
		]
        return(
            <div className="card form-group">
				<div className="card-body row">
					<div className=" col-2 right-divider">
						<h4 className="align-middle">action</h4>
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
						{(this.state.logic === "alarm") ?
						(
						<div>
							<span>Alarm MSG</span>
							<input
								type="text"
								className="form-control"
								name="alarm_msg"
								value={this.state.text}
								placeholder="Enter alarm msg which you want to get alert"
								onChange={this.handleTextChange}
							/>
						</div>
						) :
						((this.state.logic === "email") ? 
						(
						<div>
							<span>Email address</span>
							<input
								type="email"
								className="form-control"
								id="email"
								value={this.state.text}
								aria-describedby="emailHelp"
								placeholder="toiot@example.com"
								onChange={this.handleTextChange}
							/>
							<small id="emailHelp" className="form-text text-muted">
								We'll send message to this e-mail.
							</small>
						</div>
						):(
						<div></div>
						))}
					</div>
				</div>
			</div>
        )
    }
}

export default ActionCard;