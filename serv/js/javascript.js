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
        if (capitalize(cityList[i][1][j].split("-")[1]) == requestCountry) {
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
    });

    geocodeAddress(requestCountry, function(error, location) {
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

// function getMarker(pos) {
//     return null;
//     for (let i = 0; i < markers.length; i++) {
//         if (markers[i].position) {
//             return markers[i];
//         }
//     }
//     return null;
// }

// Google Markers
// async function initMarkers() {
//     const { AdvancedMarkerElement } = await google.maps.importLibrary("marker");

//     for (let i = 0; i < cityList.length; i++) {
//         for (let j = 0; j < cityList[i][1].length; j++) {
//             geocodeAddress(cityList, function(error, location) {
//             if (error) {
//                 console.error(error.message);
//             } else {
//                 var mak = getMarker();
//                 if (mak != null ) {
//                     mak.content.appendChild(document.getElementById(cityList[i][0]))
//                 } else {
//                     const content = document.createElement("div");
//                     content.style.display = "flex";
//                     content.style.flexDirection = "column"
//                     content.style.justifyContent = "center";
//                     content.appendChild(document.getElementById(cityList[i][0]));

//                     const marker = new AdvancedMarkerElement({
//                         position : location,
//                         content : content,
//                         map : map,
//                     });

//                     markers.push(marker);
//                 }
                
//             }
//             });
//         }
//     }

// }

async function simpleMarker() {
    const { AdvancedMarkerElement } = await google.maps.importLibrary("marker");
    for (let i = 0; i < cityList.length; i++) {
        for (let j = 0; j < cityList[i][1].length; j++) {
            geocodeAddress(cityList[i][1][j], function(error, location) {
                if (error) {
                    console.error(error.message);
                } else {

                    const content = document.getElementById(cityList[i][0]).cloneNode(true);
                    content.hidden = false

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
}


initMap();
simpleMarker();
//initMarkers();
