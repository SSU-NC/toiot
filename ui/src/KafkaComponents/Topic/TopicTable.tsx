import React, { Component } from 'react';
import { topicListElem } from '../../ElemInterface/ElementsInterface';
import { TOPIC_URL } from '../../defineUrl';

interface TopicTableState {
	topicList: Array<topicListElem>;
}

/*
TopicTable
- Show up topic list.
*/
class TopicTable extends Component<{}, TopicTableState> {
	state: TopicTableState = {
		topicList: [],
	};
	componentDidMount() {
		this.gettopicList();
	}

	// Get topic list from backend
	gettopicList() {
		var url = TOPIC_URL;

		fetch(url)
			.then((res) => res.json())
			.then((data) => this.setState({ topicList: data }))
			.catch((error) => console.error('Error:', error));
	}

	// Handle click event of the Remove button
	handleRemoveClick = (topic_id: number) => () => {
		var url = TOPIC_URL + '/' + topic_id;

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

	render() {
		return (
			<>
				<table className="table">
					<thead>
						<tr>
							<th scope="col">#</th>
							<th scope="col">name</th>
							<th scope="col">partitions</th>
							<th scope="col">replications</th>
							<th scope="col"></th>
						</tr>
					</thead>
					<tbody>
						{this.state.topicList.map((topic: topicListElem, idx: number) => (
							<tr>
								<th scope="row">{idx}</th>
								<td>{topic.name}</td>
								<td>{topic.partitions}</td>
								<td>{topic.replications}</td>
								<td>
									<button
										className="btn btn-default btn-sm"
										type="button"
										id="button-delete"
										onClick={this.handleRemoveClick(topic.id)}
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
			</>
		);
	}
}

export default TopicTable;
