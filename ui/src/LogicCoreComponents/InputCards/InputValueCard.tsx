import React, { Component } from 'react';
import Select from 'react-select';
import '../LogicCore.css';
import {
	valueOptionsElem,
	value_list_elem,
} from '../../ElemInterface/ElementsInterface';
import { numRange, logicElem } from '../../ElemInterface/LcElementsInterface';

interface InputValueCardProps {
	valueList: Array<value_list_elem>;
	handleInputValueCardChange: (value: logicElem) => void;
	handleRemoveInputValueCardClick: () => void;
	index: number;
}

interface InputValueCardState {
	elem: string;
	arg: {
		value: string;
		range: Array<numRange>;
	};
}

/*
InputValueCard
- Get input of value element
*/
class InputValueCard extends Component<
	InputValueCardProps,
	InputValueCardState
> {
	state: InputValueCardState = {
		elem: 'value',
		arg: {
			value: '',
			range: [{ min: 0, max: 255 }],
		},
	};
	// handle click event of the Add button
	handleAddClick = async () => {
		// Change this state and then..
		await this.setState({
			arg: {
				value: this.state.arg.value,
				range: [...this.state.arg.range, { min: 0, max: 255 }],
			},
		});
		// change parent's state
		this.props.handleInputValueCardChange(this.state);
	};

	// handle click event of the Remove button
	handleRemoveClick = (idx: number) => async () => {
		// Change this state and then..
		await this.setState({
			arg: {
				value: this.state.arg.value,
				range: this.state.arg.range.filter(
					(s: any, sidx: number) => idx !== sidx
				),
			},
		});
		// change parent's state
		this.props.handleInputValueCardChange(this.state);
	};

	// Handle number change (max, min)
	handleNumChange = (idx: number) => async (e: any) => {
		// Value's scope is updated dynamic. Scope can be added or removed freely.
		// so find changing field by using received idx and change state.
		const new_range_elem = this.state.arg.range.map(
			(rangeElem: numRange, sidx: number) => {
				if (idx !== sidx) return rangeElem;
				if (e.target.id === 'val_min')
					return { ...rangeElem, min: parseInt(e.target.value) };
				return { ...rangeElem, max: parseInt(e.target.value) };
			}
		);
		// Change this state and then..
		await this.setState({
			arg: { value: this.state.arg.value, range: new_range_elem },
		});
		// change parent's state
		this.props.handleInputValueCardChange(this.state);
	};

	// Handle value change by select
	handleValueChange = async (e: any) => {
		// Change this state and then..
		await this.setState({
			arg: { value: e.value, range: this.state.arg.range },
		});
		// change parent's state
		this.props.handleInputValueCardChange(this.state);
	};
	render() {
		let valueOptions: Array<valueOptionsElem>;
		valueOptions = this.props.valueList.map((val: value_list_elem) => {
			return {
				label: val.value_name,
				value: val.value_name,
				range: [{ min: 0, max: 255 }],
			};
		});
		return (
			<div className="card form-group">
				<div className="card-body row">
					<div className="col-2 right-divider">
						<span style={{ fontSize: '18pt', fontWeight: 500 }}>
							value #{this.props.index}
						</span>
						<button
							className="btn btn-sm float-right"
							type="button"
							id="button-addon2"
							onClick={this.props.handleRemoveInputValueCardClick}
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
					</div>
					<div className="col-3">
						<Select
							name="value"
							options={valueOptions}
							classNamePrefix="select"
							onChange={this.handleValueChange}
						/>
					</div>
					<div className="col-1"></div>
					<div className="col-4">
						{this.state.arg.range.map((range: numRange, idx: number) => (
							<div className="input-group mb-2">
								<input
									type="number"
									className="form-control"
									id="val_min"
									placeholder="min"
									value={range.min}
									onChange={this.handleNumChange(idx)}
								/>
								<span>&lt; {this.state.arg.value} &lt;</span>
								<input
									type="number"
									className="form-control"
									id="val_max"
									placeholder="max"
									value={range.max}
									onChange={this.handleNumChange(idx)}
								/>
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
											fill-rule="evenodd"
											d="M2.5 1a1 1 0 0 0-1 1v1a1 1 0 0 0 1 1H3v9a2 2 0 0 0 2 2h6a2 2 0 0 0 2-2V4h.5a1 1 0 0 0 1-1V2a1 1 0 0 0-1-1H10a1 1 0 0 0-1-1H7a1 1 0 0 0-1 1H2.5zm3 4a.5.5 0 0 1 .5.5v7a.5.5 0 0 1-1 0v-7a.5.5 0 0 1 .5-.5zM8 5a.5.5 0 0 1 .5.5v7a.5.5 0 0 1-1 0v-7A.5.5 0 0 1 8 5zm3 .5a.5.5 0 0 0-1 0v7a.5.5 0 0 0 1 0v-7z"
										/>
									</svg>
								</button>
							</div>
						))}
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

export default InputValueCard;
