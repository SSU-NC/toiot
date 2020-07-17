import React, { Component } from 'react';
import { NavLink, Link } from 'react-router-dom'; 

class Navigation extends Component {
	render() {
		return (
			<div>
				<nav className="navbar navbar-expand-lg navbar-violet fixed-top">
					<div className="container">
						<div className="navbar-collapse" id="navbarNavAltMarkup">
							<ul className="navbar-nav">
								<li class="nav-item active">
									<NavLink className="nav-item nav-link" to="/">
										HOME
									</NavLink>
								</li>
								<li class="nav-item dropdown">
									<NavLink className="nav-item nav-link dropdown-toggle" role="button" data-toggle="dropdown" to="/management">
										MANAGEMENT
									</NavLink>
									<div class="dropdown-menu" aria-labelledby="navbarDropdown">
										<Link class="dropdown-item" to="/sensor">SENSOR</Link>
										<Link class="dropdown-item" to="/node">NODE</Link>
									</div>
								</li>
								<li class="nav-item">
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