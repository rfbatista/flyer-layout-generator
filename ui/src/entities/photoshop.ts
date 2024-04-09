export class PhotoshopFile {
  props: any;
  constructor(props: any) {
    this.props = props;
  }

  get id() {
    return this.props.id;
  }

  get filename() {
    return this.props.filename;
  }

  get filepath() {
    return this.props.filepath;
  }

  get width() {
    return this.props.width;
  }

  toScale(
    maxWidth: number,
    maxHeight: number,
  ): { width: number; height: number } {
    if (this.width > this.height) {
      const scale = (100 * maxWidth) / this.width / 100;
      return { width: maxWidth, height: this.height * scale };
    } else {
      const scale = (100 * maxHeight) / this.height;
      return { width: scale * this.width, height: maxHeight };
    }
  }

  get height() {
    return this.props.height;
  }

  static from_api_list(res: any) {
    return res.data.map((data: any) => new PhotoshopFile(data));
  }
}
