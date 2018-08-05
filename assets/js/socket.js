var ansi_up = require('./ansi_up.js')

function initSocket () {
    let serverAddr = document.getElementById('server-addr').value
    console.log(serverAddr)
    let socket = io('ws://192.168.1.4:8080', {transports: ['websocket']})
    socket.on('connect', function () {
        renderLog('socket connected')
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