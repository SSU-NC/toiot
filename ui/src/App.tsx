import React, { Component } from 'react';
import { BrowserRouter as Router, Route } from 'react-router-dom';
import Nav from './Navigation';
import SensorManagement from './ManagementComponents/SensorManagement';
import NodeManagement from './ManagementComponents/NodeManagement';
import ActuatorManagement from './ManagementComponents/ActuatorManagement';
import Dashboard from './KibanaDashboard';
import Visualize from './KibanaVisualize';
import Main from './Home';
import LogicCoreManagement from './LogicCoreComponents/LogicCoreManagement';
import RegisterLogic from './LogicCoreComponents/RegisterLogic';
import SinkManagement from './ManagementComponents/SinkManagement';
//import AlertAlarm from './ManagementComponents/AlertAlarm';
import TopicManagement from './KafkaComponents/Topic/TopicManagement';
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
						{/* <AlertAlarm /> */}
						<div className="container pt-4 mt-4">
							<Route exact path="/" render={Main} />
							<Route path="/sensor" component={SensorManagement} />
							<Route path="/actuator" component={ActuatorManagement} />
							<Route path="/node" component={NodeManagement} />
							<Route path="/sink" component={SinkManagement} />
							<Route path="/topic" component={TopicManagement} />
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
