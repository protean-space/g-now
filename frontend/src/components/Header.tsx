import React, { useState } from "react";
import TagInput from "@pathofdev/react-tag-input";
import "@pathofdev/react-tag-input/build/index.css";

const Header: React.FC = () => {
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
    <div className="header-frame">
      <h3 className="header-title">G-NOW</h3>
    </div>
  );
};

export default Header;
