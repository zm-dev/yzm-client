import React from 'react';
import './progress-btn.css';
import PropTypes from 'prop-types';

export default class ProgressBtn extends React.PureComponent {
  isLoading(progress) {
    return typeof progress === 'number' && progress >= 0;
  }

  render() {
    return (
      <div className="progress_btn"
           disabled={this.isLoading(this.props.progress)}
           onClick={() => {
             !this.isLoading(this.props.progress) && this.props.onClick();
           }}>
        <span className="text">{this.props.children}</span>
        {this.isLoading(this.props.progress) && <div style={{width: `${this.props.progress}%`}} className="progress"/>}
      </div>
    );
  }
};
ProgressBtn.propTypes = {
  progress: PropTypes.number,
  onClick: PropTypes.func,
};

