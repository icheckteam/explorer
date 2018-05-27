import React, { Component } from 'react';
import { connect } from 'react-redux';

function mapStateToProps(state) {
  return {

  };
}

class AssetsContainer extends Component {
  render() {
    return (
      <div>
        AssetsContainer
      </div>
    );
  }
}

export default connect(
  mapStateToProps,
)(AssetsContainer);