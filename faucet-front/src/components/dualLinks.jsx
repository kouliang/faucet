import React from 'react';

const DualLinks = ({ link1Text, link2Text }) => {
  const handleLinkClick = (link) => {
    // 根据链接内容判断要跳转的地址
    let url;
    if (link === link1Text) {
      url = 'https://www.baidu.com';
    } else if (link === link2Text) {
      url = 'https://www.baidu.com';
    }

    // 在新标签中打开跳转的地址
    if (url) {
      window.open(url, '_blank');
    }
  };

  return (
    <div>
      <a className="link" onClick={() => handleLinkClick(link1Text)}>
        {link1Text}
      </a>
      <span className="link-separator" />
      <a className="link" onClick={() => handleLinkClick(link2Text)}>
        {link2Text}
      </a>
    </div>
  );
};

export default DualLinks;