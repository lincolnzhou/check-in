import React, {Component} from 'react';
import config from './config.json';

import style from "./greeter.less";

class Greeter extends Component {
	render() {
		return (
			<div class="sdui-container">
				<div className={style.root}>
					{config.greetText}
				</div>
			</div>
		);
	}
}

export default Greeter
