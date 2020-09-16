import React, { Component } from 'react';
import { SENSOR_URL } from '../../defineUrl';
// form : https://getbootstrap.com/docs/4.0/components/forms/?
// add, delete input : https://codesandbox.io/s/00xq32n3pn?from-embed=&file=/src/index.js

interface RegisterSensorState {
	sensor_values: Array<value_list_elem>;
	name: string;
	nameValid: boolean;
	valueValid: boolean;
}

interface value_list_elem {
	value_name: string;
}

/* 
RegisterSensor
- Show modal to register sensor
*/
class RegisterSensor extends Component<{}, RegisterSensorState> {
	state: RegisterSensorState = {
		sensor_values: [{ value_name: '' }],
		name: '',
		nameValid: false,
		valueValid: false,
	};

	// Handle node name change by typing
	handleNameChange = (e: React.ChangeEvent<HTMLInputElement>) => {
		// name valid check : user should enter sensor name
		if (e.target.value.length > 0) {
			this.setState({
				name: e.target.value,
				nameValid: true,
			});
		} else {
			this.setState({
				name: e.target.value,
				nameValid: false,
			});
		}
	};

	// Handle value's name change by typing
	handleValueChange = (idx: number) => (
		e: React.ChangeEvent<HTMLInputElement>
	) => {
		// Value list is updated dynamic. Its element can be added or removed freely.
		// so find changing field by using received idx and change state.
		const newsensor_values = this.state.sensor_values.map(
			(value: value_list_elem, sidx: number) => {
				if (idx !== sidx) return value;
				return { ...value, value_name: e.target.value };
			}
		);

		// value list valid check : User should enter more than a value and each value input field should be filled
		if (
			newsensor_values !== null &&
			!newsensor_values.some((value) => value.value_name === '') && // find empty field
			newsensor_values[idx].value_name.length > 0
		) {
			this.setState({ sensor_values: newsensor_values, valueValid: true });
		} else {
			this.setState({ sensor_values: newsensor_values, valueValid: false });
		}
	};

	// Handle click event of the Add button
	handleAddClick = () => {
		// Add a value list elem
		this.setState({
			sensor_values: [...this.state.sensor_values, { value_name: '' }],
		});
	};

	// Handle click event of the Remove button
	handleRemoveClick = (idx: number) => () => {
		// Remove #idx value list elem which user picked
		this.setState({
			sensor_values: this.state.sensor_values.filter(
				(s: any, sidx: number) => idx !== sidx
			),
		});
	};

	// Handle submit button click event
	handleSubmit = (e: React.MouseEvent<HTMLButtonElement>) => {
		e.preventDefault();

		var url = SENSOR_URL;
		var data = this.state;

		// Valid check (unvalid -> alert)
		if (!this.state.nameValid) {
			alert('Please enter sensor name.');
			return;
		}
		if (!this.state.valueValid) {
			alert('Please enter value name.');
			return;
		}

		// Check whether user really want to submit
		var submitValid: boolean;
		submitValid = window.confirm('Are you sure to register this sensor?');
		if (!submitValid) {
			return;
		}

		fetch(url, {
			method: 'POST', // or 'PUT'
			body: JSON.stringify(data),
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
		return (
			<>
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
									</div>
									<div className="form-group">
										<label>Value name</label>
										{this.state.sensor_values.map(
											(value: value_list_elem, idx: number) => (
												<div className="input-group mb-3">
													<div className="input-group-prepend">
														<span className="input-group-text">{idx}</span>
													</div>
													<input
														type="text"
														className="form-control"
														name="sensor_values"
														placeholder={'Enter value name'}
														value={value.value_name}
														onChange={this.handleValueChange(idx)}
													/>
													<div className="input-group-append">
														<button
															className="btn btn-sm"
															type="button"
															id="button-addon2"
															onClick={this.handleRemoveClick(idx)}
															style={{ background: 'pink' }}
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
										className="btn"
										onClick={this.handleAddClick}
										style={{ background: 'pink' }}
									>
										Add value
									</button>
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
