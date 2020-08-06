import React, { Component } from 'react';
import Select from 'react-select';
import '../LogicCore.css';
import { valueOptionsElem, value_list_elem } from '../ElementsInterface'

interface ValueCardProps{ 
	valueList: Array<value_list_elem>;
	handleValueCardChange: (value:any) => void;
}

interface ValueCardState{
}

class ValueCard extends Component< ValueCardProps, ValueCardState > {
    state: ValueCardState = {
    }

    render() {
		let valueOptions: Array<valueOptionsElem>;
		valueOptions = this.props.valueList.map((val: value_list_elem) => {
			return { label: val.value_name, value: val.value_name };
		});
        return(
            <div className="card form-group">
				<div className="card-body row">
					<div className="col-2 right-divider">
						<h4 className="align-center">value</h4>
					</div>
					<div className="col-3">
						<Select 
							name="value" 
							options={valueOptions}
							classNamePrefix="select" 
							onChange={this.props.handleValueCardChange}
						/>
					</div>
					<div className="col-1"></div>
					<div className="input-group mb-2 col-5">
						<input
							type="number"
							className="form-control"
							name="val_min"
							placeholder="min"
						/>
						&lt; value_name &lt;
						<input
							type="number"
							className="form-control"
							name="val_min"
							placeholder="max"
						/>
					</div>
				</div>
			</div>
        )
    }
}

export default ValueCard;