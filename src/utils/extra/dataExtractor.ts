import ytdl from "ytdl-core"
import https from 'https'
import settings from "../../../setting";
import fs from 'fs'
//atapoint.player_response.streamingData.adaptiveFormats.filter((d:any) => d.audioQuality === 'AUDIO_QUALITY_MEDIUM')

const getVideoDataPoints = async(videoID:string) => {
    try {
        const datapoint = await ytdl.getInfo(`https://www.youtube.com/watch?v=${videoID}`);
        const videoURL = datapoint.player_response.streamingData.adaptiveFormats
                                .filter((d:any) => d.mimeType.split(" ")[0] === 'audio/webm;')
                                .reduce((prev:any, curr:any) => (prev.averageBitrate > curr.averageBitrate) ? prev : curr)
        const thumbnail = datapoint.videoDetails.thumbnails.pop()
        const author = datapoint.videoDetails.author.name.replace(' - Topic', '');
        const releaseDate = new Date(datapoint.videoDetails.publishDate);
        const title = datapoint.videoDetails.title;

        return {videoURL, thumbnail, author, releaseDate, title}
    } catch (error) {
        return error
    }
}


export async function audioDownloader (id:string, url:string) {  
    const targetFile = `${settings.path}/temp/${id}.m4a`
    return await new Promise((resolve, reject) => {
      https.get(url, response => {
        const code = response.statusCode ?? 0
  
        if (code >= 400) {
          return reject(null)
        }
  
        // handle redirects
        if (code > 300 && code < 400 && !!response.headers.location) {
          return audioDownloader(response.headers.location, targetFile)
        }
  
        // save the file to disk
        const fileWriter = fs
          .createWriteStream(targetFile)
          .on('finish', () => {
            resolve(targetFile)
          })
  
        response.pipe(fileWriter)
      }).on('error', error => {
        reject(null)
      })
    })
  }

export default getVideoDataPoints;