$(".process-design-btn").click(function () {
  var dataInfo = $(this).data("design-id");
  $("#request-progress").show();
  $.post(`/upload/design/${dataInfo}/process`, function (data, status) {
    $("#request-progress").hide();
  });
});
