import React, { Component } from 'react';
import RegisterSink from './Register/RegisterSink';
import SinkTable from './Table/SinkTable';
import { sinkListElem } from '../ElemInterface/ElementsInterface';

interface SinkManagementProps {
	sinkList: Array<sinkListElem>;
}

/*
SinkManagement
- Manage sink table, register sink
*/
const SinkManagement: React.FunctionComponent<SinkManagementProps> = (
	props
) => {
	return (
		<>
			<div style={{ float: 'right' }}>
				<button
					type="button"
					className="btn"
					data-toggle="modal"
					data-target="#register-sink-modal"
					style={{ background: 'pink' }}
				>
					register sink
				</button>
				<RegisterSink />
			</div>
			<div>
				<h3>Sink</h3>
				<SinkTable sinkList={props.sinkList} />
			</div>
		</>
	);
};

export default SinkManagement;
