function adminLogin() {
    const username = document.getElementById("username").value;
    const password = document.getElementById("password").value;

    if (username === 'admin' && password === 'admin') {
        window.location.href = `/admin-login`;
    } else {
        alert("Incorrect username or password");
    }
}

function userLogin() {
    const username = document.getElementById("username").value;
    const password = document.getElementById("password").value;

    if (username !== ' ' && password !== ' ') {
        window.location.href = `/user-login`;
    } else {
        alert("Please enter username and password");
    }
}

function Back() {
    window.location.href = `/`;
}