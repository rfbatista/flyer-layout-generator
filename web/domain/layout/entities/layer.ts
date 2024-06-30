import { Point } from "fabric";
import { LayoutElement } from "./layout_element";

type Props = {
  id: number;
  element: LayoutElement;
  currentPosition: Point;
};

export class Layer {
  private p: Props;
  constructor(p: Props) {
    this.p = p;
  }

  static create(p: Props) {
    return new Layer(p);
  }

  get id() {
    return this.p.id;
  }
}
