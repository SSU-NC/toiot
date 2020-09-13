import React, { Component } from 'react';
import { alarmElem } from '../ElemInterface/ElementsInterface';
import { Alert } from 'reactstrap';
import { w3cwebsocket as W3CWebSocket } from 'websocket';
import { ALARM_URL } from '../defineUrl';
const client = new W3CWebSocket(ALARM_URL);

interface AlertAlarmState {
	alarmList: Array<alarmListElem>;
}
interface alarmListElem {
	alarm: alarmElem;
	alarm_state: boolean;
}

/*
AlertAlarm
- Alarm service by using websocket
*/
class AlertAlarm extends Component<{}, AlertAlarmState> {
	state: AlertAlarmState = {
		alarmList: [],
	};

	// Open web socket
	componentDidMount() {
		client.onopen = () => {
			console.log('Alarm WebSocket Client Connected');
		};
		client.onmessage = (message: any) => {
			// Parsing message data as json form
			var msg_json: alarmElem = JSON.parse(message.data);

			// 이미 alarm list에 alarm이 들어가있으면 pass
			// If alarm is already in alarm list,
			for (var alarm of this.state.alarmList) {
				if (JSON.stringify(msg_json) === JSON.stringify(alarm.alarm)) {
					return;
				}
			}

			this.setState({
				alarmList: [
					...this.state.alarmList,
					{ alarm: JSON.parse(message.data), alarm_state: true },
				],
			});
		};
	}

	// Handle click button of alert
	handleCloseAlert = (idx: number) => {
		// Set state to filter closed alarm
		this.setState({
			alarmList: this.state.alarmList.filter(
				(s: any, sidx: number) => idx !== sidx
			),
		});
	};

	render() {
		return (
			<>
				{this.state.alarmList.map((alarm: alarmListElem, idx: number) =>
					alarm.alarm_state ? (
						<div className="row">
							<div className="col-3"></div>
							<Alert className="col-6" color="danger">
								<button
									type="button"
									className="close"
									data-dismiss="alert"
									aria-label="Close"
									onClick={() => {
										this.handleCloseAlert(idx);
									}}
								>
									×
								</button>
								<h3 className="alert-heading">
									{alarm.alarm.sensor_name} - {alarm.alarm.sensor_id}
								</h3>
								<hr />
								<p>{alarm.alarm.msg}</p>
							</Alert>
						</div>
					) : (
						<></>
					)
				)}
			</>
		);
	}
}

export default AlertAlarm;
