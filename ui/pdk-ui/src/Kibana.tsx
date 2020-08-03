import React from 'react';
import Iframe from 'react-iframe';
import { KIBANA_URL } from './defineUrl';

function Kibana() {
	const kibanaUrl: string = KIBANA_URL;

	return (
		<div className="embed-responsive embed-responsive-16by9">
			<Iframe url={kibanaUrl} className="embed-responsive-item" />
		</div>
	);
}

export default Kibana;
