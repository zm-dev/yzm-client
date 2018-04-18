import React from 'react';
import FileItem from './FileItem';
import PropTypes from 'prop-types';

export default class FileList extends React.PureComponent {
  render() {
    return (
      <div className="file_list">
        {
          Object.keys(this.props.files).map(i => {
            const item = this.props.files[i];
            return (<FileItem key={i} fileName={item.name}
                              img={window.URL.createObjectURL(item)}/>)
          })
        }

      </div>
    );
  }
}
FileList.propTypes = {
  files: PropTypes.object.isRequired,
};
