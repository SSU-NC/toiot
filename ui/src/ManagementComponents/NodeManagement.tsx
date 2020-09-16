import React, { Component } from 'react';
import RegisterNode from './Register/RegisterNode';
import NodeTable from './Table/NodeTable';
import {
	sinkListElem,
	nodeHealthCheckElem,
} from '../ElemInterface/ElementsInterface';
import { HEALTHCHECK_URL, SINK_URL } from '../defineUrl';
// import { w3cwebsocket as W3CWebSocket } from 'websocket';
import NodeMap from './NodeMap';

// const client = new W3CWebSocket(HEALTHCHECK_URL);

interface NodeManagementState {
	sinkList: Array<sinkListElem>;
	nodeState: Array<nodeHealthCheckElem>;

	showAllValid: boolean;
}
/*
NodeManagement
- Manage node table, register node
*/

class NodeManagement extends Component<{}, NodeManagementState> {
	state: NodeManagementState = {
		sinkList: [],
		nodeState: [],

		showAllValid: true,
	};

	// Conect web socket
	// componentWillMount() {
	// 	client.onopen = () => {
	// 		console.log('WebSocket Client for Health Check Connected');
	// 	};
	// 	client.onmessage = (message: any) => {
	// 		console.log(message);
	// 		this.setState({
	// 			nodeState: JSON.parse(message.data),
	// 		});
	// 	};
	// }

	componentDidMount() {
		this.getsinkList();
	}

	// Get sink list from backend
	getsinkList() {
		var url = SINK_URL;

		fetch(url)
			.then((res) => res.json())
			.then((data) => this.setState({ sinkList: data }))
			.catch((error) => console.error('Error:', error));
	}

	handleAllClick = () => {
		this.setState({
			showAllValid: true,
		});
	}
	handleMapClick = () => {
		this.setState({
			showAllValid: false,
		});
	}

	render() {
		return (
			<div>
				<div style={{ float: 'right' }}>
					<button
						type="button"
						className="btn"
						data-toggle="modal"
						data-target="#register-node-modal"
						style={{ background: 'pink' }}
					>
						register node
					</button>
					<RegisterNode></RegisterNode>
				</div>
				<div>
					<h3>Node</h3>
					<hr />
					<div style={{ float: 'right' }}>
						<span style={{ color: 'gray' }}>● : don't know </span>
						<span style={{ color: 'lime' }}>● : stable </span>
						<span style={{ color: '#FACC2E' }}>● : unstable </span>
						<span style={{ color: 'red' }}>● : disconnect </span>
					</div>
					<span >Viewer type </span>
					<button
						type="button"
						className="btn"
						style={{ background: 'pink' }}
						onClick={this.handleAllClick}
					>
						All
					</button>
					<span> </span>
					<button
						type="button"
						className="btn"
						style={{ background: 'pink' }}
						onClick={this.handleMapClick}
					>
						Map
					</button>
					<hr/>
					{(this.state.showAllValid)?(
						<div>
						{this.state.sinkList.map((sink: sinkListElem, idx: number) => (
							<div>
								<span style={{ fontSize: '18pt', fontWeight: 500 }}>
									Sink {sink.id}
								</span>
								<button
									className="btn dropdown-toggle"
									type="button"
									data-toggle="collapse"
									data-target={'#sink' + sink.id.toString()}
									aria-controls={sink.id.toString()}
									style={{ color: 'black' }}
								></button>
								<div
									id={'sink' + sink.id.toString()}
									className="collapse"
								>
									<NodeTable
										sink_id={sink.id}
										nodeState={this.state.nodeState}
									></NodeTable>
								</div>
							</div>
						))}
					</div>
					):(
						<NodeMap nodeState={this.state.nodeState}></NodeMap>
					)}	
				</div>
			</div>
		);
	};
};


export default NodeManagement;
