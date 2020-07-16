import React, { Component } from 'react';
import { NavLink } from 'react-router-dom'; 

class Navigation extends Component {
	render() {
		return (
			<div>
				<nav className="navbar navbar-expand-lg navbar-violet fixed-top">
					<div className="container">
						<div className="navbar-collapse" id="navbarNavAltMarkup">
							<div className="navbar-nav">
								<NavLink className="nav-item nav-link" to="/">
									HOME
								</NavLink>
								<NavLink className="nav-item nav-link" to="/sensor">
									SENSOR
								</NavLink>
								<NavLink className="nav-item nav-link" to="/kibana">
									KIBANA
								</NavLink>
							</div>
						</div>
					</div>
				</nav>
			</div>

		);
	}
}

export default Navigation;