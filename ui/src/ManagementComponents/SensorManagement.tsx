import React, { Component } from 'react';
import RegisterSensor from './Register/RegisterSensor';
import SensorTable from './Table/SensorTable';

/*
SensorManagement
- Manage sensor table, register sensor
*/
class SensorManagement extends Component {
	render() {
		return (
			<>
				<div style={{ float: 'right' }}>
					<button
						type="button"
						className="btn"
						data-toggle="modal"
						data-target="#register-sensor-modal"
						style={{ background: 'pink' }}
					>
						register sensor
					</button>
					<RegisterSensor></RegisterSensor>
				</div>
				<div>
					<h3>Sensor</h3>
					<SensorTable></SensorTable>
				</div>
			</>
		);
	}
}

export default SensorManagement;
