$(document).ready(function () {
  $("#group-items-btn").click(async function GroupItems() {
    console.log("teste");
    if (!canvas.getActiveObject()) {
      return;
    }
    const formData = new FormData();
    const o = canvas.getActiveObject();
    if (o._objects) {
      for (const e of o._objects) {
        const id = e.get("element_id");
        formData.append("elements[]", id);
      }
    } else {
      const id = o.get("element_id");
      formData.append("elements[]", id);
    }
    try {
      var componentType = $("#tipo_componente").find(":selected").val();
      formData.append("type", componentType);
      var layoutId = $("#editor-form").attr("data-layout-id");
      const designId = $("#editor-form").attr("data-design-id");
      const response = await fetch(
        `/editor/design/${designId}/layout/${layoutId}/component`,
        {
          method: "POST",
          body: formData,
        },
      );
      if (response.ok) {
        const result = await response.json();
        console.log("Success:", result);
      } else {
        console.error("Error:", response.statusText);
      }
    } catch (error) {
      console.error("Error:", error);
    }
  });
});
