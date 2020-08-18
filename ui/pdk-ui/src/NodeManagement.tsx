import React, { Component } from 'react';
import RegisterNode from "./components/RegisterNode";
import NodeTable from "./components/NodeTable";
import { sensorListElem, nodeListElem, sinkListElem } from './ElemInterface/ElementsInterface';

interface NodeManagementProps {
    sensorList: Array<sensorListElem>;
    sinkList: Array<sinkListElem>
    nodeList: Array<nodeListElem>;
}

const NodeManagement: React.FunctionComponent<NodeManagementProps> = props => {
    return(
        <>
        <div style={{float:'right'}}>
            <RegisterNode sensorList={props.sensorList} sinkList={props.sinkList}></RegisterNode>
        </div>
        <div>
            <h3>Node</h3>
            <NodeTable nodeList={props.nodeList}></NodeTable>
        </div>
        </>
    );
}
/*
class NodeManagement extends Component {
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

export default NodeManagement;