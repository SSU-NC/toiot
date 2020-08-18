import React, { Component } from 'react';
import '../LogicCore.css';
import { logicElem, lcActionArg } from '../../ElemInterface/LcElementsInterface';

interface ShowActionCardProps{
	logic_elem: logicElem;
	index: number;
}

class ShowActionCard extends Component< ShowActionCardProps,{} > {
    render() {
		var action = (this.props.logic_elem.arg as lcActionArg);

        return(
            <div className="card margin-bottom">
				<div className="card-body row">
					<div className="col-2 right-divider">
						<span style={{fontSize:'15pt', fontWeight:500}} >action #{this.props.index}</span>
					</div>
					<div className="col-1"></div>
					<div>
						{(this.props.logic_elem.elem === "alarm") ?
						(
						<div>
							<span style={{fontSize:'15pt', fontWeight:450}}>Alarm MSG </span>
							<span style={{fontSize:'15pt'}}>: {action.text}</span>
						</div>
	
						) :
						((this.props.logic_elem.elem === "email") ? 
						(
						<div>
							<span style={{fontSize:'15pt', fontWeight:450}}>Email address </span>
							<span style={{fontSize:'15pt'}}>: {action.text}</span>
						</div>
						):(
						<div></div>
						))}

					</div>
				</div>
			</div>
		)
    }
}

export default ShowActionCard;