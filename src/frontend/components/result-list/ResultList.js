import React from 'react';
import './result-list.css';
import PropTypes from 'prop-types';

export default class ResultList extends React.PureComponent {
  render() {
    return (<div className="result_list_wrapper">
      <button onClick={this.props.onClose} className="close_btn"><img src={require('../../assets/close.svg')} alt=""/></button>

      <section className="body">
        <table className="result_list_table">
          <thead>
            <tr>
              <th className="pic">图片</th>
              <th>文件名</th>
              <th>识别结果</th>
            </tr>
          </thead>
          <tbody>
            {this.props.files.map((item, i) => {
              return (
                <tr key={i}>
                  <td><img src={window.URL.createObjectURL(item)} alt={item.name}/></td>
                  <td>{item.name}</td>
                  <td>{item.res}</td>
                </tr>);
            })}
          </tbody>
        </table>
      </section>
    </div>);
  }
}

ResultList.propTypes = {
  files: PropTypes.array.isRequired,
  onClose: PropTypes.func.isRequired,
};
ResultList.defaultProps = {
  onClose: () => {}
};

