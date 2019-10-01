function startChat () {
  let conn
  const log = document.getElementById("log")

  document.getElementById("form").addEventListener('submit', function (e) {
    e.preventDefault()
    const payload = {
      message: this.elements['message'].value,
      'user-id': this.elements['user-id'].value
    }
    if (!conn) { return false; }
    if (!payload['message'] || !payload['user-id']) { return false; }

    conn.send(JSON.stringify(payload))
    this.elements['message'].value = ""
  })

  if (window["WebSocket"]) {
    conn = new WebSocket("ws://" + document.location.host + "/ws")

    conn.addEventListener('message', function (e) {
    const data = JSON.parse(e.data)
    const div = document.createElement('div')
    div.style = { color: data.Color }
    div.innerHTML = `
    <span>[${data.TimeString}]</span>
    <span>&lt;${data.NickName}&gt;</span>
    <span>${data.Message}</span>`
    log.appendChild(div)
    })
  } else {
    console.log("no websocket")
  }
}
