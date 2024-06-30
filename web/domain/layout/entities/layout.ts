import { LayoutElement, LayoutElementProps } from "./layout_element";

export type LayoutProps = {
  id: number;
  image_url: string;
  design_id: number;
  width: number;
  height: number;
  elements?: Array<LayoutElementProps>;
  template: {
    distortion: {};
    created_at: string;
  };
  grid: {
    regions: any;
    SlotsX: number;
    SlotsY: number;
  };
  config: {
    grid: {
      regions: any;
      SlotsX: number;
      SlotsY: number;
    };
  };
};

export class Layout {
  private p: LayoutProps;
  private _elements: LayoutElement[];
  constructor(p: LayoutProps, elements: LayoutElement[] = []) {
    this.p = p;
    this._elements = elements;
  }

  static create(p: LayoutProps): Layout {
    const e: LayoutElement[] = [];
    if (p.elements)
      for (const element of p.elements) {
        e.push(LayoutElement.create(element));
      }
    return new Layout(p, e);
  }

  get id() {
    return this.p.id;
  }

  get imageURL() {
    return this.p.image_url;
  }

  get totalElements() {
    return this.elements.length;
  }

  get elements() {
    return this._elements;
  }
}
