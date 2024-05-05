import { PhotoshopListAPIResult } from "../api/photoshop/listPhotoshopFile";
import { appConfig } from "../config";

export type PhotoshopFileProps = {
  id: number;
  name: string;
  image_url: any;
  image_extension?: string;
  file_url: string;
  width: any;
  height: any;
  created_at: string;
  updated_at: any;
};

export class PhotoshopFile {
  props: PhotoshopFileProps;
  constructor(props: PhotoshopFileProps) {
    this.props = props;
  }

  get id() {
    return this.props.id;
  }

  get filename() {
    return this.props.name;
  }

  get imageUrl() {
    return `${appConfig.api.baeURL}/dist/${this.props.image_url}.png`;
  }

  get filepath() {
    return this.props.file_url;
  }

  get width() {
    return this.props.width;
  }

  toScale(
    maxWidth: number,
    maxHeight: number,
  ): { width: number; height: number } {
    if (this.width > this.height) {
      const scale = maxWidth / this.width;
      return { width: maxWidth, height: this.height * scale };
    } else {
      const scale = maxHeight / this.height;
      return { width: scale * this.width, height: maxHeight };
    }
  }

  get height() {
    return this.props.height;
  }

  static from_api_list(res: PhotoshopListAPIResult) {
    return res.data.map((data: any) => new PhotoshopFile(data));
  }
}

type PhotoshopElementProps = {
  id: number;
  photoshop_id: number;
  name: string;
  layer_id: string;
  text: string;
  xi: number;
  xii: number;
  yi: number;
  yii: number;
  width: number;
  height: number;
  is_group: boolean;
  group_id: number;
  level: number;
  kind: string;
  component_id: any;
  image_url: string;
  image_extension: string;
  created_at: string;
  updated_at: any;
};

export class PhotoshopElement {
  props: PhotoshopElementProps;

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
    return this.props.group_id || "0";
  }

  get isGroup() {
    return this.props.kind === "group";
  }

  get isBackground() {
    return false;
  }

  get layerId() {
    return this.props.layer_id;
  }

  get componentColor() {
    return `#${333}`;
  }

  get componentId() {
    return this.props.component_id;
  }

  get isComponent(): boolean {
    return Boolean(this.props.component_id);
  }

  static from_api_list(res: any) {
    return res.data.data.map((data: any) => new PhotoshopElement(data));
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

  get isBackground() {
    return this.element?.isBackground;
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
    ElementTree._layout("0", root, elements);
    return root;
  }

  static _layout(
    groupId: string,
    tree: ElementTree,
    elements: PhotoshopElement[],
  ) {
    const inLevel = elements.filter((e) => e.groupId === groupId);
    if (inLevel.length === 0) return tree;
    const childrens = inLevel.map((e) => new ElementTree(e));
    childrens.forEach((e) => {
      tree.addChildren(e);
      ElementTree._layout(e.element.layerId, e, elements);
    });
  }
}
