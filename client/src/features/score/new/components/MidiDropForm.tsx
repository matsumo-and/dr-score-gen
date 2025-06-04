'use client';

import React, { useRef, useState } from "react";
import { Midi } from "@tonejs/midi";

export default function MidiDropForm() {
  const [selectedFile, setSelectedFile] = useState<File | null>(null);
  const [dragActive, setDragActive] = useState(false);
  const inputRef = useRef<HTMLInputElement>(null);

  const [midiJson, setMidiJson] = useState<any>(null);

  const handleFile = async (file: File) => {
    if (file && file.name.endsWith('.mid')) {
      setSelectedFile(file);
      const arrayBuffer = await file.arrayBuffer();
      const midi = new Midi(arrayBuffer);
      const json = midi.toJSON();
      setMidiJson(json);
    } else {
      alert("MIDIファイル（.mid）のみアップロードできます。");
    }
  };

  const handleDrop = (e: React.DragEvent<HTMLDivElement>) => {
    e.preventDefault();
    setDragActive(false);
    const file = e.dataTransfer.files[0];
    handleFile(file);
  };

  const handleDragOver = (e: React.DragEvent<HTMLDivElement>) => {
    e.preventDefault();
    setDragActive(true);
  };

  const handleDragLeave = (e: React.DragEvent<HTMLDivElement>) => {
    e.preventDefault();
    setDragActive(false);
  };

  const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files?.[0];
    if (!file) return;
    handleFile(file);
  };

  return (
    <>
    <div
      className={`w-full max-w-md border-2 border-dashed rounded-xl p-8 flex flex-col items-center justify-center transition-colors cursor-pointer ${dragActive ? "border-indigo-500 bg-indigo-50" : "border-indigo-300 bg-white"}`}
      onDrop={handleDrop}
      onDragOver={handleDragOver}
      onDragLeave={handleDragLeave}
      onClick={() => inputRef.current?.click()}
    >
      <input
        ref={inputRef}
        type="file"
        accept=".mid"
        className="hidden"
        onChange={handleFileChange}
      />
      <span className="text-indigo-700 font-semibold text-lg mb-2">MIDIファイルをドラッグ＆ドロップ</span>
      <span className="text-indigo-400 text-sm">またはクリックして選択</span>
      {selectedFile && (
        <div className="mt-4 text-indigo-700">選択中: {selectedFile.name}</div>
      )}
    </div>
          {midiJson && (
            <div className="w-full max-w-2xl bg-white rounded-xl shadow p-6 mt-8 overflow-x-auto">
              <h2 className="text-lg font-bold text-indigo-700 mb-2">MIDI JSON</h2>
              <pre className="text-xs text-indigo-900 whitespace-pre-wrap break-all">
                {JSON.stringify(midiJson, null, 2)}
              </pre>
            </div>
          )}
          </>
  );
} 