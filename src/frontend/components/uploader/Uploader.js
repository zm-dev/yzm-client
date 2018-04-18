import React from 'react';
import './uploader.css';
import Drag from './Drag';
import SelectCategory from '../select-category/SelectCategory';
import ProgressBtn from '../progress-btn/ProgressBtn';
import FileItem from './FileItem';
import http from '../../utils/http';
import Complete from './Complete';
import ResultList from '../result-list/ResultList';

export default class Uploader extends React.PureComponent {
  constructor(props) {
    super(props);
    this.state = {
      files: [],
      uploadedFileNum: null,
      currentCategory: 1,
      showResultListDialog: false,
    };
  }

  async upload() {
    const self = this;
    for (const currentFile of this.state.files) {
      const form = new FormData();
      form.append('image', currentFile);
      form.append('category', String(this.state.currentCategory));
      const res = await http.post('/upload', form, {
        headers: {'Content-Type': 'multipart/form-data'},
        onUploadProgress(progressEvent) {
          currentFile.progress = progressEvent.loaded / progressEvent.total * 100;
          self.forceUpdate();
        }
      });
      currentFile.res = res.data.res;
      this.setState({uploadedFileNum: this.state.uploadedFileNum + 1});
    }
    this.setState({showResultListDialog: true});
  }

  closeResultListDialog() {
    this.setState({showResultListDialog: false});
  }

  reload() {
    this.setState({
      files: [],
      uploadedFileNum: null,
      currentCategory: 1,
      showResultListDialog: false,
    });
  }

  render() {
    return (
      <React.Fragment>
        {this.state.showResultListDialog &&
        <ResultList onClose={this.closeResultListDialog.bind(this)} files={this.state.files}/>}
        <div className="uploader_wrapper">
          {
            this.state.files.length > 0 ?
              this.state.uploadedFileNum >= this.state.files.length ?
                <Complete onReload={this.reload.bind(this)}/> :
                <React.Fragment>
                  <div className="select_category_wrapper">
                    <SelectCategory onSelect={(i) => {
                      this.setState({currentCategory: i});
                    }}/>
                  </div>
                  <div className="file_list_wrapper">
                    <div className="file_list">
                      {
                        this.state.files.map((item, i) => {
                          return (<FileItem
                            progress={item.progress}
                            key={i}
                            fileName={item.name}
                            img={window.URL.createObjectURL(item)}/>);
                        })
                      }
                    </div>
                    <ProgressBtn
                      progress={this.state.uploadedFileNum !== null ? this.state.uploadedFileNum / this.state.files.length * 100 : null}
                      onClick={this.upload.bind(this)}>
                      {typeof this.state.uploadedFileNum === 'number' && this.state.uploadedFileNum >= 0 ?
                        `上传中(${this.state.uploadedFileNum} / ${this.state.files.length})` :
                        '上传识别'}
                    </ProgressBtn>
                  </div>
                </React.Fragment> :
              <Drag onDropFiles={(files) => {
                let fileArray = [];
                for (const file of files) {
                  fileArray.push(file);
                }
                this.setState({files: fileArray})
              }}/>
          }
        </div>
      </React.Fragment>
    );
  }
}