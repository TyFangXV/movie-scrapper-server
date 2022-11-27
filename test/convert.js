const fs = require('fs')

const main = (base)  => {
    fs.writeFileSync("video.mp3", base,  'base64')
}

main('QzpcVXNlcnNcOWZhbmdcRG9jdW1lbnRzXGNvZGluZ1x3ZWJceXQtZG93bmxvYWRlci1zZXJ2ZXJcdGVtcFx2aWRvLm1wMw==')