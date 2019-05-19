window.onload = function () {
    console.log('App started')
    window.addEventListener('keydown', this.keybindings, false)
    readLCD()
    this.setInterval(readLCD, 2500)
    particlejsinit()
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

function keybindings(e) {
    // Bind Apirequests to Keys
    var key = e.keyCode;
    switch (key) {
        case 37:
            apirequest('infrared', 'arrowleft');
            break; //Left key
        case 38:
            apirequest('infrared', 'arrowup');
            break; //Up key
        case 39:
            apirequest('infrared', 'arrowright');
            break; //Right key
        case 40:
            apirequest('infrared', 'arrowdown');
            break; //Down key
        case 79:
            apirequest('shutteropen', 'fast');
            break; // o key
        case 67:
            apirequest('shutterclose', 'fast');
            break; // c key
        case 13:
            apirequest('infrared', 'enter');
            break; // enter key
        case 27:
            apirequest('infrared', 'exit');
            break; //esc key
        default:
            //alert(key); //Everything else
    }
}

function particlejsinit() {
    console.log("Loaded particles-js")
    /* ---- particles.js config ---- */

particlesJS("aa_particles", {
    "particles": {
      "number": {
        "value": 80,
        "density": {
          "enable": true,
          "value_area": 800
        }
      },
      "color": {
        "value": "#ffffff"
      },
      "shape": {
        "type": "circle",
        "stroke": {
          "width": 0,
          "color": "#000000"
        },
        "polygon": {
          "nb_sides": 5
        },
      },
      "opacity": {
        "value": 0.5,
        "random": false,
        "anim": {
          "enable": false,
          "speed": 1,
          "opacity_min": 0.1,
          "sync": false
        }
      },
      "size": {
        "value": 3,
        "random": true,
      },
      "line_linked": {
        "enable": true,
        "distance": 150,
        "color": "#ffffff",
        "opacity": 0.4,
        "width": 1
      },
      "move": {
        "enable": false,
      }
    },
    "interactivity": {
      "detect_on": "canvas",
      "events": {
        "onhover": {
          "enable": true,
          "mode": "grab"
        },
        "resize": true
      },
      "modes": {
        "grab": {
          "distance": 140,
          "line_linked": {
            "opacity": 1
          }
        },
        "bubble": {
          "distance": 400,
          "size": 40,
          "duration": 2,
          "opacity": 8,
          "speed": 3
        },
        "repulse": {
          "distance": 200,
          "duration": 0.4
        },
        "remove": {
          "particles_nb": 2
        }
      }
    },
    "retina_detect": true
  });
}