import React, { Component } from 'react';
import {
	nodeListElem,
	nodeHealthCheckElem,
} from '../../ElemInterface/ElementsInterface';
import { NODE_URL } from '../../defineUrl';
import Pagination from '../Pagination';

enum HealthColor {
	'red',
	'#FACC2E',
	'lime',
}

interface NodeTableProps {
	sink_id: number;
	nodeState: Array<nodeHealthCheckElem>;
}
interface NodeTableState {
	nodeList: Array<nodeListElem>;
	currentPage: number;
	pages: number; // num of total pages
}
/*
NodeTable
- Show up node list.
*/

class NodeTable extends Component<NodeTableProps, NodeTableState> {
	state: NodeTableState = {
		nodeList: [],
		currentPage: 1,
		pages: 0,
	};

	componentDidMount() {
		this.getnodeList(this.state.currentPage);
	}

	// Get node list from backend
	getnodeList(page: number) {
		var url =
			NODE_URL + '?sink=' + this.props.sink_id + '&page=' + page;    // 현재 페이지에서 해당 sink의 노드 list 가져오는건가

		fetch(url)
			.then((res) => res.json())
			.then((data) =>
				page === 1
					? this.setState({ nodeList: data.nodes, pages: data.pages })
					: this.setState({ nodeList: data.nodes })
			)
			.catch((error) => console.error('Error:', error));
	}

	// Handle click event of the Remove button
	handleRemoveClick = (node_id: number) => () => {
		var url = NODE_URL + '/' + node_id;

		fetch(url, {
			method: 'DELETE',
			headers: {
				'Content-Type': 'application/json',
			},
		})
			.then((res) => res.json())
			.catch((error) => console.error('Error:', error))
			.then(() => window.location.reload(false));
	};

	// Find node state(health) and represent as colors (red - yellow - green, gray)
	findNodeState = (id: number) => {
		for (let prop in this.props.nodeState) {
			if (this.props.nodeState[prop].nid === id) {
				return (
					<td
						style={{
							color: HealthColor[this.props.nodeState[prop].state],
						}}
					>
						●
					</td>
				);
			}
		}
		return <td style={{ color: 'gray' }}>●</td>;
	};

	handlePageChange = (page: number) => {
		this.setState({ currentPage: page });
		this.getnodeList(page);
	};

	render() {
		var sink_id = this.props.sink_id;
		var index;
		return (
			<>
				<table className="table">
					<thead>
						<tr>
							<th scope="col">#</th>
							<th scope="col">name</th>
							<th scope="col">id</th> 
							<th scope="col">sensors</th>
							<th scope="col">health</th>
							<th scope="col"></th>
						</tr>
					</thead>
					<tbody>
						{this.state.nodeList.map((node: nodeListElem, idx: number) => (        
							<tr>
								<th scope="row">{idx}</th>
								<td>{node.name}</td>
								<td>{node.id}</td>
								<td>{node.sensors.map((sensor: any) => sensor.name + ', ')}</td>
								{this.findNodeState(node.id)}        
								<td>
									<button
										className="btn btn-default btn-sm"
										type="button"
										id="button-delete"
										onClick={this.handleRemoveClick(node.id)}
									>
										<svg
											width="1em"
											height="1em"
											viewBox="0 0 16 16"
											className="bi bi-trash-fill"
											fill="currentColor"
											xmlns="http://www.w3.org/2000/svg"
										>
											<path
												fillRule="evenodd"
												d="M2.5 1a1 1 0 0 0-1 1v1a1 1 0 0 0 1 1H3v9a2 2 0 0 0 2 2h6a2 2 0 0 0 2-2V4h.5a1 1 0 0 0 1-1V2a1 1 0 0 0-1-1H10a1 1 0 0 0-1-1H7a1 1 0 0 0-1 1H2.5zm3 4a.5.5 0 0 1 .5.5v7a.5.5 0 0 1-1 0v-7a.5.5 0 0 1 .5-.5zM8 5a.5.5 0 0 1 .5.5v7a.5.5 0 0 1-1 0v-7A.5.5 0 0 1 8 5zm3 .5a.5.5 0 0 0-1 0v7a.5.5 0 0 0 1 0v-7z"
											/>
										</svg>
									</button>
								</td>
							</tr>
						))}
					</tbody>
				</table>
				<Pagination
					pages={this.state.pages}
					currentPage={this.state.currentPage}
					onPageChange={this.handlePageChange}
				></Pagination>
			</>
		);
	}
}

export default NodeTable;
