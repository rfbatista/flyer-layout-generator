type Props = {
  id: number;
  photoshop_id: number;
  width: number;
  height: number;
  color: any;
  xi: number;
  xii: number;
  yi: number;
  yii: number;
  created_at: string;
};

export class Component {
  props: Props;
  constructor(props: Props) {
    this.props = props;
  }

  id() {
    return this.props.id;
  }

  color() {
    return this.props.color;
  }
}
