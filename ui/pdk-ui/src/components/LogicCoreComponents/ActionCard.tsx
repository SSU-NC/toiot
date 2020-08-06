import React, { Component } from 'react';
import Select from 'react-select';
import '../LogicCore.css';

interface ActionCardProps{ 
}

interface actionOptionsElem {
	label: string;
	value: string;
}

class ActionCard extends Component< ActionCardProps, {} > {
    render() {
		let actionOptions : Array<actionOptionsElem> = [
			{label: 'alarm', value: 'alarm'},{label:'email', value: 'email'}
		]
			
		
        return(
            <div className="card form-group">
				<div className="card-body row">
					<div className=" col-2 right-divider">
						<h4 className="align-middle">action</h4>
					</div>
					<div className="col-3">
						<Select options={actionOptions} name="sensors" classNamePrefix="select" />
					</div>
					<div className="col-1"></div>
					<div className="col-5">
						<div>
							<span>Alarm MSG</span>
							<input
								type="text"
								className="form-control"
								name="alarm_msg"
								placeholder="Enter alarm msg which you want to get alert"
							/>
						</div>

						<div>
							<span>Email address</span>
							<input
								type="email"
								className="form-control"
								id="email"
								aria-describedby="emailHelp"
								placeholder="toiot@example.com"
							/>
							<small id="emailHelp" className="form-text text-muted">
								We'll send message to this e-mail.
							</small>
						</div>
					</div>
				</div>
			</div>
        )
    }
}

export default ActionCard;