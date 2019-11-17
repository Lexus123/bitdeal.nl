$(document).ready(function () {
	var coll = document.getElementsByClassName("bitdeal");
	var i;

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
			}
		});
	}
});

window.onload = function () {
	if (window.jQuery) {
		// jQuery is loaded  
		console.log("Yeah!");
	} else {
		// jQuery is not loaded
		console.log("Doesn't Work");
	}
}