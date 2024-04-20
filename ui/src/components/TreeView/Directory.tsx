import { Folder } from "lucide-react";
import { TreeView } from ".";
import { ElementTree } from "../../entities/photoshop";

export const Directory = ({
  item,
}: React.PropsWithChildren<{
  item: ElementTree;
}>): JSX.Element => {
  return (
    <li className="cursor-pointer">
      <span className=" flex flex-col transition p-2 truncate">
        <div className="flex flex-row gap-3">
          <div className="w-6 h-6">
            <Folder />
          </div>
          {item.name}
        </div>
        <TreeView tree={item} />
      </span>
    </li>
  );
};
