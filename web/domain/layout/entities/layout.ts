import { LayoutComponentProps } from "./layout_component";
import { LayoutElement, LayoutElementProps } from "./layout_element";

export type LayoutProps = {
  id: number;
  image_url: string;
  design_id: number;
  width: number;
  height: number;
  total?: number;
  done?: number;
  elements?: Array<LayoutElementProps>;
  components?: Array<LayoutComponentProps>;
  background?: LayoutComponentProps
  template?: {
    distortion: {};
    created_at: string;
  };
  grid?: {
    regions: any;
    SlotsX: number;
    SlotsY: number;
  };
  config?: {
    grid?: {
      regions?: any;
      SlotsX?: number;
      SlotsY?: number;
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
    const elements: LayoutElement[] = [];
    if (p.elements)
      for (const element of p.elements) {
        elements.push(LayoutElement.create(element));
      }

    if (p.components) {
      for (const c of p.components) {
        c.elements.forEach((e) => {
          elements.forEach((el) => {
            if (el.id === e.id) {
              el.setType(c.type);
            }
          });
        });
      }
    }

    return new Layout(p, elements);
  }

  get components(){
    return this.p.components || []
  }

  get bg(){
    return this.p.background
  }

  get id() {
    return this.p.id;
  }

  get designID() {
    return this.p.design_id;
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

  get width() {
    return this.p.width;
  }

  get height() {
    return this.p.height;
  }
}
