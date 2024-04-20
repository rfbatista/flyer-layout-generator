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

export class PhotoshopElement {
  props: any;

  constructor(props: any) {
    this.props = props;
  }

  get id() {
    return this.props.id;
  }

  get name() {
    return this.props.name;
  }

  get level() {
    return this.props.level;
  }

  get groupId() {
    return this.props.group_id || 0;
  }

  get isGroup() {
    return this.props.kind === "group";
  }

  get isBackground(){
    return this.props.is_background
  }

  get layerId() {
    return this.props.layer_id;
  }

  get componentColor() {
    return `#${this.props.component_color}`;
  }

  get componentId() {
    return this.props.component_id;
  }

  get isComponent(): boolean {
    return Boolean(this.props.component_id);
  }

  static from_api_list(res: any) {
    return res.data.map((data: any) => new PhotoshopElement(data));
  }
}

export class ElementTree {
  children: ElementTree[];
  element?: PhotoshopElement;

  constructor(element?: PhotoshopElement, children = []) {
    this.children = children;
    this.element = element;
  }

  addChildren(e: ElementTree) {
    this.children.push(e);
  }

  get isComponent() {
    return this.element?.isComponent;
  }

  get isBackground(){
    return this.element?.isBackground
  }

  get color() {
    return this.element?.componentColor;
  }

  get isDir() {
    return Boolean(this.element?.isGroup);
  }

  get name() {
    return this.element?.name;
  }

  get id() {
    return this.element?.id;
  }

  static layout(elements: PhotoshopElement[]): ElementTree {
    const root = new ElementTree();
    ElementTree._layout(0, root, elements);
    return root;
  }

  static _layout(
    groupId: number,
    tree: ElementTree,
    elements: PhotoshopElement[],
  ) {
    const inLevel = elements.filter((e) => e.groupId === groupId);
    if (inLevel.length === 0) return tree;
    const childrens = inLevel.map((e) => new ElementTree(e));
    childrens.forEach((e) => {
      tree.addChildren(e);
      ElementTree._layout(e.element?.layerId, e, elements);
    });
  }
}
