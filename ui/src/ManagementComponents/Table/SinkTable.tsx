import React, { Component } from 'react';
import { sinkListElem } from '../../ElemInterface/ElementsInterface';
import { SINK_URL } from '../../defineUrl';
import Pagination from '../Pagination';

interface SinkTableState {
	sinkList: Array<sinkListElem>;
	currentPage: number;
	pages: number; // num of total pages
}

/*
SinkTable
- Show up sink list.
*/
class SinkTable extends Component<{}, SinkTableState> {
	state: SinkTableState = {
		sinkList: [],
		currentPage: 1,
		pages: 0,
	};
	componentDidMount() {
		this.getsinkList(this.state.currentPage);
	}

	// Get sink list from backend
	getsinkList(page: number) {
		var url = SINK_URL + '?page=' + page;

		fetch(url)
			.then((res) => res.json())
			.then((data) => page === 1
			? this.setState({ sinkList: data.sinks, pages: data.pages })
			: this.setState({ sinkList: data.sinks }))
			.catch((error) => console.error('Error:', error));
	}

	// Handle click event of the Remove button
	handleRemoveClick = (sink_id: number) => () => {
		var url = SINK_URL + '/' + sink_id;

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

	handlePageChange = (page: number) => {
		this.setState({ currentPage: page });
		this.getsinkList(page);
	};

	render() {
		return (
			<>
				<table className="table">
					<thead>
						<tr>
							<th scope="col">#</th>
							<th scope="col">name</th>
							<th scope="col">id</th>
							<th scope="col">address</th>
							<th scope="col"></th>
						</tr>
					</thead>
					<tbody>
						{this.state.sinkList.map((sink: sinkListElem, idx: number) => (
							<tr>
								<th scope="row">{idx + 10 * (this.state.currentPage - 1)}</th>
								<td>{sink.name}</td>
								<td>{sink.id}</td>
								<td>{sink.addr}</td>
								<td>
									<button
										className="btn btn-default btn-sm"
										type="button"
										id="button-delete"
										onClick={this.handleRemoveClick(sink.id)}
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
												fill-rule="evenodd"
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

export default SinkTable;
