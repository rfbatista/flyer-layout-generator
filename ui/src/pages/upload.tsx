import { useEffect } from "react";
import { FileBar } from "../components/FileBar";
import { usePhotoshopFiles } from "../store/files";
import { Canvas } from "../components/Canvas";
import { TreeView } from "../components/TreeView";

export function UploadFilePage() {
  const data = usePhotoshopFiles();
  useEffect(() => {
    data.execute();
  }, []);
  return (
    <>
      <div className="w-full flex flex-col sm:flex-row flex-wrap sm:flex-nowrap flex-grow">
        <div className="w-1/4 max-w-[300px]">
          <FileBar />
        </div>
        <main role="main" className="flex-1">
          <Canvas />
        </main>
        <div className="w-1/4 max-w-[300px]">
          <TreeView />
        </div>
      </div>
    </>
  );
}
