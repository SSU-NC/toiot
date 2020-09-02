import React from 'react';
import Iframe from 'react-iframe';
import { KIBANA_DASHBOARDS_URL } from './defineUrl';

/* 
Home
- linked by HOME tab.
*/
function Home() {
	const dashboardUrl: string = KIBANA_DASHBOARDS_URL;

	return (
		<div
			className="embed-responsive embed-responsive-16by9"
			style={{
				position: 'absolute',
				width: '100%',
				height: '80%',
				left: '0px',
			}}
		>
			<Iframe
				url={dashboardUrl}
				className="embed-responsive-item"
				height="50%"
				scrolling="yes"
			></Iframe>
		</div>
	);
}

export default Home;
