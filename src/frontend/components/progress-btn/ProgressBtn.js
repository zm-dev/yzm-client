import React from 'react';
import './progress-btn.css';
import PropTypes from 'prop-types';

export default class ProgressBtn extends React.PureComponent {
  render() {
    return (
      <div className="progress_btn">
        <button disabled className="btn">{this.props.children}</button>
        <div style={{width: `${this.props.progress}%`}} className="progress" />
      </div>
    );
  }
};
ProgressBtn.propTypes = {
  progress: PropTypes.number.isRequired,
};

ProgressBtn.defaultProps = {
  progress: 0,
};

