$(document).ready(function () {
  $("#priority-list").sortable();
  $("#gen-btn").on("click", function (e) {
    orderPriorities();
    const form = document.getElementById("generate-request-form");
    const formData = new FormData(form);
    $.ajax({
      url: "/editor/create/image",
      type: "POST",
      data: formData,
      processData: false, // Prevent jQuery from automatically transforming the data into a query string
      contentType: false, // Set contentType to false for FormData
      success: function (result) {
        console.log("foi!", result);
        $("#canvas-container").html(result); // this will replace when you select the checkbox
      },
    });
  });

  $("#batch-btn").on("click", function () {
    orderPriorities();
    submitForm("/request/batch");
  });

  function submitForm(action) {
    const form = document.getElementById("generate-request-form");
    const formData = new FormData(form);

    fetch(action, {
      method: "POST",
      body: formData,
    })
      .then((response) => {
        console.log(response);
      })
      .then((data) => {
        console.log("Success:", data);
      })
      .catch((error) => {
        console.error("Error:", error);
      });
  }
});

function sort() {
  return {
    config: {
      animation: 150,
      ghostClass: "opacity-20",
      dragClass: "bg-blue-50",
    },
    init() {
      Sortable.create(this.$refs.items, this.config);
    },
  };
}
const Toast = Swal.mixin({
  toast: true,
  position: "center",
  iconColor: "white",
  customClass: {
    popup: "colored-toast",
  },
  showConfirmButton: false,
  timer: 1500,
  timerProgressBar: true,
});

document.body.addEventListener("makeToast", async function (evt) {
  if (evt.detail.level == "success") {
    await Toast.fire({
      icon: "success",
      title: "Sucesso",
    });
  } else {
    await Toast.fire({
      icon: "error",
      title: evt.detail.message,
    });
  }
});

document
  .getElementById("generate-request-form")
  .addEventListener("submit", function (event) {
    Array.from(document.getElementsByClassName("priority-item")).forEach(
      (i) => {
        const text = i.querySelector("#priority").textContent;
        i.querySelector("#hiddenInput").value = text;
      },
    );
  });

function orderPriorities() {
  Array.from(document.getElementsByClassName("priority-item")).forEach((i) => {
    const text = i.querySelector("#priority").textContent;
    i.querySelector("#hiddenInput").value = text;
  });
}
