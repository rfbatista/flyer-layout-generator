import { ContainerProps } from "./container";
import { Point } from "./point";

export type LayoutElementProps = {
  id: number;
  xi: number;
  xii: number;
  yi: number;
  yii: number;
  inner_xii: number;
  inner_yii: number;
  layer_id: string;
  width: number;
  height: number;
  kind: string;
  name: string;
  is_group: boolean;
  group_id: number;
  level: number;
  photoshop_id: number;
  image: string;
  component_id: number;
  inner_container: ContainerProps;
  outer_container: ContainerProps;
  inner_xi?: number;
  inner_yi?: number;
  type?: string;
};

export class LayoutElement {
  private p: LayoutElementProps;
  constructor(p: LayoutElementProps) {
    this.p = p;
  }

  static create(p: LayoutElementProps) {
    return new LayoutElement(p);
  }

  setType(t: string) {
    this.p.type = t;
  }

  get type() {
    return this.p.type;
  }

  get id() {
    return this.p.id;
  }

  get width() {
    return (
      this.p.outer_container.down_right.X - this.p.outer_container.upper_left.X
    );
  }

  get height() {
    return (
      this.p.outer_container.down_right.Y - this.p.outer_container.upper_left.Y
    );
  }

  get imageURL() {
    return this.p.image;
  }

  get level() {
    return this.p.level;
  }

  get left() {
    console.log(this);
    return this.p.outer_container.upper_left.X;
  }

  get top() {
    return this.p.outer_container.upper_left.Y;
  }

  get name() {
    return this.p.name;
  }

  get position() {
    return new Point(
      this.p.outer_container.upper_left.X,
      this.p.outer_container.upper_left.Y,
    );
  }

  move(x: number, y: number) {
    this.p.outer_container.down_right.X += x;
    this.p.outer_container.down_right.Y += y;
    this.p.outer_container.upper_left.X += x;
    this.p.outer_container.upper_left.Y += y;

    this.p.inner_container.down_right.X += x;
    this.p.inner_container.down_right.Y += y;
    this.p.inner_container.upper_left.X += x;
    this.p.inner_container.upper_left.Y += y;
  }
}
