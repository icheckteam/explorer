import React, { Component } from 'react';
import { connect } from 'react-redux';

function mapStateToProps(state) {
  return {

  };
}

class HomeContainer extends Component {
  render() {
    return (
      <div>
        HomeContainer
      </div>
    );
  }
}

export default connect(
  mapStateToProps,
)(HomeContainer);