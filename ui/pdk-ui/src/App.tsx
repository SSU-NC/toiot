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
import { SENSOR_URL, NODE_URL, SINK_URL, LOGICCORE_URL } from './defineUrl';
import { logicCoreElem } from './ElemInterface/LcElementsInterface';
import SinkManagement from './components/SinkManagement';
import AlertAlarm from './components/AlertAlarm';

interface AppState {
	sensorList: Array<sensorListElem>;
	nodeList: Array<nodeListElem>;
	logicCore: Array<logicCoreElem>;
	sinkList: Array<sinkListElem>;
}

class App extends Component<{}, AppState> {
	state: AppState = {
		sensorList: [],
		nodeList: [],
		logicCore: [],
		sinkList: [],
	};

	componentDidMount() {
		this.getsensorList();
		this.getnodeList();
		this.getlogicCore();
		this.getsinkList();
	}

	getsensorList() {
		var url = SENSOR_URL;

		fetch(url)
			.then((res) => res.json())
			.then((data) => {
				this.setState({ sensorList: data });
			})
			// .then(response => console.log('Success:', JSON.stringify(response)))
			.catch((error) => console.error('Error:', error));
	}

	getnodeList() {
		var url = NODE_URL;

		fetch(url)
			.then((res) => res.json())
			.then((data) => this.setState({ nodeList: data }))
			// .then(response => console.log('Success:', JSON.stringify(response)))
			.catch((error) => console.error('Error:', error));
	}

	getsinkList() {
		var url = SINK_URL;

		fetch(url)
			.then((res) => res.json())
			.then((data) => this.setState({ sinkList: data }))
			// .then(response => console.log('Success:', JSON.stringify(response)))
			.catch((error) => console.error('Error:', error));
	}

	getlogicCore() {
		//this.setState({ logicCore: logicTable_ex });

		var url = LOGICCORE_URL;

		fetch(url)
			.then((res) => res.json())
			.then((data) => this.setState({ logicCore: data }))
			// .then(response => console.log('Success:', JSON.stringify(response)))
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
							<Route
								path="/node"
								render={() => (
									<NodeManagement
										sensorList={this.state.sensorList}
										sinkList={this.state.sinkList}
										nodeList={this.state.nodeList}
									/>
								)}
							/>
							<Route
								path="/sink"
								render={() => <SinkManagement sinkList={this.state.sinkList} />}
							/>
							<Route
								path="/logicCore"
								render={() => (
									<LogicCoreManagement logicCore={this.state.logicCore} />
								)}
							/>
							<Route
								path="/registerLogic"
								render={() => (
									<RegisterLogic
										sensorList={this.state.sensorList}
										nodeList={this.state.nodeList}
									/>
								)}
							></Route>

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
