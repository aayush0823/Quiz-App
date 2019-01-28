import React, { Component } from 'react';
import { Redirect } from 'react-router-dom';
import Register from './Register';
import { BrowserRouter as Router, Switch, Route, Link } from 'react-router-dom';

var a
class ALogin extends Component {
	constructor() {
		super();
		this.state = {
formData : {
username: "",
	  password: "",
	   },
submitted: false,
		}
		this.handleUChange = this.handleUChange.bind(this);
		this.handlePChange = this.handlePChange.bind(this);
		this.handleSubmit = this.handleSubmit.bind(this);
	}

	handleSubmit(event) {
		event.preventDefault();
		fetch('http://localhost:8080/login', {
method: 'POST',
body: JSON.stringify(this.state.formData)
})
.then(response => {
		if(response.status >= 200 && response.status < 300)
		this.setState({submitted:true});
		});
}

handleUChange(event) {
	this.state.formData.username = event.target.value;
}
handlePChange(event) {
	this.state.formData.password = event.target.value;
}
render() {
	if (this.state.submitted){
		if (this.state.formData.username === "ishu" && this.state.formData.password === "ishu"){
			a="Admin/"+this.state.formData.username
				return (<Redirect to={a}/>)
		}
		else
		{
			a="LoggedIn/"+this.state.formData.username
				return (<Redirect to={a}/>)
		}
	}
	return (
			<div className="App">
			<header className="App-header" >
			<h1 className="App-title">ADMIN LOGIN</h1>
			</header>
			<br /><br /><br />
			<div className="formContainer">
			<form onSubmit={this.handleSubmit}>
			<div className="form-group">
			<label>Username</label>
			<input type="text" placeholder="Enter Username" className="form-control" value={this.state.username} onChange={this.handleUChange}/>
			</div>
			<div className="form-group">
			<label>Password</label>
			<input type="password" className="form-control" placeholder="Enter Password" value={this.state.password} onChange={this.handlePChange}/>
			</div>
			<div class="checkbox">
			<label><input type="checkbox"/> Remember me</label>
			</div>
			<button type="submit" className="btn btn-default">Submit</button>
			</form>
			</div>
			</div>
			);
}
}

export default ALogin;
