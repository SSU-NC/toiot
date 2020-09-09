import React, { Component } from 'react';
import { Link } from 'react-router-dom';
import LogicTable from './LogicTable';

/*
LogicCoreManagement
- Manage logic core table, register logic
*/
class LogicCoreManagement extends Component {
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
				<LogicTable />
			</div>
		);
	}
}
export default LogicCoreManagement;
