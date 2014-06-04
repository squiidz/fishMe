

var map;        
var defaultPlace=new google.maps.LatLng(45.833309, -73.417408);


function initialize(realPos) {

  var marker=new google.maps.Marker({
    position:realPos
  });

  var mapProp = {
      center: realPos,
      zoom: 12,
      draggable: true,
      scrollwheel: true,
      mapTypeId:google.maps.MapTypeId.ROADMAP
  };
 
  map=new google.maps.Map(document.getElementById("map-canvas"), mapProp);
  marker.setMap(map);
  
  $('#mapModal').on('show.bs.modal', function() {
    resizeMap();
  });
  
  google.maps.event.addListener(marker, 'click', function() {
    infowindow.setContent(contentString);
    infowindow.open(map, marker);  
  }); 

};

google.maps.event.addDomListener(window, 'load', initialize);

google.maps.event.addDomListener(window, "resize", resizingMap());

function resizeMap() {
   if(typeof map =="undefined") return;
   setTimeout( function(){resizingMap();} , 300);
}

function resizingMap() {
   if(typeof map =="undefined") return;
   var center = map.getCenter();
   google.maps.event.trigger(map, "resize");
   map.setCenter(center); 
}

function popMap(place) {
  var coord = place;
  var splitCoord = coord.split(",");
  var xPos = parseFloat(splitCoord[0]);
  var yPos = parseFloat(splitCoord[1]);

  var checkIsNumber = function() {
    if (isNaN(splitCoord[0])) {
      initialize(defaultPlace);
    } 
    else {
      realPos = new google.maps.LatLng(xPos, yPos);
      initialize(realPos);
    }
  }
  
  checkIsNumber();

}