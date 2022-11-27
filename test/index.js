const axios = require('axios')
const fs = require('fs')

const main = async() => {
    try {
            var startTime = performance.now()

            const {data} = await axios.get('http://127.0.0.1:3030/song/1QPH_kYHnso');


            var endTime = performance.now()
            fs.writeFileSync("vid.mp3", data.data,  'base64')
            console.log(`Call to doSomething took ${endTime - startTime} milliseconds`)
            
    } catch (error) {
            console.log(error)       
    }
    
}


main()