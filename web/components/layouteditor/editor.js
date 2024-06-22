import { LitElement, html, css } from "lit";
import fabric from "./fabric.esm";

export class LayoutEditor extends LitElement {
  static properties = {
    components: {},
  };

  static get styles() {
    return css`
      :host {
        display: block;
        width: 970px;
        height: 250px;
      }
    `;
  }

  render() {
    return html`
      <div>My LitElement Component</div>
      <canvas id="c" style="width:970px; height:250px;"></canvas>
    `;
  }

  firstUpdated() {
    console.log(fabric);
    var canvas = new fabric.Canvas("c", {
      preserveObjectStacking: true,
    });
    canvas.setHeight(250);
    canvas.setWidth(970);
    var triangle = new fabric.Triangle({
      width: 20,
      height: 30,
      fill: "blue",
      left: 50,
      top: 50,
    });

    canvas.add(circle, triangle);
    fabric.Image.fromURL(
      "/api/v1/images/3567a8ce-812c-444e-8466-00df83159f2b::bg",
      function (oImg) {
        canvas.add(oImg);
      },
    );
    fabric.Image.fromURL(
      "/api/v1/images/a6d02951-bc7f-474c-9d16-58fbd4631357::logo%20bg",
      function (oImg) {
        canvas.add(oImg);
      },
    );
    fabric.Image.fromURL(
      "/api/v1/images/824ccd8a-b112-4d77-aafd-f0198c269614::logo%20outline",
      function (oImg) {
        canvas.add(oImg);
      },
    );
    fabric.Image.fromURL(
      "/api/v1/images/b37d5fb7-4e14-46fa-9480-fbd368d6ef60::produros",
      function (oImg) {
        canvas.add(oImg);
      },
    );
    fabric.Image.fromURL(
      "/api/v1/images/94a0be4c-f343-4377-887e-4567d13a3156::Camada%201",
      function (oImg) {
        canvas.add(oImg);
      },
    );
    fabric.Image.fromURL(
      "/api/v1/images/67f8066d-faaa-421c-b511-c1974edcfde2::logotipo",
      function (oImg) {
        canvas.add(oImg);
      },
    );
    fabric.Image.fromURL(
      "/api/v1/images/0ea7800c-7fc8-4873-a145-b00774d010c8::Assinatura",
      function (oImg) {
        canvas.add(oImg);
      },
    );
    fabric.Image.fromURL(
      "/api/v1/images/32c43170-ce20-471b-ad90-0f280c4e0ddb::selinho%20copy",
      function (oImg) {
        canvas.add(oImg);
      },
    );
  }
}

customElements.define("layout-editor", LayoutEditor);
