import React from 'react';
import PropTypes from 'prop-types';
import Progress from '../progress/Progress';

export default class FileItem extends React.PureComponent {
  render() {
    return (
      <div className="file_item">
        <div className="cover">
          <img
            src={this.props.img}
            alt=""/>
          <Progress progress={10} className="progress_bar"/>
        </div>
        <p className="file_name" title={this.props.fileName}>{this.props.fileName}</p>
      </div>
    );
  }
}

FileItem.propTypes = {
  img: PropTypes.string.isRequired,
  fileName: PropTypes.string.isRequired,
};
