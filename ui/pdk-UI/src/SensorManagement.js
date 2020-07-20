import React, { Component } from 'react';
import RegisterSensor from "./components/RegisterSensor";
import SensorTable from "./components/SensorTable";

class SensorManagement extends Component {
    constructor(props) {
        super(props);
    }
    render(){
        return(
            <>
                <div style={{float:'right'}}>
                    <RegisterSensor></RegisterSensor>
                </div>
                <div>
                    <h3>Sensor</h3>
                    <SensorTable sensorList={this.props.sensorList}></SensorTable>
                </div>
            </>
        );
    }
}

export default SensorManagement;