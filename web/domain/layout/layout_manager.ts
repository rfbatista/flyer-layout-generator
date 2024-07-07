import { Canvas, Image } from "fabric";
import { LayoutElement } from "./entities/layout_element";
import { Layout } from "./entities/layout";
import { Point } from "./entities/point";
import { Layer } from "./entities/layer";

export class LayoutManager {
  editor: Canvas;
  origin: Point;
  addLayerHook: (l: Layer) => void;
  elementsOrder: number = 0;

  constructor(c: Canvas, hook: (l: Layer) => void) {
    this.editor = c;
    this.origin = new Point(0, 0);
    this.addLayerHook = hook;
  }

  setOrigin(p: Point) {
    this.origin = p;
  }

  async drawLayout(layout: Layout) {
    const elementsSorted = layout.elements.sort((a, b) => {
      return a.level - b.level;
    });
    for (const el of elementsSorted) {
    }
  }

  setCanvas(c: Canvas) {
    this.editor = c;
  }

  drawComponent() { }

}
