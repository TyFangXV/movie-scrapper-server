import https from 'https'
import fs from 'fs'
import pathM from 'path'
import settings from '../../../setting';

export const thumbnailDownloader= (id:string) => {
    const filepath = `${settings.path}/temp/${id}.png`
    return new Promise<string>((resolve, reject) => {
        https.get(`https://i.ytimg.com/vi_webp/${id}/maxresdefault.webp`, (res) => {
            if (res.statusCode === 200) {
                res.pipe(fs.createWriteStream(filepath))
                    .on('error', reject)
                    .once('close', () => resolve(filepath))
                    .on('finish', () =>  {return filepath})
            } else {
                // Consume response data to free up memory
                res.resume();
                reject(new Error(`Request Failed With a Status Code: ${res.statusCode}`));

            }
        });
    });
}