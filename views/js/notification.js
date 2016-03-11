$.get( "http://localhost:8000/notifications", function(data){
})
.done(function(data) {
  $.each(data, function(index, value) {
    $("#content").append("<tr>");	  
    $("#content").append("<td>" + value.notificationType + "</td>");	  

    $("#content").append("<td>");	  
    $.each(value.mail, function(index, value) {
      //enter
    });
    $("#content").append("</td>");	  
    $("#content").append("<td>" + value.CreatedAt+ "</td>");	  
    $("#content").append("</tr>"); 
  });
})
