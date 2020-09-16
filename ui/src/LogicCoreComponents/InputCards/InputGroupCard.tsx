import React, { Component } from 'react';
import Select from 'react-select';
import { ValueType } from 'react-select';

import {
	nodeListElem,
	groupOptionsElem,
} from '../../ElemInterface/ElementsInterface';
import '../LogicCore.css';
import { logicElem } from '../../ElemInterface/LcElementsInterface';

interface InputGroupCardProps {
	nodeList: Array<nodeListElem>;
	handleInputGroupCardChange: (group: logicElem) => void;
}
interface InputGroupCardState {
	elem: 'group';
	arg: {
		group: Array<string>;
	};
}

/*
InputGroupCard
- Get input of group element
*/
/*
class InputGroupCard extends Component<
	InputGroupCardProps,
	InputGroupCardState
> {
	state: InputGroupCardState = {
		elem: 'group',
		arg: {
			group: [],
		},
	};
 
	// Handle group change by selecting
	handleGroupChange = async (selectedOptions?: any) => {
		// Modify groupOptionsElem format as Array format
		var groups =
			selectedOptions &&
			selectedOptions.map((groupOption: groupOptionsElem) => {
				return groupOption.label;
			});
		// Change this state and then..
		await this.setState({
			arg: { group: groups },
		});
		// change parent's state
		this.props.handleInputGroupCardChange(this.state);
	};

	render() {
		let groupOptions: Array<groupOptionsElem>;

		groupOptions = this.props.nodeList.map((val: nodeListElem) => {
			return { label: val.group, value: val.group };
		});

		// Filter duplications
		var uniqueArray = groupOptions.filter((groupOption, index) => {
			return (
				index ===
				groupOptions.findIndex((groupOption2) => {
					return JSON.stringify(groupOption2) === JSON.stringify(groupOption);
				})
			);
		});
		groupOptions = uniqueArray;

		return (
			<div className="card form-group">
				<div className="card-body row">
					<div className="col-2 right-divider">
						<span style={{ fontSize: '18pt', fontWeight: 500 }}>group</span>
					</div>
					<div className="col-5">
						<Select
							isMulti
							name="group"
							options={groupOptions}
							classNamePrefix="select"
							onChange={(selectedOptions?: ValueType<groupOptionsElem>) =>
								this.handleGroupChange(selectedOptions)
							}
						/>
					</div>
				</div>
			</div>
		);
	}
}

export default InputGroupCard;
*/