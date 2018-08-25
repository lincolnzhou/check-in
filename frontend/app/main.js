import React, {Component} from 'react';
import {render} from 'react-dom';
import Header from './header.jsx';
import Greeter from './greeter.jsx';
import Calendar from './calendar.jsx';

import './sdui.less';
import './app.less';

class App extends Component {
	render() {
		return (
				<div className="sdui-app">
					<Header />
					<div class="sdui-body">
						<Greeter />
						<Calendar />
					</div>
				</div>
		);
	}
}

render(<App />, document.getElementById("root"));
