import { Checkbox } from "@/components/ui/checkbox";

interface LabeledCheckboxProps {
  id: string;
  label: string;
  description?: string;
  disabled?: boolean | false;
  value: boolean | undefined;
  onChange: (checked: boolean) => void;
}

export default function LabeledCheckbox({
  id,
  label,
  description,
  disabled,
}: LabeledCheckboxProps) {
  return (
    <div className="items-top flex space-x-2">
      <Checkbox id={id} disabled={disabled} />
      <div className="grid gap-1.5 leading-none">
        <label
          htmlFor={id}
          className="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
        >
          {label}
        </label>
        {description && (
          <p className="text-sm text-muted-foreground">{description}</p>
        )}
      </div>
    </div>
  );
}
