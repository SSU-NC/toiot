import React, { Component } from 'react';
import AlertAlarm from './components/AlertAlarm';
import Routing from './Routing';

class App extends Component {
	render() {
		return (
			<div>
				<Routing></Routing>
				<AlertAlarm></AlertAlarm>
			</div>
		);
	}
}

export default App;
