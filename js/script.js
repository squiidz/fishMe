
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
  $("#place").val(function() {
    return "Found You !";
  });
}
