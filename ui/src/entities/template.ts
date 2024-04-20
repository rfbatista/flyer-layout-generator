export class TemplatePosition {
  props: any;
  constructor(props: any) {
    this.props = props;
  }
  get xi() {
    return this.props.xi;
  }
  get yi() {
    return this.props.yi;
  }
  get xii() {
    return this.props.xii;
  }
  get yii() {
    return this.props.yii;
  }
  get width() {
    return this.props.width;
  }
  get height() {
    return this.props.height;
  }

  resize(widthRatio, heighRatio) {
    this.props.xi = this.props.xi * widthRatio;
    this.props.yi = this.props.yi * heighRatio;
    this.props.xii = this.props.xii * widthRatio;
    this.props.yii = this.props.yii * heighRatio;
    this.props.width = this.props.xii - this.props.xi
    this.props.height = this.props.yii - this.props.yi
  }
}

type TemplateProps = {
  id: number;
  name: string;
  width: number;
  height: number;
  positions: TemplatePosition[];
};

export class Template {
  props: TemplateProps;
  constructor(props: TemplateProps) {
    this.props = props;
  }
  get id() {
    return this.props.id;
  }
  get name() {
    return this.props.name;
  }
  get width() {
    return this.props.width;
  }
  get height() {
    return this.props.height;
  }
  get positions() {
    return this.props.positions;
  }

  widthRatio(newWidth: number) {
    return newWidth / this.width;
  }

  heightRatio(newHeight: number) {
    return newHeight / this.height;
  }

  toScale(
    maxWidth: number,
    maxHeight: number,
  ): { width: number; height: number } {
    if (this.width > this.height) {
      const newWidth = maxWidth;
      const newHeight = (this.width * newWidth) / this.height;
      return { width: newWidth, height: newHeight };
    } else {
      const newHeight = maxHeight;
      const newWidth = (this.height * newHeight) / this.width;
      return { width: newWidth, height: newHeight };
    }
  }

  static fromApiList(res: any) {
    const templates = [];
    for (const item of res.data) {
      const positions = [];
      for (const pos of item.positions) {
        positions.push(new TemplatePosition(pos));
      }
      templates.push(
        new Template({
          id: item.id,
          name: item.name,
          width: item.width,
          height: item.height,
          positions,
        }),
      );
    }
    return templates;
  }
}
