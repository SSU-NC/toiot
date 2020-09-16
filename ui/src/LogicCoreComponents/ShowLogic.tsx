import React, { Component } from 'react';
import { logicListElem, logicElem } from '../ElemInterface/LcElementsInterface';
import ShowValueCard from './ShowCards/ShowValueCard';
//import ShowGroupCard from './ShowCards/ShowGroupCard';
import ShowActionCard from './ShowCards/ShowActionCard';
import ShowTimeCard from './ShowCards/ShowTimeCard';
import ShowSensorCard from './ShowCards/ShowSensorCard';

interface ShowLogicProps {
	logic: logicListElem;
	index: number;
}

/*
ShowLogic
- Show logic's elements
*/
class ShowLogic extends Component<ShowLogicProps, {}> {
	render() {
		// Divide modal id per index
		var modal_id: string = 'show-logic-'
			.concat(this.props.index + '')
			.concat('-modal');
		return (
			<>
				<button
					type="button"
					className="btn"
					data-toggle="modal"
					style={{ background: 'pink' }}
					data-target={'#'.concat(modal_id)}
				>
					show logic
				</button>
				<div
					className="modal fade"
					id={modal_id}
					role="dialog"
					aria-labelledby="show-logic-modal"
				>
					<div className="modal-dialog modal-lg" role="document">
						<div className="modal-content">
							<div className="modal-header">
								<h4 className="modal-title">{this.props.logic.logic_name}</h4>
								<button
									type="button"
									className="close"
									data-dismiss="modal"
									aria-label="Close"
								>
									<span aria-hidden="true">×</span>
								</button>
							</div>
							<div className="modal-body">
								<ShowSensorCard sensor_id={this.props.logic.sensor_id} />
								{/* {this.props.logic.elems
									.filter(function (element) {
										return element.elem === 'group';
									})
									.map((groupCard: logicElem) => (
										<ShowGroupCard logic_elem={groupCard} />
									))} */}
								{this.props.logic.elems
									.filter(function (element) {
										return element.elem === 'time';
									})
									.map((timeCard: logicElem) => (
										<ShowTimeCard logic_elem={timeCard} />
									))}
								{this.props.logic.elems
									.filter(function (element) {
										return element.elem === 'value';
									})
									.map((valueCard: logicElem, idx: number) => (
										<ShowValueCard index={idx} logic_elem={valueCard} />
									))}
								{this.props.logic.elems
									.filter(function (element) {
										return element.elem === 'alarm' || element.elem === 'email';
									})
									.map((actionCard: logicElem, idx: number) => (
										<ShowActionCard index={idx} logic_elem={actionCard} />
									))}
							</div>
						</div>
					</div>
				</div>
			</>
		);
	}
}

export default ShowLogic;
