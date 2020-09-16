import React, { Component } from 'react';
import '../LogicCore.css';

interface ShowSesorCardProps {
	sensor_id: number;
}

/*
ShowSensorCard
- show sensor card
*/
class ShowSesorCard extends Component<ShowSesorCardProps, {}> {
	render() {
		return (
			<div className="card margin-bottom">
				<div className="card-body row">
					<div className="col-2 right-divider">
						<span style={{ fontSize: '15pt', fontWeight: 500 }}>sensor</span>
					</div>
					<div className="col-1"></div>
					<div>
						<span style={{ fontSize: '15pt', fontWeight: 450 }}>
							sensor id{' '}
						</span>
						<span style={{ fontSize: '15pt' }}>: {this.props.sensor_id}</span>
					</div>
				</div>
			</div>
		);
	}
}

export default ShowSesorCard;
