import React, {Component} from 'react';
import axios from "axios";
import CalHeatMap from 'cal-heatmap';

require("cal-heatmap/cal-heatmap.css");

class Calendar extends Component {
	constructor() {
		super();
	}

	componentWillMount() {
		axios.get('/api/check')
		.then(function (response) {
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

	render() {
		return (
			<div class="sdui-container heatmap">
			</div>
		);
	}
}

export default Calendar
