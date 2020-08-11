import React, { Component } from 'react';
import Select from 'react-select';
import {sensorOptionsElem,sensorListElem} from '../../ElemInterface/ElementsInterface'
import '../LogicCore.css';

interface ShowInputSensorCardProps{ 
    sensorList: Array<sensorListElem>;
	handleShowInputSensorCardChange: (sensor: any) => void;
}

class ShowInputSensorCard extends Component< ShowInputSensorCardProps, {} > {
    render() {
        let sensorOptions: Array<sensorOptionsElem>;
		sensorOptions = this.props.sensorList.map((val: sensorListElem) => {
			return { label: val.name, value: val.name, uuid: val.uuid, value_list: val.value_list};
		});
        return(
                <div className="card form-group">
					<div className="card-body row ">
						<div className="col-2 right-divider">
							<span style={{fontSize:'18pt', fontWeight:500}}>sensor</span>
							</div>
							<div className="col-5">
								<Select
									name="sensor" 
									options={sensorOptions} 
									classNamePrefix="select" 
									onChange={this.props.handleShowInputSensorCardChange}
								/>
						</div>
					</div>
                </div>		
        )
    }
}

export default ShowInputSensorCard;