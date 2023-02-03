function loadScore() {
    httpRequest = new XMLHttpRequest();

    httpRequest.onreadystatechange = function () {

        if (httpRequest.readyState == 4) {
            //the request is completed, now check its status
            if (httpRequest.status == 200) {
                parseScore(httpRequest.responseText);
            } else {
                console.log("Status error: " + httpRequest.status);
            }
        } else {
            console.log("Ignored readyState: " + httpRequest.readyState);
        }
    }

    httpRequest.open("GET", "http://localhost:8080/players/scores", true);
    httpRequest.send();
}

function parseScore(response) {

    document.getElementById("score-list").innerHTML = null;

    const scores = JSON.parse(response)?.sort((a, b) => {
        return b.score - a.score;
    }) ?? [];

    let positionValue = 0;
    let previousScore = 0;

    scores.forEach(function (score) {
        let line = document.createElement('div');
        line.classList.add("score")

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

function loadMercure() {
    const url = new URL('https://localhost/.well-known/mercure');
    url.searchParams.append('topic', 'https://example.com/my-private-topic');

    const eventSource = new EventSource(url);

    eventSource.onmessage = e => parseScore(e.data); // do something with the payload
}

loadScore()
loadMercure()