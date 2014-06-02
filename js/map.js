

var map;        
var myCenter=new google.maps.LatLng(53, -1.33);


function initialize() {

  var coord = $("#locat").val()
  var splitCoord = coord.split(",")
  var xPos = parseFloat(splitCoord[0])
  var yPos = parseFloat(splitCoord[1])
  var realPos = new google.maps.LatLng(xPos, yPos);

  var marker=new google.maps.Marker({
    position:realPos
  });

  var mapProp = {
      center: realPos,
      zoom: 8,
      draggable: true,
      scrollwheel: false,
      mapTypeId:google.maps.MapTypeId.ROADMAP
  };
 
  map=new google.maps.Map(document.getElementById("map-canvas"),mapProp);
  marker.setMap(map);
    
  google.maps.event.addListener(marker, 'click', function() {
      
    infowindow.setContent(contentString);
    infowindow.open(map, marker);
    
  }); 
};
google.maps.event.addDomListener(window, 'load', initialize);

google.maps.event.addDomListener(window, "resize", resizingMap());

$('#myMapModal').on('show.bs.modal', function() {
   //Must wait until the render of the modal appear, thats why we use the resizeMap and NOT resizingMap!! ;-)
   resizeMap();
})

function resizeMap() {
   if(typeof map =="undefined") return;
   setTimeout( function(){resizingMap();} , 300);
}

function resizingMap() {
   if(typeof map =="undefined") return;
   var center = map.getCenter();
   google.maps.event.trigger(map, "resize");
   map.setCenter(center); 
   map.setZoom(map.getZoom());
}