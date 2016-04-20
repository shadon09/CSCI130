var el = document.querySelector("#enter-area")
el.addEventListener("input", function(e){
	var xhr = new XMLHttpRequest();
	xhr.open("post", "/api/check");
	console.log(e.target.value)
	xhr.send(e.target.value)

	xhr.AddEventListener("readystatechange", function(){
		if(xhr.readyState === 4 && xhr.status())
	})
})
