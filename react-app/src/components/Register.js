import React, { Component } from 'react';

class Register extends Component {
	constructor() {
		super();
		this.state = {
formData: {
username: "",
	  name: "",
	  password: "",
	  city: "",
	  },
submitted: 0,
		}
		this.handleNChange = this.handleNChange.bind(this);
		this.handlePChange = this.handlePChange.bind(this);
		this.handleUChange = this.handleUChange.bind(this);
		this.handleCChange = this.handleCChange.bind(this);
		this.handleSubmit = this.handleSubmit.bind(this);
	}

	handleSubmit (event) {
		event.preventDefault();
		fetch('http://localhost:8080/register', {
method: 'POST',
body: JSON.stringify(this.state.formData),
})
.then(response => {
		if(response.status >= 200 && response.status < 300)
		this.setState({submitted: 1});
		else
		this.setState({submitted: -1});
		});
}

handleNChange(event) {
	this.state.formData.name = event.target.value;
}
handleUChange(event) {
	this.state.formData.username = event.target.value;
}
handlePChange(event) {
	this.state.formData.password = event.target.value;
}
handleCChange(event) {
	this.state.formData.city = event.target.value;
}

render() {

	return (
			<div className="App">
			<header className="App-header">
			<h1 className="App-title">Register A New User</h1>
			</header>
			<br/><br/>
			<div className="formContainer">
			<form onSubmit={this.handleSubmit}>
			<div className="form-group">
			<label>Name</label>
			<input type="text" placeholder="Enter Name" className="form-control" value={this.state.name} onChange={this.handleNChange}/>
			</div>
			<div className="form-group">
			<label>Username</label>
			<input type="text" placeholder="Enter Username" className="form-control" value={this.state.username} onChange={this.handleUChange}/>
			</div>
			<div className="form-group">
			<label>Password</label>
			<input type="password" placeholder="Enter Password" className="form-control" value={this.state.password} onChange={this.handlePChange}/>
			</div>
			<div className="form-group">
			<label>City</label>
			<input type="text" className="form-control" placeholder="Enter City" value={this.state.city} onChange={this.handleCChange}/>
			</div>
			<button type="submit" className="btn btn-default">Submit</button>
			</form>
			</div>

			{this.state.submitted==1 &&
				<div>
					<h2>
					YOU HAVE BEEN REGISTERED SUCCESSFULLY.
					</h2>
					</div>
			}
	{this.state.submitted==-1 &&
		<div>
			<h2>
			USER ALREADY EXISTS.
			</h2>
			</div>
	}

	</div>
		);
}
}
export default Register;
