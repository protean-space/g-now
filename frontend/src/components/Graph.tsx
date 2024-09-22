// src/components/Graph.tsx
import React, { useRef, useEffect } from "react";
import ForceGraph3D from "3d-force-graph";
import SpriteText from "three-spritetext";

interface GraphProps {
  data: {
    nodes: Array<{ id: string }>;
    links: Array<{ source: string; target: string }>;
  };
}

const Graph: React.FC<GraphProps> = ({ data }) => {
  const graphRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    if (graphRef.current) {
      const Graph = ForceGraph3D();
      const graph = Graph(graphRef.current);

      // 動的に渡されたグラフデータを使用
      graph.graphData(data);

      // リンクのカスタマイズ
      graph.linkWidth((link: any) => link.value || 1.8);
      graph.linkColor(() => "#FFFFFF");

      // ノードを文字列に設定
      graph.nodeThreeObject((node: any) => {
        const text = new SpriteText(node.id);
        text.color = "#00ff00";
        text.textHeight = 6;
        return text;
      });

      // ウィンドウがリサイズされたときの処理
      window.addEventListener("resize", () => {
        graph.width(window.innerWidth);
        graph.height(window.innerHeight);
      });

      // 初期サイズ設定
      graph.width(window.innerWidth);
      graph.height(window.innerHeight);
    }

    return () => {
      window.removeEventListener("resize", () => {
        if (graphRef.current) {
          const graph = ForceGraph3D()(graphRef.current);
          graph.width(window.innerWidth);
          graph.height(window.innerHeight);
        }
      });
    };
  }, [data]);  // dataに依存して再レンダリング

  return <div ref={graphRef} style={{ width: "100%", height: "100%" }} />;
};

export default Graph;