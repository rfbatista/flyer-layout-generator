import { Canvas, Image } from "fabric";
import { LayoutElement } from "./entities/layout_element";
import { Layout } from "./entities/layout";

export class LayoutManager {
  editor: Canvas;

  constructor(c: Canvas) {
    this.editor = c;
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
      left: l.left,
      top: l.top,
      element: l,
      id: l.id,
      selectable: true,
    });
    this.editor.add(i);
  }
}
