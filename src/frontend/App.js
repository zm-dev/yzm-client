import React, {Component} from 'react';
import Uploader from './components/uploader/Uploader';
import './App.css';

class App extends Component {
  render() {
    return (
      <div className="container">
        <div className="left">
          <div className="logo">
            <img src={require('./assets/logo.svg')} alt="A16 验证码识别"/>
          </div>
          <h1>A16 验证码识别</h1>
          <div className="copy">
            copyright &copy; 知明
          </div>
        </div>
        <div className="right">
          <Uploader/>
        </div>
      </div>
    );
  }
}

export default App;
