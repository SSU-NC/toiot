import React, { Component } from 'react';
import RegisterTopic from './RegisterTopic';
import TopicTable from './TopicTable';

/*
TopicManagement
- Manage topic table, register topic
*/
class TopicManagement extends Component {
	render() {
		return (
			<>
				<div style={{ float: 'right' }}>
					<button
						type="button"
						className="btn"
						data-toggle="modal"
						data-target="#register-topic-modal"
						style={{ background: 'pink' }}
					>
						register topic
					</button>
					<RegisterTopic></RegisterTopic>
				</div>
				<div>
					<h3>Topic</h3>
					<TopicTable></TopicTable>
				</div>
			</>
		);
	}
}

export default TopicManagement;
