import React, { Component } from 'react';
import { timeRange, logicElem } from '../../ElemInterface/LcElementsInterface';

import '../LogicCore.css';

interface InputTimeCardProps {
	handleInputTimeCardChange: (time_range: logicElem) => void;
}

interface InputTimeCardState {
	elem: 'time';
	arg: { range: Array<timeRange> };
}

/*
InputTimeCard
- Get input of time element
*/
class InputTimeCard extends Component<InputTimeCardProps, InputTimeCardState> {
	state: InputTimeCardState = {
		elem: 'time',
		arg: { range: [{ start: '00:10:10', end: '23:59:59' }] },
	};
	// handle click event of the Add button
	handleAddClick = async () => {
		// Change this state and then..
		await this.setState({
			arg: {
				range: [
					...this.state.arg.range,
					{ start: '00:10:10', end: '23:59:59' },
				],
			},
		});
		// change parent's state
		this.props.handleInputTimeCardChange(this.state);
	};

	// handle click event of the Remove button
	handleRemoveClick = (idx: number) => async () => {
		// Time scope is updated dynamic. Scope can be added or removed freely.
		// so find changing field by using received idx and change state.
		// Change this state and then..
		await this.setState({
			arg: {
				range: this.state.arg.range.filter(
					(s: any, sidx: number) => idx !== sidx
				),
			},
		});
		// change parent's state
		this.props.handleInputTimeCardChange(this.state);
	};

	// Handle time change
	handleTimeChange = (idx: number) => async (
		e: React.ChangeEvent<HTMLInputElement>
	) => {
		// Change this state and then..
		const new_range = this.state.arg.range.map(
			(range: timeRange, sidx: number) => {
				if (idx !== sidx) return range;
				if (e.target.id === 'start-time-input')
					return { ...range, start: e.target.value };
				return { ...range, end: e.target.value };
			}
		);
		// Change this state and then..
		await this.setState({ arg: { range: new_range } });
		// change parent's state
		this.props.handleInputTimeCardChange(this.state);
	};

	render() {
		return (
			<div className="card form-group">
				<div className="card-body row">
					<div className="col-2 right-divider">
						<span style={{ fontSize: '18pt', fontWeight: 500 }}>time</span>
					</div>
					<div className="col-6">
						<div className="row">
							<div className="input-group row margin-left">
								{this.state.arg.range.map((range: timeRange, idx: number) => (
									<div className="input-group margin-bottom">
										<div className="col">
											<input
												className="form-control"
												type="time"
												id="start-time-input"
												value={range.start}
												onChange={this.handleTimeChange(idx)}
											/>
										</div>
										<span>~</span>
										<div className="col">
											<input
												className="form-control col"
												type="time"
												id="end-time-input"
												value={range.end}
												onChange={this.handleTimeChange(idx)}
											/>
										</div>
										<button
											className="btn btn-sm"
											type="button"
											id="button-addon2"
											onClick={this.handleRemoveClick(idx)}
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
									</div>
								))}
							</div>
						</div>
					</div>
					<div className="col">
						<button
							type="button"
							className="btn float-right"
							style={{ background: 'pink' }}
							onClick={this.handleAddClick}
						>
							Add scope
						</button>
					</div>
				</div>
			</div>
		);
	}
}

export default InputTimeCard;
