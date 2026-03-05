"use client";

type Props = {
  changelog: string;
  commitCount: number;
  from: string;
  to: string;
};

export default function ChangelogOutput({
  changelog,
  commitCount,
  from,
  to,
}: Props) {
  function handleCopy() {
    navigator.clipboard.writeText(changelog);
  }

  function handleDownload() {
    const blob = new Blob([changelog], { type: "text/markdown" });
    const url = URL.createObjectURL(blob);
    const a = document.createElement("a");
    a.href = url;
    a.download = "CHANGELOG.md";
    a.click();
    URL.revokeObjectURL(url);
  }

  const meta = [
    `${commitCount} commits`,
    from && to ? `${from} → ${to}` : from ? `from ${from}` : to ? `to ${to}` : null,
  ]
    .filter(Boolean)
    .join(" · ");

  return (
    <div className="flex flex-col gap-3">
      {meta && <p className="text-xs text-[#8a9ba8]">{meta}</p>}

      <div className="border border-[#3a4a52] bg-[#2b3a42] p-6 font-mono text-sm text-[#e8e8e8] max-h-[500px] overflow-y-auto whitespace-pre-wrap">
        {changelog}
      </div>

      <div className="flex gap-3">
        <button
          onClick={handleCopy}
          className="border border-[#00acd7] text-[#00acd7] px-4 py-2 text-sm font-medium hover:bg-[#00acd7]/10 transition-colors"
        >
          Copy
        </button>
        <button
          onClick={handleDownload}
          className="border border-[#00acd7] text-[#00acd7] px-4 py-2 text-sm font-medium hover:bg-[#00acd7]/10 transition-colors"
        >
          Download .md
        </button>
      </div>
    </div>
  );
}
