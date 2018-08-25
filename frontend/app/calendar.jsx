import React, {Component} from 'react';
import CalHeatMap from 'cal-heatmap';

require("cal-heatmap/cal-heatmap.css");

class Calendar extends Component {
	componentWillMount() {
		console.log("componentWillMount")
	}

	componentDidMount() {
		console.log("componentDidMount")
		var datas = '{"1534694400": 15,"1534867200": 25,"1535472000": 10}';
		var cal = new CalHeatMap();
		var dt = new Date();
		dt.setDate(dt.getDate() + 1);
		cal.init({
			data: JSON.parse(datas),
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
	}

	render() {
		return (
			<div class="sdui-container heatmap">
			</div>
		);
	}
}

export default Calendar
