
var timeouted = false; //var for storing if we had a timeout, 

window.onload = function () {
  console.log("App started")
  setInterval(heartbeat, 5000);
};

function stdrequest(key) {
  // Standard Request for activating functions
  fetch("app/send/" + key, {
    methode: 'POST',
    credentials: 'same-origin'
  })
}

function heartbeat() {
  /**
   * Sends an Ping to the Backend to check for avalability
   * If the Ping fails, the user is notified with an alert. 
   */
  var xhr = new XMLHttpRequest();
  xhr.timeout = 3000;
  onerror = () => {
    if (!timeouted) {
      document.body.style.background = 'red';
      swal("Beamer nicht mehr erreichbar!", "Netzwerkverbindung überprüfen oder neustarten!", "error");
    }
    timeouted = true;
  };
  xhr.ontimeout = onerror;
  xhr.onerror = () => {
    if (xhr.status == 401) {
      onerror();
    }
  };
  xhr.onload = () => {
    //if we are here then nothing special happened and everything is ok
    //check if we had a timeout before and message the user
    if (timeouted) {
      document.body.style.background = document.body.style.getPropertyValue('--backgroundcolor');
      swal("Beamer verbunden!", "Die Verbindung zu Beamer wurde erfolgreich wiederhergestellt.", "success");
    }
    timeouted = false;
  };
  xhr.open('GET', '/ping', true);
  xhr.send();
}
