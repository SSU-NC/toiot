import React, { Component } from 'react';
import '../LogicCore.css';
import {
	numRange,
	logicElem,
	lcValueArg,
} from '../../ElemInterface/LcElementsInterface';

interface ShowValueCardProps {
	logic_elem: logicElem;
	index: number;
}

/*
ShowValueCard
- show value card
*/
class ShowValueCard extends Component<ShowValueCardProps, {}> {
	render() {
		var range = (this.props.logic_elem.arg as lcValueArg).range;
		var value_name = (this.props.logic_elem.arg as lcValueArg).value;

		return (
			<div className="card margin-bottom">
				<div className="card-body row">
					<div className="col-2 right-divider">
						<span style={{ fontSize: '15pt', fontWeight: 500 }}>
							value #{this.props.index}
						</span>
					</div>
					<div className="col-1"></div>
					<div>
						{range.map((range: numRange, idx: number) => (
							<div>
								<span style={{ fontSize: '15pt', fontWeight: 450 }}>
									[range #{idx}]
								</span>
								<br />
								<span style={{ fontSize: '15pt' }}>
									: {range.min} &lt; {value_name} &lt; {range.max}
								</span>
							</div>
						))}
					</div>
				</div>
			</div>
		);
	}
}

export default ShowValueCard;
