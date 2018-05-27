import React, { Component } from 'react';
import { connect } from 'react-redux';

function mapStateToProps(state) {
  return {

  };
}

class BlocksContainer extends Component {
  render() {
    return (
      <div>
        BlocksContainer
      </div>
    );
  }
}

export default connect(
  mapStateToProps,
)(BlocksContainer);