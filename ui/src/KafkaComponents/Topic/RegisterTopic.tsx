import React, { Component } from 'react';
import Select from 'react-select';
import { TOPIC_URL } from '../../defineUrl';
import { topicListElem, topicOptionsElem } from '../../ElemInterface/ElementsInterface';
// form : https://getbootstrap.com/docs/4.0/components/forms/?
// add, delete input : https://codesandbox.io/s/00xq32n3pn?from-embed=&file=/src/index.js

interface RegisterTopicState {
	name: string;
	partitions: number;
	replications: number;
	nameValid: boolean;
	partitionsValid: boolean;
	replicationsValid: boolean;
}

/* 
RegisterTopic
- Show modal to register topic
*/
class RegisterTopic extends Component<{}, RegisterTopicState> {
	state: RegisterTopicState = {
		name: '',
        partitions: 0,
        replications: 0,
        nameValid: false,
		partitionsValid: false,
		replicationsValid: false,
		
	};

	// Handle node name change by typing
	handleNameChange = (e: React.ChangeEvent<HTMLInputElement>) => {
		// name valid check : user should enter topic name
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

    handlePartitionsChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        // if (e.target.value > 0) {
        //     this.setState({
        //         name: e.target.value,
        //     });
        // } else {
        //     this.setState({
        //         name: e.target.value,
        //     });
        // }
        this.setState({
            partitions: parseInt(e.target.value),
        })
    }
    handleReplicationsChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        this.setState({
            replications: parseInt(e.target.value),
        })   
    }

	// Handle submit button click event
	handleSubmit = (e: React.MouseEvent<HTMLButtonElement>) => {
		e.preventDefault();

		var url = TOPIC_URL;
		var data = {
            name: this.state.name,
            partitions: this.state.partitions,
            replications: this.state.replications,
        };

		// Valid check (unvalid -> alert)
		if (!this.state.nameValid) {
			alert('Please enter topic.');
			return;
		}

		// Check whether user really want to submit
		var submitValid: boolean;
		submitValid = window.confirm('Are you sure to register this topic?');
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
					id="register-topic-modal"
					role="dialog"
					aria-labelledby="register-topic-modal"
				>
					<div className="modal-dialog" role="document">
						<div className="modal-content">
							<div className="modal-header">
								<h4 className="modal-title" id="register-topic-modal">
									Register topic
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
										<label>Topic name</label>
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
                                        <label>Partitions</label>
                                        <input
                                            type="number"
                                            className="form-control"
                                            id="partitions"
                                            value={this.state.partitions}
                                            onChange={this.handlePartitionsChange}
                                        />
                                    </div>
                                    <div className="form-group">
                                        <label>Replications</label>
                                        <input
                                            type="number"
                                            className="form-control"
                                            id="replications"
                                            value={this.state.replications}
                                            onChange={this.handleReplicationsChange}
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

export default RegisterTopic;
