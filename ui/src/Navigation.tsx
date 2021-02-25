import React, { Component } from 'react';
import { NavLink, Link } from 'react-router-dom';

/* 
Navigation
- Navigation bar 
*/
class Navigation extends Component {
	render() {
		return (
			<>
				<nav className="navbar navbar-expand-lg navbar-dark bg-dark">
					<Link
						className="navbar-brand"
						style={{ fontSize: '20pt', fontWeight: 'bold', color: 'pink' }}
						to="/"
					>
						ToIoT
					</Link>
					<div className="container">
						<div className="navbar-collapse" id="navbarNavAltMarkup">
							<ul className="navbar-nav">
								<li className="nav-item active">
									<NavLink className="nav-item nav-link" to="/">
										HOME
									</NavLink>
								</li>
								<li className="nav-item dropdown">
									<NavLink
										className="nav-item nav-link dropdown-toggle"
										role="button"
										data-toggle="dropdown"
										to="/management"
									>
										MANAGEMENT
									</NavLink>
									<div
										className="dropdown-menu"
										aria-labelledby="navbarDropdown"
										style={{ background: 'pink' }}
									>
										<Link
											className="dropdown-item"
											to="/sensor"
											style={{ background: 'pink' }}
										>
											Sensor
										</Link>
										<Link
											className="dropdown-item"
											to="/node"
											style={{ background: 'pink' }}
										>
											Node
										</Link>
										<Link
											className="dropdown-item"
											to="/sink"
											style={{ background: 'pink' }}
										>
											Sink
										</Link>
										<Link
											className="dropdown-item"
											to="/actuator"
											style={{ background: 'pink' }}
										>
											Actuator
										</Link>
									</div>
								</li>
								<li className="nav-item dropdown">
									<NavLink
										className="nav-item nav-link dropdown-toggle"
										role="button"
										data-toggle="dropdown"
										to="/kafka"
									>
										KAFKA
									</NavLink>
									<div
										className="dropdown-menu"
										aria-labelledby="navbarDropdown"
										style={{ background: 'pink' }}
									>
										<Link
											className="dropdown-item"
											to="/topic"
											style={{ background: 'pink' }}
										>
											Topic
										</Link>
										{/* <Link
											className="dropdown-item"
											to="/logicService"
											style={{ background: 'pink' }}
										>
											Logic Service
										</Link> */}
									</div>
								</li>
								<li className="nav-item dropdown">
									<NavLink
										className="nav-item nav-link dropdown-toggle"
										role="button"
										data-toggle="dropdown"
										to="/management"
									>
										SERVICE
									</NavLink>
									<div
										className="dropdown-menu"
										aria-labelledby="navbarDropdown"
										style={{ background: 'pink' }}
									>
										<Link
											className="dropdown-item"
											to="/logicCore"
											style={{ background: 'pink' }}
										>
											Logic core
										</Link>
									</div>
								</li>
								<li className="nav-item dropdown">
									<NavLink
										className="nav-item nav-link dropdown-toggle"
										role="button"
										data-toggle="dropdown"
										to="/kibana"
									>
										KIBANA
									</NavLink>
									<div
										className="dropdown-menu"
										aria-labelledby="navbarDropdown"
										style={{ background: 'pink' }}
									>
										<Link
											className="dropdown-item"
											to="/visualize"
											style={{ background: 'pink' }}
										>
											Visualize
										</Link>

										<Link
											className="dropdown-item"
											to="/dashboard"
											style={{ background: 'pink' }}
										>
											Dashboard
										</Link>
									</div>
								</li>
							</ul>
						</div>
					</div>
				</nav>
			</>
		);
	}
}

export default Navigation;
