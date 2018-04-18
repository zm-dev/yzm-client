import React from 'react';
import './uploader.css';
import Drag from './Drag';
import FileList from './FileList';
import ProgressBtn from '../progress-btn/ProgressBtn';

export default class Uploader extends React.PureComponent {
  constructor(props) {
    super(props);
    this.state = {
      files: []
    };
  }

  render() {
    return (
      <div className="uploader_wrapper">
        {
          this.state.files.length > 0 ?
            <div className="file_list_wrapper">
              <FileList files={this.state.files}/>
              <ProgressBtn progress={30}>上传识别</ProgressBtn>
            </div> :
            <Drag onDropFiles={(files) => {
              this.setState({files})
            }}/>
        }
      </div>
    );
  }
}