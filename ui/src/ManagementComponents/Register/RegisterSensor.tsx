import React, { Component } from 'react';
import { SENSOR_URL } from '../../defineUrl';
// form : https://getbootstrap.com/docs/4.0/components/forms/?
// add, delete input : https://codesandbox.io/s/00xq32n3pn?from-embed=&file=/src/index.js

interface RegisterSensorState { 
	sensor_values: Array<value_list_elem>;        // value name list
	name: string;              // 센서 이름
	nameValid: boolean;        // 센서 이름이 제대로 입력되었는지
	valueValid: boolean;       // value name이 제대로 입력되었는지
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
	
	// handleNameChange() 메서드는 Sensor name 폼에 입력된 값을 state 객체에 저장
	// Handle node name change by typing
	handleNameChange = (e: React.ChangeEvent<HTMLInputElement>) => {
		// name valid check : user should enter sensor name
		if (e.target.value.length > 0) {              // e.target.value.length는 입력값
			this.setState({							  
				name: e.target.value,                 // state의 name을 입력값으로 바꾸고 
				nameValid: true,                      // nameValid true로 바꿈 (제대로 이름이 입력되었으므로)
			});
		} else {                                      // 센서 이름 제대로 입력 안됨
			this.setState({
				name: e.target.value,
				nameValid: false,
			});
		}
	};

	// handleValueChange() 메서드는 Value name 폼에 입력된 값을 state 객체에 저장
	// Handle value's name change by typing
	handleValueChange = (idx: number) => (                 // idx는 render()에서 파라미터로 받음 (value name 갯수)
		e: React.ChangeEvent<HTMLInputElement>
	) => {
		// Value list is updated dynamic. Its element can be added or removed freely.
		// so find changing field by using received idx and change state.
		const newsensor_values = this.state.sensor_values.map(
			(value: value_list_elem, sidx: number) => {             
				if (idx !== sidx) return value;                           
				return { ...value, value_name: e.target.value };      // value_list_elem의 마지막에 입력받은 value name 추가 (...value는 해당 value_list_elem의 모든 정보를 전달)
			}
		); 

		// value list valid check : User should enter more than a value and each value input field should be filled
		if (
			newsensor_values !== null &&
			!newsensor_values.some((value) => value.value_name === '') && // find empty field
			newsensor_values[idx].value_name.length > 0
		) { 
			this.setState({ sensor_values: newsensor_values, valueValid: true });        // sensor_values update 해주고 value name 제대로 입력 받았으므로 valueValid true
		} else {
			this.setState({ sensor_values: newsensor_values, valueValid: false }); 
		}
	};

	// handleAddClick() 메서드는 value name 폼을 state 객체에 추가
	// Handle click event of the Add button
	handleAddClick = () => {
		// Add a value list elem
		this.setState({
			sensor_values: [...this.state.sensor_values, { value_name: '' }],    // value_list_elem에 빈 value_name 추가
		});
	};

	// handleReamoveClick() 메서드는 해당 value name을 state 객체에서 삭제
	// Handle click event of the Remove button
	handleRemoveClick = (idx: number) => () => {
		// Remove #idx value list elem which user picked   idx는 user가 선택한 value_name 번호
		this.setState({
			sensor_values: this.state.sensor_values.filter(      
				(s: any, sidx: number) => idx !== sidx            // sidx는 value_list_elem의 index 의미 즉, user가 선택한 value_name 번호인 idx값과 sidx값이 다른 요소들로 새로운 배열 생성해 리턴
			),                                                    // 따라서 sensor_vlaues에는 idx에 해당하는 요소를 삭제한 배열이 저장됨  참고: https://niceman.tistory.com/77
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
