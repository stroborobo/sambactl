/**
 * main.js
 */

function submitForm(event) {
	event.preventDefault();
	event.stopPropagation();

	$.post('/user/', {
		username: document.user.username.value,
		password: document.user.password.value
	}, onSuccessSubmit);

	return false;
}

function onSuccessSubmit() {
	document.location = '/';
}

function resetForm(event) {
	document.user.username.disabled = false;
}

function editUser(event) {
	event.preventDefault();
	event.stopPropagation();

	var username = event.data.username;
	document.user.username.value = username;
	document.user.username.disabled = true;

	return false;
}

function delUser(event) {
	event.preventDefault();
	event.stopPropagation();

	var tr = $(event.target).parents('tr.user');
	var username = event.data.username;
	$.post('/user/', {username: username, del: 1});
	tr.remove();

	return false;
}

function onKeyDownUsername(event) {
	// Nur Buchstaben
	if ((event.keyCode >= 65 && event.keyCode <= 90) ||
	    (event.keyCode >= 97 && event.keyCode <= 122) ||
	    (event.keyCode >= 37 && event.keyCode <= 40) ||
	    event.keyCode == 8 || event.keyCode == 46 || event.keyCode == 9)
	    //(event.keyCode >= 48 && event.keyCode <= 57))
		return true;
	return false;
}

$(document.user.username).on('keydown', onKeyDownUsername);
$('button.submit').on('click', submitForm);
$('button.reset').on('click', resetForm);
$('tr.user').each(function(i, row) {
	row = $(row);
	var username = row.attr('data-username')
	var edit = $('.edit', row);
	var del = $('.delete', row);
	edit.on('click', {username: username}, editUser);
	del.on('click', {username: username}, delUser);
});

