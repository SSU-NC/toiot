import React, { useImperativeHandle, Component } from 'react';
import { BrowserRouter as Router, Route } from 'react-router-dom';
import Nav from './Navigation';
import Dashboard from './Dashboard';
import SensorManagement from './SensorManagement';
import NodeManagement from './NodeManagement';
import Kibana from './Kibana';

class App extends Component {
	constructor(props) {
		super(props);
		this.state = {
			sensorlist: [],
			nodelist: [],
			// rasp: []
		};
	}
	componentDidMount() {
		this.getSensorList();
		this.getNodeList();
	}

	getSensorList() {
		var url = 'http://220.70.2.160:8080/sensor/info';

		fetch(url)
			.then((res) => res.json())
			.then((data) => this.setState({ sensorlist: data }))
			// .then(response => console.log('Success:', JSON.stringify(response)))
			.catch((error) => console.error('Error:', error));
	}

	getNodeList() {
		var url = 'http://220.70.2.160:8080/node/info';

		fetch(url)
			.then((res) => res.json())
			.then((data) => this.setState({ sensorlist: data }))
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
							{/* <Route
								path="/management"
								render={() => (
									<SensorManagement
										sensorList={this.state.sensorlist}
										nodeList={this.state.nodeList}
									/>
								)}
							/> */}
							<Route
								path="/sensor"
								render={() => (
									<SensorManagement
										sensorList={this.state.sensorlist}
										nodeList={this.state.nodeList}
									/>
								)}
							/>
							<Route
								path="/node"
								render={() => (
									<NodeManagement
										sensorList={this.state.sensorlist}
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
