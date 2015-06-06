/**
 * Created by shenung on 6/5/2015.
 */
var Timer = 30,
    myNode = document.querySelector('button');
var timerId = null;

myNode.addEventListener('click', function() {
    myNode.value='self destruct engaged';
    timerId = setInterval ( function() {
        var visualTime = document.querySelector('p');
        Timer = Timer-1;
        visualTime.innerHTML = Timer+' seconds!';
        visualTime.style.color = 'red';
        visualTime.style.zIndex = '1';
        visualTime.style.fontSize = '30px';
        if(Timer === 25){
            document.querySelector('#at25').style.visibility = 'visible';
        }
        else if(Timer === 20){
            document.querySelector('#at25').style.visibility = 'hidden';
            document.querySelector('#at20').style.visibility = 'visible';
        }
        else if(Timer === 15){
            document.querySelector('#at20').style.visibility = 'hidden';
            document.querySelector('#at15').style.visibility = 'visible';
        }
        else if(Timer === 10){
            document.querySelector('#at15').style.visibility = 'hidden';
            document.querySelector('#at10').style.visibility = 'visible';
        }
        else if(Timer === 5){
            document.querySelector('.start').style.visibility = 'hidden';
            document.querySelector('#at10').style.visibility = 'hidden';
            document.querySelector('#at5').style.visibility = 'visible';
            console.log("its working");
            document.querySelector('.almostFail').style.visibility = 'visible';
        }
        else if(Timer === 0)
        {
            document.querySelector('#at5').style.visibility = 'hidden';
            document.querySelector('.almostFail').style.visibility = 'hidden';
            document.querySelector('.fail').style.visibility = 'visible';
            abort();
        }
    }, 1000 );
});

document.querySelector('#abort').addEventListener('click',abort);

function abort(){
    clearInterval(timerId);
    timerId = null;
}