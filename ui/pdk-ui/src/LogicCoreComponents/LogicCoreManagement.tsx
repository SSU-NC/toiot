import React, { Component } from 'react';
import RegisterLogic from './RegisterLogic';
import {
	sensorListElem,
	nodeListElem,
} from '../ElemInterface/ElementsInterface';
import { Link } from 'react-router-dom';
import { logicCoreElem } from '../ElemInterface/LcElementsInterface';
import LogicCoreTable from './LogicCoreTable';

interface LogicCoreManagementProps {
	sensorList: Array<sensorListElem>;
	nodeList: Array<nodeListElem>;
	logicCore: Array<logicCoreElem>;
}

class LogicCoreManagement extends Component<LogicCoreManagementProps> {
	render() {
		return (
			<div>
				<Link to="/registerLogic">
					<button
						type="button"
						className="btn float-right"
						style={{ background: 'pink' }}
					>
						register logic
					</button>
				</Link>
				<h3>Logic Core</h3>
				<LogicCoreTable logicCore={this.props.logicCore} />
			</div>
		);
	}
}
export default LogicCoreManagement;
