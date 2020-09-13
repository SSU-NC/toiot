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
import SinkManagement from './components/SinkManagement';
import AlertAlarm from './components/AlertAlarm';

/* 
App
- Routing
- Show navigation bar (Nav)
- Alert alarm service
*/
class App extends Component {
	render() {
		return (
			<div>
				<Router>
					<div>
						<Nav></Nav>
						<AlertAlarm />
						<div className="container pt-4 mt-4">
							<Route exact path="/" render={Main} />
							<Route path="/sensor" component={SensorManagement} />
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
