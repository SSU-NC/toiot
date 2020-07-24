import React, { Component } from 'react';
import RegisterNode from "./components/RegisterNode";
import NodeTable from "./components/NodeTable";
import { sensorListElem, nodeListElem } from './components/ElementsInterface';

interface SensorManagementProps {
    sensorList: Array<sensorListElem>;
    nodeList: Array<nodeListElem>;
}

const SensorManagement: React.FunctionComponent<SensorManagementProps> = props => {
    return(
        <>
        <div style={{float:'right'}}>
            <RegisterNode sensorList={props.sensorList}></RegisterNode>
        </div>
        <div>
            <h3>Node</h3>
            <NodeTable nodeList={props.nodeList}></NodeTable>
        </div>
        </>
    );
}
/*
class SensorManagement extends Component {
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
}*/

export default SensorManagement;