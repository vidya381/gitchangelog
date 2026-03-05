"use client";

import { useState } from "react";
import Nav from "@/components/Nav";
import RepoForm, { FormValues } from "@/components/RepoForm";
import ChangelogOutput from "@/components/ChangelogOutput";
import { generateChangelog } from "@/lib/api";

type Result = {
  changelog: string;
  commit_count: number;
  from: string;
  to: string;
};

export default function Home() {
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [result, setResult] = useState<Result | null>(null);

  async function handleSubmit(form: FormValues) {
    setLoading(true);
    setError(null);
    setResult(null);

    try {
      const data = await generateChangelog({
        repo_url: form.repoURL,
        token: form.token || undefined,
        from_tag: form.fromTag || undefined,
        to_tag: form.toTag || undefined,
        format: "markdown",
      });
      setResult(data);
    } catch (err) {
      setError(err instanceof Error ? err.message : "Something went wrong");
    } finally {
      setLoading(false);
    }
  }

  return (
    <div className="min-h-screen bg-[#1a1a1a]">
      <Nav />
      <main className="max-w-3xl mx-auto px-6 py-12 flex flex-col gap-8">
        <h1 className="text-xl font-medium text-[#e8e8e8]">
          Generate a changelog from any GitHub repository.
        </h1>
        <RepoForm onSubmit={handleSubmit} loading={loading} error={error} />
        {result && (
          <ChangelogOutput
            changelog={result.changelog}
            commitCount={result.commit_count}
            from={result.from}
            to={result.to}
          />
        )}
      </main>
    </div>
  );
}
