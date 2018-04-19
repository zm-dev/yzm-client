import React from 'react';
import PropTypes from 'prop-types';
import Idle from '../icon/Idle';

export default class Uploader extends React.PureComponent {
  constructor(props) {
    super(props);
    this.state = {
      onDrag: false,
      fileInput: null,
    };
  }

  onDragEnter() {
    this.setState({onDrag: true});
  }

  onDragLeave() {
    this.setState({onDrag: false});
  }

  onDrop(e) {
    e.preventDefault();
    this.props.onDropFiles(e.dataTransfer.files);
  }

  onChange() {
    this.props.onDropFiles(this.state.fileInput.files);
  }

  render() {
    return (
      <div className={`drag${this.state.onDrag ? ' on_drag' : ''}`}>
        <input onChange={this.onChange.bind(this)} ref={(fileInput) => {
          this.state.fileInput = fileInput
        }} className="file_input" type="file"/>
        <div onClick={() => {
          this.state.fileInput.click();
        }} className="border">
          <div
            className="drag_box"
            onDragEnter={this.onDragEnter.bind(this)}
            onDragLeave={this.onDragLeave.bind(this)}
            onDragOver={(e) => {
              e.preventDefault();
            }}
            onDrop={this.onDrop.bind(this)}
          />
          <Idle/>
          <h2>{this.state.onDrag ? 'Drop' : 'Drag & Drop'}</h2>
          <p className="info">拖入图片或点击上传</p>
        </div>
      </div>
    );
  }
}
Uploader.propTypes = {
  onDropFiles: PropTypes.func.isRequired,
};
