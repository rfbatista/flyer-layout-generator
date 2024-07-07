import { FabricObject } from "fabric";
import { LayoutElement } from "./layout_element";
import { Point } from "./point";
import { Container } from "./container";

type Props = {
  id: number;
  element: LayoutElement;
  canvasPosition: Point;
  addedOrder: number;
  origin: Point;
  type?: string;
  object: FabricObject;
};

export class Layer {
  private p: Props;
  moved: boolean = false;
  scaled: boolean = false;
  initCanvasPosition: Point;
  initContainer: Container;
  currentContainer: Container;
  constructor(p: Props) {
    this.p = p;
    this.initContainer = structuredClone(
      Container.create({
        width: p.element.width,
        height: p.element.height,
      }),
    );
    this.currentContainer = structuredClone(
      Container.create({
        width: p.element.width,
        height: p.element.height,
      }),
    );
    this.initCanvasPosition = structuredClone(p.canvasPosition);
  }

  static create(p: Props) {
    return new Layer(p);
  }

  get id() {
    return this.p.id;
  }

  get x() {
    return this.p.canvasPosition.x;
  }

  get layoutPosition() {
    return new Point(
      this.p.canvasPosition.x - this.p.origin.x,
      this.p.canvasPosition.y - this.p.origin.y,
    );
  }

  get y() {
    return this.p.canvasPosition.y;
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

  setPosition(x: number, y: number) {
    console.log({
      origin: this.p.origin,
      canvasInitialPosition: this.initCanvasPosition,
      layoutInitialPosition: this.p.element.position,
      newCanvasPosition: new Point(x, y),
    });
    this.moved = true;
    this.p.canvasPosition.x = x;
    this.p.canvasPosition.y = y;
  }

  positionDTO() {
    return {
      OriginalX: this.initCanvasPosition.x,
      OriginalY: this.initCanvasPosition.y,
      position: {
        X: Math.floor(this.layoutPosition.x),
        Y: Math.floor(this.layoutPosition.y),
      },
    };
  }

  setNewSize(width: number, height: number) {
    this.scaled = true;
    this.currentContainer = Container.create({
      width: Math.ceil(width),
      height: Math.ceil(height),
    });
    // console.log({ init: this.initContainer, cur: this.currentContainer });
  }

  print() {
    console.log({
      init: this.initContainer,
      current: this.currentContainer,
    });
  }
}
