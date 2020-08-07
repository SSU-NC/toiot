import React from 'react';
import Iframe from 'react-iframe';
import { KIBANA_URL } from './defineUrl';

function Main() {
	const kibanaUrl: string = KIBANA_URL;

	return (
		<div className="embed-responsive embed-responsive-16by9"
		style={{position:"absolute", width:"50%" , height:"50%", left :"0px"}} >	
			<Iframe url={kibanaUrl} className="embed-responsive-item"  
			 
			scrolling="yes" >
			</Iframe>
		</div>
		
	);
}

export default Main;
