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
	console.log(data);

	var responses = [];

	data.responsetimes.forEach((response) => {
		responses.push(response.requesttime);
	})

	var myChart = Highcharts.chart('container', {
		title: {
			text: '/api/getprices response tijd distributie'
		},

		chart: {
			backgroundColor: "rgba(0,0,0,0.0)",
		},

		xAxis: [{
			title: {
				text: 'Data punt (#)'
			},
			alignTicks: true
		}, {
			title: {
				text: 'Response tijd (ms)'
			},
			alignTicks: true,
			opposite: true
		}],

		yAxis: [{
			title: { text: 'Response tijd' }
		}, {
			title: { text: 'Bell curve' },
			opposite: true
		}],

		series: [{
			name: 'Bell curve',
			type: 'bellcurve',
			xAxis: 1,
			yAxis: 1,
			baseSeries: 1,
			zIndex: -1
		}, {
			name: 'Request',
			type: 'scatter',
			data: responses,
			marker: {
				radius: 1.5
			}
		}]
	});
}