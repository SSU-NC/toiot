import React from 'react';
import Iframe from 'react-iframe';
import { KIBANA_DASHBOARDS_URL } from './defineUrl';

/*
DashBoard
- link with KIBANA - Dashboard tab
*/
function DashBoard() {
	const dashboardUrl: string = KIBANA_DASHBOARDS_URL;

	return (
		<div className="embed-responsive embed-responsive-16by9">
			<Iframe
				url={dashboardUrl}
				className="embed-responsive-item"
				width="650"
				height="850"
				scrolling="yes"
			/>
		</div>
	);
}

export default DashBoard;
