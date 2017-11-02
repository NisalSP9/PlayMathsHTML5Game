var canvas;
var canvas2;
var ctx;
var ctx2;
var x = -550;
var y = 0;
var x2 = -1000;
var y2 = -2;

function init() {
    canvas = document.getElementById("canvas");
    ctx = canvas.getContext("2d");
    draw();

    setInterval(draw, 10);

}


function init2() {


    canvas2 = document.getElementById("canvas2");
    ctx2 = canvas2.getContext("2d");
    draw2();

    setInterval(draw2, 10);

}


function draw() {

    ctx.clearRect(0, 0, 1366, 300);
    var img = new Image();

    if (x % 2 == 0) {
        img.src = "img/h2.png";
    } else {
        img.src = "img/h1.png";
    }

    ctx.drawImage(img, x, -40, 300, 225);

    x += 1;

    replay();

}

function replay() {

    if (x > 1300) {
        x = -550;
        draw();


    }

}


function draw2() {

    ctx2.clearRect(0, 0, 1366, 300);
    var img2 = new Image();

    if (x2 % 2 == 0) {
        img2.src = "img/suv.png";
    } else {
        img2.src = "img/suv2.png";
    }

    ctx2.drawImage(img2, x2, y2, 300, 225);

    x2 += 1;

    replay2();

}

function replay2() {
    if (x2 > 1358) {
        x2 = -250;
        draw2();

    }

}

function nextMtd() {

    window.location.href = "/components/city.html";

}


function testMtd() {

    window.location.href = "/components/test_login/test_login.html";

}


