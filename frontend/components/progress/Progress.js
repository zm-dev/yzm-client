import React from 'react';
import './progress.css';
import PropTypes from 'prop-types';

export default class Progress extends React.PureComponent{
  render () {
    return (
      <div className="progress_bar">
        <div style={{width: `${this.props.progress}%`}} className="progress" />
      </div>
    );
  }
}
Progress.propTypes = {
  progress: PropTypes.number.isRequired,
};
Progress.defaultProps = {
  progress: 0,
};
