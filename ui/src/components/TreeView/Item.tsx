import { FileIcon, Image, Puzzle } from "lucide-react";
import { ElementTree } from "../../entities/photoshop";
import { usePhotoshopFiles } from "../../store/files";
import { Spinner } from "flowbite-react";

export const Item: React.FC<{ item: ElementTree }> = ({ item }) => {
  const { isSelected, selectItem, isLoading } = usePhotoshopFiles((d) => ({
    selectItem: d.selectElement,
    isSelected: d.elementsSelected.includes(item.element?.id),
    isLoading: d.isLoading,
  }));
  return (
    <li className="cursor-pointer">
      <span
        className={`hover:bg-gray-100 transition flex flex-row gap-3  p-2 truncate ${isSelected && "bg-indigo-100"}`}
        onClick={() => selectItem(item.id)}
      >
        <div className="w-5 h-5">
          {item.isComponent ? (
            isLoading && isSelected ? (
              <Spinner />
            ) : (
              <Puzzle color={item.color} />
            )
          ) : isLoading && isSelected ? (
            <Spinner />
          ) : item.isBackground ? (
            <Image />
          ) : (
            <FileIcon />
          )}
        </div>
        {item.name}
      </span>
    </li>
  );
};
