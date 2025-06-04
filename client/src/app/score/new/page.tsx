import React from "react";
import MidiDropForm from "@/features/score/new/components/MidiDropForm";

export default function NewScorePage() {

  return (
    <div className="h-full flex flex-col items-center justify-center py-16 gap-8">
      <h1 className="text-3xl font-bold text-indigo-800 mb-4">サンプル楽譜</h1>
      <p className="text-indigo-600 mb-4">以下はサンプルのドラム譜です。</p>
      <div className="bg-white rounded-xl shadow p-6 flex flex-col items-center mb-8">
        {/* SVGで簡単なドラム譜を表示 */}
        <svg width="320" height="120" viewBox="0 0 320 120" fill="none" xmlns="http://www.w3.org/2000/svg">
          <rect x="0" y="0" width="320" height="120" rx="16" fill="#EEF2FF"/>
          {/* 五線 */}
          {[30, 45, 60, 75, 90].map((y) => (
            <line key={y} x1="24" y1={y} x2="296" y2={y} stroke="#6366F1" strokeWidth="2" />
          ))}
          {/* ハイハット（×印） */}
          <line x1="50" y1="25" x2="60" y2="35" stroke="#F59E42" strokeWidth="2" />
          <line x1="60" y1="25" x2="50" y2="35" stroke="#F59E42" strokeWidth="2" />
          {/* スネア（●） */}
          <circle cx="100" cy="60" r="7" fill="#6366F1" />
          {/* バスドラム（○） */}
          <circle cx="150" cy="90" r="7" fill="#fff" stroke="#6366F1" strokeWidth="2" />
          {/* タム（●） */}
          <circle cx="200" cy="45" r="7" fill="#6366F1" />
          {/* クラッシュシンバル（×印） */}
          <line x1="250" y1="15" x2="260" y2="25" stroke="#F59E42" strokeWidth="2" />
          <line x1="260" y1="15" x2="250" y2="25" stroke="#F59E42" strokeWidth="2" />
        </svg>
      </div>
      <MidiDropForm  />
    </div>
  );
} 