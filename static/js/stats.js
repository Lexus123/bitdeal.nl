document.addEventListener("DOMContentLoaded", function (event) {
	statsCall()
		.then((res) => res.json())
		.then((data) => createChart(data))
		.catch((err) => console.log(err))
});

async function statsCall() {
	return await fetch('/api/getstats', {
		method: 'GET',
		headers: new Headers(),
	});
}

function createChart(data) {
	var scatterData = [];

	data.responsetimes.forEach(response => {
		var singleScatter = {
			name: response.requesttime,
			x: response.zscore,
			y: response.cdf
		};
		scatterData.push(singleScatter);
	})

	Highcharts.chart('container', {
		chart: {
			type: 'scatter',
			zoomType: 'xy',
			backgroundColor: "rgba(0,0,0,0.0)",
		},
		accessibility: {
			description: 'Response tijd van /api/getprices'
		},
		title: {
			text: 'Response tijd van /api/getprices'
		},
		subtitle: {
			text: 'Bron: /api/getprices response tijden'
		},
		xAxis: {
			title: {
				enabled: true,
				text: 'Standaard afwijkingen'
			},
			startOnTick: true,
			endOnTick: true,
			showLastLabel: true
		},
		yAxis: {
			title: {
				text: 'Kans'
			}
		},
		legend: {
			layout: 'vertical',
			align: 'left',
			verticalAlign: 'top',
			x: 100,
			y: 70,
			floating: true,
			backgroundColor: Highcharts.defaultOptions.chart.backgroundColor,
			borderWidth: 1
		},
		plotOptions: {
			scatter: {
				marker: {
					radius: 5,
					states: {
						hover: {
							enabled: true,
							lineColor: 'rgb(100,100,100)'
						}
					}
				},
				states: {
					hover: {
						marker: {
							enabled: false
						}
					}
				},
				tooltip: {
					headerFormat: '<b>{point.key}ms</b><br>',
					pointFormat: '{point.x} sigma afwijkingen<br>{point.y} kans'
				}
			}
		},
		series: [{
			name: `Requests (${data.count} totaal)`,
			color: 'rgba(223, 83, 83, .5)',
			data: scatterData,
		}]
	});
}