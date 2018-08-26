import React, {Component} from 'react';
import CalHeatMap from 'cal-heatmap';
import { getCheckList, postCheck } from "./server";

require("cal-heatmap/cal-heatmap.css");

class Calendar extends Component {
	constructor() {
		super();
	}

	componentWillMount() {
		getCheckList().then(function (response) {
			var cal = new CalHeatMap();
			var dt = new Date();
			cal.init({
				data: response.data.data,
				itemSelector: ".heatmap",
				domain: "month",
				subDomain: "x_day",
				subDomainTextFormat: "%d",
				cellSize: 25,
				range: 6,
				domainMargin: [0, 10, 0, 0],
				domainDynamicDimension: false,
				label: {
					position: "top"
				},
				start: new Date(2018, 6, 11),
				highlight: ["now", dt],
				tooltip: true,
			});
	  	})
		.catch(function (error) {
			console.log(error);
		});
	}

	componentDidMount() {
	}

	postCheck() {
		postCheck().then(function (response) {

		}).catch(function(error) {
			console.log(error)
		})
	}

	render() {
		return (
			<div class="sdui-container">
				<div class="operate">
					<button class="btn btn-primary" onClick={this.postCheck}>签到</button>
				</div>
				<div class="heatmap"></div>
			</div>
		);
	}
}

export default Calendar
