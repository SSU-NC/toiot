import React, { Component } from 'react';


class RegisterAlarm extends Component {

    render() {
        return (
            <div>
                <form>
                    <div class="form-group">
						<label for="alarm_name">Alarm name</label>
						<input
							type="text"
							class="form-control"
							name="alarm_name"
							placeholder="name"
							// value={this.state.node_name}
							// onChange={this.handleNameChange}
						/>
					</div>
                    <div class="form-group">
						<label for="alarm_msg">Alarm message</label>
						<input
							type="text"
							class="form-control"
							name="alarm_msg"
							placeholder="Enter alarm msg which you want to get alert"
							// value={this.state.node_name}
							// onChange={this.handleNameChange}
						/>
					</div>
                    <div class="form-group">
                        <label for="inputEmail">Email address</label>
                        <input type="email" class="form-control" id="email" aria-describedby="emailHelp" placeholder="Enter email"/>
                        <small id="emailHelp" class="form-text text-muted">We'll send message to this e-mail.</small>
                    </div>
                    <button
                        type="submit"
                        class="btn btn-primary"
                        data-dismiss="modal"
                    >
						Submit
					</button>

                </form>
            </div>
        );
    }
}

export default RegisterAlarm;