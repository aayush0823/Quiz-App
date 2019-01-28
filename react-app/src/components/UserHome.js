import React,{Component} from 'react';
import { BrowserRouter as Router, Switch, Route, Link } from 'react-router-dom';
var a;
var b;
class UserHome extends Component{

	constructor()
	{
		super();
		this.state ={
random : ""
		}
	}

	render(){
		a='/LoggedIn/' +this.props.uname + '/anime';
		var la='/LoggedIn/' +this.props.uname + '/anime/leaderboard';
		b='/LoggedIn/' +this.props.uname+ '/politics';
		var lb='/LoggedIn/' +this.props.uname + '/politics/leaderboard';
		return(
				<div class="App">
				<center>
				<header className="App-header">
				<h1 className="App-title">WELCOME {this.props.uname}</h1>
				</header>
				<h2>Let's Get Things Started</h2>
				<br/><br/>
				<h4>SELECT YOUR INTEREST:</h4>
				<ul>
				<li>
				<Link to={a}>ANIME</Link>
				</li>
				<li>
				<Link to={b}>POLITICS</Link>
				</li>
				</ul>
				<br/><br/>
				<h4>LEADER-BOARD</h4>
				<ul>
				<li>
				<Link to={la}>ANIME</Link>
				</li>
				<li>
				<Link to={lb}>POLITICS</Link>
				</li>
				</ul>
				</center>
				</div>
				)
	}
}
export default UserHome;
