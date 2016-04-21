var list = document.createElement('ul');
for(var i=0; i<4; i++){
	var item = document.createElement('li');
        item.appendChild(document.createTextNode(i));
        list.appendChild(item);
}
document.getElementById('my-list').appendChild(list);

var el = document.getElementById('my-list');

el.addEventListener('click', function(e){
	if(e.target.tagName === 'LI'){
		alert('Dont touch this');
	}
});
