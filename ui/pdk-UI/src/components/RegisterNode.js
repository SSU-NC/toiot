import React, { Component } from 'react';
import Modal from 'react-modal';
import Select from 'react-select';

// react-select : https://github.com/JedWatson/react-select

class RegisterNode extends Component {
	constructor(props) {
		super(props);
		this.state = {
			node_name: '',
			location: '',
			sensors: [],
		};

		this.handleNameChange = this.handleNameChange.bind(this);
		this.handleLocationChange = this.handleLocationChange.bind(this);
		this.handleSensorsChange = this.handleSensorsChange.bind(this);
		this.handleSubmit = this.handleSubmit.bind(this);
	}
	// openModal(){
	//     setShow(true);
	// }
	// closeModal(){
	//     setShow(false);
	// }
	handleNameChange(e) {
		this.setState({
			[e.target.name]: e.target.value,
		});
	}
	handleLocationChange(e) {
		this.setState({
			[e.target.name]: e.target.value,
		});
	}
	handleSensorsChange = (sensors) => {
		this.setState({ sensors });
	};
	handleSubmit(e) {
		e.preventDefault();

		var url = 'http://220.70.2.160:8080/node';
		var data = this.state;
		var sensor_uuid = data.sensors.map((val) => {
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
	}

	render() {
		let sensorOptions = this.props.sensorList.map((val) => {
			return { label: val.name, value: val.name, uuid: val.uuid };
		});

		return (
			<>
				{/*    <button type="button" class="btn btn-primary btn-lg" onClick={openModal}>register node</button>
        //     <Modal
        //         isOpen={modalIsOpen}
        //         aria-labelledby="register-node"
        //     >
        //         <Modal.Header closeButton>
        //             <Modal.Title id="register-node">
        //                 Register node
        //             </Modal.Title>
        //         </Modal.Header>
        //         <form >
        //             <Modal.Body>
        //                 <div class="form-group">
        //                     <label for="node_name">Node name</label>
        //                     </div>
        //                 <div class="form-group">
        //                     <label for="location">Location</label>
        //                     <input type="text" class="form-control" name="location" placeholder="location" value={this.state.location} onChange={this.handleLocationChange}/>
        //                 </div>
        //                 <div class="form-group">
        //                     <label for="select_sensor">Select sensors</label>
        //                     <Select
        //                         isMulti
        //                         class="form-control"
        //                         name="sensors"
        //                         options={sensorOptions}
        //                         className="basic-multi-select"
        //                         classNamePrefix="select"
        //                         value={this.state.sensors}
        //                         onChange={this.handleSensorsChange}
        //                     />
        //                 </div>
        //             </Modal.Body>
        //             <Modal.Footer>
        //                 <button type="submit" class="btn btn-primary" data-dismiss="modal" onClick={this.handleSubmit}>Submit</button>
        //                 <button type="reset" class="btn btn-default" data-dismiss="modal">Cancel</button>
        //             </Modal.Footer>
        //         </form>
        //     </Modal> */}

				<button
					type="button"
					class="btn btn-primary"
					data-toggle="modal"
					data-target="#register-node"
				>
					register node
				</button>
				<div
					class="modal fade"
					id="register-node"
					tabindex="-1"
					role="dialog"
					aria-labelledby="register-node"
				>
					<div class="modal-dialog" role="document">
						<div class="modal-content">
							<div class="modal-header">
								<h4 class="modal-title" id="register-node">
									Register node
								</h4>
								<button
									type="button"
									class="close"
									data-dismiss="modal"
									aria-label="Close"
								>
									<span aria-hidden="true">×</span>
								</button>
							</div>
							<div class="modal-body">
								<form>
									<div class="form-group">
										<label for="node_name">Node name</label>
										<input
											type="text"
											class="form-control"
											name="node_name"
											placeholder="name"
											value={this.state.node_name}
											onChange={this.handleNameChange}
										/>
									</div>
									<div class="form-group">
										<label for="location">Location</label>
										<input
											type="text"
											class="form-control"
											name="location"
											placeholder="location"
											value={this.state.location}
											onChange={this.handleLocationChange}
										/>
									</div>
									<div class="form-group">
										<label for="select_sensor">Select sensors</label>
										<Select
											isMulti
											class="form-control"
											name="sensors"
											options={sensorOptions}
											className="basic-multi-select"
											classNamePrefix="select"
											value={this.state.sensors}
											onChange={this.handleSensorsChange}
										/>
									</div>
									<div class="modal-footer">
										<button
											type="submit"
											class="btn btn-primary"
											data-dismiss="modal"
											onClick={this.handleSubmit}
										>
											Submit
										</button>
										<button
											type="reset"
											class="btn btn-default"
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
