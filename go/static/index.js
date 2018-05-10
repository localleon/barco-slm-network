window.onload = function () {
    console.log('App started')
    readLCD()
    this.setInterval(readLCD, 2500)
}

function apirequest(cmd, opt) {
    // Standard Request to the Backend
    var url;
    if (opt == '') {
        url = '/api/' + cmd;
    } else {
        url = '/api/' + cmd + '/' + opt;
    }
    fetch(url, {
        method: 'GET',
        credentials: 'same-origin',
    })
    readLCD();
}

function readLCD() {
    // Fetches LCD Content from Backend
    fetch('/api/lcdread')
        .then(function (response) {
            return response.json();
        })
        .then(function (json) {
            parseLCD(json)
        });
}

function parseLCD(json) {
    // Parse json Response and write to html element
    var lcdfirst = document.getElementById("lcdfirst")
    var lcdsecond = document.getElementById("lcdsecond")
    lcdfirst.innerHTML = json.First
    lcdsecond.innerHTML = json.Second
}