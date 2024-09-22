import { useEffect, useRef } from "react";
import { useWebCam } from "../hooks/useWebCam"

export const CameraControl = () => {
    const { stream, getMedia } = useWebCam();
    let videoSrcRef = useRef<HTMLVideoElement>(null);
    let canvasRef = useRef<HTMLCanvasElement>(null);
    let photoRef = useRef<HTMLImageElement>(null);
    const width = 320;    // We will scale the photo width to this
    const height = 100;     // This will be computed based on the input stream


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

        canvasRef.current.width = width;
        canvasRef.current.height = height;
        canvasRef.current.getContext("2d")?.drawImage(videoSrcRef.current, 0, 0, width, height);

        const data = canvasRef.current.toDataURL("image/png");
        photoRef.current.setAttribute("src", data);

    }



    return (
        <>
            <div className="camera">
                <video id="video" ref={videoSrcRef}>Video stream not available.</video>
                <button id="startbutton" onClick={handleClick}>Start recording</button>
                <button onClick={takePhoto}>Take photo</button>
            </div>

            <canvas id="canvas" ref={canvasRef}> </canvas>
            <div className="output">
                <img id="photo" ref={photoRef} alt="The screen capture will appear in this box." />
            </div>
        </>
    )
}