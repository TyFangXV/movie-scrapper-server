import { Response } from 'express'
import ytdl from 'ytdl-core'
import fs from 'fs'
import path from 'path'
import settings from '../../../setting'
const { exec } = require("child_process");
import Downloader from 'node-downloader-helper'
import { thumbnailDownloader } from '../extra/meta-data'
import getVideoDataPoints, { audioDownloader } from '../extra/dataExtractor'

const convertBase64 = (path:string) => {
    // read binary data from file
    const bitmap = fs.readFileSync(path);
    // convert the binary data to base64 encoded string
    return bitmap.toString('base64');
};




const soundDownloader = async(songID: string, res: Response) => {
    try {
        const dataPoints = await getVideoDataPoints(songID) as any;
        const downloadedAudioPath = `${settings.path}/temp/${songID}.mp3`;
        if(dataPoints)
        {
            console.log("downloading song");
            
            const downloader = Downloader.DownloaderHelper


            console.log("done downloading song");
            
            if(downloadedAudioPath)
            {
                const pathToSong = downloadedAudioPath;
                //download the thumbnail of the video     
                const thumbnailPath = await thumbnailDownloader(songID);
                    exec(`ffmpeg -i ${pathToSong} -id3v2_version 3 -write_id3v1 1 ${path.join(settings.path + "/temp/out.mp3")}`, (err: any, _: any) => {
                        if(err) throw Error(err)
                        //edit the meta data of the song                                        
                        exec(`ffmpeg -i ${path.join(settings.path + "/temp/out.mp3")} -i ${thumbnailPath} -c copy -map 0 -map 1 -metadata:s:v title="Album cover" -metadata:s:v comment="Cover (Front)" ${path.join(settings.path + '/temp/vido.mp3')}`, (err: any, _: any)=> {
                            if(err) throw Error(err)
                            res.send({
                                status: "Done",
                                data: convertBase64(path.join(settings.path + '/temp/vido.mp3'))
                            });
        
                            //delete the file
/*                            fs.unlinkSync(path.join(settings.path + '/temp/vido.mp3'))
                            fs.unlinkSync(path.join(settings.path + "/temp/out.mp3"))
                            fs.unlinkSync(pathToSong);
                            fs.unlinkSync(thumbnailPath);*/
                        })
                    })
            }else{
                res.send({status : 'err', data : "Couldnt download the audio"})
            }
        }else{
            res.send({status : "err", data : "Couldnt get the data point"})
        }

    } catch (error) {
        res.send({ status: "err", data: error })
    }
}

export default soundDownloader;