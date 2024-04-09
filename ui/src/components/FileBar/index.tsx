import { useState } from "react";
import { api } from "../../infra/api";
import React from "react";
import { toast } from "react-toastify";
import { usePhotoshopFiles } from "../../store/files";

export function DocUpload() {
  const [file, setFile] = useState<File>();

  function handleChange(event: React.ChangeEvent<HTMLInputElement>) {
    if (!event.target.files) {
      console.error("nenhum arquivo selecionado");
      return;
    }
    setFile(event.target.files[0]);
  }

  async function handleSubmit(event: any) {
    event.preventDefault();
    if (!file) {
      console.error("arquivo não selecionado");
      return;
    }
    const formData = new FormData();
    formData.append("file", file);
    formData.append("fileName", file.name);
    const config = {
      headers: {
        "content-type": "multipart/form-data",
      },
    };
    await api
      .post("/api/v1/photoshop", formData, config)
      .then((response) => {
        toast.success("Arquivo carregado");
        console.log(response.data);
      })
      .catch((e) => toast.error("Falha ao carregar arquivo"));
  }
  return (
    <form onSubmit={handleSubmit}>
      <input
        className="block w-full text-sm text-gray-900 border border-gray-300 rounded-lg cursor-pointer bg-gray-50 dark:text-gray-400 focus:outline-none dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400"
        onChange={handleChange}
        type="file"
      />
      <button
        type="submit"
        className="w-full my-2 py-2.5 px-5 me-2 mb-2 text-sm font-medium text-gray-900 focus:outline-none bg-white rounded-lg border border-gray-200 hover:bg-gray-100 hover:text-blue-700 focus:z-10 focus:ring-4 focus:ring-gray-100 dark:focus:ring-gray-700 dark:bg-gray-800 dark:text-gray-400 dark:border-gray-600 dark:hover:text-white dark:hover:bg-gray-700"
      >
        Carregar arquivo
      </button>
    </form>
  );
}

export function DocList() {
  const { files, getImages } = usePhotoshopFiles((data) => ({
    files: data.files,
    getImages: data.getImageFromPsd,
  }));
  return (
    <div>
      {files.map((d) => {
        return (
          <div key={d.id} className="mt-2">
            <button
              type="button"
              className="w-full text-gray-900 bg-white border border-gray-300 focus:outline-none hover:bg-gray-100 focus:ring-4 focus:ring-gray-100 font-medium rounded-lg text-sm px-5 py-2.5 me-2 mb-2 dark:bg-gray-800 dark:text-white dark:border-gray-600 dark:hover:bg-gray-700 dark:hover:border-gray-600 dark:focus:ring-gray-700"
              onClick={() => getImages(d)}
            >
              {d.filename}
            </button>
          </div>
        );
      })}
    </div>
  );
}

export function FileBar() {
  return (
    <>
      <div className="w-full max-w-72 border-gray-200 border px-3 h-screen pt-3">
        <DocUpload />
        <hr />
        <DocList />
      </div>
    </>
  );
}
