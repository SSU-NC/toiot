import React, { Component } from 'react';
import Select from 'react-select';
import { ValueType } from "react-select";

import { nodeListElem, groupOptionsElem } from '../../ElemInterface/ElementsInterface'
import '../LogicCore.css';
import { logicElem } from '../../ElemInterface/LcElementsInterface';

interface InputGroupCardProps{ 
	nodeList: Array<nodeListElem>;
    handleInputGroupCardChange: (group: logicElem) => void;
}
interface InputGroupCardState {
	elem: "group",
	arg: {
		group: Array<string>;
	}
}

class InputGroupCard extends Component< InputGroupCardProps, InputGroupCardState > {
	state: InputGroupCardState ={
		elem: "group",
		arg: {
			group: [],
		}
	}
	handleGroupChange = async(selectedOptions?: any) => {
		var groups = (selectedOptions && selectedOptions.map((groupOption: groupOptionsElem) => {return groupOption.label}));
		await this.setState({
			arg:{ group: groups},
		})
		this.props.handleInputGroupCardChange(this.state);
	}
	
    render() {
        let groupOptions: Array<groupOptionsElem>;
		groupOptions = this.props.nodeList.map((val: nodeListElem) => {
			return { label: val.location, value: val.location };
		});
        return(
            <div className="card form-group">
				<div className="card-body row">
					<div className="col-2 right-divider">
						<span style={{fontSize:'18pt', fontWeight:500}}>group</span>
					</div>
					<div className="col-5">
						<Select 
							isMulti 
							name="group" 
							options={groupOptions} 
							classNamePrefix="select"
							onChange={(selectedOptions?: ValueType<groupOptionsElem>)=>(this.handleGroupChange(selectedOptions ))}
						/>
					</div>
				</div>
			</div>
        )
    }
}

export default InputGroupCard;