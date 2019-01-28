import React,{Component} from 'react';
import { Redirect } from 'react-router-dom';
class BoardPolitics extends Component{
	constructor()	{
		super();
		this.state = {
data: [],
		}
	}
	componentDidMount() {
		const request = new Request('http://127.0.0.1:8080/scoreboard/politics');
		fetch(request)
			.then( response => response.json())
			.then(data => this.setState({data: data}));
	}
	render(){
		return(
				<div className="App">
				<header className="App-header">
				<h1 className="App-title">LEADER-BOARD</h1>
				</header>
				<table className="table">
				<thead>
				<tr>
				<th> Username </th>
				<th>INDIAN POLITICS</th>
				<th>US POLITICS</th>
				</tr>
				</thead>
				<tbody>{this.state.data && this.state.data.map(function(item, key) {
						return (
								<tr key = {key}>
								<td><b>{item.username}</b></td>
								<td>{item.ip} </td>
								<td>{item.usp}</td>
								</tr>
						       )
						})}
		</tbody>
			</table>
			</div>
			);
	}
}

export default BoardPolitics;
