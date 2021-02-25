import React, { Component } from 'react';
import { ACTUATOR_URL } from '../../defineUrl';

interface RegisterActuatorState {
    name: string;
    nameValid: boolean;
}

class RegisterActuator extends Component<{}, RegisterActuatorState> {
    state: RegisterActuatorState = {
        name: '',
        nameValid: false,
    };

    handleNameChange = (e: React.ChangeEvent<HTMLInputElement>) => {
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
        console.log(this.state.name + '@@2!!@!@');
    };

    handleSubmit = (e: React.MouseEvent<HTMLButtonElement>) => {
        e.preventDefault();

        var url = ACTUATOR_URL;
        var data = this.state;

        if (!this.state.nameValid) {
            alert('Please enter actuator name.');
			return;
		}

        var submitValid: boolean;
		submitValid = window.confirm('Are you sure to register this actuator?');
		if (!submitValid) {
			return;
        }
        alert('data : '+data+', url : ' + url);
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
					id="register-actuator-modal"
					role="dialog"
					aria-labelledby="register-actuator-modal"
				>
					<div className="modal-dialog" role="document">
						<div className="modal-content">
							<div className="modal-header">
								<h4 className="modal-title" id="register-actuator-modal">
									Register actuator
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
										<label>Actuator name</label>
										<input
											type="text"
											className="form-control"
											name="name"
											placeholder="name"
											value={this.state.name}
											onChange={this.handleNameChange}
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

export default RegisterActuator;