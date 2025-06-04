import Image from "next/image";

export default function Home() {
  return (
    <div className="min-h-screen flex flex-col bg-gradient-to-br from-blue-50 to-indigo-100 font-[family-name:var(--font-geist-sans)]">
      {/* ヘッダー */}
      <header className="flex items-center justify-between px-8 py-4 bg-white/80 shadow-md backdrop-blur-md">
        <div className="flex items-center gap-2">
          <Image src="/drum-favicon.svg" alt="Logo" width={36} height={36} />
          <span className="text-2xl font-bold tracking-tight text-indigo-700">Dr.ScoreGen</span>
        </div>
        <nav className="flex gap-6 text-indigo-700 font-medium">
          <a href="#" className="hover:underline">ホーム</a>
          <a href="#" className="hover:underline">マイスコア</a>
          <a href="#" className="hover:underline">ヘルプ</a>
        </nav>
        <button className="bg-indigo-600 text-white px-4 py-2 rounded-lg shadow hover:bg-indigo-700 transition">ログイン</button>
      </header>

      {/* メイン */}
      <main className="flex-1 flex flex-col items-center justify-center py-16 gap-12">
        {/* メインビジュアル */}
        <div className="flex flex-col items-center gap-4">
          <h1 className="text-4xl sm:text-5xl font-bold text-indigo-800 mb-2">あなたの音楽をスコアに</h1>
          <p className="text-lg text-indigo-600 mb-6">Dr.ScoreGenで簡単に楽譜を作成・管理しましょう</p>
          <button className="bg-gradient-to-r from-indigo-500 to-blue-500 text-white px-8 py-3 rounded-full text-lg font-semibold shadow-lg hover:scale-105 transition">新しいスコアを作成</button>
        </div>

        {/* 最近のスコア一覧（ダミー） */}
        <section className="w-full max-w-2xl bg-white/80 rounded-xl shadow p-6">
          <h2 className="text-xl font-semibold text-indigo-700 mb-4">最近のスコア</h2>
          <ul className="divide-y divide-indigo-100">
            <li className="py-3 flex items-center justify-between">
              <span className="font-medium text-indigo-900">ピアノ練習曲 第1番</span>
              <span className="text-sm text-indigo-400">2024/06/01</span>
            </li>
            <li className="py-3 flex items-center justify-between">
              <span className="font-medium text-indigo-900">バンドスコア「青春」</span>
              <span className="text-sm text-indigo-400">2024/05/28</span>
            </li>
            <li className="py-3 flex items-center justify-between">
              <span className="font-medium text-indigo-900">オリジナル曲「夜空」</span>
              <span className="text-sm text-indigo-400">2024/05/20</span>
            </li>
          </ul>
        </section>
      </main>

      {/* フッター */}
      <footer className="text-center text-indigo-500 py-4 bg-white/60 text-sm">
        © 2024 Dr.ScoreGen. All rights reserved.
      </footer>
    </div>
  );
}
