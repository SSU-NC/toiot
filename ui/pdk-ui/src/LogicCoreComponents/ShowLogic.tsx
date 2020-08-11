import React, { Component } from 'react';
import { logicCoreElem } from '../ElemInterface/LcElementsInterface';
// form : https://getbootstrap.com/docs/4.0/components/forms/?
// add, delete input : https://codesandbox.io/s/00xq32n3pn?from-embed=&file=/src/index.js

interface ShowLogicProps {
    logic: logicCoreElem;
}

class ShowLogic extends Component<ShowLogicProps, {}> {
    render() {
        return (
            <>
            <button 
                type="button" 
                className="btn" 
                data-toggle="modal" 
                style={{ background: 'pink' }}
                data-target="#show-logic-modal"
            >
                show logic
            </button>
            <div className="modal fade" id="show-logic-modal" role="dialog" aria-labelledby="show-logic-modal">
                <div className="modal-dialog modal-lg" role="document">
                    <div className="modal-content">
                        <div className="modal-header"> 
                        <h4 className="modal-title" id="show-logic-modal">{this.props.logic.logic_name}</h4>
                            <button 
                                type="button" 
                                className="close" 
                                data-dismiss="modal" 
                                aria-label="Close"
                            >
                                <span aria-hidden="true">Ã—</span>
                            </button>
                        </div>
                    </div>
                </div>
            </div>
            </>
        );
    }
}

export default ShowLogic;