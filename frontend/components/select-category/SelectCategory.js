import React from 'react';
import './select-category.css';
import PropTypes from 'prop-types';

export default class SelectCategory extends React.PureComponent{
  constructor (props) {
    super(props);
    this.state = {
      current: 0,
    };
    this.props.onSelect(this.state.current);
  }
  render() {
    let categories = [];
    for (let i = 0; i < 5; i++) {
      categories.push(<div onClick={() => {
        this.setState({current: i});
        this.props.onSelect(i);
      }} key={i} className={`item${this.state.current === i ? ' active' : ''}`}>第{i + 1}类</div>)
    }
    return (
      <div className="select_category">
        <span>
        {categories}
        </span>
      </div>);
  }
}


SelectCategory.propTypes = {
  onSelect: PropTypes.func
};
SelectCategory.defaultProps = {
  onSelect: () => {}
};
