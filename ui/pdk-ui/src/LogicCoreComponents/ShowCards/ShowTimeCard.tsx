import React, { Component } from 'react';
import '../LogicCore.css';
import { timeRange, logicElem, lcTimeArg } from '../../ElemInterface/LcElementsInterface';

interface ShowTimeCardProps{
	logic_elem: logicElem;
}

class ShowTimeCard extends Component< ShowTimeCardProps,{} > {
    render() {
		var range = (this.props.logic_elem.arg as lcTimeArg).range;

        return(
            <div className="card margin-bottom">
				<div className="card-body row">
					<div className="col-2 right-divider">
						<span style={{fontSize:'15pt', fontWeight:500}} >time</span>
					</div>
					<div className="col-1"></div>
					<div>
						{range.map((range: timeRange, idx: number) => (
							<div>
								<span style={{fontSize:'15pt', fontWeight:450}}>range #{idx}</span>
								<br/>
								<span style={{fontSize:'15pt'}}>: {range.start} ~ {range.end}</span>
							</div>
						))}
					</div>
				</div>
			</div>
		)
    }
}

export default ShowTimeCard;