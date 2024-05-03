import React from "react";
import ReactDOM from "react-dom/client";
import "react-toastify/dist/ReactToastify.css";
import "./index.css";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import { ToastContainer } from "react-toastify";
import { UploadFilePage } from "./pages/Upload.tsx";
import { NavBar } from "./components/Navbar/index.tsx";
import { appConfig } from "./config.ts";
import { ListTemplatePage } from "./pages/ListTemplates.tsx";
import { CreateTemplatePage } from "./pages/CreateTemplate.tsx";
import { GeneratePartPage } from "./pages/generate/GeneratePart.tsx";
import { GeneratedImagesPage } from "./pages/GeneratedImages.tsx";

const router = createBrowserRouter([
  {
    path: appConfig.paths.gerador,
    element: (
      <>
        <header className="top-0 flex items-center py-2">
          <NavBar />
        </header>
        <UploadFilePage />
      </>
    ),
  },
  {
    path: appConfig.paths.templates,
    element: (
      <>
        <header className="top-0 flex items-center py-2">
          <NavBar />
        </header>
        <ListTemplatePage />
      </>
    ),
  },
  {
    path: appConfig.paths.createTemplates,
    element: (
      <>
        <header className="top-0 flex items-center py-2">
          <NavBar />
        </header>
        <CreateTemplatePage />
      </>
    ),
  },
  {
    path: appConfig.paths.generateParts,
    element: (
      <>
        <header className="top-0 flex items-center py-2">
          <NavBar />
        </header>
        <GeneratePartPage />
      </>
    ),
  },
  {
    path: appConfig.paths.designs,
    element: (
      <>
        <header className="top-0 flex items-center py-2">
          <NavBar />
        </header>
        <GeneratedImagesPage />
      </>
    ),
  },
]);

ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <div>
      <RouterProvider router={router} />
      <ToastContainer position="bottom-center" />
    </div>
  </React.StrictMode>,
);
