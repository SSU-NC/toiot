import React, { Component } from 'react';
import RegisterNode from "./components/RegisterNode";

class SensorManagement extends Component {
    constructor(props) {
        super(props);
    }
    render(){
        return(
            <div style={{float:'right'}}>
                <RegisterNode sensorList={this.props.sensorList}></RegisterNode>
            </div>
        );
    }
}

export default SensorManagement;