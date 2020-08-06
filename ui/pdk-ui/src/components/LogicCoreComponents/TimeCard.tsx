import React, { Component } from 'react';
import Select from 'react-select';
import { nodeListElem, groupOptionsElem } from '../ElementsInterface'
import '../LogicCore.css';

interface TimeCardProps{ 
    //onTimeCardSubmit: (time: any) => void;
}

class TimeCard extends Component< TimeCardProps, {} > {
    render() {
        return(
            <div className="card form-group">
            <div className="card-body row">
                <div className="col-2 right-divider">
                    <h4 className="align-middle">time</h4>
                </div>
                <div className="col-5">
                    <div className="form-group row">
                        <div className="input-group row margin-left">
                            <div className="col">
                                <input
                                    className="form-control"
                                    type="time"
                                    value="20:10:12"
                                    id="start-time-input"
                                />
                            </div>
                            <span>~</span>
                            <div className="col">
                                <input
                                    className="form-control col"
                                    type="time"
                                    value="23:59:59"
                                    id="end-time-input"
                                />
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div> 
        )
    }
}

export default TimeCard;