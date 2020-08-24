import React, { Component } from 'react';
import { logicCoreElem } from '../ElemInterface/LcElementsInterface';
import ShowLogic from './ShowLogic';
import { LOGICCORE_URL } from '../defineUrl';

interface logicTableProps {
	logicCore: Array<logicCoreElem>;
}

class logicTable extends Component<logicTableProps> {
	handleRemoveClick = (id: string) => () => {
		var url = LOGICCORE_URL;

		fetch(url, {
			method: 'DELETE',
			body: JSON.stringify({ id: id }),
			headers: {
				'Content-Type': 'application/json',
			},
		})
			.then((res) => res.json())
			.catch((error) => console.error('Error:', error));
	};

	render() {
		return (
			<>
				<table className="table">
					<thead>
						<tr>
							<th scope="col">#</th>
							<th scope="col">name</th>
							<th scope="col">logic info</th>
							<th scope="col"></th>
						</tr>
					</thead>
					<tbody>
						{this.props.logicCore.map((logic: logicCoreElem, idx: number) => (
							<tr>
								<th scope="row">{idx}</th>
								<td>{logic.logic_name}</td>
								<td>
									<ShowLogic index={idx} logic={logic} />
								</td>
								<td>
									<button
										className="btn btn-default btn-sm"
										type="button"
										id="button-delete"
										onClick={this.handleRemoveClick(logic.id)}
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

export default logicTable;
