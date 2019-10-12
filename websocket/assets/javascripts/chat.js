function startChat () {
  _changeNickNameFeature()
  _changeTextColorFeature()

  let conn
  const log = document.getElementById("log")

  document.getElementById("form").addEventListener('submit', function (e) {
    e.preventDefault()
    if (_stopSendOnInput()) { return false }
    const payload = {
      message: this.elements['message'].value,
      'user-id': this.elements['user-id'].value
    }
    if (!conn) { return false }
    if (!payload['message'] || !payload['user-id']) { return false }

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

function _stopSendOnInput () {
  return (document.activeElement.id && document.activeElement.id === 'nickname')
}

function _changeNickNameFeature () {
  const nickShowElm = document.getElementById('chat-nickname-show')
  const nickInputElm = document.getElementById('chat-nickname-input')
  const nickShow = nickShowElm.firstElementChild
  const nickInput = nickInputElm.firstElementChild
  const updateNick = () => {
    nickShow.innerText = nickInput.value
    nickInput.blur()
    nickShowElm.style.display = 'block'
    nickInputElm.style.display = 'none'
  }
  nickShowElm.addEventListener('click', function (e) {
    nickShowElm.style.display = 'none'
    nickInputElm.style.display = 'block'
    nickInput.addEventListener('focus', function (e) { e.target.select() })
    nickInput.focus()
  })
  nickInput.addEventListener('keyup', function (e) {
    if (e.keyCode === 13) { updateNick() }
  })
  nickInput.addEventListener('click', function (e) { updateNick() })
}

function _changeTextColorFeature () {
  document.getElementById('btn-chat-color').addEventListener('click', function (e) {
    e.preventDefault()
    const elm = document.getElementById('color')
    const textColors = ['#000000', '#4464AD', '#843B62' , '#36453B', '#FF785A']
    const nextIndex = function () {
      const currentIndex = textColors.includes(elm.value)
        ? textColors.indexOf(elm.value)
        : 0
      if (currentIndex === (textColors.length - 1)) {
        return 0
      } else {
        return currentIndex + 1
      }
    }()
    elm.value = textColors[nextIndex]
    e.target.style.backgroundColor = textColors[nextIndex]
  })
}
