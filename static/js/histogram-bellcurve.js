/*
 Highcharts JS v7.2.1 (2019-10-31)

 (c) 2010-2019 Highsoft AS
 Author: Sebastian Domas

 License: www.highcharts.com/license
*/
(function (a) { "object" === typeof module && module.exports ? (a["default"] = a, module.exports = a) : "function" === typeof define && define.amd ? define("highcharts/modules/histogram-bellcurve", ["highcharts"], function (c) { a(c); a.Highcharts = c; return a }) : a("undefined" !== typeof Highcharts ? Highcharts : void 0) })(function (a) {
	function c(a, b, c, f) { a.hasOwnProperty(b) || (a[b] = f.apply(null, c)) } a = a ? a._modules : {}; c(a, "mixins/derived-series.js", [a["parts/Globals.js"], a["parts/Utilities.js"]], function (a, b) {
		var h = b.defined, c = a.Series,
			g = a.addEvent; return {
				hasDerivedData: !0, init: function () { c.prototype.init.apply(this, arguments); this.initialised = !1; this.baseSeries = null; this.eventRemovers = []; this.addEvents() }, setDerivedData: a.noop, setBaseSeries: function () { var a = this.chart, b = this.options.baseSeries; this.baseSeries = h(b) && (a.series[b] || a.get(b)) || null }, addEvents: function () {
					var a = this; var b = g(this.chart, "afterLinkSeries", function () {
						a.setBaseSeries(); a.baseSeries && !a.initialised && (a.setDerivedData(), a.addBaseSeriesEvents(), a.initialised =
							!0)
					}); this.eventRemovers.push(b)
				}, addBaseSeriesEvents: function () { var a = this; var b = g(a.baseSeries, "updatedData", function () { a.setDerivedData() }); var h = g(a.baseSeries, "destroy", function () { a.baseSeries = null; a.initialised = !1 }); a.eventRemovers.push(b, h) }, destroy: function () { this.eventRemovers.forEach(function (a) { a() }); c.prototype.destroy.apply(this, arguments) }
			}
	}); c(a, "modules/histogram.src.js", [a["parts/Globals.js"], a["parts/Utilities.js"], a["mixins/derived-series.js"]], function (a, b, c) {
		function h(a) {
			return function (k) {
				for (var d =
					1; a[d] <= k;)d++; return a[--d]
			}
		} var g = b.arrayMax, r = b.arrayMin, m = b.isNumber, n = b.objectEach; b = a.seriesType; var d = a.correctFloat; a = a.merge; var e = { "square-root": function (a) { return Math.ceil(Math.sqrt(a.options.data.length)) }, sturges: function (a) { return Math.ceil(Math.log(a.options.data.length) * Math.LOG2E) }, rice: function (a) { return Math.ceil(2 * Math.pow(a.options.data.length, 1 / 3)) } }; b("histogram", "column", {
			binsNumber: "square-root", binWidth: void 0, pointPadding: 0, groupPadding: 0, grouping: !1, pointPlacement: "between",
			tooltip: { headerFormat: "", pointFormat: '<span style="font-size: 10px">{point.x} - {point.x2}</span><br/><span style="color:{point.color}">\u25cf</span> {series.name} <b>{point.y}</b><br/>' }
		}, a(c, {
			setDerivedData: function () { var a = this.baseSeries.yData; a.length && (a = this.derivedData(a, this.binsNumber(), this.options.binWidth), this.setData(a, !1)) }, derivedData: function (a, b, e) {
				var k = g(a), c = d(r(a)), p = [], l = {}, f = []; e = this.binWidth = this.options.pointRange = d(m(e) ? e || 1 : (k - c) / b); for (b = c; b < k && (this.userOptions.binWidth ||
					d(k - b) >= e || 0 >= d(c + p.length * e - b)); b = d(b + e))p.push(b), l[b] = 0; 0 !== l[c] && (p.push(d(c)), l[d(c)] = 0); var q = h(p.map(function (a) { return parseFloat(a) })); a.forEach(function (a) { a = d(q(a)); l[a]++ }); n(l, function (a, b) { f.push({ x: Number(b), y: a, x2: d(Number(b) + e) }) }); f.sort(function (a, b) { return a.x - b.x }); return f
			}, binsNumber: function () { var a = this.options.binsNumber, b = e[a] || "function" === typeof a && a; return Math.ceil(b && b(this.baseSeries) || (m(a) ? a : e["square-root"](this.baseSeries))) }
		})); ""
	}); c(a, "modules/bellcurve.src.js",
		[a["parts/Globals.js"], a["parts/Utilities.js"], a["mixins/derived-series.js"]], function (a, b, c) {
			function f(a) { var b = a.length; a = a.reduce(function (a, b) { return a + b }, 0); return 0 < b && a / b } function g(a, b) { var c = a.length; b = m(b) ? b : f(a); a = a.reduce(function (a, c) { c -= b; return a + c * c }, 0); return 1 < c && Math.sqrt(a / (c - 1)) } function h(a, b, c) { a -= b; return Math.exp(-(a * a) / (2 * c * c)) / (c * Math.sqrt(2 * Math.PI)) } var m = b.isNumber; b = a.seriesType; var n = a.correctFloat; a = a.merge; b("bellcurve", "areaspline", {
				intervals: 3, pointsInInterval: 3,
				marker: { enabled: !1 }
			}, a(c, {
				setMean: function () { this.mean = n(f(this.baseSeries.yData)) }, setStandardDeviation: function () { this.standardDeviation = n(g(this.baseSeries.yData, this.mean)) }, setDerivedData: function () { 1 < this.baseSeries.yData.length && (this.setMean(), this.setStandardDeviation(), this.setData(this.derivedData(this.mean, this.standardDeviation), !1)) }, derivedData: function (a, b) {
					var c = this.options.intervals, d = this.options.pointsInInterval, e = a - c * b; c = c * d * 2 + 1; d = b / d; var f = [], g; for (g = 0; g < c; g++)f.push([e, h(e,
						a, b)]), e += d; return f
				}
			})); ""
		}); c(a, "masters/modules/histogram-bellcurve.src.js", [], function () { })
});
//# sourceMappingURL=histogram-bellcurve.js.map