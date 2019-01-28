import React, { Component } from 'react';
import Register from './Register';
import Login from './Login';
import Home from './Home';
import AdLogin from './AdLogin';
import UserHome from './UserHome'
import Admin from './Admin'
import ViewPeople from './ViewPeople'
import DeletePerson from './DeletePerson'
import Politics from './Politics'
import Anime from './Anime'
import DeathNote from './DeathNote'
import CodeGeass from './CodeGeass'
import IndianPolitics from './IndianPolitics'
import USPolitics from './USPolitics'
import DNScore from './DNScore'
import CGScore from './CGScore'
import IPScore from './IPScore'
import USScore from './USScore'
import BoardAnime from './BoardAnime'
import BoardPolitics from './BoardPolitics'


import { BrowserRouter as Router, Switch, Route, Link } from 'react-router-dom';

class App extends Component {
	render() {
		return (
				<div>
				<Router>
				<div>
				<nav className="navbar navbar-default">
				<div className="container-fluid">
				<div className="navbar-header">
				<Link className="navbar-brand" to={'/'}>Quiz UP</Link>
				</div>
				<ul className="nav navbar-nav">
				<li><Link to={'/login'}>Login</Link></li>
				<li><Link to={'/Register'}>Register</Link></li>
				<li><Link to={'/admin'}>Admin</Link></li>
				</ul>
				</div>
				</nav>
				<Switch>
				<Route exact path='/' component={Home} />
				<Route exact path='/login' component={Login} />
				<Route exact path='/Register' component={Register} />
				<Route exact path='/admin' component={AdLogin} />
				<Route exact path='/LoggedIn/:username' render={({match})=><UserHome uname={match.params.username} />} />
				<Route exact path='/Admin/:username' render={({match})=><Admin uname={match.params.username} />} />
				<Route exact path='/Admin/:username/view' render={({match})=><ViewPeople uname={match.params.username} />} />
				<Route exact path='/Admin/:username/delete' render={({match})=><DeletePerson uname={match.params.username} />} />
				<Route exact path='/LoggedIn/:username/anime' render={({match})=><Anime uname={match.params.username} />} />
				<Route exact path='/LoggedIn/:username/anime/deathNote' render={({match})=><DeathNote uname={match.params.username} />} />
				<Route exact path='/LoggedIn/:username/anime/codeGeass' render={({match})=><CodeGeass uname={match.params.username} />} />
				<Route exact path='/LoggedIn/:username/politics' render={({match})=><Politics uname={match.params.username} />} />
				<Route exact path='/LoggedIn/:username/politics/IndianPolitics' render={({match})=><IndianPolitics uname={match.params.username} />} />
				<Route exact path='/LoggedIn/:username/politics/USPolitics' render={({match})=><USPolitics uname={match.params.username} />} />
				<Route exact path='/LoggedIn/:username/anime/deathNote/score' render={({match})=><DNScore uname={match.params.username} />} />
				<Route exact path='/LoggedIn/:username/anime/codeGeass/score' render={({match})=><CGScore uname={match.params.username} />} />
				<Route exact path='/LoggedIn/:username/politics/indian/score' render={({match})=><IPScore uname={match.params.username} />} />
				<Route exact path='/LoggedIn/:username/politics/us/score' render={({match})=><USScore uname={match.params.username} />} />
				<Route exact path='/LoggedIn/:username/anime/leaderboard' render={({match})=><BoardAnime uname={match.params.username} />} />
				<Route exact path='/LoggedIn/:username/politics/leaderboard' render={({match})=><BoardPolitics uname={match.params.username} />} />
				<Route exact path='/adminpanel' component={()=> window.location='http://127.0.0.1:9000/admin'}/>
				</Switch>
				</div>
				</Router>
				</div>
				);
	}
}
export default App;
