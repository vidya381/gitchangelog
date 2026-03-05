export default function Nav() {
  return (
    <nav className="border-b border-[#3a4a52] px-6 py-3 flex items-center justify-between">
      <span className="text-[#e8e8e8] font-medium">gitchangelog</span>
      <a
        href="https://github.com/vidya381/gitchangelog"
        target="_blank"
        rel="noopener noreferrer"
        className="text-[#00acd7] text-sm hover:text-[#0099be] transition-colors"
      >
        GitHub ↗
      </a>
    </nav>
  );
}
