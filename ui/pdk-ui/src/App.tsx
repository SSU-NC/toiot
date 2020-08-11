import React, { Component } from 'react';
import { BrowserRouter as Router, Route } from 'react-router-dom';
import Nav from './Navigation';
import Dashboard from './Dashboard';
import SensorManagement from './SensorManagement';
import NodeManagement from './NodeManagement';
import Kibana from './Kibana';
import RegisterAlarm from './components/RegisterAlarm';
import LogicCoreManagement from './LogicCoreManagement';
import RegisterLogic from './LogicCoreComponents/RegisterLogic'
import { sensorListElem, nodeListElem } from './ElemInterface/ElementsInterface';
import { SENSOR_URL, NODE_URL } from './defineUrl';
import { logicCoreElem } from './ElemInterface/LcElementsInterface';

interface appState {
	sensorList: Array<sensorListElem>;
	nodeList: Array<nodeListElem>;
	logicCore: Array<logicCoreElem>
}

class App extends Component<{}, appState> {
	state: appState = {
		sensorList: [],
		nodeList: [],
		logicCore:[],
		// rasp: []
	};

	componentDidMount() {
		this.getsensorList();
		this.getnodeList();
	}

	getsensorList() {
		var url = SENSOR_URL;

		fetch(url)
			.then((res) => res.json())
			.then((data) => {
				console.log(typeof data);
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

	render() {
		return (
			<div>
				<Router>
					<div>
						<Nav></Nav>
						<div className="container pt-4 mt-4">
							<Route exact path="/" render={Dashboard} />
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
										nodeList={this.state.nodeList}
									/>
								)}
							/>
							<Route
								path="/alarm"
								render={() => (
									<RegisterAlarm sensorList={this.state.sensorList} />
								)}
							/>
							<Route 
							path="/logicCore" 
							render={() => (
									<LogicCoreManagement 
										sensorList={this.state.sensorList} 
										nodeList={this.state.nodeList}
										logicCore={this.state.logicCore}
									/>
								)}  
							/>
							<Route path="/kibana" component={Kibana} />
							<Route path="/registerLogic" 
								render={() => (
									<RegisterLogic sensorList={this.state.sensorList} nodeList={this.state.nodeList}/>
								)}>
							</Route>
						</div>
					</div>
				</Router>
			</div>
		);
	}
}

/*
컴포넌트에 전달할 속성이 있을 경우 render 사용
  * <Route exact path="/" render={() => <CardContainer location='cards.json' member={true}/>} />
속성이 없다면 component 사용
  * <Route path="/about" component={About} />
*/
export default App;
