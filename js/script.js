

function getLocation()  {
  if (navigator.geolocation)  {
    navigator.geolocation.getCurrentPosition(showPosition);
  }
  else  {
    alert("Geolocation is not supported by this browser.");
  }
}

function showPosition(position) {
  $("#findme").val(function() {
    return position.coords.latitude + ", " + position.coords.longitude;
  });
}

var map;
function watchLocation() {
  var mapOptions = {
    zoom: 8,
    center: new google.maps.LatLng(parseFloat($("#catchPos").val()))
  };
  map = new google.maps.Map(document.getElementById('map-canvas'),
      mapOptions);
}

google.maps.event.addDomListener(window, 'load', initialize);
