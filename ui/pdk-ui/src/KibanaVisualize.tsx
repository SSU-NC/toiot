import React from 'react';
import Iframe from 'react-iframe';
import { KIBANA_VISUALIZE_URL } from './defineUrl';

function Visualize() {
	const visualizeUrl: string = KIBANA_VISUALIZE_URL;

	return (
		<div className="embed-responsive embed-responsive-16by9">
			<Iframe url={visualizeUrl} className="embed-responsive-item"  width="650" height = "850"
			scrolling = "yes"
			/>
		</div>
	);
}

export default Visualize;
