$(document).ready(function () {
	var eurInput = document.getElementById("eur");
	var btcInput = document.getElementById("btc");
	var buyButton = document.getElementById("buy-button");
	var sellButton = document.getElementById("sell-button");
	var eurLabel = document.getElementById("eur-label");
	var btcLabel = document.getElementById("btc-label");

	eurInput.addEventListener("keyup", debouncedEur);
	btcInput.addEventListener("keyup", debouncedBtc);

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
	} else {
		bitdealCall("sell", "eur", document.getElementById("eur").value);
	}
}

function getBtcPrices() {
	if (getCalculatorState() === "buy") {
		bitdealCall("buy", "btc", document.getElementById("btc").value);
	} else {
		bitdealCall("sell", "btc", document.getElementById("btc").value);
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
			applyNewPrices(response);
		},
		error: function (jqXHR, textStatus, errorThrown) {
			console.log(textStatus, errorThrown);
		}
	});
}

function applyNewPrices(priceData) {
	console.log(priceData);

	if (priceData.currency === "eur") {
		document.getElementById("btc").value = priceData.bestamount;
	} else {
		document.getElementById("eur").value = priceData.bestamount;
	}

	for (exchangeIndex = 0; exchangeIndex < priceData.exchangerates.length; exchangeIndex++) {
		var bitdeal = document.getElementById(priceData.exchangerates[exchangeIndex].exchange);
		var bitdealChildren = bitdeal.children;

		// DESKTOP
		// Company name update
		var bitdealName = bitdealChildren[0].children[0];
		bitdealName.innerHTML = priceData.exchangerates[exchangeIndex].exchange;

		// Badge update
		var bitdealBadge = bitdealChildren[1].children[0];
		bitdealBadge.classList.remove("desc__pill--deal");
		bitdealBadge.classList.remove("desc__pill--review");
		bitdealBadge.innerHTML = "";
		if (priceData.mostreviews === priceData.exchangerates[exchangeIndex].reviews) {
			bitdealBadge.innerHTML = "Meeste reviews!";
			bitdealBadge.classList.add("desc__pill--review");
		}
		if (priceData.bestamount === priceData.exchangerates[exchangeIndex].amount) {
			bitdealBadge.innerHTML = "Beste deal!";
			bitdealBadge.classList.add("desc__pill--deal");
		}

		// Amount update
		var bitdealAmount = bitdealChildren[3].children[0];
		bitdealAmount.innerHTML = "₿ " + priceData.exchangerates[exchangeIndex].amount;

		// Rate update
		var bitdealRate = bitdealChildren[3].children[1];
		bitdealRate.innerHTML = "(€ " + priceData.exchangerates[exchangeIndex].rate + " / BTC)";

		// Link update

		console.log(bitdealChildren[1].children);
	}
}

const debouncedEur = debounce(getEurPrices, 300);
const debouncedBtc = debounce(getBtcPrices, 300);
