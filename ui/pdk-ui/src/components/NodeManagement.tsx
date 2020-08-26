import React, { Component } from 'react';
import RegisterNode from './Register/RegisterNode';
import NodeTable from './Table/NodeTable';
import {
	sensorListElem,
	nodeListElem,
	sinkListElem,
	nodeHealthCheckElem,
} from '../ElemInterface/ElementsInterface';
import { HEALTHCHECK_URL } from '../defineUrl';
import { w3cwebsocket as W3CWebSocket } from 'websocket';

const client = new W3CWebSocket(HEALTHCHECK_URL);

interface NodeManagementProps {
	sensorList: Array<sensorListElem>;
	sinkList: Array<sinkListElem>;
	nodeList: Array<nodeListElem>;
}
interface NodeManagementState {
	nodeState: Array<nodeHealthCheckElem>;
}
interface GroupedNodeListElem {
	sink_id: number;
	node_list: Array<nodeListElem>;
}

function groupBySinkid(
	nodeList: Array<nodeListElem>,
	sinkList: Array<sinkListElem>
) {
	let groupedNodeList: Array<GroupedNodeListElem>;
	groupedNodeList = sinkList.map((sink) => {
		return { sink_id: sink.id, node_list: [] };
	});

	for (var node of nodeList) {
		for (var group of groupedNodeList) {
			if (node.sink_id === group.sink_id) {
				group.node_list.push(node);
			}
		}
	}
	return groupedNodeList;
}

class NodeManagement extends Component<
	NodeManagementProps,
	NodeManagementState
> {
	state: NodeManagementState = {
		nodeState: [],
	};

	componentWillMount() {
		client.onopen = () => {
			console.log('WebSocket Client Connected');
		};
		client.onmessage = (message: any) => {
			console.log(message);
			this.setState({
				nodeState: JSON.parse(message.data),
			});
		};
	}
	render() {
		var groupedNodeList = groupBySinkid(
			this.props.nodeList,
			this.props.sinkList
		);

		return (
			<>
				<div style={{ float: 'right' }}>
					<RegisterNode
						sensorList={this.props.sensorList}
						sinkList={this.props.sinkList}
					></RegisterNode>
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
/*
class NodeManagement extends Component {
    render(){
        return(
            <>
            <div style={{float:'right'}}>
                <RegisterNode sensorList={this.props.sensorList}></RegisterNode>
            </div>
            <div>
                <h3>Node</h3>
                <NodeTable nodeList={this.props.nodeList}></NodeTable>
            </div>
            </>
        );
    }
}*/

export default NodeManagement;
