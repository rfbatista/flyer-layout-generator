import { appConfig } from "../config";

export class DesignImage {
  props: any;
  constructor(props: any) {
    this.props = props;
  }

  get src() {
    return appConfig.api.baeURL + this.props.image_url;
  }

  static create(res: any): DesignImage {
    return new DesignImage(res.data.data);
  }

  static from_api_list(res: any): DesignImage[] {
    return res.data.map((d: any) => new DesignImage(d));
  }
}
