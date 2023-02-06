const form = document.querySelector('#validation');
const serverAddr = window.location.protocol + '//' + window.location.hostname + ":8080";

form.addEventListener('submit', function (e) {
    // prevent the form from submitting
    e.preventDefault();

    const userID = form.querySelector("select#username").value;
    if (userID.replace(/\s/g, '').length === 0) {

        alert('Veuillez selectionner une valeur');
        return;
    }

    const tagValue = form.querySelector("input#tagId").value;
    const tagSecret = form.querySelector("input#tagSecret").value;

    verifyTag(tagValue, tagSecret);
    validateTag(tagValue, userID, tagSecret)
});


function loadUserList() {
    httpRequest = new XMLHttpRequest();

    httpRequest.onreadystatechange = function () {
        if (httpRequest.readyState === 4) {
            if (httpRequest.status == 200) {
                parseUser(httpRequest.responseText);
            } else {
                alert("Status error: " + httpRequest.status);
            }
        }
    }
    try {
        httpRequest.open("GET", serverAddr + "/players", true);
        httpRequest.send();
    } catch (e) {
        alert(JSON.stringify(e));
    }
}

function parseUser(response) {

    const select = document.getElementById("username");

    select.innerHTML = null;


    const newOption = document.createElement('option');
    const optionText = document.createTextNode("");

    newOption.appendChild(optionText);
    newOption.disabled = true;
    newOption.selected = true;

    select.appendChild(newOption);


    const userList = JSON.parse(response)?.sort((a, b) => 0.5 - Math.random()) ?? [];


    userList.forEach(function (user) {

        const newOption = document.createElement('option');
        const optionText = document.createTextNode(user.username);

        newOption.appendChild(optionText);
        newOption.setAttribute('value', user.id);

        select.appendChild(newOption);
    });
}

function getTagInfo() {
    let paramString = window.location.toString().split('?')[1];
    let queryString = new URLSearchParams(paramString);

    let tagId = '';
    let tagSecret = '';

    for (let pair of queryString.entries()) {
        if (pair[0] === 'id') {
            tagId = pair[1];
        } else if (pair[0] === 'sc') {
            tagSecret = pair[1];
        }
    }
    if (tagId.replace(/\s/g, '').length === 0
        || tagSecret.replace(/\s/g, '').length === 0) {
        document.querySelector('.wrapper').innerHTML = 'Erreur sur ce tag';

        throw new Error("Erreur sur ce tag");
    }

    document.querySelector('span#tagNumber').innerHTML = tagId;
    document.querySelector('input#tagId').value = tagId;
    document.querySelector('input#tagSecret').value = tagSecret;

    verifyTag(tagId, tagSecret);
}

function verifyTag(tagId, tagSecret) {
    httpRequestTag = new XMLHttpRequest();

    httpRequestTag.onreadystatechange = function () {
        if (httpRequestTag.readyState === 4) {
            if (httpRequestTag.status == 404) {
                const msg = `Tag non trouvé`;
                document.querySelector('.wrapper').innerHTML = msg;

                throw new Error(msg);
            } else if (httpRequestTag.status == 200) {
                response = JSON.parse(httpRequestTag.responseText);
                if (response.valid !== true) {
                    const msg = `Nop, ya un code secret sur les tags, petit malin va !`;
                    document.querySelector('.wrapper').innerHTML = msg;

                    throw new Error(msg);
                }
            } else {
                document.querySelector('.wrapper').innerHTML = 'Erreur sur ce tag';

                throw new Error("Erreur sur ce tag");
            }
        }
    }
    try {
        httpRequestTag.open("POST", `${serverAddr}/tags/${tagId}/verify`, true);

        httpRequestTag.setRequestHeader("Content-Type", "application/json");
        httpRequestTag.setRequestHeader("Accept", "application/json");
        httpRequestTag.send(JSON.stringify({'secret': tagSecret}));
    } catch (e) {
        alert(JSON.stringify(e));
        throw e
    }
}

function validateTag(tagId, playerId, tagSecret) {
    httpRequest = new XMLHttpRequest();

    httpRequest.onreadystatechange = function () {
        if (httpRequest.readyState === 4) {
            if (httpRequest.status == 200) {
                response = JSON.parse(httpRequest.responseText);

                //todo couleur et formatage

                const msg = `Bien joué <span class="surlign">${response.username}</span>!<br><br>
                             Ton nouveau score: <span class="surlign">${response.score} points</span>`;
                document.querySelector('.wrapper').innerHTML = msg;

            } else if (httpRequest.status == 403) {
                const msg = `Nop, ya un code secret sur les tags, petit malin va !`;
                document.querySelector('.wrapper').innerHTML = msg;

                throw new Error(msg);
            }else if (httpRequest.status == 409) {
                const msg = `Tu as déjà validé ce tag`;
                document.querySelector('.wrapper').innerHTML = msg;

                throw new Error(msg);
            } else {
                document.querySelector('.wrapper').innerHTML = 'Erreur sur ce tag';

                throw new Error("Erreur sur ce tag");
            }
        }
    }
    try {
        httpRequest.open("POST", `${serverAddr}/players/${playerId}/validated_tags`, true);

        httpRequest.setRequestHeader("Content-Type", "application/json");
        httpRequest.setRequestHeader("Accept", "application/json");
        httpRequest.send(JSON.stringify({'player_id': parseInt(playerId), 'tag_id': parseInt(tagId), 'secret': tagSecret}));
    } catch (e) {
        alert(JSON.stringify(e));
        throw e
    }
}

getTagInfo()
loadUserList()