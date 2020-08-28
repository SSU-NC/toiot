import React, { Component } from 'react';
import '../LogicCore.css';
import { logicElem, lcGroupArg } from '../../ElemInterface/LcElementsInterface';

interface ShowGroupCardProps {
	logic_elem: logicElem;
}
/*
ShowGroupCard
- show group card
*/
class ShowGroupCard extends Component<ShowGroupCardProps, {}> {
	render() {
		var groups = (this.props.logic_elem.arg as lcGroupArg).group;

		return (
			<div className="card margin-bottom">
				<div className="card-body row">
					<div className="col-2 right-divider">
						<span style={{ fontSize: '15pt', fontWeight: 500 }}>group</span>
					</div>
					<div className="col-1"></div>
					<div>
						{groups.map((group: string, idx: number) => (
							<div>
								<span style={{ fontSize: '15pt' }}>{group},</span>
							</div>
						))}
					</div>
				</div>
			</div>
		);
	}
}

export default ShowGroupCard;
