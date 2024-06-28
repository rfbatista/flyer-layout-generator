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
  constructor(p: Props) {
    this.p = p;
  }

  static create(p: Props): Design {
    return new Design(p);
  }

  get id() {
    return this.p.id;
  }

  get imageURL() {
    return this.p.image_path;
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
