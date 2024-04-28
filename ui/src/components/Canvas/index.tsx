import { useEffect, useRef, useState } from "react";
import { usePhotoshopFiles } from "../../store/files";

export function Canvas() {
  const { photoshopfile, boardSize } = usePhotoshopFiles((data) => ({
    photoshopfile: data.activeFile,
    boardSize: data.mainBoardSize,
  }));

  const canvasRef = useRef(null);
  useEffect(() => {
    const canvas = canvasRef.current;
    if (!canvas) return;
    const image = new Image();
    const context = canvas.getContext("2d");
    image.src = photoshopfile?.imageUrl || "";
    console.log(photoshopfile?.imageUrl);
    image.onload = () => {
      context.drawImage(
        image,
        0,
        0,
        photoshopfile?.toScale(boardSize.width, boardSize.height)?.width,
        photoshopfile?.toScale(boardSize.width, boardSize.height)?.height,
      );
    };
  }, [photoshopfile]);

  return (
    <canvas
      ref={canvasRef}
      width={boardSize.width}
      height={boardSize.height}
    ></canvas>
  );
}
