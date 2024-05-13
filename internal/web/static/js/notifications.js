document.body.addEventListener("request-notification", function (evt) {
  if (evt.detail.level === "info") {
    alert(evt.detail.message);
  }
})
