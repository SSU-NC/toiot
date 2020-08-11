import React, { Component } from 'react';
import RegisterLogic from "./LogicCoreComponents/RegisterLogic";
//import LogicCoreTable from "./components/LogicCoreTable";
import { sensorListElem, nodeListElem } from './ElemInterface/ElementsInterface';
import { Link, Route, BrowserRouter as Router } from "react-router-dom"
import { logicCoreElem } from './ElemInterface/LcElementsInterface';
import LogicCoreTable from './LogicCoreComponents/LogicCoreTable';

interface LogicCoreManagementProps {
    sensorList: Array<sensorListElem>;
    nodeList: Array<nodeListElem>;
    logicCore: Array<logicCoreElem>;
}

class LogicCoreManagement extends Component<LogicCoreManagementProps> {
    render() {
    return(
        <div>
            <Link to='/registerLogic'>
                <button 
                    type="button" 
                    className="btn float-right"
                    style={{background:'pink'}}
                >register logic</button>
            </Link>
            <h3>Logic Core</h3>
            <LogicCoreTable logicCore={this.props.logicCore}/>
        </div>
    );
    }
}
export default LogicCoreManagement;