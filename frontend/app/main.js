import React from 'react';
import {render} from 'react-dom';
import Greeter from './greeter';

import './sdui.less';
import './app.less';
require("cal-heatmap/cal-heatmap.css");

import CalHeatMap from 'cal-heatmap';

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
render(<Greeter />, document.getElementById('greeter'));
