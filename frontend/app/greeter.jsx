import React, {Component} from 'react';
import { getHitCount } from "./server";

import style from "./greeter.less";

class Greeter extends Component {
	constructor() {
		super();
		this.state = {
			hit: "",
		}
	}

	componentWillMount() {
		var that = this;
		getHitCount().then(function (response) {
			that.setState({
				hit: response.data.data,
			})
	  	})
		.catch(function (error) {
			console.log(error);
		});
	}

	render() {
		return (
			<div class="sdui-container">
				<div className={style.root}>
					<span>欢迎回来！总访问量：{this.state.hit}</span>
				</div>
			</div>
		);
	}
}

export default Greeter
