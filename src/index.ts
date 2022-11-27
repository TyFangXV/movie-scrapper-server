import express from 'express'
import ytdl from 'ytdl-core';
import soundDownloader from './utils/download/song';
import getVideoDataPoints from './utils/extra/dataExtractor';


const app = express();

app.get('/song/:videoID', (req, res) => {
    const songID = req.params.videoID;

    if (songID) {
        //download the video, convert it to buffer and send it as a base64 then delete the video
        soundDownloader(songID, res)
    } else {
        res.status(400).send({ msg: "No id provided" })
    }

})


app.get("/data/song/:id", async(req, res) => {
    const id = req.params.id;

    if(id)
    {
        res.send(await getVideoDataPoints(id))
    }else{
        res.send("No")
    }
})


app.listen(3030, () => console.log("Server up"))