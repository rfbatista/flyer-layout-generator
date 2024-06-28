type Props = {
  id: number;
  client: {
    id: number;
    name: string;
    created_at: string;
    updated_at: string;
    deleteed_at: string;
  };
  advertiser: {
    id: number;
    name: string;
    created_at: string;
    updated_at: string;
    deleteed_at: string;
  };
  name: string;
  created_at: string;
  updated_at: string;
  deleteed_at: string;
};

export class Project {
  private props: Props;

  constructor(props: Props) {
    this.props = props;
  }

  static create(props: Props) {
    return new Project(props);
  }

  get id() {
    return this.props.id;
  }

  get name() {
    return this.props.name;
  }

  get clientName() {
    return this.props.client.name;
  }

  get advertiserName() {
    return this.props.advertiser.name;
  }

  get createdAtText(): string {
    let data = new Date(this.props.created_at);
    let dataFormatada =
      data.getDate() + "/" + (data.getMonth() + 1) + "/" + data.getFullYear();
    return dataFormatada;
  }
}
