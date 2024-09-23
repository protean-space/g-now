import React from 'react';

interface ModalProps {
  isOpen: boolean;
  onClose: () => void;
  relatedNews: Array<{
    title: string;
    contents: string;
    ArticleUrl: string;
  }>;
}

const Modal: React.FC<ModalProps> = ({ isOpen, onClose, relatedNews }) => {
  if (!isOpen) return null;

  return (
    <div className="modal-overlay" style={modalOverlayStyle}>
      <div className="modal-content" style={modalContentStyle}>
        <button onClick={onClose} style={closeButtonStyle}>Close</button>
        <h3>Related News</h3>
        <div>
          {relatedNews.length > 0 ? (
            relatedNews.map((newsItem, index) => (
              <div key={index}>
                <h5>{newsItem.title}</h5>
                <p>{newsItem.contents}</p>
                <a href={newsItem.ArticleUrl} target="_blank" rel="noopener noreferrer">Read more</a>
              </div>
            ))
          ) : (
            <p>No related news found.</p>
          )}
        </div>
      </div>
    </div>
  );
};

// スタイル（必要に応じてCSSファイルで定義することも可能）
const modalOverlayStyle: React.CSSProperties = {
  position: 'fixed',
  top: 0,
  left: 0,
  width: '100%',
  height: '100%',
  backgroundColor: 'rgba(0, 0, 0, 0.7)',
  display: 'flex',
  justifyContent: 'center',
  alignItems: 'center',
  zIndex: 1000
};

const modalContentStyle: React.CSSProperties = {
  backgroundColor: '#fff',
  padding: '20px',
  borderRadius: '8px',
  maxWidth: '500px',
  width: '100%',
  boxSizing: 'border-box'
};

const closeButtonStyle: React.CSSProperties = {
  position: 'absolute',
  top: '10px',
  right: '10px',
  background: 'red',
  color: 'white',
  border: 'none',
  padding: '5px 10px',
  cursor: 'pointer'
};

export default Modal;