import { Canvas, Image } from "fabric";
import { LayoutElement } from "./entities/layout_element";
import { Layout } from "./entities/layout";
import { Point } from "./entities/point";

export class LayoutManager {
  editor: Canvas;
  origin: Point;

  constructor(c: Canvas) {
    this.editor = c;
    this.origin = new Point(0, 0);
  }

  setOrigin(p: Point) {
    this.origin = p;
  }

  async drawLayout(layout: Layout) {
    const elementsSorted = layout.elements.sort((a, b) => {
      return a.level - b.level;
    });
    for (const el of elementsSorted) {
      console.log(el.level);
      await this.drawElement(el);
    }
    console.log(this.editor.toObject());
  }

  setCanvas(c: Canvas) {
    this.editor = c;
  }

  drawComponent() {}

  async drawElement(l: LayoutElement) {
    const i = await Image.fromURL(l.imageURL);
    i.set({
      left: l.left + this.origin.x,
      top: l.top + this.origin.y,
      element: l,
      id: l.id,
      selectable: true,
    });
    this.editor.add(i);
  }
}
