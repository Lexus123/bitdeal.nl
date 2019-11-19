$(document).ready(function () {
	document.getElementById("eur").addEventListener("keyup", debouncedEur);
	document.getElementById("btc").addEventListener("keyup", debouncedBtc);

	var buyButton = document.getElementById("buy-button");
	var sellButton = document.getElementById("sell-button");

	buyButton.addEventListener("click", () => {
		if (buyButton.classList.contains("switch-button__left--non-active")) {
			buyButton.classList.toggle("switch-button__left--non-active");
			sellButton.classList.toggle("switch-button__right--non-active");
		}
	});

	sellButton.addEventListener("click", () => {
		if (sellButton.classList.contains("switch-button__right--non-active")) {
			buyButton.classList.toggle("switch-button__left--non-active");
			sellButton.classList.toggle("switch-button__right--non-active");
		}
	});

	var coll = document.getElementsByClassName("bitdeal");

	if (isViewportHandAndSmaller()) {
		for (i = 0; i < coll.length; i++) {
			coll[i].addEventListener("click", () => {
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

const MOBILE_WIDTH = 991;

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
	console.log(document.getElementById("eur").value);
}

function getBtcPrices() {
	console.log(document.getElementById("btc").value);
}

const debouncedEur = debounce(getEurPrices, 300);
const debouncedBtc = debounce(getBtcPrices, 300);