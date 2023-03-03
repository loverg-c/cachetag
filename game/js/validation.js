const form = document.querySelector('#validation');
const serverAddr = window.location.protocol + '//' + window.location.hostname + ":8080";

function getCookie (name) {
    let value = `; ${document.cookie}`;
    let parts = value.split(`; ${name}=`);
    if (parts.length === 2) return parts.pop().split(';').shift();
}

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
    httpRequestUserList = new XMLHttpRequest();

    httpRequestUserList.onreadystatechange = function () {
        if (httpRequestUserList.readyState === 4) {
            if (httpRequestUserList.status == 200) {
                parseUser(httpRequestUserList.responseText);
            } else {
                alert("Status error: " + httpRequestUserList.status);
            }
        }
    }
    try {
        httpRequestUserList.open("GET", serverAddr + "/players", true);
        httpRequestUserList.send();
    } catch (e) {
        alert(JSON.stringify(e));
    }
}

function parseUser(response) {

    const select = document.getElementById("username");

    select.innerHTML = null;


    const defaultOption = document.createElement('option');
    const optionText = document.createTextNode("");

    defaultOption.appendChild(optionText);
    defaultOption.disabled = true;
    defaultOption.selected = true;

    select.appendChild(defaultOption);


    const userList = JSON.parse(response)?.sort((a, b) => 0.5 - Math.random()) ?? [];

    const optCached = localStorage['preferedPlayer'] || getCookie("preferedPlayer") || undefined;

    userList.forEach(function (user) {

        const newOption = document.createElement('option');
        const optionText = document.createTextNode(user.username);

        if (optCached !== undefined && parseInt(optCached) === user.id) {
            newOption.selected = true;
            defaultOption.selected = false;
        }

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
    httpRequestValidateTag = new XMLHttpRequest();

    httpRequestValidateTag.onreadystatechange = function () {
        if (httpRequestValidateTag.readyState === 4) {
            if (httpRequestValidateTag.status == 200) {
                response = JSON.parse(httpRequestValidateTag.responseText);

                const msg = `Bien joué <span class="surlign">${response.username}</span>!<br><br>
                             Ton nouveau score: <span class="surlign">${response.score} points</span>`;
                document.querySelector('.wrapper').innerHTML = msg;
                localStorage['preferedPlayer'] = playerId; // only strings
                document.cookie = `preferedPlayer=${playerId}`;

            } else if (httpRequestValidateTag.status == 403) {
                const msg = `Nop, ya un code secret sur les tags, petit malin va !`;
                document.querySelector('.wrapper').innerHTML = msg;

                throw new Error(msg);
            }else if (httpRequestValidateTag.status == 409) {
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
        httpRequestValidateTag.open("POST", `${serverAddr}/players/${playerId}/validated_tags`, true);

        httpRequestValidateTag.setRequestHeader("Content-Type", "application/json");
        httpRequestValidateTag.setRequestHeader("Accept", "application/json");
        httpRequestValidateTag.send(JSON.stringify({'player_id': parseInt(playerId), 'tag_id': parseInt(tagId), 'secret': tagSecret}));
    } catch (e) {
        alert(JSON.stringify(e));
        throw e
    }
}

getTagInfo()
loadUserList()