let map;
let jsVariable = document.getElementById("locs");
let cityList = []; // [LOCATIONS]

function capitalize(str) {
    if (typeof str !== 'string' || str.length === 0) {
        return str;
    }
    return str.split(' ').map(word => {
        return word.charAt(0).toUpperCase() + word.slice(1).toLowerCase();
    }).join(' ');
}

cityList = jsVariable.getAttribute("data-variable").substring(1, jsVariable.getAttribute("data-variable").length-1).split(" ");

for (let j = 0; j < cityList.length; j++) {
    cityList[j] = cityList[j].split("-")[0].replaceAll("_", " ") + ", " + capitalize(cityList[j].split("-")[1]);
    cityList[j] = capitalize(cityList[j]);
}


function geocodeAddress(address, callback) {
    var geocoder = new google.maps.Geocoder();
    geocoder.geocode({ 'address': address }, function(results, status) {
      if (status === 'OK') {
        var location = {
          lat: results[0].geometry.location.lat(),
          lng: results[0].geometry.location.lng()
        };
        callback(null, location);
      } else {
        callback(new Error('Geocode was not successful for the following reason: ' + status));
      }
    });
}


// Google Map init 
async function initMap() {
    const { Map } = await google.maps.importLibrary("maps");


    map = new Map(document.getElementById("map"), {
        center: {lat: -34.397, lng: 150.644},
        zoom: 3,
        mapId: "1",
        streetViewControl: false,
        mapTypeControl: false,
    });

    geocodeAddress("Paris, France", function(error, location) {
        if (error) {
            console.error(error.message);
        } else {
            map.setCenter(location);
        }
    });

    //geocode({ address: requestCountry });

    //map.setCenter(geocode({ address: requestCountry }))
}

var markers = [];

async function simpleMarker() {
    const { AdvancedMarkerElement } = await google.maps.importLibrary("marker");
    for (let i = 0; i < cityList.length; i++) {
        geocodeAddress(cityList[i], function(error, location) {
            if (error) {
                console.error(error.message);
            } else {

                const content = document.createElement("div");
                content.style.display = "flex";
                content.style.flexDirection = "column";
                content.style.justifyContent = "center";

                const markerImg = document.createElement("img");
                markerImg.src = "serv/assets/images/marker.png";
                markerImg.style.justifyContent = "center";
                markerImg.style.alignItems = "center";
                markerImg.style.marginLeft = "auto";
                markerImg.style.marginRight = "auto";
                markerImg.height = "30";
                markerImg.width = "30";

                const artist = document.getElementById("artist").cloneNode(true);
                artist.style.display = "flex";

                content.appendChild(markerImg);
                content.insertBefore(artist, content.children[0]);

                const marker = new AdvancedMarkerElement({
                    position : location,
                    content : content,
                    map : map,
                });
        
                markers.push(marker);
            }
        });
    }
}


initMap();
simpleMarker();
//initMarkers();
