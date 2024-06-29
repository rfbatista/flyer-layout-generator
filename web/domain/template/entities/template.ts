type Props = {
  id: number;
  name: string;
  width: number;
  height: number;
  distortion: {};
  created_at: string;
};

export class Template {
  private p: Props;

  constructor(p: Props) {
    this.p = p;
  }

  static create = (p: Props): Template => {
    return new Template(p);
  };

  get id() {
    return this.p.id;
  }

  get name() {
    return this.p.name;
  }

  get width() {
    return this.p.width;
  }

  get height() {
    return this.p.height;
  }
}
