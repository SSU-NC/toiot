import React from 'react';
import Iframe from 'react-iframe'
// 220.70.2.1:5601
function Kibana() {
	var kibanaUrl = "http://220.70.2.1:5601";

	return (
		<div className="embed-responsive embed-responsive-16by9">
			<Iframe url={kibanaUrl}
				className="embed-responsive-item"
				frameborder="0" 
				align="middle"
			/>
		</div>
	);
}

export default Kibana;
