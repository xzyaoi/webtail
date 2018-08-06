var ansi_up = require('./ansi_up.js')

let connected = false

function initSocket () {
    let serverAddr = document.getElementById('server-addr').value
    let socket = io('ws://' + serverAddr, {transports: ['websocket']})
    socket.on('connect', function () {
        connected = true
        toggleStatus()
    })
    socket.on('logevent', function(message) {
        renderLog(message.Data)
    })
}

function renderLog (log) {
    let data = ansi_up.escape_for_html(log)
    data = ansi_up.ansi_to_html(data)
    let div = document.createElement('div')
    let p = document.createElement('p')
    div.className = 'line'
    p.innerHTML = data
    div.appendChild(p)
    document.getElementById('log').appendChild(div)
}

function toggleStatus () {
    if (connected) {
        document.getElementById('connect_button').innerHTML = "Disconnect"
        document.getElementById('connect_button').classList.remove("pure-button-primary")
        document.getElementById('connect_button').classList.add("button-error")
    } else {
        document.getElementById('connect_button').innerHTML = "Connect"
        document.getElementById('connect_button').classList.remove("pure-button-error")
        document.getElementById('connect_button').classList.add("button-primary")    }
}

function preparePage () {
    document.getElementById('connect_button').addEventListener('click',function (e) {
        initSocket()
        e.preventDefault()
    })
}

preparePage()