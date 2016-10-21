var express = require("express");
var bodyParser = require("body-parser");
var app = express();
 
app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: true }));
 
//var routes = require("./routes/routes.js")(app);
 
var people = new Array();
function Persona(ID,Fn,Ln,City,State){
	this.ID = ID;
	this.Firstname = Fn;
	this.Lastname = Ln;
	this.City = City;
	this.State = State;
}

app.get("/people", function(req,res){
	res.send(people);
})
app.get("/people/:id", function(req,res){
	for (var i = people.length - 1; i >= 0; i--) {
		if (people[i].ID==req.params.id) {
			res.send(people[i]);
			break;
		}
	}
})
app.post("/people/:id", function(req,res){
	people.push(new Persona(req.params.id,req.body.Firstname,req.body.Lastname,req.body.City,req.body.State));
	res.send(people);
})
app.delete("/people/:id", function(req,res){
	for (var i = people.length - 1; i >= 0; i--) {
		if (people[i].ID==req.params.id) {
			people.splice(i,i+1);
			res.send(people);
			break;
		}
	}
})

var server = app.listen(3000, function () {
    console.log("Olá mundo da porta %s", server.address().port);
});