import React, { Component } from 'react';
import Select from 'react-select';
import { nodeListElem, groupOptionsElem } from '../ElementsInterface'
import '../LogicCore.css';
import { logicElem } from '../LcElementsInterface';

interface GroupCardProps{ 
	nodeList: Array<nodeListElem>;
    handleGroupCardChange: (group: logicElem) => void;
}
interface GroupCardState {
	elem: "group",
	arg: {
		group: Array<string>;
	}
}

class GroupCard extends Component< GroupCardProps, GroupCardState > {
	state: GroupCardState ={
		elem: "group",
		arg: {
			group: [],
		}
	}
	handleGroupChange = async(e: any) => {
		await this.setState({
			arg:{group :[ e.target.value ]},
		})
		this.props.handleGroupCardChange(this.state);
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
							onChange={this.handleGroupChange} 
						/>
					</div>
				</div>
			</div>
        )
    }
}

export default GroupCard;