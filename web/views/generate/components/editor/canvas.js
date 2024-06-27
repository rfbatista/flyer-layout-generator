import { calculateProportionalSize } from "./scale";

const data = JSON.parse($("#canvas-container").attr("json"));
const containerWidth = $("#canvas-container").width();
const containerHeight = $("#canvas-container").height();

$("#canvas-editor").css("width", `${data.width}px`);
$("#canvas-editor").css("height", `${data.height}px`);
$("#c").css("width", `${data.width}px`);
$("#c").css("height", `${data.height}px`);

const size = calculateProportionalSize(
  data.width,
  data.height,
  containerWidth,
  containerHeight,
);

var canvas = new fabric.Canvas("c", {});
window.canvas = canvas;

canvas.setHeight(size.height);
canvas.setWidth(size.width);
console.log(data)

for (const c of data.elements) {
  let i = 0;
  fabric.Image.fromURL(c.image, function (oImg) {
    oImg.scaleToWidth(c.width);
    oImg.scaleToHeight(c.height);
    function scaleImageAndReposition(image, scale) {
      console.log(c, scale);
      const originalLeft = c.OuterContainer.UpperLeft.X;
      const originalTop = c.OuterContainer.UpperLeft.Y;
      image.set({
        left: originalLeft * scale,
        top: originalTop * scale,
        element_id: c.id,
        id: c.id,
        name: c.name,
      });
      const neww = scale * c.width;
      image.scaleToWidth(neww);
      canvas.add(oImg);
      i += 1;
      const cid = `camada-${c.id}`;
      var $newButton = $("<sl-button>", {
        text: c.name,
        id: cid,
      });
      $("#editor-form-camadas").append($newButton);
      $($newButton).on("click", function () {
        canvas.getObjects().forEach(function (o) {
          if (o.element_id == c.id) {
            canvas.setActiveObject(o);
            canvas.requestRenderAll();
            console.log("id", o.id);
          }
        });
        console.log("clicked");
      });
    }
    scaleImageAndReposition(oImg, size.scale);
  });
}

export { canvas };
