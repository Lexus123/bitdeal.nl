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

function timeConverter(UNIX_timestamp) {
	var a = new Date(UNIX_timestamp * 1000);
	var months = ['Jan', 'Feb', 'Maa', 'Apr', 'Mei', 'Jun', 'Jul', 'Aug', 'Sep', 'Okt', 'Nov', 'Dec'];
	var year = a.getFullYear();
	var month = months[a.getMonth()];
	var date = a.getDate();
	var hour = a.getHours();
	var min = a.getMinutes();
	var sec = a.getSeconds();
	var time = date + ' ' + month + ' ' + year + ' ' + hour + ':' + min + ':' + sec;
	return time;
}

function createChart(data) {
	var scatterData = [];
	var mostRecentData = [];

	data.responsetimes.forEach(response => {
		var singleScatter = {
			name: `${response.requesttime}ms om ${timeConverter(response.created)}`,
			x: response.zscore,
			y: parseInt((response.cdf * 100).toFixed())
		};
		scatterData.push(singleScatter);
	})

	mostRecentData.push(scatterData[scatterData.length - 1]);

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
			startOnTick: false,
			min: -1.5,
			endOnTick: false,
			max: 3,
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
					headerFormat: `<b>{point.key}</b><br>`,
					pointFormat: '{point.x} sigma afwijkingen van het gemiddelde<br>{point.y} % van de requests is sneller'
				}
			}
		},
		series: [{
			name: `Requests (${data.count} totaal)`,
			color: 'rgba(223, 83, 83, .5)',
			data: scatterData,
		},
		{
			name: `Meest recente request`,
			color: 'rgba(70, 200, 83, 1)',
			data: mostRecentData,
		}]
	});
}