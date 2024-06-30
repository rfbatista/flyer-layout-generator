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
    const promises = [];
    for (const el of elementsSorted) {
      console.log(el.level);
      promises.push(this.drawElement(el));
    }
    await Promise.all(promises);
  }

  setCanvas(c: Canvas) {
    this.editor = c;
  }

  drawComponent() {}

  async drawElement(l: LayoutElement) {
    const i = await Image.fromURL(l.imageURL);
    const opts = {
      left: l.left + this.origin.x,
      top: l.top + this.origin.y,
      element: l,
      id: l.id,
      selectable: false,
      order: this.elementsOrder,
    };
    i.set(opts);
    this.editor.add(i);
    this.elementsOrder += 1;
    this.addLayerHook(
      Layer.create({
        element: l,
        id: opts.id,
        currentPosition: new Point(opts.left, opts.top),
        addedOrder: this.elementsOrder,
        object: i,
      }),
    );
  }
}
