import React, { Component } from 'react';
import Select from 'react-select';
import { SINK_URL, TOPIC_URL } from '../../defineUrl';
import { topicListElem, topicOptionsElem } from '../../ElemInterface/ElementsInterface';
// form : https://getbootstrap.com/docs/4.0/components/forms/?
// add, delete input : https://codesandbox.io/s/00xq32n3pn?from-embed=&file=/src/index.js

interface RegisterSinkState {
	topicList: Array<topicListElem>;
	name: string;
	topic_id: number;
	ip: string;
	nameValid: boolean;
	topicValid: boolean;
	ipValid: boolean;
}

/* 
RegisterSink
- Show modal to register sink
*/
class RegisterSink extends Component<{}, RegisterSinkState> {
	state: RegisterSinkState = {
		topicList: [],
		name: '',
		topic_id: 0,
		ip: '',
		nameValid: false,
		topicValid: false,
		ipValid: false,
	};

	componentDidMount() {
		this.gettopicList();
	}

	// Get topic list from backend
	gettopicList() {
		var url = TOPIC_URL;

		fetch(url)
			.then((res) => res.json())
			.then((data) => this.setState({ topicList: data }))
			.catch((error) => console.error('Error:', error));
	}

	// Handle node name change by typing
	handleNameChange = (e: React.ChangeEvent<HTMLInputElement>) => {
		// name valid check : user should enter sink name
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

	// Handle selected sensor change by selecting sensors
	handleTopicChange = (topic: any) => {
		// sensor valid check : user should select sensor
		if (topic !== null) {
			this.setState({
				topic_id: topic.id,
				topicValid: true,
			});
		} else {
			this.setState({
				topic_id: topic.id,
				topicValid: false,
			});
		}
	};


	// Handle ip:port change by typing
	handleIpChange = (e: React.ChangeEvent<HTMLInputElement>) => {
		// Regular expression of check the ip:port format
		const ipportRegExp = /^[0-9]{1,3}.[0-9]{1,3}.[0-9]{1,3}.[0-9]{1,3}:[0-9]{4,5}/;

		// ip:port valid check : user should enter correct format of ip:port.
		if (e.target.value.match(ipportRegExp)) {
			this.setState({
				ip: e.target.value,
				ipValid: true,
			});
		} else {
			this.setState({
				ip: e.target.value,
				ipValid: false,
			});
		}
	};

	// Handle submit button click event
	handleSubmit = (e: React.MouseEvent<HTMLButtonElement>) => {
		e.preventDefault();

		var url = SINK_URL;
		var data = this.state;

		// Valid check (unvalid -> alert)
		if (!this.state.nameValid) {
			alert('Please enter sink.');
			return;
		}
		if (!this.state.topicValid) {
			alert('Please select topic.');
			return;
		}
		if (!this.state.ipValid) {
			alert('Please enter valid type of ip:port.');
			return;
		}

		// Check whether user really want to submit
		var submitValid: boolean;
		submitValid = window.confirm('Are you sure to register this sink?');
		if (!submitValid) {
			return;
		}

		fetch(url, {
			method: 'POST', // or 'PUT'
			body: JSON.stringify({
				name: this.state.name,
				addr: this.state.ip,
				topic_id: this.state.topic_id,
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
		let topicOptions: Array<topicOptionsElem>;
		topicOptions = this.state.topicList.map((val: topicListElem) => {
			return {
				label: val.name,
				value: val.name,
				id: val.id,
			};
		});
		return (
			<>
				<div
					className="modal fade"
					id="register-sink-modal"
					role="dialog"
					aria-labelledby="register-sink-modal"
				>
					<div className="modal-dialog" role="document">
						<div className="modal-content">
							<div className="modal-header">
								<h4 className="modal-title" id="register-sink-modal">
									Register sink
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
										<label>Sink name</label>
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
										<label>Select topic</label>
										<Select
											name="topic"
											options={topicOptions}
											classNamePrefix="select"
											onChange={this.handleTopicChange}
										/>
									</div>
									<div className="form-group">
										<label>ip:port</label>
										<input
											type="text"
											className="form-control"
											name="ip"
											placeholder="ex) 123.123.123:8080"
											value={this.state.ip}
											onChange={this.handleIpChange}
										/>
									</div>
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

export default RegisterSink;
