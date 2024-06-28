type Props = {
  label: string;
  name: string;
};

export function Input({ label = "" }: Props) {
  return (
    <div className="flex w-full max-w-xs flex-col gap-1 text-slate-700 dark:text-slate-300">
      <label className="w-fit pl-0.5 text-sm">
        {label}
      </label>
      <input
        id="textInputDefault"
        type="text"
        className="w-full rounded-xl border border-slate-300 bg-slate-100 px-2 py-2.5 text-sm focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-blue-700 disabled:cursor-not-allowed disabled:opacity-75 dark:border-slate-700 dark:bg-slate-800/50 dark:focus-visible:outline-blue-600"
        name="name"
        placeholder="Enter your name"
      />
    </div>
  );
}
