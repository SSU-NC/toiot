import React, { Component } from 'react';
import RegisterLogicCore from "./components/RegisterLogicCore";
//import LogicCoreTable from "./components/LogicCoreTable";
import { sensorListElem, nodeListElem } from './components/ElementsInterface';
import { Link, Route, BrowserRouter as Router } from "react-router-dom"
interface LogicCoreManagementProps {
    sensorList: Array<sensorListElem>;
	nodeList: Array<nodeListElem>;
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
        </div>
    );
    }
}
export default LogicCoreManagement;