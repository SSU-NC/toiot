import React, { Component } from 'react';
import Select from 'react-select';
import { logicElem } from '../../ElemInterface/LcElementsInterface';
import '../../LogicCore.css';

interface ShowActionCardProps { 
	handleShowActionCardChange: (value: logicElem) => void;
	handleRemoveShowActionCardClick: () => void;
	index:number;
}

interface ShowActionCardState { 
	elem: string;
	arg : {
		text: string;
	}
}
interface actionOptionsElem {
	label: string;
	value: string;
}

class ShowActionCard extends Component< ShowActionCardProps, ShowActionCardState > {
	state: ShowActionCardState = {
		elem: '',
		arg: { text: ''}
	}
	handleActionChange = async(e: any) => {
		await this.setState({
			elem: e.value,
		})
		this.props.handleShowActionCardChange(this.state);
	}

	handleTextChange = async(e: React.ChangeEvent<HTMLInputElement>) => {
		await this.setState({
			arg: {text: e.target.value},
		})
		this.props.handleShowActionCardChange(this.state);
	}

    render() {
		let actionOptions : Array<actionOptionsElem> = [
			{label: 'alarm', value: 'alarm'}, {label:'email', value: 'email'}
		]
        return(
            <div className="card form-group">
				<div className="card-body row">
					<div className=" col-2 right-divider">
						<span style={{fontSize:'18pt', fontWeight:500}} >action #{this.props.index}</span>
						<button className="btn btn-sm float-right" type="button" id="button-addon2" onClick={this.props.handleRemoveShowActionCardClick}>
							<svg width="1em" height="1em" viewBox="0 0 16 16" className="bi bi-trash-fill" fill="currentColor" xmlns="http://www.w3.org/2000/svg">
								<path fill-rule="evenodd" d="M2.5 1a1 1 0 0 0-1 1v1a1 1 0 0 0 1 1H3v9a2 2 0 0 0 2 2h6a2 2 0 0 0 2-2V4h.5a1 1 0 0 0 1-1V2a1 1 0 0 0-1-1H10a1 1 0 0 0-1-1H7a1 1 0 0 0-1 1H2.5zm3 4a.5.5 0 0 1 .5.5v7a.5.5 0 0 1-1 0v-7a.5.5 0 0 1 .5-.5zM8 5a.5.5 0 0 1 .5.5v7a.5.5 0 0 1-1 0v-7A.5.5 0 0 1 8 5zm3 .5a.5.5 0 0 0-1 0v7a.5.5 0 0 0 1 0v-7z"/>
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
						{(this.state.elem === "alarm") ?
						(
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
						
						) :
						((this.state.elem === "email") ? 
						(
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
						):(
						<div></div>
						))}
					</div>
				</div>
			</div>
        )
    }
}

export default ShowActionCard;