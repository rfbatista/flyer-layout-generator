import { FabricObject } from "fabric";
import { LayoutElement } from "./layout_element";
import { Point } from "./point";

type Props = {
  id: number;
  element: LayoutElement;
  currentPosition: Point;
  addedOrder: number;
  type?: string;
  object: FabricObject;
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

  get type() {
    return this.p.element.type;
  }

  get name() {
    return this.p.element.name;
  }

  get addedOrder() {
    return this.p.addedOrder;
  }

  get obj() {
    return this.p.object;
  }
}
