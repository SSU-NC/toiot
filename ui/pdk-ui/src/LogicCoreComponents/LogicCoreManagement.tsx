import React, { Component } from 'react';
import { Link } from 'react-router-dom';
import { logicCoreElem } from '../ElemInterface/LcElementsInterface';
import LogicCoreTable from './LogicCoreTable';

interface LogicCoreManagementProps {
	logicCore: Array<logicCoreElem>;
}

/*
LogicCoreManagement
- Manage logic core table, register logic
*/
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
