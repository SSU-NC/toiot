import React, { useImperativeHandle, Component } from 'react';
import { BrowserRouter as Router, Route } from 'react-router-dom';
import Nav from './Navigation';
import Dashboard from './Dashboard';
import SensorManagement from './SensorManagement';
import NodeManagement from './NodeManagement';
import Kibana from './Kibana';

import RegisterAlarm from './components/RegisterAlarm';

interface appState {
	sensorList: any;
	nodeList: any;
}

class App extends Component<{}, appState> {
	state: appState = {
		sensorList: [],
		nodeList: [],
		// rasp: []
	};

	componentDidMount() {
		this.getsensorList();
		this.getnodeList();
	}

	getsensorList() {
		var url = 'http://220.70.2.160:8080/sensor';

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
		var url = 'http://220.70.2.160:8080/node';

		fetch(url)
			.then((res) => res.json())
			.then((data) => this.setState({ nodeList: data }))
			// .then(response => console.log('Success:', JSON.stringify(response)))
			.catch((error) => console.error('Error:', error));
	}

	render() {
		return (
			// <div>
			// 	<RegisterAlarm />
			// </div>
			<div>
				<Router>
					<div>
						<Nav></Nav>
						<div className="container pt-4 mt-4">
							<Route exact path="/" render={Dashboard} />
							{/* <Route
								path="/management"
								render={() => (
									<SensorManagement
										sensorList={this.state.sensorList}
										nodeList={this.state.nodeList}
									/>
								)}
							/> */}
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
							<Route path="/kibana" component={Kibana} />
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
