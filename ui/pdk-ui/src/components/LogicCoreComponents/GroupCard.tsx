import React, { Component } from 'react';
import Select from 'react-select';
import { nodeListElem, groupOptionsElem } from '../ElementsInterface'
import '../LogicCore.css';

interface GroupCardProps{ 
	nodeList: Array<nodeListElem>;
    handleGroupCardChange: (group: any) => void;
}

class GroupCard extends Component< GroupCardProps, {} > {
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
							onChange={this.props.handleGroupCardChange} 
						/>
					</div>
				</div>
			</div>
        )
    }
}

export default GroupCard;