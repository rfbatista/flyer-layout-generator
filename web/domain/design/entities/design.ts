import { Layout, LayoutProps } from "../../layout/entities/layout";

type Props = {
  id: number;
  name: string;
  filepath: string;
  layout_id: number;
  image_path: string;
  image_url: string;
  width: number;
  height: number;
  created_at: string;
  project_id: number;
  is_processed: boolean;
};

export class Design {
  private p: Props;
  private _layout?: Layout;
  constructor(p: Props, layout?: Layout) {
    this.p = p;
    this._layout = layout;
  }

  static create(p: Props & { layout?: LayoutProps }): Design {
    if (p.layout) {
      return new Design(p, Layout.create(p.layout));
    }
    return new Design(p);
  }

  get id() {
    return this.p.id;
  }

  get imageURL() {
    return this.p.image_path;
  }

  get layout(): Layout | undefined {
    return this._layout;
  }

  get name() {
    return this.p.name;
  }

  get isProcessed(): boolean {
    return this.p.is_processed;
  }

  get projectId(): number {
    return this.p.project_id;
  }
}
