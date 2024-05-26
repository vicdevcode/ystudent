import { Button } from "@/components/ui/button";
import {
  Command,
  CommandEmpty,
  CommandGroup,
  CommandInput,
  CommandItem,
  CommandList,
} from "@/components/ui/command";
import { FormControl } from "@/components/ui/form";
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover";
import { cn } from "@/lib/utils";
import { Check, ChevronsUpDown } from "lucide-react";
import { FC } from "react";

interface ComboboxProps {
  value: string;
  array: {
    id: string;
    name: string;
  }[];
  placeholder: string;
  notfound: string;
  setValue: (value: string) => void;
}

const Combobox: FC<ComboboxProps> = ({
  value,
  array,
  placeholder,
  notfound,
  setValue,
}) => {
  return (
    <Popover>
      <PopoverTrigger asChild>
        <FormControl>
          <Button
            variant="outline"
            role="combobox"
            className={cn(
              "w-[200px] justify-between",
              !value && "text-muted-foreground",
            )}
          >
            {value ? array.find((el) => el.id === value)?.name : placeholder}
            <ChevronsUpDown className="ml-2 h-4 w-4 shrink-0 opacity-50" />
          </Button>
        </FormControl>
      </PopoverTrigger>
      <PopoverContent className="w-[200px] p-0">
        <Command>
          <CommandInput placeholder="Выберите факультет" />
          <CommandList>
            <CommandEmpty>{notfound}</CommandEmpty>
            <CommandGroup>
              {array.map((el) => (
                <CommandItem
                  value={el.name}
                  key={el.id}
                  onSelect={() => setValue(el.id)}
                >
                  <Check
                    className={cn(
                      "mr-2 h-4 w-4",
                      el.id === value ? "opacity-100" : "opacity-0",
                    )}
                  />
                  {el.name}
                </CommandItem>
              ))}
            </CommandGroup>
          </CommandList>
        </Command>
      </PopoverContent>
    </Popover>
  );
};

export default Combobox;
