import React,{Component} from 'react';
import { BrowserRouter as Router, Switch, Route, Link } from 'react-router-dom';
var a;
var b;
class Politics extends Component{

	constructor() {
		super();
		this.state = {
data: [],
		}
	}
	componentDidMount() {
		const request = new Request('http://127.0.0.1:8080/LoggedIn/politics');
		fetch(request)
			.then( response => response.json())
			.then(data => this.setState({data: data}));
	}
	render(){
		a='/LoggedIn/' +this.props.uname + '/politics/IndianPolitics';
		b='/LoggedIn/' +this.props.uname+ '/politics/USPolitics';
		return(
				<div class="App">
				<center>
				<header className="App-header">
				<h1 className="App-title">WELCOME {this.props.uname}</h1>
				</header>
				<h2>Politics seems to be your field of Interest</h2>
				<br/><br/>
				<h4>SELECT YOUR ONE:</h4>
				<ul>
				<li>
				<Link to={a}>INDIAN-POLITICS -- YOUR MAX SCORE:{this.state.data.ip}</Link>
				</li>
				<li>
				<Link to={b}>US-POLITICS -- YOUR MAX SCORE:{this.state.data.usp}</Link>
				</li>
				</ul>
				<br/><br/>
				</center>
				</div>
		      )
	}
}
export default Politics;
