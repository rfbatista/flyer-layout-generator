import { Canvas } from "fabric";
import { create } from "zustand";

type Props = {
  editor?: Canvas;
  isReady: boolean;
  setEditor: (f: Canvas) => void;
};

const useEditorStore = create<Props>((set) => ({
  editor: undefined,
  isReady: false,
  setEditor: (f: Canvas) => {
    console.log("set");
    set({ editor: f, isReady: true });
  },
}));

export { useEditorStore };
