//account = 10000;
$(document).ready(function() {
    $.ajax({
        url: "/js"
    }).then(function(data) {
       $('.title').append(data.title);
       //$('.title').append(account);
    });
});