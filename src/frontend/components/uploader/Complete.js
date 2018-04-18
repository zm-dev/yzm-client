import React from 'react';
import Done from '../icon/Done';
import PropTypes from "prop-types";

export default class Complete extends React.PureComponent {
  render() {
    return (
      <div className="complete">
        <Done/>
        <h2>完成</h2>
        <p className="info">验证码已识别完成</p>
        {
          this.props.downloadUrl ?
            <div>
              <a onClick={this.props.onDownloaded} href={this.props.downloadUrl}>点击下载mapping.txt</a>
              <button onClick={this.props.onReload}>重新上传</button>
            </div> :
            <button onClick={this.props.onReload}>重新上传</button>

        }
      </div>
    );
  }
}

Complete.propTypes = {
  onReload: PropTypes.func.isRequired,
  onDownloaded: PropTypes.func.isRequired,
  downloadUrl: PropTypes.string,
};
Complete.defaultProps = {
  onReload: () => {
  },
  onDownloaded: () => {
  }
};
