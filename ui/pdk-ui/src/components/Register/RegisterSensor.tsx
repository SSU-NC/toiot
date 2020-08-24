import React, { Component } from 'react';
import { SENSOR_URL } from '../../defineUrl';
// form : https://getbootstrap.com/docs/4.0/components/forms/?
// add, delete input : https://codesandbox.io/s/00xq32n3pn?from-embed=&file=/src/index.js

interface RegisterSensorState {
	value_list: Array<value_list_elem>;
	name: string;
}

interface value_list_elem {
	value_name: string;
}

class RegisterSensor extends Component<{}, RegisterSensorState> {
	state: RegisterSensorState = {
		value_list: [{ value_name: '' }],
		name: '',
	};

	handleNameChange = (e: React.ChangeEvent<HTMLInputElement>) => {
		this.setState({
			name: e.target.value,
		});
	};

	handleValueChange = (idx: number) => (
		e: React.ChangeEvent<HTMLInputElement>
	) => {
		const newvalue_list = this.state.value_list.map(
			(value: value_list_elem, sidx: number) => {
				if (idx !== sidx) return value;
				return { ...value, value_name: e.target.value };
			}
		);
		this.setState({ value_list: newvalue_list });
	};

	// handle click event of the Add button
	handleAddClick = () => {
		this.setState({
			value_list: [...this.state.value_list, { value_name: '' }],
		});
	};

	// handle click event of the Remove button
	handleRemoveClick = (idx: number) => () => {
		this.setState({
			value_list: this.state.value_list.filter(
				(s: any, sidx: number) => idx !== sidx
			),
		});
	};

	handleSubmit = (e: React.MouseEvent<HTMLButtonElement>) => {
		e.preventDefault();

		var url = SENSOR_URL;
		var data = this.state;

		fetch(url, {
			method: 'POST', // or 'PUT'
			body: JSON.stringify(data),
			headers: {
				'Content-Type': 'application/json',
			},
		})
			.then((res) => res.json())
			.then((response) => console.log('Success:', JSON.stringify(response)))
			.catch((error) => console.error('Error:', error));
	};

	render() {
		return (
			<>
				<button
					type="button"
					className="btn btn-primary"
					data-toggle="modal"
					data-target="#register-sensor-modal"
				>
					register sensor
				</button>
				<div
					className="modal fade"
					id="register-sensor-modal"
					role="dialog"
					aria-labelledby="register-sensor-modal"
				>
					<div className="modal-dialog" role="document">
						<div className="modal-content">
							<div className="modal-header">
								<h4 className="modal-title" id="register-sensor-modal">
									Register sensor
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
							<form>
								<div className="modal-body">
									<div className="form-group">
										<label>Sensor name</label>
										<input
											type="text"
											className="form-control"
											name="name"
											placeholder="name"
											value={this.state.name}
											onChange={this.handleNameChange}
										/>
										<div className="invalid-feedback">
											This sensor name is already exist.
										</div>
									</div>
									<div className="form-group">
										<label>Value name</label>
										{this.state.value_list.map(
											(value: value_list_elem, idx: number) => (
												<div className="input-group mb-3">
													<div className="input-group-prepend">
														<span className="input-group-text">{idx}</span>
													</div>
													<input
														type="text"
														className="form-control"
														name="value_list"
														placeholder={'Enter value name'}
														value={value.value_name}
														onChange={this.handleValueChange(idx)}
													/>
													<div className="input-group-append">
														<button
															className="btn btn-primary btn-sm"
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
												</div>
											)
										)}
									</div>
									<button
										type="button"
										className="btn btn-primary"
										onClick={this.handleAddClick}
									>
										Add value
									</button>
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
										type="button"
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
			</>
		);
	}
}

export default RegisterSensor;
