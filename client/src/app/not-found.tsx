import PrimaryButton from "./components/PrimaryButton";

export default function NotFound() {
  return (
      <main className="h-full flex-1 flex flex-col items-center justify-center gap-8">
        <h1 className="text-5xl font-bold text-indigo-700">404</h1>
        <p className="text-lg text-indigo-600">お探しのページは見つかりませんでした。</p>
        <PrimaryButton title="ホームに戻る" href="/" />
      </main>
  );
} 