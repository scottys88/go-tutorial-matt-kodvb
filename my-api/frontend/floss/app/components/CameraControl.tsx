import { SyntheticEvent, useEffect, useRef, useState } from "react";
import { useWebCam } from "../hooks/useWebCam"

export const CameraControl = () => {
    const { stream, getMedia } = useWebCam();
    let videoSrcRef = useRef<HTMLVideoElement>(null);
    let canvasRef = useRef<HTMLCanvasElement>(null);
    let photoRef = useRef<HTMLImageElement>(null);
    const [videoDimensions, setVideoDimensions] = useState({
        width: 320,
        height: 0
    })
    const [streaming, setStreaming] = useState(false)


    useEffect(() => {
        if (!videoSrcRef.current) {
            return;
        }

        if (stream) {
            videoSrcRef.current.srcObject = stream;
            videoSrcRef.current.play();
        }

    }, [stream, videoSrcRef])

    const handleClick = (e: React.MouseEvent<HTMLButtonElement>) => {
        e.preventDefault();
        getMedia();

    }

    const takePhoto = (e: React.MouseEvent<HTMLButtonElement>) => {

        e.preventDefault();

        if (!canvasRef.current || !stream || !videoSrcRef.current || !photoRef.current) {
            return;
        }


        const context = canvasRef.current?.getContext("2d");

        if (!context) {
            return;
        }

        console.log(canvasRef.current.getContext("2d"))

        canvasRef.current.width = videoDimensions.width;
        canvasRef.current.height = videoDimensions.height;
        canvasRef.current.getContext("2d")?.drawImage(videoSrcRef.current, 0, 0, videoDimensions.width, videoDimensions.height);

        const data = canvasRef.current.toDataURL("image/png");
        photoRef.current.setAttribute("src", data);

    }

    const handleCanPlay = (e: SyntheticEvent<HTMLVideoElement>) => {
        console.log(e.nativeEvent.target);

        if (!videoSrcRef.current || !canvasRef.current) {
            return;
        }

        if (!streaming) {
            const height = (videoSrcRef.current.videoHeight / videoSrcRef.current.videoWidth) * videoDimensions.width;

            setVideoDimensions(prev => ({ ...prev, height: height }))

            videoSrcRef.current.setAttribute("width", `${videoDimensions.width}`);
            videoSrcRef.current.setAttribute("height", `${height}`);
            canvasRef.current.setAttribute("width", `${videoDimensions.width}`);
            canvasRef.current.setAttribute("height", `${height}`);
            setStreaming(true);
        }
    }

    function clearphoto() {
        if(!canvasRef.current || !photoRef.current) {
            return;
        }

        const context = canvasRef.current.getContext("2d");

        if(!context) {
            return;
        }

        context.fillStyle = "#AAA";
        context.fillRect(0, 0, canvasRef.current.width, canvasRef.current.height);
      
        const data = canvasRef.current.toDataURL("image/png");
        photoRef.current.setAttribute("src", data);
      }
      



    return (
        <>
            <div className="camera">
                <video id="video" ref={videoSrcRef} onCanPlay={e => handleCanPlay(e)}>Video stream not available.</video>
                <button id="startbutton" onClick={handleClick}>Start recording</button>
                <button onClick={takePhoto}>Take photo</button>
                <button onClick={clearphoto}>Clear photo</button>
            </div>

            <canvas id="canvas" ref={canvasRef}> </canvas>
            <div className="output">
                <img id="photo" ref={photoRef} alt="The screen capture will appear in this box." />
            </div>
        </>
    )
}