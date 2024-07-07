import { Canvas, Image, Rect } from "fabric";
import { create } from "zustand";
import { Layer } from "../../../domain/layout/entities/layer";
import { Layout } from "../../../domain/layout/entities/layout";
import { LayoutElement } from "../../../domain/layout/entities/layout_element";
import { Point } from "../../../domain/layout/entities/point";
import { LayoutManager } from "../../../domain/layout/layout_manager";
import { updateElementPositionAPI } from "../../../domain/layout/api/update_element_position";
import { updateElementSizeAPI } from "../../../domain/layout/api/update_element_size";

type Props = {
  editor?: Canvas;
  isReady: boolean;
  haveInitiated: boolean;
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
  onObjectMoving: (opt: any) => void;
  activeItems: Layer[];
  layoutManager?: LayoutManager;
  addLayer: (l: Layer) => void;
  setLayoutManager: (l: LayoutManager) => void;
  onUnselect: () => void;
  updateLayers: (l: Layer[]) => void;
  onScaling: (opt: any) => void;
  origin: Point;
  elementsOrder: number;
  save: () => Promise<void>;
  init: (
    editor: Canvas,
    layout: Layout,
    width: number,
    height: number,
  ) => Promise<void>;
  isLoading: boolean;
  drawElement: (l: LayoutElement) => Promise<void>;
};

const useEditorStore = create<Props>((set, get) => ({
  isLoading: false,
  origin: new Point(0, 0),
  elementsOrder: 0,
  isDragging: false,
  selection: true,
  lastPosX: 0,
  lastPosY: 0,
  editor: undefined,
  isReady: false,
  layers: [],
  activeItems: [],
  layoutManager: undefined,
  haveInitiated: false,
  setEditor: (f: Canvas) => {
    set({ editor: f, isReady: true });
  },
  init: async (
    editor: Canvas,
    layout: Layout,
    width: number,
    height: number,
  ) => {
    if (get().haveInitiated) return;
    set({ haveInitiated: true });
    editor.setWidth(width);
    editor.setHeight(height);
    editor.on("mouse:down", get().onMouseDown);
    editor.on("mouse:move", get().onMouseMove);
    editor.on("mouse:up", get().onMouseUp);
    editor.on("object:moving", get().onObjectMoving);
    editor.on("object:scaling", get().onScaling);
    editor.on("before:selection:cleared", get().onUnselect);
    const viewport = new Rect({
      left: 0,
      top: 0,
      width: layout.width,
      height: layout.height,
      selectable: false,
      fill: "white",
      lockMovementX: true,
      lockMovementY: true,
    });
    editor.add(viewport);
    editor.centerObject(viewport);
    editor.zoomToPoint(viewport.getCenterPoint(), 0.5);
    set({ origin: new Point(viewport.left, viewport.top), editor: editor });
    const elementsSorted = layout.elements.sort((a, b) => {
      return a.level - b.level;
    });
    for (const el of elementsSorted) {
      await get().drawElement(el);
    }
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
    set({ activeItems: [] });
  },
  updateLayers: (l: Layer[]) => {
    set({ layers: l });
  },
  onScaling: (opt) => {
    const idx = get().layers.findIndex((l) => l.id === opt.target.id);
    const layers = get().layers;
    const rect = opt.target.getBoundingRect();
    layers[idx].setNewSize(rect.width, rect.height);
    layers[idx].setPosition(opt.target.left, opt.target.top);
    set({ layers });
  },
  onObjectMoving: (opt) => {
    // console.log(opt)
    const idx = get().layers.findIndex((l) => l.id === opt.target.id);
    const layers = get().layers;
    layers[idx].setPosition(opt.target.left, opt.target.top);
    set({ layers });
  },
  drawElement: async (l: LayoutElement): Promise<void> => {
    const editor = get().editor;
    if (!editor) return;
    const i = await Image.fromURL(l.imageURL);
    const opts = {
      left: l.left + get().origin.x,
      top: l.top + get().origin.y,
      element: l,
      id: l.id,
      selectable: true,
      order: get().elementsOrder,
      lockSkewingX: true,
      lockSkewingY: true,
      lockRotation: true,
      lockScalingFlip: true,
      lockScalingY: false,
    };
    i.scaleToWidth(l.width);
    i.scaleToHeight(l.height);
    i.set(opts);
    editor.add(i);
    set({ elementsOrder: get().elementsOrder + 1 });
    get().addLayer(
      Layer.create({
        element: l,
        origin: get().origin,
        id: opts.id,
        canvasPosition: new Point(opts.left, opts.top),
        addedOrder: get().elementsOrder,
        object: i,
      }),
    );
  },
  save: async () => {
    try {
      set({ isLoading: true });
      const layers = get().layers;
      for (const l of layers) {
        if (l.moved) {
          await updateElementPositionAPI(1, l.id, l.positionDTO());
        }
        if (l.scaled) {
          l.print()
          await updateElementSizeAPI(1, l.id, {
            width: l.currentContainer.witdth,
            height: l.currentContainer.height,
          });
        }
      }
      set({ isLoading: false });
    } catch (e) {
      set({ isLoading: false });
      console.error(e);
    }
  },
}));

export { useEditorStore };
