import React, { Component } from 'react';
import Select from 'react-select';
import { sensorListElem, sensorOptionsElem } from './ElementsInterface';
import { NODE_URL } from '../defineUrl';

// react-select : https://github.com/JedWatson/react-select

interface RegisterNodeState {
	node_name: string;
	location: string;
	sensors: Array<sensorOptionsElem>;
}

interface RegisterNodeProps {
	sensorList: Array<sensorListElem>;
}

class RegisterNode extends Component<RegisterNodeProps, RegisterNodeState> {
	state: RegisterNodeState = {
		node_name: '',
		location: '',
		sensors: [],
	};
	handleNameChange = (e: React.ChangeEvent<HTMLInputElement>) => {
		this.setState({
			node_name: e.target.value,
		});
	};
	handleLocationChange = (e: React.ChangeEvent<HTMLInputElement>) => {
		this.setState({
			location: e.target.value,
		});
	};
	handleSensorsChange = (sensors: any) => {
		//sensors: Array<sensorOptionsElem> ?? ?? ??..
		this.setState({
			sensors,
		});
	};
	handleSubmit = (e: React.MouseEvent<HTMLButtonElement>) => {
		e.preventDefault();

		var url = NODE_URL;
		var data = this.state;
		var sensor_uuid = data.sensors.map((val: sensorOptionsElem) => {
			return { uuid: val.uuid };
		});

		console.log(
			JSON.stringify({
				name: data.node_name,
				location: data.location,
				sensors: sensor_uuid,
			})
		);

		fetch(url, {
			method: 'POST', // or 'PUT'
			body: JSON.stringify({
				name: data.node_name,
				location: data.location,
				sensors: sensor_uuid,
			}),
			headers: {
				'Content-Type': 'application/json',
			},
		})
			.then((res) => res.json())
			.then((response) => console.log('Success:', JSON.stringify(response)))
			.catch((error) => console.error('Error:', error));
	};

	render() {
		let sensorOptions: Array<sensorOptionsElem>;
		sensorOptions = this.props.sensorList.map((val: sensorListElem) => {
			return { label: val.name, value: val.name, uuid: val.uuid };
		});

		return (
			<>
				<button
					type="button"
					className="btn btn-primary"
					data-toggle="modal"
					data-target="#register-node"
				>
					register node
				</button>
				<div
					className="modal fade"
					id="register-node"
					//tabindex="-1"
					role="dialog"
					aria-labelledby="register-node"
				>
					<div className="modal-dialog" role="document">
						<div className="modal-content">
							<div className="modal-header">
								<h4 className="modal-title" id="register-node">
									Register node
								</h4>
								<button
									type="button"
									className="close"
									data-dismiss="modal"
									aria-label="Close"
								>
									<span aria-hidden="true">×</span>
								</button>
							</div>
							<div className="modal-body">
								<form>
									<div className="form-group">
										<label>Node name</label>
										<input
											type="text"
											className="form-control"
											name="node_name"
											placeholder="name"
											value={this.state.node_name}
											onChange={this.handleNameChange}
										/>
									</div>
									<div className="form-group">
										<label>Location</label>
										<input
											type="text"
											className="form-control"
											name="location"
											placeholder="location"
											value={this.state.location}
											onChange={this.handleLocationChange}
										/>
									</div>
									<div className="form-group">
										<label>Select sensors</label>
										<Select
											isMulti
											className="form-control basic-multi-select"
											name="sensors"
											options={sensorOptions}
											classNamePrefix="select"
											value={this.state.sensors}
											onChange={this.handleSensorsChange}
										/>
									</div>
									<div className="modal-footer">
										<button
											type="submit"
											className="btn btn-primary"
											data-dismiss="modal"
											onClick={this.handleSubmit}
										>
											Submit
										</button>
										<button
											type="reset"
											className="btn btn-default"
											data-dismiss="modal"
										>
											Cancel
										</button>
									</div>
								</form>
							</div>
						</div>
					</div>
				</div>
			</>
		);
	}
}

export default RegisterNode;
