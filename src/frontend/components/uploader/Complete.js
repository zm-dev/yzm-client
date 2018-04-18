import React from 'react';
import Done from '../icon/Done';
import ResultList from "../result-list/ResultList";
import PropTypes from "prop-types";

export default class Complete extends React.PureComponent {
  render() {
    return (
      <div className="complete">
        <Done/>
        <h2>完成</h2>
        <p className="info">验证码已识别完成</p>
        <button onClick={this.props.onReload}>重新上传</button>
      </div>
    );
  }
}

Complete.propTypes = {
  onReload: PropTypes.func.isRequired,
};
Complete.defaultProps = {
  onReload: () => {
  }
};
