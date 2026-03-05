"use client";

import { useState } from "react";

const inputClass =
  "w-full bg-transparent border border-[#3a4a52] text-[#e8e8e8] placeholder-[#8a9ba8] px-3 py-2 text-sm focus:outline-none focus:border-[#00acd7] transition-colors";

export type FormValues = {
  repoURL: string;
  fromTag: string;
  toTag: string;
  token: string;
};

type Props = {
  onSubmit: (values: FormValues) => void;
  loading: boolean;
  error: string | null;
};

export default function RepoForm({ onSubmit, loading, error }: Props) {
  const [form, setForm] = useState<FormValues>({
    repoURL: "",
    fromTag: "",
    toTag: "",
    token: "",
  });

  function handleSubmit(e: React.FormEvent) {
    e.preventDefault();
    onSubmit(form);
  }

  return (
    <form onSubmit={handleSubmit} className="flex flex-col gap-4">
      <div>
        <input
          type="text"
          placeholder="https://github.com/owner/repo"
          value={form.repoURL}
          onChange={(e) => setForm({ ...form, repoURL: e.target.value })}
          className={`${inputClass} ${error ? "border-[#ef4444]" : ""}`}
          required
        />
        {error && <p className="mt-1 text-xs text-[#ef4444]">{error}</p>}
      </div>

      <div className="grid grid-cols-1 sm:grid-cols-2 gap-4">
        <input
          type="text"
          placeholder="From tag (optional)"
          value={form.fromTag}
          onChange={(e) => setForm({ ...form, fromTag: e.target.value })}
          className={inputClass}
        />
        <input
          type="text"
          placeholder="To tag (optional)"
          value={form.toTag}
          onChange={(e) => setForm({ ...form, toTag: e.target.value })}
          className={inputClass}
        />
      </div>

      <div className="relative">
        <input
          type="password"
          placeholder="GitHub token (optional)"
          value={form.token}
          onChange={(e) => setForm({ ...form, token: e.target.value })}
          className={`${inputClass} pr-8`}
        />
        <span className="absolute right-3 top-[10px] text-[#8a9ba8] text-sm">
          🔒
        </span>
        <p className="mt-1 text-xs text-[#8a9ba8]">
          Used to increase GitHub API rate limits. Never stored.
        </p>
      </div>

      <div className="flex justify-end">
        <button
          type="submit"
          disabled={loading}
          className="bg-[#00acd7] text-white px-4 py-2 text-sm font-medium hover:bg-[#0099be] transition-colors disabled:opacity-50"
        >
          {loading ? "Generating..." : "Generate"}
        </button>
      </div>
    </form>
  );
}
