$(function(){

  "use strict";
  
  var tableSource = $("#table-template").html();
  var tableTemplate = Handlebars.compile(tableSource)

  var rowSource = $("#row-template").html();
  var rowTemplate = Handlebars.compile(rowSource)
  
  $.getJSON("/certificates", function(data){

  	   $("#certsTable-container").html(tableTemplate(data));
  	   var table = $("#certsTable").DataTable();

       var host = "http://" + location.host;
   
       var socket = glue(host,{ baseURL: "/ws/"});

  	   socket.onMessage(function(data){

	       var count = ($("#notificationCount").data("count") + 1);

	       $("#notificationCount").data("count", count);
           var text = " update to table"
	       if(count > 1){
	           text = " updates to table"
	       }

	       $("#notificationCount").text(count + text);

	   	   var row = JSON.parse(data);
	       var newRow = $(rowTemplate(row));

	   	   var node = table.row.add(newRow).draw().node();
	   	   $(node).addClass("active");
       });

       $("#notificationCount").click(function(){
       	    $("#notificationCount").data("count", 0);
	        $("#notificationCount").text("");
	        $("tr").removeClass("active");
       });
		
  });
   
});