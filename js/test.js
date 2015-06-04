//http://planetoftheweb.com/presentations/fresno/w01d02.html#/5

//var myLinks = [{link : 'google.com', icon : 'google'}, {link : 'yahoo.com', icon : 'yahoo'}, {link : 'github.com', icon : 'github'}, {link : 'facebook.com', icon : 'facebook-official'}];
//var out = "<ul>";
//for (var i = 0; i < myLinks.length; i++)
//{
//  out += ("<li><a href = " + myLinks[i].link + "><i class='fa fa-'>" + myLinks[i].icon + "</i></a></li>");
//}
//out += "</ul>";
//console.log(out);
//----------------------------------------------------------------------
//var links = ['http://go.com', 'go', 'http://planetoftheweb.com', 'planetoftheweb'];
//console.log('<ul>');
//for(var i = 0; i< links.length; i+=2)
//{
//    console.log("<li><a href = " + links[i] + "> </a>" + links[i+1] + "</li>");
//}
//console.log('</ul>');
//----------------------------------------------------------------------
// var info = {
//   full_name : "Shen",
//   age : 23,
//   computer : false,
//   movies: ['San Andreas', '22 Jump Street']
// };
  
// console.log(info.full_name);
// console.log(info.age);

// if(info.compuer){
//   console.log("Computer Upgraded!");
// }
// else{
//   console.log("Nothing new");
// }
// console.log('');

// for(var i = 0; i < info.movies.length; i++){
//   console.log(info.movies[i]);
// }
//----------------------------------------------------------------------
//var links = {
//  smedia: ['facebook.com', 'googleplus.com']
//};
//
//for(var index in links.smedia){
//  console.log("\n<a href=" + links.smedia[index] + "</a>");
//}
//----------------------------------------------------------------------
//var myNode = document.querySelectorAll('ul>li');
//for(var i = 0; i < myNode.length; i++){
//  myNode[i].className = 'myClass';
//  myNode[i].style.backgroundColor = 'red';
//}
// for(var index in myNode){
//   myNode[index].className= 'myClass';
// }
//----------------------------------------------------------------------

//var myNode = document.querySelector('.item');
//
//for(var i = 1; i < 10; i++){
//  var myElement = document.createElement('img');
//  myElement.src = 'http://lorempixel.com/120/120/sports/'+i;
//  myElement.alt = 'Sports Image '+i;
//  myNode.appendChild(myElement);
//  myElement.style['-moz-display'] = 'flex';
//  myElement.style['-moz-flex-direction'] = 'row';
//  myElement.style['-moz-border-radius'] = '50% 50%';
//  myElement.style['-moz-margin-left'] = '10px';
//----------------------------------------------------------------------

//document.querySelector('.boxes').addEventListener('click', function(e){ if(e.target.className !== 'boxes'){
// e.target.id='pink';}}, false);

//----------------------------------------------------------------------
//function myFunction(){
//  var randItem = document.querySelector('div');
//  var randList=['http://google.com','http://yahoo.com','http://hotmail.com','http://gmail.com'];
//  var randLink = document.createElement('a');
//  randLink.href = randList[Math.floor(Math.random()*(randList.length))];
//  randLink.innerHTML = randLink.href;
//  randItem.appendChild(randLink);
//  console.log(randLink.href);
//}

//-----------------8 BALL------------------------------------------------


//-----------------------------------------------------------------------
//function printDay(){
//  var weekdays = ['Sunday', 'Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday'];
//  var day = new Date();
//  
//  var newDay = document.querySelector('body');
//  var hTag = document.createElement('h1');
//  hTag.innerHTML = weekdays[day.getDay()];
//  newDay.appendChild(hTag); 
//}