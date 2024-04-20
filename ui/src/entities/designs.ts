import { appConfig } from "../config";

export class Design {
  props: any;
  constructor(props: any) {
    this.props = props;
  }

  get src(){
    return appConfig.api.baeURL + "/static/" + this.filename
  }

  get filename() {
    return this.props.filename;
  }

  get createdAt() {
    return this.props.created_at;
  }

  static from_api_list(res: any): Design[] {
    return res.data.map((d: any) => new Design(d));
  }
}
