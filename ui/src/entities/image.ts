export class PhotoshopImage {
  props: any;
  constructor(props: any) {
    this.props = props;
  }

  get src(){
    return "http://localhost:8080" + this.props.image
  }

  static from_api_list(res: any): PhotoshopImage[] {
    return res.data.map((d: any) => new PhotoshopImage(d));
  }
}
