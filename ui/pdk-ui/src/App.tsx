import React, { Component } from 'react';
import { BrowserRouter as Router, Route } from 'react-router-dom';
import Nav from './Navigation';
import SensorManagement from './components/SensorManagement';
import NodeManagement from './components/NodeManagement';
import Dashboard from './KibanaDashboard';
import Visualize from './KibanaVisualize';
import Main from './Home';
import LogicCoreManagement from './LogicCoreComponents/LogicCoreManagement';
import RegisterLogic from './LogicCoreComponents/RegisterLogic';
import {
	sensorListElem,
	nodeListElem,
	sinkListElem,
} from './ElemInterface/ElementsInterface';
import { SENSOR_URL, NODE_URL, SINK_URL } from './defineUrl';
import SinkManagement from './components/SinkManagement';
import AlertAlarm from './components/AlertAlarm';

interface AppState {
	sensorList: Array<sensorListElem>;
	nodeList: Array<nodeListElem>;

	sinkList: Array<sinkListElem>;
}

/* 
App
- Routing
- Show navigation bar (Nav)
- Alert alarm service
*/
class App extends Component<{}, AppState> {
	state: AppState = {
		sensorList: [],
		nodeList: [],
		sinkList: [],
	};

	// Get sensor list, node list, logic core, sink list
	componentDidMount() {
		this.getsensorList();
		this.getnodeList();
		this.getsinkList();
	}

	// Get sensor list from backend
	getsensorList() {
		var url = SENSOR_URL;

		fetch(url)
			.then((res) => res.json())
			.then((data) => {
				this.setState({ sensorList: data });
			})
			.catch((error) => console.error('Error:', error));
	}

	// Get node list from backend
	getnodeList() {
		var url = NODE_URL;

		fetch(url)
			.then((res) => res.json())
			.then((data) => this.setState({ nodeList: data }))
			.catch((error) => console.error('Error:', error));
	}

	// Get sink list from backend
	getsinkList() {
		var url = SINK_URL;

		fetch(url)
			.then((res) => res.json())
			.then((data) => this.setState({ sinkList: data }))
			.catch((error) => console.error('Error:', error));
	}

	render() {
		return (
			<div>
				<Router>
					<div>
						<Nav></Nav>
						<AlertAlarm />
						<div className="container pt-4 mt-4">
							<Route exact path="/" render={Main} />
							<Route
								path="/sensor"
								render={() => (
									<SensorManagement sensorList={this.state.sensorList} />
								)}
							/>
							<Route path="/node" component={NodeManagement} />
							<Route path="/sink" component={SinkManagement} />
							<Route path="/logicCore" component={LogicCoreManagement} />
							<Route path="/registerLogic" component={RegisterLogic} />
							<Route path="/visualize" component={Visualize} />
							<Route path="/dashboard" component={Dashboard} />
						</div>
					</div>
				</Router>
			</div>
		);
	}
}

export default App;
