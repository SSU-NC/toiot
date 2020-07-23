import React, { Component } from 'react';
import Select from 'react-select';
import { sensorListElem } from './ElementsInterface';

interface RegisterAlarmState {
	sensor: any;
}

interface RegisterAlarmProps {
	sensorList: Array<sensorListElem>;
}

interface sensorOptionsElem {
	label: string;
	value: string;
	uuid: string;
}

class RegisterAlarm extends Component<RegisterAlarmProps, RegisterAlarmState> {

	state: RegisterAlarmState = {
		sensor: {},
	};

	render() {
		let sensorOptions = this.props.sensorList.map((val: sensorListElem) => {
			return { label: val.name, value: val.name, uuid: val.uuid };
		});

		return (
			<div>
				<h4>
					Register Alarm Service
				</h4>
				<form>
					<div className="form-group">
						<label>Select sensors</label>
						<Select
							className="form-control"
							name="sensors"
							options={sensorOptions}
							classNamePrefix="select"
							value={this.state.sensor}
							//onChange={this.handleSensorsChange}
						/>
					</div>
					<div className="form-group">
						<label>Alarm name</label>
						<input
							type="text"
							className="form-control"
							name="alarm_name"
							placeholder="name"
							// value={this.state.alarm_name}
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
							// value={this.state.alarm_msg}
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
							placeholder="iotoi@example.com"
						/>
						<small id="emailHelp" className="form-text text-muted">
							We'll send message to this e-mail.
						</small>
					</div>
					<div className="form-group row">
						<div className="col-5">
							<label className="col-form-label">Start time</label>
							<div>
								<input className="form-control" type="time" placeholder="00:00:00" id="example-time-input"/>
							</div>
						</div>
						<p>
							<br/>~
						</p>
						<div className="col-5">
							<label className="col-form-label">End time</label>
							<div>
								<input className="form-control" type="time" placeholder="23:59:59" id="example-time-input"/>
							</div>
						</div>
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
