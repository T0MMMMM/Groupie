let map;
let requestCountry = document.getElementById("country").getAttribute("data-variable");
let jsVariable = document.getElementsByClassName("js-variable");
let cityList = []; // [ NAME , [LOCATIONS]]

function capitalize(str) {
    if (typeof str !== 'string' || str.length === 0) {
        return str;
    }
    return str.split(' ').map(word => {
        return word.charAt(0).toUpperCase() + word.slice(1).toLowerCase();
    }).join(' ');
}

for (let i = 0; i < jsVariable.length; i++) {
    cityList.push([]);
    cityList[i].push(jsVariable[i].getAttribute("data-variable").split(":")[0]);
    cityList[i].push(jsVariable[i].getAttribute("data-variable").split(":")[1].substring(1, jsVariable[i].getAttribute("data-variable").split(":")[1].length-1).split(" "));
    for (let j = 0; j < cityList[i][1].length; j++) {
        if (capitalize(cityList[i][1][j].split("-")[1].replaceAll("_", " ")) == requestCountry) {
            cityList[i][1][j] = cityList[i][1][j].split("-")[0].replaceAll("_", " ") + ", " + capitalize(cityList[i][1][j].split("-")[1]);
            cityList[i][1][j] = capitalize(cityList[i][1][j]);
        } else {
            cityList[i][1][j] = null;
        }
    }
    cityList[i][1] = cityList[i][1].filter(element => { return element !== null; });
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
        zoom: 6,
        mapId: "1",
        streetViewControl: false,
        mapTypeControl: false,
    });

    geocodeAddress(requestCountry, function(error, location) {
        if (error) {
            console.error(error.message);
        } else {
            alert(requestCountry);
            map.setCenter(location);
        }
    });
}


var markers = {};

function getMarker(pos) {
    for (let i = 0; i < Object.keys(markers).length; i++) {
        if (Object.keys(markers)[i] == pos) {
            return markers[Object.keys(markers)[i]];
        }
    }
    return null;
}

async function simpleMarker() {
    const { AdvancedMarkerElement } = await google.maps.importLibrary("marker");
    for (let i = 0; i < cityList.length; i++) {
        for (let j = 0; j < cityList[i][1].length; j++) {
            geocodeAddress(cityList[i][1][j], function(error, location) {
                if (error) {
                    console.error(error.message);
                } else {

                    var existantMarker = getMarker(cityList[i][1][j]);
                    
                    if (existantMarker != null) {

                        const div = existantMarker.content;
                        const artist = document.getElementById(cityList[i][0]).cloneNode(true);
                        artist.style.display = "flex";
                        div.insertBefore(artist, div.children[0]);
                        existantMarker.content = div;

                    } else {

                        const content = document.createElement("div");
                        content.style.display = "flex";
                        content.style.flexDirection = "column";
                        content.style.justifyContent = "center";

                        const artist = document.getElementById(cityList[i][0]).cloneNode(true);
                        artist.style.display = "flex";

                        const markerImg = document.createElement("img");
                        markerImg.src = "serv/assets/images/marker.png";
                        markerImg.style.justifyContent = "center";
                        markerImg.style.alignItems = "center";
                        markerImg.style.marginLeft = "auto";
                        markerImg.style.marginRight = "auto";
                        markerImg.height = "30";
                        markerImg.width = "30";

                        content.appendChild(markerImg);
                        content.insertBefore(artist, content.children[0]);

                        const marker = new AdvancedMarkerElement({
                            position : location,
                            content : content,
                            map : map,
                        });
                
                        markers[cityList[i][1][j]] = marker;

                    }

                }
            });
        }
    }
}


initMap();
simpleMarker();
//initMarkers();
