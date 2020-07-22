import React, { Component } from 'react';

class RegisterAlarm extends Component {
	render() {
		return (
			<div>
				<form>
					<div className="form-group">
						<label>Alarm name</label>
						<input
							type="text"
							className="form-control"
							name="alarm_name"
							placeholder="name"
							// value={this.state.node_name}
							// onChange={this.handleNameChange}
						/>
					</div>
					<div className="form-group">
						<label>Alarm message</label>
						<input
							type="text"
							className="form-control"
							name="alarm_msg"
							placeholder="Enter alarm msg which you want to get alert"
							// value={this.state.node_name}
							// onChange={this.handleNameChange}
						/>
					</div>
					<div className="form-group">
						<label>Email address</label>
						<input
							type="email"
							className="form-control"
							id="email"
							aria-describedby="emailHelp"
							placeholder="Enter email"
						/>
						<small id="emailHelp" className="form-text text-muted">
							We'll send message to this e-mail.
						</small>
					</div>
					<button
						type="submit"
						className="btn btn-primary"
						data-dismiss="modal"
					>
						Submit
					</button>
				</form>
			</div>
		);
	}
}

export default RegisterAlarm;
