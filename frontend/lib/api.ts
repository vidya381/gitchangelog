export async function generateChangelog(payload: {
  repo_url: string;
  token?: string;
  from_tag?: string;
  to_tag?: string;
  format: "markdown" | "plaintext";
}) {
  const res = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/api/changelog`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(payload),
  });
  if (!res.ok) throw new Error(await res.text());
  return res.json();
}
