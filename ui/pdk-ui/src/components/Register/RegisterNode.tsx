import React, { Component } from 'react';
import Select from 'react-select';
import {
	sensorListElem,
	sensorOptionsElem,
	sinkListElem,
	locationElem,
	sinkOptionsElem,
} from '../../ElemInterface/ElementsInterface';
import { NODE_URL } from '../../defineUrl';
// react-select : https://github.com/JedWatson/react-select

interface RegisterNodeState {
	node_name: string;
	group: string;
	location: locationElem;
	sink_id: number;
	sensors: Array<sensorOptionsElem>;
	nameValid: boolean;
	groupValid: boolean;
	sensorValid: boolean;
	sinkValid: boolean;
}

interface RegisterNodeProps {
	sensorList: Array<sensorListElem>;
	sinkList: Array<sinkListElem>;
}

class RegisterNode extends Component<RegisterNodeProps, RegisterNodeState> {
	state: RegisterNodeState = {
		node_name: '',
		group: '',
		location: {
			lon: 0,
			lat: 0,
		},
		sink_id: 0,
		sensors: [],

		nameValid: false,
		groupValid: false,
		sensorValid: false,
		sinkValid: false,
	};
	handleNameChange = (e: React.ChangeEvent<HTMLInputElement>) => {
		if (e.target.value.length > 0) {
			this.setState({
				node_name: e.target.value,
				nameValid: true,
			});
		} else {
			this.setState({
				node_name: e.target.value,
				nameValid: false,
			});
		}
	};
	handleGroupChange = (e: React.ChangeEvent<HTMLInputElement>) => {
		if (e.target.value.length > 0) {
			this.setState({
				group: e.target.value,
				groupValid: true,
			});
		} else {
			this.setState({
				group: e.target.value,
				groupValid: false,
			});
		}
	};
	handleLatChange = (e: React.ChangeEvent<HTMLInputElement>) => {
		this.setState({
			location: { ...this.state.location, lat: parseFloat(e.target.value) },
		});
	};
	handleLonChange = (e: React.ChangeEvent<HTMLInputElement>) => {
		this.setState({
			location: { ...this.state.location, lon: parseFloat(e.target.value) },
		});
	};
	handleSensorsChange = (sensors: any) => {
		//sensors: Array<sensorOptionsElem> ?? ?? ??..
		if (sensors !== null || sensors !== []) {
			this.setState({
				sensors,
				sensorValid: true,
			});
		} else {
			this.setState({
				sensors,
				sensorValid: false,
			});
		}
	};
	handleSinkChange = (sink: any) => {
		//sensors: Array<sensorOptionsElem> ?? ?? ??..
		if (sink !== null) {
			this.setState({
				sink_id: sink.id,
				sinkValid: true,
			});
		} else {
			this.setState({
				sink_id: sink.id,
				sinkValid: false,
			});
		}
	};
	handleSubmit = (e: React.MouseEvent<HTMLButtonElement>) => {
		e.preventDefault();

		var url = NODE_URL;
		var data = this.state;
		var sensor_uuid = data.sensors.map((val: sensorOptionsElem) => {
			return { uuid: val.uuid };
		});

		if (!this.state.nameValid) {
			alert('Please enter node name.');
			return;
		}
		if (!this.state.groupValid) {
			alert('Please enter group.');
			return;
		}
		if (!this.state.sensorValid) {
			alert('Please select more than a sensor.');
			return;
		}
		if (!this.state.sinkValid) {
			alert('Please enter sink.');
			return;
		}

		var submitValid: boolean;
		submitValid = window.confirm('Are you sure to register this node?');
		if (!submitValid) {
			return;
		}

		console.log(
			JSON.stringify({
				name: data.node_name,
				group: data.group,
				location: {
					lat: data.location.lat,
					lon: data.location.lon,
				},
				sensors: sensor_uuid,
			})
		);

		fetch(url, {
			method: 'POST', // or 'PUT'
			body: JSON.stringify({
				name: data.node_name,
				group: data.group,
				location: {
					lat: data.location.lat,
					lon: data.location.lon,
				},
				sink_id: data.sink_id,
				sensors: sensor_uuid,
			}),
			headers: {
				'Content-Type': 'application/json',
			},
		})
			.then((res) => res.json())
			.then((response) => console.log('Success:', JSON.stringify(response)))
			.catch((error) => console.error('Error:', error))
			.then(() => window.location.reload(false));
	};

	render() {
		let sensorOptions: Array<sensorOptionsElem>;
		sensorOptions = this.props.sensorList.map((val: sensorListElem) => {
			return {
				label: val.name,
				value: val.name,
				uuid: val.uuid,
				value_list: val.value_list,
			};
		});
		let sinkOptions: Array<sinkOptionsElem>;
		sinkOptions = this.props.sinkList.map((val: sinkListElem) => {
			return { label: val.name, value: val.name, id: val.id };
		});
		return (
			<>
				<button
					type="button"
					className="btn"
					data-toggle="modal"
					data-target="#register-node"
					style={{ background: 'pink' }}
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
										<label>group</label>
										<input
											type="text"
											className="form-control"
											name="group"
											placeholder="group"
											value={this.state.group}
											onChange={this.handleGroupChange}
										/>
									</div>
									<div className="form-group">
										<label>location - latitude</label>
										<input
											type="number"
											className="form-control col-3 margin-right"
											name="lat"
											value={this.state.location.lat}
											onChange={this.handleLatChange}
										/>
									</div>
									<div className="form-group">
										<label>location - longitude</label>
										<input
											type="number"
											className="form-control col-3"
											name="lon"
											value={this.state.location.lon}
											onChange={this.handleLonChange}
										/>
									</div>
									<div className="form-group">
										<label>Select sensors</label>
										<Select
											isMulti
											className="basic-multi-select"
											name="sensors"
											options={sensorOptions}
											classNamePrefix="select"
											value={this.state.sensors}
											onChange={this.handleSensorsChange}
										/>
									</div>
									<div className="form-group">
										<label>Select sink</label>
										<Select
											className="basic-select"
											name="sink"
											options={sinkOptions}
											classNamePrefix="select"
											onChange={this.handleSinkChange}
										/>
									</div>
									<div className="modal-footer">
										<button
											type="submit"
											className="btn"
											onClick={this.handleSubmit}
											style={{ background: 'pink' }}
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
