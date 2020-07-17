import React, { Component } from 'react';
import RegisterSensor from "./components/RegisterSensor";
import RegisterNode from "./components/RegisterNode";

class SensorManagement extends Component {
    constructor(props) {
        super(props);
    }
    render(){
        return(
            <div style={{float:'right'}}>
                <RegisterSensor></RegisterSensor>
            </div>
        );
    }
}

export default SensorManagement;