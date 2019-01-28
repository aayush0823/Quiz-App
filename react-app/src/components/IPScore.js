import React,{Component} from 'react';
import { BrowserRouter as Router, Switch, Route, Link } from 'react-router-dom';

class IPScore extends Component{
	constructor()	{
		super();
		this.state = {
data: "",
		}
	}
	componentDidMount() {
		const request = new Request('http://127.0.0.1:8080/LoggedIn/politics/indian/score');
		fetch(request)
			.then( response => response.json())
			.then(data => this.setState({data: data}));
	}

	render(){
		var a='/LoggedIn/' +this.props.uname;
		return(
				<div className="App">
				<h1>Your Score: {this.state.data.score} </h1>
				<Link to={a}>HOME</Link>
				</div>
		      );
	}
}

export default IPScore;
