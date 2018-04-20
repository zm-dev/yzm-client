import React from 'react';
import './uploader.css';
import Drag from './Drag';
import SelectCategory from '../select-category/SelectCategory';
import ProgressBtn from '../progress-btn/ProgressBtn';
import FileItem from './FileItem';
import http from '../../utils/http';
import Complete from './Complete';
import ResultList from '../result-list/ResultList';
import Notifications, {notify} from 'react-notify-toast';

export default class Uploader extends React.PureComponent {
  constructor(props) {
    super(props);
    this.state = {
      files: [],
      uploadedFileNum: null,
      currentCategory: 1,
      showResultListDialog: false,
      downloadUrl: null,
    };
  }

  static getExtName(name) {
    return name.split('.').pop();
  }

  static isAllowFile(name) {
    const extName = Uploader.getExtName(name);
    return extName === 'jpg' || extName === 'jpeg' || extName === 'png' || extName === 'zip';
  }

  async upload() {
    const self = this;
    this.setState({uploadedFileNum: 0});
    for (const currentFile of this.state.files) {
      const form = new FormData();
      form.append('image', currentFile);
      form.append('category', String(this.state.currentCategory));
      const isZip = Uploader.getExtName(currentFile.name) === 'zip';
      try {
        const res = await http.post(isZip ? '/batch_upload' : '/upload', form, {
          headers: {'Content-Type': 'multipart/form-data'},
          onUploadProgress(progressEvent) {
            const {loaded, total} = progressEvent;
            currentFile.progress = loaded / total * 100;
            self.forceUpdate();
          }
        });
        if (!isZip) {
          currentFile.res = res.data.res;
          currentFile.category = res.data.category;
        } else {
          this.setState({downloadUrl: res.data.download_url});
        }
        this.setState({uploadedFileNum: this.state.uploadedFileNum + 1});
      } catch (e) {
        notify.show(currentFile.name + ' ' + e.response.data.message, 'error', 3000);
        currentFile.res = e.response.data.message;
        this.setState({uploadedFileNum: this.state.uploadedFileNum + 1});
      }
    }
    const exceptZipFiles = this.state.files.filter(f => Uploader.getExtName(f.name) !== 'zip');
    if (exceptZipFiles.length > 0) {
      this.setState({showResultListDialog: true});
    }
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
        <Notifications/>
        {this.state.showResultListDialog &&
        <ResultList onClose={this.closeResultListDialog.bind(this)} files={this.state.files}/>}
        <div className="uploader_wrapper">
          {
            this.state.files.length > 0 ?
              this.state.uploadedFileNum >= this.state.files.length ?
                <Complete onDownloaded={() => {
                  this.setState({downloadUrl: null});
                }
                } downloadUrl={this.state.downloadUrl} onReload={this.reload.bind(this)}/> :
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
                            img={Uploader.getExtName(item.name) === 'zip' ? require('../../assets/zip.png') : window.URL.createObjectURL(item)}/>);
                        })
                      }
                    </div>
                    <ProgressBtn
                      progress={
                        this.state.uploadedFileNum !== null ?
                          this.state.uploadedFileNum / this.state.files.length * 100 :
                          null}
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
                  if (Uploader.isAllowFile(file.name))
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