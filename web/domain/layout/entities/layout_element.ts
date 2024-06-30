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
  InnerContainer: {
    UpperLeft: {
      X: number;
      Y: number;
    };
    DownRight: {
      X: number;
      Y: number;
    };
  };
  OuterContainer: {
    UpperLeft: {
      X: number;
      Y: number;
    };
    DownRight: {
      X: number;
      Y: number;
    };
  };
  inner_xi?: number;
  inner_yi?: number;
};

export class LayoutElement {
  private p: LayoutElementProps;
  constructor(p: LayoutElementProps) {
    this.p = p;
  }

  static create(p: LayoutElementProps) {
    return new LayoutElement(p);
  }

  get id() {
    return this.p.id;
  }

  get imageURL() {
    return this.p.image;
  }

  get level(){
    return this.p.level
  }

  get left() {
    return this.p.OuterContainer.UpperLeft.X;
  }

  get top() {
    return this.p.OuterContainer.UpperLeft.Y;
  }
}
