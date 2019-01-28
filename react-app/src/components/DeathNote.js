import React, { Component } from 'react';
import { BrowserRouter as Router, Switch, Route, Link } from 'react-router-dom';
import { Redirect } from 'react-router-dom';
class DeathNote extends Component {
	constructor() {
		super();
		this.state = {
data: [],
      formData : {
ans1: "",
      ans2: "",
      ans3: "",
      ans4: "",
      ans5: "",
      ans6: "",
      ans7: "",
      ans8: "",
      },
submitted: false,
		}
		this.handleA1Change = this.handleA1Change.bind(this);
		this.handleA2Change = this.handleA2Change.bind(this);
		this.handleA3Change = this.handleA3Change.bind(this);
		this.handleA4Change = this.handleA4Change.bind(this);
		this.handleA5Change = this.handleA5Change.bind(this);
		this.handleA6Change = this.handleA6Change.bind(this);
		this.handleA7Change = this.handleA7Change.bind(this);
		this.handleA8Change = this.handleA8Change.bind(this);
		this.handleSubmit = this.handleSubmit.bind(this);
	}
	handleA1Change(event) {
		this.state.formData.ans1 = event.target.value;
	}

	handleA2Change(event) {
		this.state.formData.ans2 = event.target.value;
	}

	handleA3Change(event) {
		this.state.formData.ans3 = event.target.value;
	}

	handleA4Change(event) {
		this.state.formData.ans4 = event.target.value;
	}

	handleA5Change(event) {
		this.state.formData.ans5 = event.target.value;
	}

	handleA6Change(event) {
		this.state.formData.ans6 = event.target.value;
	}

	handleA7Change(event) {
		this.state.formData.ans7 = event.target.value;
	}

	handleA8Change(event) {
		this.state.formData.ans8 = event.target.value;
	}

	handleSubmit(event) {
		event.preventDefault();
		fetch('http://localhost:8080/LoggedIn/anime/deathnote/score', {
method: 'POST',
body: JSON.stringify(this.state.formData)
})
.then(response => {
		if(response.status >= 200 && response.status < 300)
		this.setState({submitted:true});
		});
}
componentDidMount() {
	const request = new Request('http://127.0.0.1:8080/LoggedIn/anime/deathnote');
	fetch(request)
		.then( response => response.json())
		.then(data => this.setState({data: data}));
}
render() {
	if (this.state.submitted){
		var a="/LoggedIn/" + this.props.uname + "/anime/deathNote/score"
			return (<Redirect to={a}/>)
	}
	return (
			<div className="App">
			<header className="App-header">
			<h1 className="App-title">Welcome to The World Of Death-Note</h1>
			</header>
			<h2>Let's See How much have You Followed</h2>
			<div className="App">
			<header className="App-header">
			<h1 className="App-title">Questions</h1>
			</header>
			<h4><b>First 4-Single Correct & Last 4-Multiple Choice(Options To Be Space Separeted)</b></h4>
			<table className="table">
			<thead>
			<tr>
			<th>Question</th>
			<th>Option-1</th>
			<th>Option-2</th>
			<th>Option-3</th>
			</tr>
			</thead>
			<tbody>{this.state.data && this.state.data.map(function(item, key) {
					return (
							<tr key = {key}>
							<td><b>{item.question}</b></td>
							<td>{item.opt1} </td>
							<td>{item.opt2}</td>
							<td>{item.opt3}</td>
							</tr>
					       )
					})}
	</tbody>
		</table><br/><br/>
		<div className="formContainer">
		<form onSubmit={this.handleSubmit}>
		<div className="form-group">
		<label>Question-1</label>
		<input type="text" className="form-control" onChange={this.handleA1Change}/>
		</div>
		<div className="form-group">
		<label>Question-2</label>
		<input type="text" className="form-control" onChange={this.handleA2Change}/>
		</div>
		<div className="form-group">
		<label>Question-3</label>
		<input type="text" className="form-control" onChange={this.handleA3Change}/>
		</div>
		<div className="form-group">
		<label>Question-4</label>
		<input type="text" className="form-control" onChange={this.handleA4Change}/>
		</div>
		<div className="form-group">
		<label>Question-5</label>
		<input type="text" className="form-control" onChange={this.handleA5Change}/>
		</div>
		<div className="form-group">
		<label>Question-6</label>
		<input type="text" className="form-control" onChange={this.handleA6Change}/>
		</div>
		<div className="form-group">
		<label>Question-7</label>
		<input type="text" className="form-control" onChange={this.handleA7Change}/>
		</div>
		<div className="form-group">
		<label>Question-8</label>
		<input type="text" className="form-control" onChange={this.handleA8Change}/>
		</div>
		<button type="submit" className="btn btn-default">Submit</button>
		</form>
		</div>
		</div>
		</div>
		);
}
}
export default DeathNote;
