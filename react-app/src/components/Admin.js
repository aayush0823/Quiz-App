import React,{Component} from 'react';
import { BrowserRouter as Router, Switch, Route, Link } from 'react-router-dom';
var a;
var b;
class Admin extends Component{

	constructor()
	{
		super();
		this.state ={
random : ""
		}
	}

	render(){
		a='/Admin/' +this.props.uname + '/view';
		b='/Admin/' +this.props.uname+ '/delete';
		return(
				<div className="App">
				<center>
				<header className="App-header">
				<h1 className="App-title">WELCOME ADMIN  {this.props.uname}</h1>
				</header>
				<h2>YOU HAVE ALL EXCESS FOR THE APP</h2>
				<br/><br/>
				<h4>EXCERCISE YOUR AUTHORITIES:</h4>
				<ul>
				<li>
				<a href="/adminpanel">Admin Panel</a>
				</li>
				</ul>
				<br/><br/>
				</center>
				</div>
		      )
	}
}
export default Admin;
