import React, { Component } from 'react';
import { NavLink, Link } from 'react-router-dom'; 

class Navigation extends Component {
	render() {
		return (
			<div>
				<nav className="navbar navbar-expand-lg navbar-dark bg-dark">
					<Link className="navbar-brand" style={{color:'pink'}} to="#">ToIoT</Link>
					<div className="container">
						<div className="navbar-collapse" id="navbarNavAltMarkup">
							<ul className="navbar-nav">
								<li className="nav-item active">
									<NavLink className="nav-item nav-link" to="/">
										HOME
									</NavLink>
								</li>
								<li className="nav-item dropdown">
									<NavLink className="nav-item nav-link dropdown-toggle" role="button" data-toggle="dropdown" to="/management">
										MANAGEMENT
									</NavLink>
									<div className="dropdown-menu" aria-labelledby="navbarDropdown">
										<Link className="dropdown-item" to="/sensor">SENSOR</Link>
										<Link className="dropdown-item" to="/node">NODE</Link>
									</div>
								</li>
								<li className="nav-item dropdown">
									<NavLink className="nav-item nav-link dropdown-toggle" role="button" data-toggle="dropdown" to="/management">
										SERVICE
									</NavLink>
									<div className="dropdown-menu" aria-labelledby="navbarDropdown">
										<Link className="dropdown-item" to="/alarm">ALARM</Link>
									</div>
								</li>
								<li className="nav-item">
									<NavLink className="nav-item nav-link" to="/kibana">
										KIBANA
									</NavLink>
								</li>
							</ul>
						</div>
					</div>
				</nav>
			</div>
		);
	}
}

export default Navigation;