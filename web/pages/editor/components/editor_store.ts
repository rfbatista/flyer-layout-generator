import { Canvas } from "fabric";
import { create } from "zustand";
import { LayoutManager } from "../../../domain/layout/layout_manager";
import { Layer } from "../../../domain/layout/entities/layer";

type Props = {
  editor?: Canvas;
  isReady: boolean;
  setEditor: (f: Canvas) => void;
  isDragging: boolean;
  lastPosX: number;
  lastPosY: number;
  selection: boolean;
  layers: Layer[];
  onMouseDown: (opt: any) => void;
  onMouseMove: (opt: any) => void;
  onMouseUp: () => void;
  addActiveItem: (l: Layer) => void;
  activeItems: Layer[];
  layoutManager?: LayoutManager;
  addLayer: (l: Layer) => void;
  setLayoutManager: (l: LayoutManager) => void;
  onUnselect: () => void;
};

const useEditorStore = create<Props>((set, get) => ({
  isDragging: false,
  selection: true,
  lastPosX: 0,
  lastPosY: 0,
  editor: undefined,
  isReady: false,
  layers: [],
  activeItems: [],
  layoutManager: undefined,
  setEditor: (f: Canvas) => {
    set({ editor: f, isReady: true });
  },
  setLayoutManager: (l: LayoutManager) => {
    set({ layoutManager: l });
  },
  addActiveItem: (l: Layer) => {
    const editor = get().editor;
    if (editor) {
      editor.setActiveObject(l.obj);
      editor.renderAll();
      set((s) => ({ activeItems: [...s.activeItems, l] }));
    }
  },
  onMouseDown: (opt: any) => {
    var evt = opt.e;
    if (evt.altKey === true) {
      const editor = get().editor;
      if (editor) {
        set({
          isDragging: true,
          selection: true,
          lastPosX: evt.clientX,
          lastPosY: evt.clientY,
        });
      }
    }
  },
  addLayer: (l: Layer) => {
    set((s) => ({ layers: [...s.layers, l] }));
  },
  onMouseUp: () => {
    const editor = get().editor;
    if (editor) {
      // on mouse up we want to recalculate new interaction
      // for all objects, so we call setViewportTransform
      editor.setViewportTransform(editor.viewportTransform);
      set({ isDragging: false, selection: true });
    }
  },
  onMouseMove: (opt: any) => {
    const editor = get().editor;
    if (editor) {
      if (get().isDragging) {
        var e = opt.e;
        var vpt = editor.viewportTransform;
        vpt[4] += e.clientX - get().lastPosX;
        vpt[5] += e.clientY - get().lastPosY;
        editor.requestRenderAll();
        set({ lastPosX: e.clientX, lastPosY: e.clientY });
      }
    }
  },
  onUnselect: () => {
    console.log("aqui");
    set({ activeItems: [] });
  },
}));

export { useEditorStore };
