import { useEffect, useState } from "react"

export const useWebCam = () => {
    const [stream, setStream] = useState<MediaStream>()

    const defaulConstraints = {
        video: { frameRate: { ideal: 10, max: 15 } },
      };



    async function getMedia(constraints = defaulConstraints) {
        let stream: MediaStream;
      
        try {
          stream = await navigator.mediaDevices.getUserMedia(constraints);
          /* use the stream */
          setStream(stream);
        } catch (err) {
          /* handle the error */
        }
      }

    return  {
        stream,
        getMedia,

    }
}