import Image from "next/image";

export default function Header() {
  return (
    <header className="flex items-center justify-between px-8 py-4 bg-white/80 shadow-md backdrop-blur-md">
      <div className="flex items-center gap-2">
        <Image src="/music-logo.svg" alt="Logo" width={36} height={36} />
        <span className="text-2xl font-bold tracking-tight text-indigo-700">Dr.ScoreGen</span>
      </div>
      <nav className="flex gap-6 text-indigo-700 font-medium">
        <a href="#" className="hover:underline">ホーム</a>
        <a href="#" className="hover:underline">マイスコア</a>
        <a href="#" className="hover:underline">ヘルプ</a>
      </nav>
      <button className="bg-indigo-600 text-white px-4 py-2 rounded-lg shadow hover:bg-indigo-700 transition">ログイン</button>
    </header>
  );
} 