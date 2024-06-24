var isVisible = false;
$("#editor-switch").on("sl-change", function () {
  isVisible = !isVisible;
  if (isVisible) {
    $("#editor-form-container").show();
    $("#layout-image").hide();
    $("#canvas-editor-container").show();
  } else {
    $("#editor-form-container").hide();
    $("#canvas-editor-container").hide();
    $("#layout-image").show();
  }
});
