import { CircleX, Image, Puzzle } from "lucide-react";
import React, { useEffect } from "react";
import { Canvas } from "../components/Canvas";
import { FileBar } from "../components/FileBar";
import { TreeView } from "../components/TreeView";
import { usePhotoshopFiles } from "../store/files";
import { usePhotoshopStore } from "../store/photoshop";

export function UploadFilePage() {
  const mainRef = React.useRef<any>();
  const data = usePhotoshopFiles((d) => ({
    init: d.init,
    elementsSelected: d.elementsSelected,
    activeTree: d.activeTree,
    createComponent: d.createComponent,
    setMainBoardSize: d.setMainBoardSize,
    setBackground: d.setBackground,
  }));
  const { tree, onCreateComponent, onSetBackground, onRemoveComponent } =
    usePhotoshopStore();

  useEffect(() => {
    data.init();
  }, []);

  useEffect(() => {
    mainRef.current &&
      data.setMainBoardSize({
        width: mainRef.current.offsetWidth,
        height: mainRef.current.clientHeight,
      });
    const observer = new ResizeObserver((entries) => {
      data.setMainBoardSize({
        width: entries[0].contentRect.width,
        height: entries[0].contentRect.height,
      });
    });
    observer.observe(mainRef.current);
    return () => mainRef.current && observer.unobserve(mainRef.current);
  }, [mainRef.current, mainRef.current?.offsetWidth]);

  return (
    <>
      <div className="w-full flex flex-col sm:flex-row flex-wrap sm:flex-nowrap flex-grow">
        <div className="w-1/4 max-w-[300px]">
          <FileBar />
        </div>
        <main role="main" className="flex-1" ref={mainRef}>
          <Canvas />
        </main>
        <div className="w-1/4 max-w-[450px] border-gray-200 border px-3 h-full pb-5 pt-3">
          <div className="my-2 flex justify-center items-center">
            <div className="flex">
              <button
                type="button"
                className={`nav-left`}
                onClick={onCreateComponent}
              >
                <Puzzle size={"20px"} className="mr-2" />
              </button>
              <button
                type="button"
                className="nav-mid"
                onClick={onSetBackground}
              >
                <Image size={"20px"} className="mr-2" />
              </button>
              <button
                type="button"
                className="nav-right"
                onClick={onRemoveComponent}
              >
                <CircleX size={"20px"} className="mr-2" />
              </button>
            </div>
          </div>
          <TreeView tree={tree} />
        </div>
      </div>
    </>
  );
}
