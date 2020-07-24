import React, { Component } from 'react';
import Select from 'react-select';
import { sensorListElem, sensorOptionsElem, value_list_elem } from './ElementsInterface';

interface RegisterAlarmState {
	alarm_name: string;
	alarm_msg: string;
	email: string;
	sensor: sensorOptionsElem;
	value_list: Array<value_list_elem>;
}

interface RegisterAlarmProps {
	sensorList: Array<sensorListElem>;
}

class RegisterAlarm extends Component<RegisterAlarmProps, RegisterAlarmState> {
	state: RegisterAlarmState = {
		alarm_name: '',
		alarm_msg:'',
		email:'',
		sensor: {label: '', value: '', uuid: '' }, // 알람받을 sensor 1개
		value_list: [] // 알람받을 value 값 list
	};

	handleNameChange = (e: React.ChangeEvent<HTMLInputElement>) => {
		this.setState({
			alarm_name: e.target.value,
		});
	};
	handleMsgChange = (e: React.ChangeEvent<HTMLInputElement>) => {
		this.setState({
			alarm_msg: e.target.value,
		});
	};
	handleEmailChange = (e: React.ChangeEvent<HTMLInputElement>) => {
		this.setState({
			email: e.target.value,
		});
	};
	handleSensorChange = (sensor: any) => {
		//sensors: Array<sensorOptionsElem> 쓰면 실행 안됨..
		this.setState({
			sensor,
		});
	};
	handleRemoveClick = (idx: number) => () => {
        this.setState({
            value_list: this.state.value_list.filter((s: any, sidx:number) => idx !== sidx)
        });
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
							onChange={this.handleSensorChange}
						/>
					</div>
					<div className="form-group">
						<label>Alarm name</label>
						<input
							type="text"
							className="form-control"
							name="alarm_name"
							placeholder="name"
							value={this.state.alarm_name}
							onChange={this.handleNameChange}
						/>
					</div>
					<div className="form-group">
						<label>Alarm message</label>
						<input
							type="text"
							className="form-control"
							name="alarm_msg"
							placeholder="Enter alarm msg which you want to get alert"
							value={this.state.alarm_msg}
							onChange={this.handleMsgChange}
						/>
					</div>
					<div className="form-group">
						<label>Email address</label>
						<input
							type="email"
							className="form-control"
							id="email"
							aria-describedby="emailHelp"
							placeholder="toiot@example.com"
							value={this.state.email}
							onChange={this.handleEmailChange}
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
					{/* <div className="form-group">
                        <label>Value name</label>
                        {this.state.value_list.map((value: value_list_elem, idx: number) => (
                            <div className="input-group mb-3">
                                <div className="input-group-prepend">
                                    <span className="input-group-text">{idx}</span>
                                </div>
                                <input type="text" className="form-control" name="value_list" placeholder={"Enter value name"} value={value.value_name} onChange={this.handleValueChange(idx)}/>
                                <div className="input-group-append">
                                    <button className="btn btn-primary btn-sm" type="button" id="button-addon2" onClick={this.handleRemoveClick(idx)}>
                                        <svg width="1em" height="1em" viewBox="0 0 16 16" className="bi bi-trash-fill" fill="currentColor" xmlns="http://www.w3.org/2000/svg">
                                            <path fill-rule="evenodd" d="M2.5 1a1 1 0 0 0-1 1v1a1 1 0 0 0 1 1H3v9a2 2 0 0 0 2 2h6a2 2 0 0 0 2-2V4h.5a1 1 0 0 0 1-1V2a1 1 0 0 0-1-1H10a1 1 0 0 0-1-1H7a1 1 0 0 0-1 1H2.5zm3 4a.5.5 0 0 1 .5.5v7a.5.5 0 0 1-1 0v-7a.5.5 0 0 1 .5-.5zM8 5a.5.5 0 0 1 .5.5v7a.5.5 0 0 1-1 0v-7A.5.5 0 0 1 8 5zm3 .5a.5.5 0 0 0-1 0v7a.5.5 0 0 0 1 0v-7z"/>
                                        </svg>
                                    </button>
                                </div>
                            </div>
                        ))}
                    </div> */}
					<button
						type="submit"
						className="btn"
						data-dismiss="modal"
						style={{background:'pink'}}
					>
						Submit
					</button>
				</form>
			</div>
		);
	}
}

export default RegisterAlarm;
