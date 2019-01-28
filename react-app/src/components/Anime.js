import React,{Component} from 'react';
import { BrowserRouter as Router, Switch, Route, Link } from 'react-router-dom';
var a;
var b;
class Anime extends Component{

	constructor() {
		super();
		this.state = {
data: [],
		}
	}
	componentDidMount() {
		const request = new Request('http://127.0.0.1:8080/LoggedIn/anime');
		fetch(request)
			.then( response => response.json())
			.then(data => this.setState({data: data}));
	}
	render(){
		a='/LoggedIn/' +this.props.uname + '/anime/deathNote';
		b='/LoggedIn/' +this.props.uname+ '/anime/codeGeass';
		return(
				<div class="App">
				<center>
				<header className="App-header">
				<h1 className="App-title">WELCOME {this.props.uname}</h1>
				</header>
				<h2>You Look To Be Pretty Fond Of Anime</h2>
				<br/><br/>
				<h4>SELECT YOUR ONE:</h4>
				<ul>
				<li>
				<Link to={a}>DEATH-NOTE -- YOUR MAX SCORE : {this.state.data.dn} </Link>
				</li>
				<li>
				<Link to={b}>CODE-GEASS -- YOUR MAX SCORE : {this.state.data.cg} </Link>
				</li>
				</ul>
				<br/><br/>
				</center>
				</div>
		      )
	}
}
export default Anime;
