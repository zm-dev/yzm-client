import React from 'react';
import './uploader.css';
import Drag from './Drag';
import SelectCategory from '../select-category/SelectCategory';
import ProgressBtn from '../progress-btn/ProgressBtn';
import FileItem from './FileItem';
import http from '../../utils/http';

export default class Uploader extends React.PureComponent {
  constructor(props) {
    super(props);
    this.state = {
      files: [],
      totalProgress: null,
      currentCategory: 1,
    };
  }

  upload() {
    this.setState({totalProgress: 0});
    Object.keys(this.state.files).forEach(async (i) => {
      const form = new FormData();
      const currentFile = this.state.files[i];
      form.append('image', currentFile);
      form.append('category', this.state.currentCategory);
      const self = this;
      const res = await http.post('/upload', form, {
        headers: {'Content-Type': 'multipart/form-data'},
        onUploadProgress(progressEvent) {
          currentFile.progress = progressEvent.loaded / progressEvent.total * 100;
          self.forceUpdate();
        }
      });
      this.setState({totalProgress: this.state.totalProgress + 1});
    });
  }

  render() {
    return (
      <div className="uploader_wrapper">
        {
          this.state.files.length > 0 ?
            <React.Fragment>
              <div className="select_category_wrapper">
                <SelectCategory onSelect={(i) => {
                  this.setState({currentCategory: i});
                }}/>
              </div>
              <div className="file_list_wrapper">
                <div className="file_list">
                  {
                    Object.keys(this.state.files).map(i => {
                      const item = this.state.files[i];
                      return (<FileItem
                        progress={item.progress}
                        key={i}
                        fileName={item.name}
                        img={window.URL.createObjectURL(item)}/>);
                    })
                  }
                </div>
                <ProgressBtn
                  progress={this.state.totalProgress}
                  onClick={this.upload.bind(this)}>
                  {typeof this.state.totalProgress === 'number' && this.state.totalProgress >= 0 ?
                    `上传中(${this.state.totalProgress} / ${this.state.files.length})` :
                    '上传识别'}
                </ProgressBtn>
              </div>
            </React.Fragment> :
            <Drag onDropFiles={(files) => {
              this.setState({files})
            }}/>
        }
      </div>
    );
  }
}