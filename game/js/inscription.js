const form = document.querySelector('#registration');
const serverAddr = window.location.protocol + '//' + window.location.hostname + ":8080";

form.addEventListener('submit', function (e) {
    // prevent the form from submitting
    e.preventDefault();

    const username = form.querySelector("input#username").value;

    if (username.replace(/\s/g, '').length === 0) {
        alert('Votre nom est vide');
        return;
    }

    registerUser(username.trim());
});


function registerUser(username) {
    let httpRequest = new XMLHttpRequest();

    httpRequest.onreadystatechange = () => onRegisterUserResponse(httpRequest);

    try {
        httpRequest.open("POST", serverAddr + "/players", true);

        httpRequest.setRequestHeader("Content-Type", "application/json");
        httpRequest.setRequestHeader("Accept", "application/json");
        httpRequest.send(JSON.stringify({'username': username}));
    } catch (e) {
        alert(JSON.stringify(e));
    }
}


function onRegisterUserResponse(httpRequest) {
    if (httpRequest.readyState === 4) {
        if (httpRequest.status === 409) {
            if (JSON.parse(httpRequest.responseText).message === 'Player already exist')
            alert('Nom de joueur déjà existant');
            return;
        }
        if (httpRequest.status === 200) {
            const username = JSON.parse(httpRequest.responseText).username;
            const welcome = document.createElement('div');

            welcome.id = 'welcoming-message';
            welcome.innerHTML = `Bienvenue <span class="surlign">${username}</span> !!!<br><br>
            Le jeu peut maintenant commencer.<br><br>
            Scan tout les <span class="surlign">tags NFC</span> que tu trouvera.<br><br>
            Un formulaire s'affichera, sélectionne à ce moment <span class="surlign">${username}</span> dans la liste et valide.<br><br>
            Tu gagnera instantanement un nombre de point associé au tag.<br><br>
            Un tag ne te fera gagner des points <span class="surlign">qu'une seule fois</span>. <br><br><br>
            Que le meilleur gagne !!!
`;

            const container = document.querySelector(".wrapper:first-child");
            container.innerHTML = null;
            container.appendChild(welcome);
        } else {
            alert("Status error: " + httpRequest.status);
            alert(JSON.stringify(httpRequest.response));
        }
    }
}
