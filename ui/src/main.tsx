import React from "react";
import ReactDOM from "react-dom/client";
import "react-toastify/dist/ReactToastify.css";
import "./index.css";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import { ToastContainer } from "react-toastify";
import { UploadFilePage } from "./pages/upload.tsx";

const router = createBrowserRouter([
  {
    path: "/",
    element: <UploadFilePage />,
  },
]);

ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <div>
      <RouterProvider router={router} />
      <ToastContainer />
    </div>
  </React.StrictMode>,
);
