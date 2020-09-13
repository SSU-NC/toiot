import React, { Component } from 'react';
import RegisterNode from './Register/RegisterNode';
import NodeTable from './Table/NodeTable';
import {
	nodeListElem,
	sinkListElem,
	nodeHealthCheckElem,
} from '../ElemInterface/ElementsInterface';
import { HEALTHCHECK_URL, NODE_URL, SINK_URL } from '../defineUrl';
import { w3cwebsocket as W3CWebSocket } from 'websocket';
import NodeMap from './NodeMap';

const client = new W3CWebSocket(HEALTHCHECK_URL);

interface NodeManagementState {
	nodeList: Array<nodeListElem>;
	sinkList: Array<sinkListElem>;
	nodeState: Array<nodeHealthCheckElem>;
}
interface GroupedNodeListElem {
	sink_id: number;
	node_list: Array<nodeListElem>;
}

// Grouping node list as sink id.
function groupBySinkid(
	nodeList: Array<nodeListElem>,
	sinkList: Array<sinkListElem>
) {
	let groupedNodeList: Array<GroupedNodeListElem>;

	// Initialize Grouped node list as sink id.
	groupedNodeList = sinkList.map((sink) => {
		return { sink_id: sink.id, node_list: [] };
	});

	// Fill node_list field of Grouped node list.
	for (var node of nodeList) {
		for (var group of groupedNodeList) {
			if (node.sink_id === group.sink_id) {
				group.node_list.push(node);
			}
		}
	}
	return groupedNodeList;
}

/*
NodeManagement
- Manage node table, register node
*/
class NodeManagement extends Component<{}, NodeManagementState> {
	state: NodeManagementState = {
		nodeList: [],
		sinkList: [],
		nodeState: [],
	};

	// Conect web socket
	componentWillMount() {
		client.onopen = () => {
			console.log('WebSocket Client for Health Check Connected');
		};
		client.onmessage = (message: any) => {
			console.log(message);
			this.setState({
				nodeState: JSON.parse(message.data),
			});
		};
	}

	componentDidMount() {
		this.getnodeList();
		this.getsinkList();
	}

	// Get node list from backend
	getnodeList() {
		var url = NODE_URL;

		fetch(url)
			.then((res) => res.json())
			.then((data) => this.setState({ nodeList: data }))
			.catch((error) => console.error('Error:', error));
	}

	// Get sink list from backend
	getsinkList() {
		var url = SINK_URL;

		fetch(url)
			.then((res) => res.json())
			.then((data) => this.setState({ sinkList: data }))
			.catch((error) => console.error('Error:', error));
	}

	render() {
		var groupedNodeList = groupBySinkid(
			this.state.nodeList,
			this.state.sinkList
		);

		return (
			<>
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
					<NodeMap
						nodeList={this.state.nodeList}
						nodeState={this.state.nodeState}
					></NodeMap>
					<div>
						{groupedNodeList.map((group: GroupedNodeListElem, idx: number) => (
							<div>
								<span style={{ fontSize: '18pt', fontWeight: 500 }}>
									Sink {group.sink_id}
								</span>
								<button
									className="btn dropdown-toggle"
									type="button"
									data-toggle="collapse"
									data-target={'#sink' + group.sink_id.toString()}
									aria-controls={group.sink_id.toString()}
									style={{ color: 'black' }}
								></button>
								<div
									id={'sink' + group.sink_id.toString()}
									className="collapse"
								>
									<NodeTable
										nodeList={group.node_list}
										nodeState={this.state.nodeState}
									></NodeTable>
								</div>
							</div>
						))}
					</div>
				</div>
			</>
		);
	}
}

export default NodeManagement;
