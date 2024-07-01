export type LayoutRequestJobProps = {
  id: number;
  design_id: number;
  layout_id?: number;
  request_id: number;
  template_id: number;
  created_at: string;
  started_at: string;
  config: {
    minimium_component_size: number;
    minimium_text_size: number;
    slots_x: number;
    slots_y: number;
    grid: {
      regions: any;
      SlotsX: number;
      SlotsY: number;
    };
    padding: number;
    priorities: {
      celebridade: number;
      cta: number;
      grafismo: number;
      ilustracao: number;
      marca: number;
      modelo: number;
      oferta: number;
      packshot: number;
      "plano-de-fundo": number;
      produto: number;
      "texto-legal": number;
    };
  };
  error_at: string;
  log: string;
};

export class LayoutRequestJob {
  private p: LayoutRequestJobProps;
  constructor(p: LayoutRequestJobProps) {
    this.p = p;
  }

  static create(p: LayoutRequestJobProps) {
    return new LayoutRequestJob(p);
  }

  get id() {
    return this.p.id;
  }

  get layoutID(){
    return this.p.layout_id
  }
}
