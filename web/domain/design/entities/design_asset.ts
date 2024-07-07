export const DesignAssetType = {
  TEXT: 6,
};

export const DesignAssetPropertyKey = Object.freeze({
  CONTENT: "text",
  FONT_NAME: "font_name",
  FONT_SIZE: "font_size",
});

export interface DesignAssetProps {
  id: number;
  design_id: number;
  project_id: number;
  type: number;
  Width: number;
  Height: number;
  properties?: Property[];
}

interface Property {
  key: string;
  value: string;
}

export class DesignAsset {
  private p: DesignAssetProps;
  constructor(p: DesignAssetProps) {
    this.p = p;
  }

  static create(p: DesignAssetProps): DesignAsset {
    return new DesignAsset(p);
  }

  get id() {
    return this.p.id;
  }

  get desinID(){
    return this.p.design_id
  }

  get type() {
    return this.p.type;
  }

  get text() {
    if (!this.p.properties) return;
    const textProp = this.p.properties.filter(
      (p) => p.key === DesignAssetPropertyKey.CONTENT,
    );
    if (textProp.length === 0) return;
    const text = textProp[0].value;
    return text;
  }

  get fonts(): string[] {
    if (!this.p.properties) return [];
    const fontsName = this.p.properties.filter(
      (p) => p.key === DesignAssetPropertyKey.FONT_NAME,
    );
    return fontsName.map((t) => t.value);
  }

  get fontSize(){
    if (!this.p.properties) return;
    const fontSize = this.p.properties.filter(
      (p) => p.key === DesignAssetPropertyKey.FONT_SIZE,
    );
    if (fontSize.length === 0) return;
    const text = fontSize[0].value;
    return text;
  }
}
