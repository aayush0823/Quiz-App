import React, { Component } from 'react';
import './Home.css';
import { BrowserRouter as Router, Switch, Route, Link } from 'react-router-dom';
class Home extends Component {
	render() {
		return (
				<div className="App">
				<header className="App-header">
				<h1 className="App-title">Welcome to My Quiz App</h1>
				</header>
				<h2>Let's See How much can You Answer</h2>
				<br/><br/><br/><br/><br/><br/>
				</div>
		       );
	}
}
export default Home;
