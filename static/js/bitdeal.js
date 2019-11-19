$(document).ready(function () {
	document.getElementById("eur").addEventListener("keyup", debouncedEur);
	document.getElementById("btc").addEventListener("keyup", debouncedBtc);

	var buyButton = document.getElementById("buy-button");
	var sellButton = document.getElementById("sell-button");
	var eurLabel = document.getElementById("eur-label");
	var btcLabel = document.getElementById("btc-label");

	buyButton.addEventListener("click", () => {
		if (buyButton.classList.contains("switch-button__left--non-active")) {
			buyButton.classList.toggle("switch-button__left--non-active");
			sellButton.classList.toggle("switch-button__right--non-active");
			eurLabel.innerHTML = "Betaal";
			btcLabel.innerHTML = "Ontvang";
			calculatorState = "buy";
		}
	});

	sellButton.addEventListener("click", () => {
		if (sellButton.classList.contains("switch-button__right--non-active")) {
			buyButton.classList.toggle("switch-button__left--non-active");
			sellButton.classList.toggle("switch-button__right--non-active");
			eurLabel.innerHTML = "Ontvang";
			btcLabel.innerHTML = "Betaal";
			calculatorState = "sell";
		}
	});

	var coll = document.getElementsByClassName("bitdeal");
	var i;

	if (isViewportHandAndSmaller()) {
		for (i = 0; i < coll.length; i++) {
			coll[i].addEventListener("click", function () {
				var content = this.children;
				for (c = 0; c < content.length; c++) {
					if (content[c].classList.contains("content")) {
						if (content[c].style.display === "block") {
							content[c].style.display = "none";
						} else {
							content[c].style.display = "block";
						}
					};

					if (content[c].classList.contains("arrow")) {
						content[c].classList.toggle("active-arrow");
					};
				}
			});
		}

		for (c = 0; c < coll[1].children.length; c++) {
			if (coll[1].children[c].classList.contains("content")) {
				coll[1].children[c].style.display = "block";
			};

			if (coll[1].children[c].classList.contains("arrow")) {
				coll[1].children[c].classList.add("active-arrow");
			};
		}
	}
});

// window.onload = function () {
// 	if (window.jQuery) {
// 		// jQuery is loaded  
// 		console.log("Yeah!");
// 	} else {
// 		// jQuery is not loaded
// 		console.log("Doesn't Work");
// 	}
// }

var MOBILE_WIDTH = 991;
var calculatorState = "buy";

function getCalculatorState() {
	return calculatorState;
}

function getViewportWidth() {
	return Math.max(document.documentElement.clientWidth, window.innerWidth || 0);
}

function isViewportHandAndSmaller() {
	return getViewportWidth() <= MOBILE_WIDTH;
}

function debounce(func, wait = 100) {
	let timeout;
	return function (...args) {
		clearTimeout(timeout);
		timeout = setTimeout(() => {
			func.apply(this, args);
		}, wait);
	};
}

function getEurPrices() {
	if (getCalculatorState() === "buy") {
		bitdealCall("buy", "eur", document.getElementById("eur").value);
		console.log("Koop zoveel EUR:")
		console.log(document.getElementById("eur").value);
	} else {
		bitdealCall("sell", "eur", document.getElementById("eur").value);
		console.log("Verkoop zoveel EUR:")
		console.log(document.getElementById("eur").value);
	}
}

function getBtcPrices() {
	if (getCalculatorState() === "buy") {
		bitdealCall("buy", "btc", document.getElementById("btc").value);
		console.log("Koop zoveel BTC:")
		console.log(document.getElementById("btc").value);
	} else {
		bitdealCall("sell", "btc", document.getElementById("btc").value);
		console.log("Verkoop zoveel BTC:")
		console.log(document.getElementById("btc").value);
	}
}

function bitdealCall(type, currency, amount) {
	$.ajax({
		url: "/api/getprices",
		type: "POST",
		contentType: "application/json; charset=utf-8",
		dataType: "json",
		data: JSON.stringify({
			"type": type,
			"currency": currency,
			"amount": amount
		}),
		success: function (response) {
			console.log(response);
		},
		error: function (jqXHR, textStatus, errorThrown) {
			console.log(textStatus, errorThrown);
		}
	});
}

const debouncedEur = debounce(getEurPrices, 300);
const debouncedBtc = debounce(getBtcPrices, 300);
