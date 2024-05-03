import { Label, TextInput } from "flowbite-react";

export const SlotsCanvasInputControls = () => {
  return (
    <div>
      <Label htmlFor="name" value="Dimensões" />
      <div className="grid grid-cols-2 gap-x-3">
        <div>
          <div className="mb-2 block">
            <Label htmlFor="name" value="Largura" />
          </div>
          <TextInput
            name="width"
            id="width"
            type="number"
            placeholder=""
            required
            onChange={handleWidthInputChange}
            value={canvasState.width}
          />
        </div>
        <div>
          <div className="mb-2 block">
            <Label htmlFor="name" value="Altura" />
          </div>
          <TextInput
            id="height"
            name="height"
            type="number"
            placeholder=""
            required
            onChange={handleHeightInputChange}
            value={canvasState.height}
          />
        </div>
      </div>
    </div>
  );
};

export const SlotsCanvasActionsControls = () => {
  return (
    <div className="flex justify-center items-center mt-5">
      <div className="my-2 flex justify-center items-center">
        <div className="flex">
          <button
            type="button"
            className={`nav-left`}
            onClick={() => addRectangle()}
          >
            <Scan size={"20px"} className="mr-2" />
          </button>
          <button type="button" className="nav-mid">
            <MousePointer2 size={"20px"} className="mr-2" />
          </button>
          <button
            type="button"
            className="nav-right"
            onClick={() => removeSelected()}
          >
            <Trash2 size={"20px"} className="mr-2" />
          </button>
        </div>
      </div>
    </div>
  );
};

