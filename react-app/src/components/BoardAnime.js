import React,{Component} from 'react';
import { Redirect } from 'react-router-dom';
class BoardAnime extends Component{
	constructor()	{
		super();
		this.state = {
data: [],
		}
	}
	componentDidMount() {
		const request = new Request('http://127.0.0.1:8080/scoreboard/anime');
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
				<th>Username</th>
				<th>DEATH-NOTE</th>
				<th>CODE-GEASS</th>
				</tr>
				</thead>
				<tbody>{this.state.data && this.state.data.map(function(item, key) {
						return (
								<tr key = {key}>
								<td><b>{item.username}</b></td>
								<td>{item.dn} </td>
								<td>{item.cg}</td>
								</tr>
						       )
						})}
		</tbody>
			</table>
			</div>
			);
	}
}

export default BoardAnime;
