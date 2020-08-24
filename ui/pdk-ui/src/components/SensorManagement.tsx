import React, { Component } from 'react';
import RegisterSensor from './Register/RegisterSensor';
import SensorTable from './Table/SensorTable';

interface SensorManagementProps {
	sensorList: any;
}

class SensorManagement extends Component<SensorManagementProps> {
	render() {
		return (
			<>
				<div style={{ float: 'right' }}>
					<RegisterSensor></RegisterSensor>
				</div>
				<div>
					<h3>Sensor</h3>
					<SensorTable sensorList={this.props.sensorList}></SensorTable>
				</div>
			</>
		);
	}
}

export default SensorManagement;
