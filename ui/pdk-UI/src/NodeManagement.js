import React, { Component } from 'react';
import RegisterNode from "./components/RegisterNode";
import NodeTable from "./components/NodeTable";

class SensorManagement extends Component {
    constructor(props) {
        super(props);
    }
    render(){
        return(
            <>
            <div style={{float:'right'}}>
                <RegisterNode sensorList={this.props.sensorList}></RegisterNode>
            </div>
            <div>
                <h3>Node</h3>
                <NodeTable nodeList={this.props.nodeList}></NodeTable>
            </div>
            </>
        );
    }
}

export default SensorManagement;