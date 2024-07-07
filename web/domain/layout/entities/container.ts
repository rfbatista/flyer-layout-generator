export type ContainerProps = {
  upper_left: {
    X: number;
    Y: number;
  };
  down_right: {
    X: number;
    Y: number;
  };
};

type Props = {
  width: number;
  height: number;
};

export class Container {
  private p: Props;

  constructor(p: Props) {
    this.p = p;
  }

  static create(p: Props): Container {
    return new Container(p);
  }

  get witdth() {
    return this.p.width;
  }

  get height() {
    return this.p.height;
  }
}
