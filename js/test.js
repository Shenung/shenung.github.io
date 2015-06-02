//var myLinks = [{link : 'google.com', icon : 'google'}, {link : 'yahoo.com', icon : 'yahoo'}, {link : 'github.com', icon : 'github'}, {link : 'facebook.com', icon : 'facebook-official'}];
//var out = "<ul>";
//for (var i = 0; i < myLinks.length; i++)
//{
//  out += ("<li><a href = " + myLinks[i].link + "><i class='fa fa-'>" + myLinks[i].icon + "</i></a></li>");
//}
//out += "</ul>";
//console.log(out);
//----------------------------------------------------------------------
var links = ['http://go.com', 'go', 'http://planetoftheweb.com', 'planetoftheweb'];
console.log('<ul>');
for(var i = 0; i< links.length; i+=2)
{
    console.log("<li><a href = " + links[i] + "> </a>" + links[i+1] + "</li>");
}
console.log('</ul>');