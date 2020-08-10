import React, { Component } from 'react';
import Select from 'react-select';
import {sensorOptionsElem,sensorListElem} from '../ElementsInterface'
import '../LogicCore.css';

interface SensorCardProps{ 
    sensorList: Array<sensorListElem>;
	handleSensorCardChange: (sensor: any) => void;
}

class SensorCard extends Component< SensorCardProps, {} > {
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
									onChange={this.props.handleSensorCardChange}
								/>
						</div>
					</div>
                </div>		
        )
    }
}

export default SensorCard;