const form = document.querySelector('#validation');
const serverAddr = window.location.protocol + '//' + window.location.hostname + ":8080";


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


    const userList = JSON.parse(response)?.sort((a, b) => {
            return ('' + a.username).localeCompare(b.username);
    }) ?? [];



    userList.forEach(function (user) {

        const newOption = document.createElement('option');
        const optionText = document.createTextNode(user.username);

        newOption.appendChild(optionText);
        newOption.setAttribute('value',user.value);

        select.appendChild(newOption);
    });
}


loadUserList()