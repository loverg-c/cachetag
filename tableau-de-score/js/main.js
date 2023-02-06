const serverAddr = window.location.protocol + '//' + window.location.hostname;

function loadMax() {
    httpRequestNbTag = new XMLHttpRequest();

    httpRequestNbTag.onreadystatechange = function () {
        if (httpRequestNbTag.readyState === 4) {
            if (httpRequestNbTag.status == 200) {
                parseMax(httpRequestNbTag.responseText);
            } else {
                alert("Status error: " + httpRequestNbTag.status);
            }
        }
    }
    try {
        httpRequestNbTag.open("GET", serverAddr + ":8080/tags", true);
        httpRequestNbTag.send();
    } catch (e) {
        alert(JSON.stringify(e));
    }
}


function parseMax(response) {
    const tagList = JSON.parse(response);

    const maxScore = tagList.reduce((partialSum, a) => partialSum + a.score, 0);

    const info = document.getElementById("max-tag");
    info.innerHTML = `Le score maximum possible est de <span class="surlign">${maxScore} points</span>`;
}

function loadScore() {
    httpRequest = new XMLHttpRequest();

    httpRequest.onreadystatechange = function () {
        if (httpRequest.readyState === 4) {
            if (httpRequest.status == 200) {
                parseScore(httpRequest.responseText);
            } else {
                alert("Status error: " + httpRequest.status);
            }
        }
    }
    try {
        httpRequest.open("GET", serverAddr + ":8080/players/scores", true);
        httpRequest.send();
    } catch (e) {
        alert(JSON.stringify(e));
    }
}

function parseScore(response) {
    document.getElementById("score-list").innerHTML = null;

    const scores = JSON.parse(response)?.sort((a, b) => {
        if (b.score === a.score) {
            return ('' + a.username).localeCompare(b.username);
        }
        return b.score - a.score;
    }) ?? [];

    let positionValue = 0;
    let previousScore = 0;

    scores.forEach(function (score) {
        let line = document.createElement('div');
        line.classList.add("score");

        if (previousScore !== score.score) {
            positionValue++;
        }

        let position = document.createElement('div');
        position.innerHTML = positionValue.toString().padStart(2, '0');
        position.classList.add("position");
        line.appendChild(position);

        if (positionValue === 1) {
            line.classList.add('first');
        } else if (positionValue === 2) {
            line.classList.add('second');
        } else if (positionValue === 3) {
            line.classList.add('third');
        }

        let username = document.createElement('div');
        username.innerHTML = score.username;
        username.classList.add("username");
        line.appendChild(username);

        let scoreValue = document.createElement('div');
        scoreValue.innerHTML = score.score;
        previousScore = score.score;
        scoreValue.classList.add("score-value");
        line.appendChild(scoreValue);

        document.getElementById("score-list").appendChild(line);
    });
}

function displayNotification() {
    const alertPopup = document.querySelector('#alert-popup');

    alertPopup.classList.remove('hidden');
    alertPopup.classList.add('show');
    setTimeout(function () {
        alertPopup.classList.add('hidden');
        alertPopup.classList.remove('show');
    }, 3000);
}

function loadMercure() {
    const url = new URL(serverAddr + ":8081/.well-known/mercure");

    url.searchParams.append('topic', serverAddr + '/tableau-de-score');
    const eventSource = new EventSource(url);
    eventSource.onmessage = ({data}) => {
        displayNotification();
        parseScore(data);
    }
}

function init() {
    loadMax();
    loadScore();
    loadMercure();

}

init();