type Props = {
  id: number
  name: string
  clientName: string
  advertiserName: string
  createdAt: Date
}

export class Project {
  private props: Props

  constructor(props: Props) {
    this.props = props
  }

  static create(props: Props) {
    return new Project(props)
  }

  get name() {
    return this.props.name
  }

  get clientName() {
    return this.props.clientName
  }

  get advertiserName() {
    return this.props.advertiserName
  }

  get createdAt(): string {
    return this.createdAt.toString()
  }
}
