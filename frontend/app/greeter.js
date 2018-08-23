import React, {Component} from 'react';
import config from './config.json';

import style from "./greeter.css";

class Greeter extends Component {
	render() {
		return (
				<div className={style.root}>
					{config.greetText}
				</div>
				);
	}
}

export default Greeter
