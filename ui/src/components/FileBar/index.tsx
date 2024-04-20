import { useState } from "react";
import { api } from "../../infra/api";
import React from "react";
import { toast } from "react-toastify";
import { usePhotoshopFiles } from "../../store/files";
import { Button } from "flowbite-react";
import { useNavigate } from "react-router-dom";
import { appConfig } from "../../config";

export function DocUpload() {
  const [file, setFile] = useState<File>();
  const [isLoading, setLoading] = useState<boolean>(false);
  const data = usePhotoshopFiles();

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
    setLoading(true);
    await api
      .post("/api/v1/photoshop", formData, config)
      .then(() => {
        toast.success("Arquivo carregado");
        data.execute();
      })
      .catch(() => toast.error("Falha ao carregar arquivo"));
    setLoading(false);
  }
  return (
    <form onSubmit={handleSubmit} className="mt-2">
      <input
        className="block w-full text-sm text-gray-900 border border-gray-300 rounded-lg cursor-pointer bg-gray-50 dark:text-gray-400 focus:outline-none dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400"
        onChange={handleChange}
        type="file"
      />
      <Button
        isProcessing={isLoading}
        className="px-2 w-full my-2 hover:bg-gray-50"
        type="submit"
      >
        Carregar arquivo
      </Button>
    </form>
  );
}

export function DocList() {
  const { files, selectPhotoshop } = usePhotoshopFiles((data) => ({
    files: data.files,
    selectPhotoshop: data.selectPhotoshop,
  }));
  return (
    <div>
      {files.map((d) => {
        return (
          <div key={d.id} className="mt-2">
            <button
              type="button"
              className="w-full text-gray-900 bg-white border border-gray-300 focus:outline-none hover:bg-gray-100 focus:ring-4 focus:ring-gray-100 font-medium rounded-lg text-sm px-5 py-2.5 me-2 mb-2 dark:bg-gray-800 dark:text-white dark:border-gray-600 dark:hover:bg-gray-700 dark:hover:border-gray-600 dark:focus:ring-gray-700"
              onClick={() => selectPhotoshop(d)}
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
  const navigate = useNavigate();
  return (
    <>
      <div className="w-full max-w-72 border-gray-200 border px-3 h-screen pt-3">
        <Button
          className="w-full my-2"
          onClick={() => navigate(appConfig.paths.generateParts)}
        >
          Gerar peças
        </Button>
        <hr />
        <DocUpload />
        <hr />
        <DocList />
      </div>
    </>
  );
}
