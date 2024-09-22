import React, { useState } from "react";
import TagInput from "@pathofdev/react-tag-input";
import "@pathofdev/react-tag-input/build/index.css";

const TagForm: React.FC = () => {
  const [tags, setTags] = useState<string[]>([]);
  const maxTags = 3; // 3つの単語まで

  // バリデーション
  const handleTagChange = (newTags: string[]) => {
    if (newTags.length <= maxTags) {
      setTags(newTags);
    } else {
      alert(`${maxTags}つの単語までしか入力することができません！`);
    }
  };

  // 検索処理
  const handleSearch = () => {
    if (tags.length > 0) {
      // 検索(post)処理の実装部分。
      alert(`Searching for: ${tags.join(", ")}`);
    } else {
      alert("単語を入力してください。");
    }
  };

  return (
    <div
      style={{
        maxWidth: "70vh",
        margin: "0 auto",
        padding: "10px 40px 10px 40px",
        position: "fixed", // 固定
        top: "2%", // 画面の上部に固定
        left: "0", // 画面の左端に固定
        right: "0", // 画面の右端に固定
        backgroundColor: "#374654", // 背景色を白に設定（スクロール時の視認性向上のため）
        zIndex: 1000, // 他の要素の上に表示
      }}
    >
      <h3>文献 - 単語検索</h3>
      <div style={{ display: "flex", alignItems: "center" }}>
        <TagInput
          tags={tags}
          onChange={handleTagChange}
          placeholder="単語を入力してください"
        />
        <button
          onClick={handleSearch}
          style={{
            marginLeft: "8px",
            padding: "8px 8px",
            fontSize: "16px",
            cursor: "pointer",
          }}
        >
          Search
        </button>
      </div>
      <p
        style={{ fontSize: "10px", color: "yellow" }}
      >{`${maxTags}つの単語まで入力できます`}</p>
    </div>
  );
};

export default TagForm;
